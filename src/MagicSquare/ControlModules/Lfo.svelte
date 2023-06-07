<script lang="ts">
    import { prevent_default } from 'svelte/internal';
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'
  import { LfoDestination } from './LfoDestination'
  import { LfoShape } from './LfoShape'

  let langVal: Lang 
  lang.subscribe(val => langVal = val)
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

  enum Lfo {
    one = "1",
    two = "2",
    three = "3",
    four = "4"
  }

  let lfo: Lfo = Lfo.one

  function handleLfoDestChange(e: any, lfo: Lfo) {
    var input = document.getElementById('magic_square_input_lfo_1_dest')
    input.value = e.target.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  function handleLfoShapeChange(e: any, lfo: Lfo) {
    var input = document.getElementById('magic_square_input_lfo_1_shape')
    input.value = e.target.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<div class="h-full pt-5 pb-5 flex flex-col justify-between items-stretch">
  <div id="magic_square_lfo_select">
    <button on:click = {() => lfo = Lfo.one}
            class:active = {lfo1Active}
            class:selected = {lfo === Lfo.one}
            class="pt-2 pb-2 pr-3 pl-3">
      {Lfo.one}
    </button>
    <button on:click = {() => lfo = Lfo.two}
            class:active = {lfo2Active}
            class:selected = {lfo === Lfo.two}
            class="pt-2 pb-2 pr-3 pl-3">
      {Lfo.two}
    </button>
    <button on:click = {() => lfo = Lfo.three}
            class:active = {lfo3Active}
            class:selected = {lfo === Lfo.three}
            class="pt-2 pb-2 pr-3 pl-3">
      {Lfo.three}
    </button>
    <button on:click = {() => lfo = Lfo.four}
            class:active = {lfo4Active}
            class:selected = {lfo === Lfo.four}
            class="pt-2 pb-2 pr-3 pl-3">
      {Lfo.four}
    </button>
  </div>
  <div class="grow pt-5 pb-5 flex flex-col justify-between items-stretch">
    <div class="grow flex flex-col justify-around items-center"
         class:hidden = {lfo !== Lfo.one}>
      <div class="w-full pl-5 pr-5 flex flex-col justify-around items-stretch">
        <label for="lfo_1_dest_select"
               class="destination_label w-full flex justify-between text-left">
          <div>
            {i18n.t("destination", langVal)}
          </div>
        </label>
        <select id="lfo_1_dest_select"
                value={lfo1Dest}
                on:input={e => e.stopPropagation()}
                on:change={(e) => {
                  handleLfoDestChange(e, Lfo.one)
                }}>
          <optgroup label={i18n.t("rotation", langVal)}>
            <option> {i18n.t("pitchBase", langVal)} </option>
            <option> {i18n.t("pitchSpread", langVal)} </option>
            <option> {i18n.t("pitchX", langVal)} </option>
            <option> {i18n.t("pitchY", langVal)} </option>
            <option> {i18n.t("rollBase", langVal)} </option>
            <option> {i18n.t("rollSpread", langVal)} </option>
            <option> {i18n.t("rollX", langVal)} </option>
            <option> {i18n.t("rollY", langVal)} </option>
            <option> {i18n.t("yawBase", langVal)} </option>
            <option> {i18n.t("yawSpread", langVal)} </option>
            <option> {i18n.t("yawX", langVal)} </option>
            <option> {i18n.t("yawY", langVal)} </option>
          </optgroup>
          <optgroup label={i18n.t("radius", langVal)}>
            <option> {i18n.t("minimum", langVal)} </option>
            <option> {i18n.t("step", langVal)} </option>
          </optgroup>
          <optgroup label={i18n.t("translation", langVal)}>
            <option value={LfoDestination.translationX}> 
              X 
            </option>
            <option value={LfoDestination.translationY}> 
              Y 
            </option>
          </optgroup>
        </select>
      </div>
      <slot name="lfo1"/>
      <div class="w-full pl-5 pr-5 flex flex-col justify-around items-stretch">
        <label for="lfo_1_dest_select"
               class="destination_label w-full flex justify-between text-left">
          <div>
            {i18n.t("shape", langVal)}
          </div>
        </label>
        <select id="lfo_1_dest_select"
                on:input={e => e.stopPropagation()}
                on:change={e => handleLfoShapeChange(e, Lfo.one)}>
          {#each Object.keys(LfoShape) as shapeKey}
            <option>
              {LfoShape[shapeKey]}
            </option>
          {/each}
        </select>
      </div>
    </div>
    <div class="grow flex justify-around items-center"
         class:hidden = {lfo !== Lfo.two}>
      <slot name="lfo2"/>
    </div>
    <div class="grow flex justify-around items-center"
         class:hidden = {lfo !== Lfo.three}>
      <slot name="lfo3"/>
    </div>
    <div class="grow flex justify-around items-center"
         class:hidden = {lfo !== Lfo.four}>
      <slot name="lfo4"/>
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"
  
  .destination_label
    width: 80%
    font-weight: text.$fw-l

  .hidden
    display: none

  .selected
    background-color: color.$blue-8

  .active
    border-color: color.$red-5
</style>
