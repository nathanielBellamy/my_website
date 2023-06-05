<script lang="ts">
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'

  let langVal: Lang 
  lang.subscribe(val => langVal = val)
  let i18n = new I18n("magicSquare/mouseTracking")

  enum Lfo {
    one = "1",
    two = "2",
    three = "3",
    four = "4"
  }

  let lfo: Lfo = Lfo.one
</script>

<div>
  <div id="magic_square_lfo_select">
    {#each Object.keys(Lfo) as lfoKey}
      <button on:click = {() => lfo = Lfo[lfoKey]}
              class:selected = {lfo === Lfo[lfoKey]}>
        {Lfo[lfoKey]}
      </button>
    {/each}
  </div>
  <slot name="lfoSliders">
    LFO: {lfo}
  </slot>
</div>

<style lang="sass">
  @use "./../../styles/color"

  .selected
    background-color: color.$blue-8
</style>
