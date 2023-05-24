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

  var rotationSliders: RangeInput[] = [
    {
      id: "magic_square_input_x_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Roll Spread",
      min: -0.3,
      max: 0.3,
      step: 0.01
    },
    {
      id: "magic_square_input_y_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Pitch Spread",
      min: -0.3,
      max: 0.3,
      step: 0.1
    },
    {
      id: "magic_square_input_z_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Yaw Spread",
      min: -0.3,
      max: 0.3,
      step: 0.1
    },
    {
      id: "magic_square_input_x_axis_x_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "X Axis - Roll Coeff",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_x_axis_y_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "X Axis - Pitch Coeff",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_x_axis_Z_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "X Axis - Yaw Coeff",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_y_axis_x_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Y Axis - Roll Coeff",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_y_axis_y_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Y Axis - Pitch Coeff",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_y_axis_Z_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Y Axis - Yaw Coeff",
      min: -1,
      max: 1,
      step: 0.01
    }
  ]

  function handleRotationChange(e: any, id: string, idx: number) {
    rotationSliders[idx].value = e.detail.value
    rotationSliders = rotationSliders
    var input = document.getElementById(id)
    input.value = e.detail.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<ControlModule title="ROTATION">
  <div class="flex flex-col space-between">
    <div id="radius_sliders" 
         class="flex flex-col space-between">
      {#each rotationSliders as  {id, label, min, max, initialValue, value}, idx}
        <label for={id}>
          {label}
        </label>
        <Range id={`${id}_range`}
               min={min}
               max={max}
               initialValue={initialValue}
               value={value}
               on:change={(e) => handleRotationChange(e, id, idx)}/>
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
