<script lang="ts">
  import { onDestroy } from 'svelte'
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'
  import { Lfo } from './Lfo'
  import { LfoDestination } from './LfoDestination'
  import { LfoShape } from './LfoShape'
  import { WasmInputId } from '../WasmInputId';

  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)
  onDestroy(unsubLang)

  let i18n = new I18n("magicSquare/lfo")

  export let lfo1Active: boolean = false
  export let lfo1Dest: LfoDestination = LfoDestination.none
  export let lfo1Shape: LfoShape = LfoShape.sine

  export let lfo2Active: boolean = false
  export let lfo2Dest: LfoDestination = LfoDestination.none
  export let lfo2Shape: LfoShape = LfoShape.sine

  export let lfo3Active: boolean = false
  export let lfo3Dest: LfoDestination = LfoDestination.none
  export let lfo3Shape: LfoShape = LfoShape.sine

  export let lfo4Active: boolean = false
  export let lfo4Dest: LfoDestination = LfoDestination.none
  export let lfo4Shape: LfoShape = LfoShape.sine

  let lfo: Lfo = Lfo.one

  // destination
  $: currDest = intoCurrDestination(lfo)

  function intoCurrDestination(lfo: Lfo) {
    switch (lfo) {
      case Lfo.two:
        return lfo2Dest
      case Lfo.three:
        return lfo3Dest
      case Lfo.four:
        return lfo4Dest
      case Lfo.one:
      default:
        return lfo1Dest
    }
  }

  function handleLfoDestChange(e: any, lfo: Lfo) {
    const input_id = intoDestInputId(lfo)
    var input = document.getElementById(input_id)
    if (!!input) {
      input.value = e.target.value
      input.dispatchEvent(new Event('input', {bubbles: true}))
    }
  }

  function intoDestInputId(lfo: Lfo): string {
    switch (lfo) {
      case Lfo.two:
        return WasmInputId.lfo2Dest
      case Lfo.three:
        return WasmInputId.lfo3Dest
      case Lfo.four:
        return WasmInputId.lfo4Dest
      case Lfo.one:
      default:
        return WasmInputId.lfo1Dest
    }
  }

  // shape
  $: currShape = intoCurrShape(lfo)

  function intoCurrShape(lfo: Lfo) {
    switch (lfo) {
      case Lfo.two:
        return lfo2Shape
      case Lfo.three:
        return lfo3Shape       
      case Lfo.four:
        return lfo4Shape
      case Lfo.one:
      default:
        return lfo1Shape
    }
  }

  function handleLfoShapeChange(e: any, lfo: Lfo) {
    const input_id = intoShapeInputId(lfo)
    var input = document.getElementById(input_id)
    if (!!input) {
      input.value = e.target.value
      input.dispatchEvent(new Event('input', {bubbles: true}))
    }
  }

  function intoShapeInputId(lfo: Lfo): string {
    switch (lfo) {
      case Lfo.two:
        return WasmInputId.lfo2Shape
      case Lfo.three:
        return WasmInputId.lfo3Shape
      case Lfo.four:
        return WasmInputId.lfo4Shape
      case Lfo.one:
      default:
        return WasmInputId.lfo1Shape
    }
  }

  function handleLfoDoubleClick(lfo: Lfo) {
    // TODO
    // activate and deactivate lfo from double clicking the button to select it
  }

  function lowerCaseFirstLetter(ld: LfoDestination): string {
    const oldStr = ld.toString()
    const newFirst = oldStr[0].toLowerCase()
    return `${newFirst}${oldStr.substring(1)}`
  }
</script>

