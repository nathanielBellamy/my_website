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
      label: "Spread Roll",
      min: -0.3,
      max: 0.3,
      step: 0.01
    },
    {
      id: "magic_square_input_y_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Spread Pitch",
      min: -0.3,
      max: 0.3,
      step: 0.1
    },
    {
      id: "magic_square_input_z_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Spread Yaw",
      min: -0.3,
      max: 0.3,
      step: 0.1
    },
    {
      id: "magic_square_input_x_axis_x_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "X Roll",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_x_axis_y_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "X Pitch",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_x_axis_Z_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "X Yaw",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_y_axis_x_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Y Roll",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_y_axis_y_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Y Pitch",
      min: -1,
      max: 1,
      step: 0.01
    },
    {
      id: "magic_square_input_y_axis_Z_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Y Yaw",
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
  {#each rotationSliders as  {id, label, min, max, initialValue, value}, idx}
    <div class="rotation_input flex flex-col">
      <label class="rotation_input_label" 
             for={id}>
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
    </div>
  {/each}
</ControlModule>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .hidden_input
    display: none

  .rotation_input
    width: 100%
    &_label
      font-weight: text.$fw-l
      font-size: text.$fs-m
      width: 100%
      text-align: left
      padding-left: 10px
</style>
