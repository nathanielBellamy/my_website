<script lang="ts">
  import { TransformOrder } from "./TransformOrder"
  import { WasmInputId } from "../WasmInputId"

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
    <div class="grow flex flex-col justify-between items-stretch m-5">
      <select class="shape_select p-5">
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
    </div>
  </div>
  <div class="grow flex flex-col justify-between items-stretch">
    <div class="title flex items-stretch pl-5 underline">
      transform order
    </div>
    <button class="grow flex justify-around items-center"
            class:selected={transformOrder === TransformOrder.rotateThenTranslate}
            on:click={() => handleTransformOrderClick(TransformOrder.rotateThenTranslate)}>
      Rotate -> Translate
    </button>
    <button class="grow flex justify-around items-center"
            class:selected={transformOrder === TransformOrder.translateThenRotate}
            on:click={() => handleTransformOrderClick(TransformOrder.translateThenRotate)}>
      Translate -> Rotate 
    </button>
    <slot name="transformOrder"/>
  </div>
  <div class="grow">
    <!-- TODO   -->
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
