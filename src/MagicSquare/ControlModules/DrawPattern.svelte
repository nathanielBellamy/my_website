<script lang="ts">
  import { WasmInputId } from '../WasmInputId';
  import { DrawPatternType } from './DrawPattern'
  
  export let drawPatternType: DrawPatternType
  export let drawPatternCount: number

  function handleDrawPatternTypeChange(e: any) {
    var input = document.getElementById(WasmInputId.drawPatternType)
    drawPatternType = e.target.value
    input.value = e.target.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<section class="h-full pl-5 pr-5 flex flex-col justify-between items-stretch">
  <div class="text-left flex flex-col justify-betwen items-stretch">
    <div class="title flex items-stretch underline">
      type
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
  </div>
  <div class="grow flex flex-col justify-between items-stretch">
    <slot name="countAndSpeed" />
  </div>
  <slot name="hiddenInput" />
</section>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"
  .title
    color: color.$blue-7
    font-size: text.$fs-ml
    font-weight: text.$fw-l

  .selected
    background-color: color.$blue-8
</style>
