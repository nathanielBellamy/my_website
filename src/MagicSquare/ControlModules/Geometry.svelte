<script lang="ts">
  import { TransformOrder } from "./TransformOrder"
  import { WasmInputId } from "../WasmInputId"
  import { Shape } from "./Shape"

  export let transformOrder: TransformOrder

  function handleTransformOrderClick(tr_or: TransformOrder) {
    console.log(tr_or)
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
    <div class="pl-5 pr-5 grid grid-cols-4">
      <select class="col-span-3 shape_select">
        <optgroup label="2d">
          <option>
            Triangle
          </option>
          <option>
            Square
          </option>
          <option>
            Pentagon
          </option>
          <option>
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
          <option>
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
        <button>
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