<div class="h-full pb-5 flex flex-col justify-between items-stretch">
  <div id="magic_square_lfo_select"
       class="lfo_select pb-2 flex flex-col justify-between items-stretch">
    <button on:click = {() => lfo = Lfo.one}
            class:active = {lfo1Active}
            class:selected = {lfo === Lfo.one}
            class="grid grid-cols-5 ml-5 mr-5 text-sm border-4 rounded-xl">
      <p> {Lfo.one} </p>
      <p class="col-span-4"> 
        {i18n.t(lowerCaseFirstLetter(lfo1Dest), langVal)} 
      </p>
    </button>
    <button on:click = {() => lfo = Lfo.two}
            class:active = {lfo2Active}
            class:selected = {lfo === Lfo.two}
            class="grid grid-cols-5 ml-5 mr-5 text-sm border-4 rounded-xl">
      <p> {Lfo.two} </p>
      <p class="col-span-4"> 
        {i18n.t(lowerCaseFirstLetter(lfo2Dest), langVal)} 
      </p>
    </button>
    <button on:click = {() => lfo = Lfo.three}
            class:active = {lfo3Active}
            class:selected = {lfo === Lfo.three}
            class="grid grid-cols-5 ml-5 mr-5 text-sm border-4 rounded-xl">
      <p> {Lfo.three} </p>
      <p class="col-span-4"> 
        {i18n.t(lowerCaseFirstLetter(lfo3Dest), langVal)} 
      </p>
    </button>
    <button on:click = {() => lfo = Lfo.four}
            class:active = {lfo4Active}
            class:selected = {lfo === Lfo.four}
            class="grid grid-cols-5 ml-5 mr-5 text-sm border-4 rounded-xl">
      <p> {Lfo.four} </p>
      <p class="col-span-4"> 
        {i18n.t(lowerCaseFirstLetter(lfo4Dest), langVal)}
      </p>
    </button>
  </div>
  <div class="grow pt-1 flex flex-col justify-between items-stretch">
    <div class="grow flex flex-col justify-around items-center">
      <div class="w-full pl-5 pr-5 flex flex-col justify-around items-stretch">
        <label for="lfo_dest_select"
               class="destination_label w-full flex justify-between text-left">
          <div>
            {i18n.t("destination", langVal)}
          </div>
        </label>
        <select id="lfo_dest_select"
                value={currDest}
                on:input={e => e.stopPropagation()}
                on:change={(e) => {
                  handleLfoDestChange(e, lfo)
                }}>
          <optgroup label={i18n.t("rotation", langVal)}>
            <option value={LfoDestination.pitchBase}> 
              {i18n.t("pitchBase", langVal)} 
            </option>
            <option value={LfoDestination.pitchSpread}> 
              {i18n.t("pitchSpread", langVal)} 
            </option>
            <option value={LfoDestination.pitchX}> 
              {i18n.t("pitchX", langVal)} 
            </option>
            <option value={LfoDestination.pitchY}> 
              {i18n.t("pitchY", langVal)}
            </option>
            <option value={LfoDestination.rollBase}> 
              {i18n.t("rollBase", langVal)}
            </option>
            <option value={LfoDestination.rollSpread}> 
              {i18n.t("rollSpread", langVal)} 
            </option>
            <option value={LfoDestination.rollX}> 
              {i18n.t("rollX", langVal)} 
            </option>
            <option value={LfoDestination.rollY}> 
              {i18n.t("rollY", langVal)} 
            </option>
            <option value={LfoDestination.yawBase}> 
              {i18n.t("yawBase", langVal)} 
            </option>
            <option value={LfoDestination.yawSpread}> 
              {i18n.t("yawSpread", langVal)} 
            </option>
            <option value={LfoDestination.yawX}>
              {i18n.t("yawX", langVal)}
            </option>
            <option value={LfoDestination.yawY}> 
              {i18n.t("yawY", langVal)} 
            </option>
          </optgroup>
          <optgroup label={i18n.t("radius", langVal)}>
            <option value={LfoDestination.radiusBase}> 
              {i18n.t("radiusBase", langVal)} 
            </option>
            <option value={LfoDestination.radiusStep}> 
              {i18n.t("radiusStep", langVal)} 
            </option>
          </optgroup>
          <optgroup label={i18n.t("translation", langVal)}>
            <option value={LfoDestination.translationXBase}> 
              {i18n.t("translationXBase", langVal)}
            </option>
            <option value={LfoDestination.translationXSpread}> 
              {i18n.t("translationXSpread", langVal)}
            </option>
            <option value={LfoDestination.translationYBase}> 
              {i18n.t("translationYBase", langVal)}
            </option>
            <option value={LfoDestination.translationYSpread}> 
              {i18n.t("translationYSpread", langVal)}
            </option>
          </optgroup>
        </select>
      </div>
      {#if lfo === Lfo.one}
        <slot name="lfo1"/>
      {:else if lfo === Lfo.two}
        <slot name="lfo2"/>
      {:else if lfo === Lfo.three}
        <slot name="lfo3"/>
      {:else if lfo === Lfo.four}
        <slot name="lfo4"/>
      {/if}
      <div class="w-full pl-5 pr-5 flex flex-col justify-around items-stretch">
        <label for="lfo_shape_select"
               class="destination_label w-full flex justify-between text-left">
          <div>
            {i18n.t("shape", langVal)}
          </div>
        </label>
        <select id="lfo_shape_select"
                value={currShape}
                on:input={e => e.stopPropagation()}
                on:change={e => handleLfoShapeChange(e, lfo)}>
          <option value={LfoShape.linear}> 
            Linear 
          </option>
          <option value={LfoShape.sine}> 
            Sine 
          </option>
        </select>
      </div>
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"
  
  .destination_label
    width: 80%
    font-weight: text.$fw-l

  .selected
    background-color: color.$blue-7
    box-shadow: 0 0

  .active
    border-color: color.$red-7

  .lfo_select
    flex-grow: 0.2
    border-bottom: 5px double color.$blue-7
</style>
