<script lang="ts">
  import Range from "../../lib/Range.svelte"

  //TODO: generalize RangeInput
  interface RangeInput {
    id: string,
    label: string,
    min: number,
    max: number,
    step: number,
    initialValue: number,
    value: number
  }

  var radiusSliders: RangeInput[] = [
    {
      id: "magic_square_input_radius_min",
      label: "MIN",
      min: 0.1,
      max: 1,
      step: 0.01,
      initialValue: 0.5,
      value: 0.5
    },
    {
      id: "magic_square_input_radius_step",
      label: "STEP",
      min: 0.01,
      max: 0.5,
      step: 0.01,
      initialValue:  0.5,
      value: 0.5
    }
  ]

  function handleRadiusChange(e: any, id: string, idx: number) {
    radiusSliders[idx].value = e.detail.value
    radiusSliders = radiusSliders
    var input = document.getElementById(id)
    input.value = e.detail.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<div class="flex flex-col space-between">
  <div id="radius_sliders" 
       class="flex flex-col space-between items-center">
    {#each radiusSliders as  {id, label, min, max, initialValue, value}, idx}
      <label class="radius_input_label flex justify-between" 
             for={id}>
        <div> {label} </div>
        <div> {value} </div>
      </label>
      <input id={id}
             class="radius_input"
             type="range"
             min={min}
             max={max}
             bind:value={value}
             step={.01}/>
    {/each}
  </div>
</div>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .radius_input
    width: 90%
    &_label
      width: 100%
      font-weight: text.$fw-l
      font-size: text.$fs-m
      padding-left: 10px
      padding-right: 10px

  .hidden_input
    display: none
</style>
