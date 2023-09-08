<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import { WasmInputId } from "../WasmInputId"
  import { ShapeTag } from "./Shape"
  import type { Shape } from "./Shape"
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'

  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)
  let i18n = new I18n("magicSquare/geometry")

  import { msStoreSettings } from '../../stores/msStoreSettings'
  import type { MsStoreSettings } from '../../stores/msStoreSettings'
  let msStoreSettingsVal: MsStoreSettings
  const unsubMsStoreSettings = msStoreSettings.subscribe((val: MsStoreSettings) => msStoreSettingsVal = val)

  import { currSquare, SquareType } from '../../stores/currSquare'
  let currSquareVal: SquareType
  const unsubCurrSquare = currSquare.subscribe((val: SquareType) => currSquareVal = val)

  export let shapes: Shape[]

  let shapeIdx: number = 0
  let n: number = 0 

  let idxA: number
  let idxB: number

  $: idxLeft = idxA < idxB ? idxA : idxB
  $: idxRight = idxA > idxB ? idxA : idxB

  function handleShapeIndexSelect(e: any, new_idx: number) {
    e.stopPropagation()
    shapeIdx = new_idx
    updateStoreShapeIdx(new_idx)
    n = shapes[shapeIdx].c
    setShapeSelectValue(shapeIdx)
  }

  function setShapeSelectValue(shapeIdx: number) {
    var sel = document.getElementById('magic_square_shape_select')
    sel.value = JSON.stringify({t: shapes[shapeIdx].t, c: n})
  }

  function updateStoreShapeIdx(shapeIdx: number) {
    msStoreSettings.update((prevSettings: MsStoreSettings) => {
      switch (currSquareVal) {
        case SquareType.magic:
          prevSettings.msGeometryShapeIdx = shapeIdx
          break
        case SquareType.public:
          prevSettings.psGeometryShapeIdx = shapeIdx
          break
      }
      return prevSettings
    })
  }

  function handleShapeSelect(e: any) {
    e.stopPropagation()
    var input = document.getElementById(WasmInputId.shapes)
    const new_shape: Shape = JSON.parse(e.target.value)

    if (shapes[shapeIdx].t === ShapeTag.misc && new_shape.t === ShapeTag.ngon) { 
      // if prevShapeTag == misc && newShapeTag == nGon
      n = 3
      new_shape.c = n
    } else {
      n = new_shape.c
    }

    shapes[shapeIdx] = new_shape
    shapes = [...shapes]

    input.value = JSON.stringify({shape: {t: new_shape.t, c: n}, index: shapeIdx})
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  function handleRangeIndexChange(e: any, id: string) {
    e.stopPropagation()
    switch (id) {
      case 'a':
        idxA = parseInt(e.target.value)
        break
      case 'b':
        idxB = parseInt(e.target.value)
        break
      default:
        break
    }
    updateStoreRange(idxA, idxB)
  }

  function updateStoreRange(idxA: number, idxB: number) {
    msStoreSettings.update((prevSettings: MsStoreSettings) => {
      switch (currSquareVal) {
        case SquareType.magic:
          prevSettings.msGeometryIdxA = idxA
          prevSettings.msGeometryIdxB = idxB
          break
        case SquareType.public:
          prevSettings.psGeometryIdxA = idxA
          prevSettings.psGeometryIdxB = idxB
          break
      }
      return prevSettings
    })
  }

  function setRange() {
    let width: number = idxRight - idxLeft;
    let input = document.getElementById(WasmInputId.shapes)
    if (!!width) {
      const new_shape = shapes[shapeIdx]
      shapes.forEach((_: Shape, idx: number) => {
        if (idx >= idxLeft && idx <= idxRight) {
          shapes[idx] = new_shape
          shapes = [...shapes]
          input.value = JSON.stringify({shape: shapes[shapeIdx], index: idx})
          input.dispatchEvent(new Event('input', {bubbles: true}))
        }
      })
    }
  }

  function handleNInput() {
    shapes[shapeIdx].c = n
    shapes = [...shapes]
    var input = document.getElementById(WasmInputId.shapes)
    input.value = JSON.stringify({shape: shapes[shapeIdx], index: shapeIdx})
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  onMount(() => {
    switch (currSquareVal) {
      case SquareType.magic:
        idxA = msStoreSettingsVal.msGeometryIdxA
        idxB = msStoreSettingsVal.msGeometryIdxB
        shapeIdx = msStoreSettingsVal.msGeometryShapeIdx
        n = shapes[msStoreSettingsVal.msGeometryShapeIdx].c
        break
      case SquareType.public:
        idxA = msStoreSettingsVal.psGeometryIdxA
        idxB = msStoreSettingsVal.psGeometryIdxB
        shapeIdx = msStoreSettingsVal.psGeometryShapeIdx
        n = shapes[msStoreSettingsVal.psGeometryShapeIdx].c
        break
      case SquareType.none:
        idxA = 0
        idxB = 15
        shapeIdx = 0
        n = 0
        break
    }
  })

  onDestroy(() => {
    unsubCurrSquare()
    unsubLang()
    unsubMsStoreSettings()
  })
</script>

<div class="h-full flex flex-col justify-between items-stretch gap-2">
  <div class="grow flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      {i18n.t("shape", langVal)}
    </div>
    <slot name="shapes"/>
    <div class="grow pt-2 pl-5 pr-5 grid grid-cols-4 grid-rows-4 gap-0">
      {#each {length: 16} as _, i}
        <button class="flex justify-around items-center"
                on:click={(e) => handleShapeIndexSelect(e, i)}
                class:selected={shapeIdx === i}>
          {i + 1}
        </button>
      {/each}
    </div>
    <div class="pt-5 pl-5 pr-5 grid grid-cols-4">
      <select id="magic_square_shape_select"
              class="col-span-3 shape_select"
              value={JSON.stringify({t: shapes[shapeIdx].t, c: n})}
              on:input={(e) => e.stopPropagation()}
              on:change={handleShapeSelect}>
        <optgroup label="Misc.">
          <option value={JSON.stringify({t: ShapeTag.misc, c: 0})}>
            {i18n.t("star", langVal)}
          </option>
          <option value={JSON.stringify({t: ShapeTag.misc, c: 1})}>
            {i18n.t("coolS", langVal)}
          </option>
        </optgroup>
        <optgroup label="2d">
          <option label="Ngon"
                  value={JSON.stringify({t: ShapeTag.ngon, c: n})}>
            {i18n.t("ngon", langVal)}
          </option>
        </optgroup>
        <optgroup label="3d">
          <option value={JSON.stringify({t: ShapeTag.platoThree, c: 4})}>
            {i18n.t("tetrahedron", langVal)}
          </option>
          <option value={JSON.stringify({t: ShapeTag.platoThree, c: 6})}>
            {i18n.t("cube", langVal)}
          </option>
          <option value={JSON.stringify({t: ShapeTag.platoThree, c: 8})}>
            {i18n.t("octahedron", langVal)}
          </option>
          <option value={JSON.stringify({t: ShapeTag.platoThree, c: 12})}>
            {i18n.t("dodecahedron", langVal)}
          </option>
          <option value={JSON.stringify({t: ShapeTag.platoThree, c: 20})}>
            {i18n.t("icosahedron", langVal)}
          </option>
        </optgroup>
      </select>
    </div>
    <div class="pl-5 pr-5 flex flex-col justify-between items-stretch">
      <label class="slider_label flex justify-between" 
             class:disabled={shapes[shapeIdx].t !== ShapeTag.ngon}
             for="ngon_n">
        <div> {"n"} </div>
        <div> {n} </div>
      </label>
      <input id="ngon_n"
             bind:value={n}
             disabled={shapes[shapeIdx].t !== ShapeTag.ngon}
             on:input={handleNInput}
             type="range"
             min={3}
             max={30}
             step={1}/>
    </div>
    <div class="geometry_range rounded-md p-2 m-2 grid grid-cols-2 grid-rows-2 gap-2">
      <select bind:value={idxA}
              class="flex justify-around items-center"
              on:input={(e) => e.stopPropagation()}
              on:change={(e) => handleRangeIndexChange(e, 'a')}>
        {#each {length: 16} as _, idx}
          <option selected={idxA === idx}
                  value={idx}>
            {idx + 1}
          </option>
        {/each}
      </select>
      <select bind:value={idxB}
              class="flex justify-around items-center"
              on:input={(e) => e.stopPropagation()}
              on:change={(e) => handleRangeIndexChange(e, 'b')}>
        {#each {length: 16} as _, idx}
          <option selected={idxB === idx}
                  value={idx}>
            {idx + 1}
          </option>
        {/each}
      </select>
      <button class="col-span-2 flex justify-around items-center"
              on:click={setRange}>
       {i18n.t("range", langVal)}
      </button>
    </div>
  </div>
  <div class="title flex items-stretch pl-5 underline">
    {i18n.t("radius", langVal)}
  </div>
  <div class="grow flex flex-col justify-between items-stretch">
    <slot name="radiusSliders"/>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  @import "../styles/control_module_title.sass"

  .disabled
    color: #666

  .title
    @include control_module_title

  .shape_select
    font-size: 1.3em
    cursor: pointer

  .selected
    background-color: color.$blue-7

  .slider_label
    width: 100%
    font-weight: text.$fw-l
    font-size: text.$fs-m
    padding-right: 5%

  .geometry_range
    border: 2px solid color.$blue-7
</style>
