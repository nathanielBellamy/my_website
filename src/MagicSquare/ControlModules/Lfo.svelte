<script lang="ts">
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'

  let langVal: Lang 
  lang.subscribe(val => langVal = val)
  let i18n = new I18n("magicSquare/lfo")

  enum Lfo {
    one = "1",
    two = "2",
    three = "3",
    four = "4"
  }

  enum Shape {
    sawtooth = "sawtooth",
    sine = "sine",
    square = "square"
  }

  let lfo: Lfo = Lfo.one
</script>

<div class="h-full pt-5 pb-5 flex flex-col justify-between items-stretch">
  <div id="magic_square_lfo_select">
    {#each Object.keys(Lfo) as lfoKey}
      <button on:click = {() => lfo = Lfo[lfoKey]}
              class:selected = {lfo === Lfo[lfoKey]}
              class="pt-2 pb-2 pr-3 pl-3">
        {Lfo[lfoKey]}
      </button>
    {/each}
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
        <select id="lfo_1_dest_select">
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
            <option> X </option>
            <option> Y </option>
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
        <select id="lfo_1_dest_select">
          {#each Object.keys(Shape) as shapeKey}
            <option>
              {Shape[shapeKey]}
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
</style>
