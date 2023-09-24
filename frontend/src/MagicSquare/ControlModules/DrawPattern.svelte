<script lang="ts">
  import { onDestroy } from 'svelte'
  import { WasmInputId } from '../WasmInputId'
  import { DrawPatternType } from './DrawPattern'
  import { TransformOrder } from './TransformOrder'
  
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'
  let i18n = new I18n("magicSquare/drawPattern")
  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)

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

  onDestroy(unsubLang)
</script>

<section class="h-full pb-5 pr-5 flex flex-col justify-between items-stretch gap-2">
  <div class="title">
    {i18n.t("order", langVal)}
  </div>
  <div class="transform_order flex flex-col justify-between items-stretch">
    <slot name="transformOrder"/>
    <div class="grow pl-5 pr-5 flex flex-col justify-around items-stretch gap-3">
      <button class="grow text-left flex justify-around items-center"
              class:selected={transformOrder === TransformOrder.rotateThenTranslate}
              on:click={() => handleTransformOrderClick(TransformOrder.rotateThenTranslate)}>
        <ol>
          <li>
            1. {i18n.t("rotate", langVal)}
          </li>
          <li>
            2. {i18n.t("translate", langVal)}
          </li>
        </ol>
      </button>
      <button class="grow text-left flex justify-around items-center"
              class:selected={transformOrder === TransformOrder.translateThenRotate}
              on:click={() => handleTransformOrderClick(TransformOrder.translateThenRotate)}>
        <ol type="1">
          <li>
            1. {i18n.t("translate", langVal)}
          </li>
          <li>
            2. {i18n.t("rotate", langVal)}
          </li>
        </ol>
      </button>
    </div>
  </div>
  <div class="title">
    {i18n.t("animation", langVal)}
  </div>
  <div class="grow pl-5 text-left flex flex-col justify-betwen items-stretch">
    <div class="slider_label flex items-stretch">
      {i18n.t("direction", langVal)}
    </div>
    <select class="w-full"
            value={drawPatternType}
            on:change={handleDrawPatternTypeChange}
            on:input={(e) => e.stopPropagation()}>
      <option value={DrawPatternType.out}>
        {i18n.t("out", langVal)}
      </option>
      <option value={DrawPatternType.in}>
        {i18n.t("in", langVal)}
      </option>
      <option value={DrawPatternType.fix}>
        {i18n.t("fix", langVal)}
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
  @import "../styles/control_module_title.sass"

  .title
    @include control_module_title

  .selected
    background-color: color.$blue-8

  .transform_order
    flex-grow: 0.25

  .slider_label
      width: 100%
      font-weight: text.$fw-l
      font-size: text.$fs-m
      padding-right: 5%
</style>
