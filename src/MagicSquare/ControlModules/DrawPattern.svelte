<script lang="ts">
  import { WasmInputId } from '../WasmInputId'
  import { DrawPatternType } from './DrawPattern'
  import { TransformOrder } from './TransformOrder'
  
  export let drawPatternType: DrawPatternType
  export let transformOrder: TransformOrder

  function handleDrawPatternTypeChange(e: any) {
    var input = document.getElementById(WasmInputId.drawPatternType)
    drawPatternType = e.target.value
    input.value = e.target.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  function handleTransformOrderClick(tr_or: TransformOrder) {
    transformOrder = tr_or
    var input = document.getElementById(WasmInputId.transformOrder)
    input.value = tr_or
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<section class="h-full pl-5 pr-5 flex flex-col justify-between items-stretch">
  <div class="transform_order flex flex-col justify-between items-stretch">
    <slot name="transformOrder"/>
    <div class="grow p-5 flex flex-col justify-around items-stretch gap-8">
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
    </div>
  </div>
  <div class="grow text-left flex flex-col justify-betwen items-stretch">
    <div class="title flex items-stretch underline">
      direction
    </div>
    <select class="w-full"
            value={drawPatternType}
            on:change={handleDrawPatternTypeChange}
            on:input={(e) => e.stopPropagation()}>
      <option value={DrawPatternType.out}>
        Out
      </option>
      <option value={DrawPatternType.in}>
        In
      </option>
      <option value={DrawPatternType.fix}>
        Fix
      </option>
    </select>
    <div class="grow flex flex-col justify-between items-stretch">
      <slot name="countAndSpeed" />
    </div>
  </div>
  <slot name="hiddenInput" />
</section>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"
  .title
    color: color.$blue-7
    font-size: text.$fs-m
    font-weight: text.$fw-l

  .selected
    background-color: color.$blue-8

  .transform_order
    flex-grow: 0.25
</style>
