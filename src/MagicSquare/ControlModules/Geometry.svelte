<script lang="ts">
  import { TransformOrder } from "./TransformOrder"
  import { WasmInputId } from "../WasmInputId"
  import { intoShape, Shape } from "./Shape"

  export let transformOrder: TransformOrder
  export let shapes: Shape[]

  let shapeIndex: number = 0
  let shape: Shape

  $: shapes[shapeIndex] = shape

  function handleShapeIndexSelect(e: any, new_idx: number) {
    e.stopPropagation()
    console.log(shapes)
    shapeIndex = new_idx
    shape = shapes[new_idx]
  }

  function handleShapeSelect(e: any) {
    e.stopPropagation()
    var input = document.getElementById(WasmInputId.shapes)
    shape = intoShape(e.target.value.trim())
    shapes[shapeIndex] = shape
    shapes = [...shapes]
    input.value = JSON.stringify({shape, index: shapeIndex})
    input.dispatchEvent(new Event('input', {bubbles: true}))
  } 

  function handleTransformOrderClick(tr_or: TransformOrder) {
    transformOrder = tr_or
    var input = document.getElementById(WasmInputId.transformOrder)
    input.value = tr_or
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<div class="h-full flex flex-col justify-between items-stretch">
  <div class="grow flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      shape
    </div>
    <slot name="shapes"/>
    <div class="pl-5 pr-5 grid grid-cols-4">
      <select class="col-span-3 shape_select"
              bind:value={shape}
              on:input={(e) => e.stopPropagation()}
              on:change={handleShapeSelect}>
        <optgroup label="2d">
          <option value={Shape.triangle}>
            Triangle
          </option>
          <option>
            Square
          </option>
          <option>
            Pentagon
          </option>
          <option value={Shape.hexagon}>
            Hexagon
          </option>
          <option>
            Heptagon
          </option>
          <option>
            Octagon
          </option>
        </optgroup>
        <optgroup label="3d">
          <option>
            Tetrahedron
          </option>
          <option>
            Cube
          </option>
          <option>
            Octohedron
          </option>
          <option>
            Dodecahedron
          </option>
          <option value={Shape.icosahedron}>
            Icosahedron
          </option>
        </optgroup>
      </select>
      <button class="col-span-1">
        all
      </button>
    </div>
    <div class="pt-2 pl-5 pr-5 grid grid-cols-4 grid-rows-4 gap-0">
      {#each {length: 16} as _, i}
        <button on:click={(e) => handleShapeIndexSelect(e, i)}
                class:selected={shapeIndex === i}>
          {i + 1}
        </button>
      {/each}
    </div>
  </div>
  <div class="grow flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      transform order
    </div>
    <div class="pl-2 pr-2 flex justify-around">
      <button class="grow flex justify-around items-center"
              class:selected={transformOrder === TransformOrder.rotateThenTranslate}
              on:click={() => handleTransformOrderClick(TransformOrder.rotateThenTranslate)}>
        R -> T
      </button>
      <button class="grow flex justify-around items-center"
              class:selected={transformOrder === TransformOrder.translateThenRotate}
              on:click={() => handleTransformOrderClick(TransformOrder.translateThenRotate)}>
        T -> R
      </button>
    </div>
    <slot name="transformOrder"/>
  </div>
  <div class="grow flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      radius
    </div>
    <slot name="radiusSliders"/>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .title
    color: color.$blue-7
    font-size: text.$fs-ml
    font-weight: text.$fw-l

  .shape_select
    font-size: 1.3em
    cursor: pointer

  .selected
    background-color: color.$blue-7
</style>
