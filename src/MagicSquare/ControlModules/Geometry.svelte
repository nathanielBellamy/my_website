<script lang="ts">
  import { WasmInputId } from "../WasmInputId"
  import { ShapeTag } from "./Shape"
  import type { Shape } from "./Shape"

  export let shapes: Shape[]

  let shapeIndex: number = 0
  $: shapeJson = JSON.stringify({shape: shapes[shapeIndex], index: shapeIndex})
  let n: number = shapes[shapeIndex].c

  let idx_a: number = 0
  let idx_b: number = 4

  $: idxLeft = idx_a < idx_b ? idx_a : idx_b
  $: idxRight = idx_a > idx_b ? idx_a : idx_b

  function handleShapeIndexSelect(e: any, new_idx: number) {
    e.stopPropagation()
    shapeIndex = new_idx
    n = shapes[shapeIndex].c
    var sel = document.getElementById('magic_square_shape_select')
    sel.value = JSON.stringify({t: shapes[shapeIndex].t, c: n})
  }

  function handleShapeSelect(e: any) {
    e.stopPropagation()
    var input = document.getElementById(WasmInputId.shapes)
    const new_shape: Shape = JSON.parse(e.target.value)
    shapes[shapeIndex] = new_shape
    shapes = [...shapes]
    input.value = JSON.stringify({shape: {t: new_shape.t, c: new_shape.c}, index: shapeIndex})
    input.dispatchEvent(new Event('input', {bubbles: true}))
  } 

  function handleRangeIndexChange(e: any, id: string) {
    e.stopPropagation()
    switch (id) {
      case 'a':
        idx_a = parseInt(e.target.value)
        break
      case 'b':
        idx_b = parseInt(e.target.value)
        break
      default:
        break
    }
  }

  function setRange() {
    let width: number = idxRight - idxLeft;
    let input = document.getElementById(WasmInputId.shapes)
    if (!!width) {
      const new_shape = shapes[shapeIndex]
      shapes.forEach((_: Shape, idx: number) => {
        if (idx >= idxLeft && idx <= idxRight) {
          shapes[idx] = new_shape
          shapes = [...shapes]
          input.value = JSON.stringify({shape: shapes[shapeIndex], index: idx})
          input.dispatchEvent(new Event('input', {bubbles: true}))
        }
      })
    }
  }

  function handleNInput() {
   shapes[shapeIndex].c = n
   shapes = [...shapes]
   var input = document.getElementById(WasmInputId.shapes)
   input.value = JSON.stringify({shape: shapes[shapeIndex], index: shapeIndex})
   input.dispatchEvent(new Event('input', {bubbles: true}))
 }
</script>

<div class="grow pb-5 flex flex-col justify-between items-stretch">
  <div class="grow flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      shape
    </div>
    <slot name="shapes"/>
    <div class="grow pt-2 pl-5 pr-5 grid grid-cols-4 grid-rows-4 gap-0">
      {#each {length: 16} as _, i}
        <button class="flex justify-around items-center"
                on:click={(e) => handleShapeIndexSelect(e, i)}
                class:selected={shapeIndex === i}>
          {i + 1}
        </button>
      {/each}
    </div>
    <div class="pt-5 pl-5 pr-5 grid grid-cols-4">
      <select id="magic_square_shape_select"
              class="col-span-3 shape_select"
              value={JSON.stringify({t: shapes[shapeIndex].t, c: n})}
              on:input={(e) => e.stopPropagation()}
              on:change={handleShapeSelect}>
        <optgroup label="2d">
          <option label="Ngon"
                  value={JSON.stringify({t: ShapeTag.ngon, c: n})}>
            Ngon
          </option>
        </optgroup>
        <optgroup label="3d">
          <option>
            Tetrahedron
          </option>
          <option value={JSON.stringify({t: ShapeTag.platoThree, c: 6})}>
            Cube
          </option>
          <option>
            Octohedron
          </option>
          <option>
            Dodecahedron
          </option>
          <option value={JSON.stringify({t: ShapeTag.platoThree, c: 20})}>
            Icosahedron
          </option>
        </optgroup>
      </select>
    </div>
    <div class="pl-5 pr-5 flex flex-col justify-between items-stretch">
      <label class="slider_label flex justify-between" 
             class:disabled={shapes[shapeIndex].t !== ShapeTag.ngon}
             for="ngon_n">
        <div> {"n"} </div>
        <div> {n} </div>
      </label>
      <input id="ngon_n"
             bind:value={n}
             disabled={shapes[shapeIndex].t !== ShapeTag.ngon}
             on:input={handleNInput}
             type="range"
             min={3}
             max={30}
             step={1}/>
    </div>
    <div class="pt-5 pr-5 pl-5 grid grid-cols-4 grid-rows-1 gap-2">
      <button class="col-span-2 flex justify-around items-center"
              on:click={setRange}>
        Range
      </button>
      <select bind:value={idx_a}
              class="flex justify-around items-center"
              on:input={(e) => e.stopPropagation()}
              on:change={(e) => handleRangeIndexChange(e, 'a')}>
        {#each {length: 16} as _, idx}
          <option selected={idx_a === idx}
                  value={idx}>
            {idx + 1}
          </option>
        {/each}
      </select>
      <select bind:value={idx_b}
              class="flex justify-around items-center"
              on:input={(e) => e.stopPropagation()}
              on:change={(e) => handleRangeIndexChange(e, 'b')}>
        {#each {length: 16} as _, idx}
          <option selected={idx_b === idx}
                  value={idx}>
            {idx + 1}
          </option>
        {/each}
      </select>
    </div>
  </div>
  <div class="grow pt-2 flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      radius
    </div>
    <slot name="radiusSliders"/>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .disabled
    color: #666

  .title
    color: color.$blue-7
    font-size: text.$fs-ml
    font-weight: text.$fw-l

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
</style>
