<script lang="ts">
  import { onMount } from 'svelte'
  import { WasmInputId } from "../WasmInputId"
  import { intoShape, Shape } from "./Shape"

  export let shapes: Shape[]

  let shapeIndex: number = 0

  function handleShapeIndexSelect(e: any, new_idx: number) {
    e.stopPropagation()
    shapeIndex = new_idx
  }

  function handleShapeSelect(e: any) {
    e.stopPropagation()
    var input = document.getElementById(WasmInputId.shapes)
    shapes[shapeIndex] = intoShape(e.target.value.trim())
    shapes = [...shapes]
    input.value = JSON.stringify({shape: shapes[shapeIndex], index: shapeIndex})
    input.dispatchEvent(new Event('input', {bubbles: true}))
  } 

  function handleAllClick() {
    const new_shapes = shapes.map(_ => shapes[shapeIndex])
    shapes = [...new_shapes]
    var input = document.getElementById(WasmInputId.shapes)
    input.value = JSON.stringify({shape: shapes[shapeIndex], index: 16})
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

</script>

<div class="h-full flex flex-col justify-between items-stretch">
  <div class="grow flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      shape
    </div>
    <slot name="shapes"/>
    <div class="pt-5 pl-5 pr-5 grid grid-cols-4">
      <select class="col-span-3 shape_select"
              on:input={(e) => e.stopPropagation()}
              on:change={handleShapeSelect}>
        <optgroup label="2d">
          <option selected={shapes[shapeIndex] === Shape.triangle}
                  value={Shape.triangle}>
            Triangle
          </option>
          <option>
            Square
          </option>
          <option>
            Pentagon
          </option>
          <option selected={shapes[shapeIndex] === Shape.hexagon}
                  value={Shape.hexagon}>
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
          <option selected={shapes[shapeIndex] === Shape.cube}
                  value={Shape.cube}>
            Cube
          </option>
          <option>
            Octohedron
          </option>
          <option>
            Dodecahedron
          </option>
          <option selected={shapes[shapeIndex] === Shape.icosahedron}
                  value={Shape.icosahedron}>
            Icosahedron
          </option>
        </optgroup>
      </select>
      <button class="col-span-1"
              on:click={handleAllClick}>
        all
      </button>
    </div>
    <div class="grow pt-2 pl-5 pr-5 grid grid-cols-4 grid-rows-4 gap-0">
      {#each {length: 16} as _, i}
        <button class="flex justify-around items-center"
                on:click={(e) => handleShapeIndexSelect(e, i)}
                class:selected={shapeIndex === i}>
          {i + 1}
        </button>
      {/each}
    </div>
  </div>
  <div class="grow pt-5 flex flex-col justify-between items-stretch">
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
