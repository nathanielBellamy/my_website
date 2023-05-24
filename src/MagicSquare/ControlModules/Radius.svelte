<script lang="ts">
  import ControlModule from "../ControlModule.svelte"
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

<ControlModule title="RADIUS">
  <div class="flex flex-col space-between">
    <div id="radius_sliders" 
         class="flex flex-col space-between">
      {#each radiusSliders as  {id, label, min, max, initialValue, value}, idx}
        <label for={id}>
          {label}
        </label>
        <Range id={`${id}_range`}
               min={min}
               max={max}
               initialValue={initialValue}
               value={value}
               on:change={(e) => handleRadiusChange(e, id, idx)}/>
        <input id={id}
               class="hidden_input"/>
      {/each}
    </div>
  </div>
</ControlModule>

<style lang="sass">
  .hidden_input
    display: none
</style>
