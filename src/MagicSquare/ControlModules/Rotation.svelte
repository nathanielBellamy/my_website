<script lang="ts">
  import ControlModule from "../ControlModule.svelte"

  enum Freedom {
    Pitch,
    Roll,
    Yaw
  }

  //TODO: generalize RangeInput
  interface RangeInput {
    id: string,
    label: string,
    min: number,
    max: number,
    step: number,
    initialValue: number,
    value: number,
    freedom: Freedom
  }
 
  var rotationSliders: RangeInput[] = [
    // pitch
    {
      id: "magic_square_input_y_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Spread",
      min: -0.3,
      max: 0.3,
      step: 0.1,
      freedom: Freedom.Pitch
    },
    {
      id: "magic_square_input_x_axis_y_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Mouse X",
      min: -1,
      max: 1,
      step: 0.01,
      freedom: Freedom.Pitch
    },
    {
      id: "magic_square_input_y_axis_y_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Mouse Y",
      min: -1,
      max: 1,
      step: 0.01,
      freedom: Freedom.Pitch
    },
    // roll
    {
      id: "magic_square_input_x_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Spread",
      min: -0.3,
      max: 0.3,
      step: 0.01,
      freedom: Freedom.Roll
    },
    {
      id: "magic_square_input_x_axis_x_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Mouse X",
      min: -1,
      max: 1,
      step: 0.01,
      freedom: Freedom.Roll
    },
    {
      id: "magic_square_input_y_axis_x_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Mouse Y",
      min: -1,
      max: 1,
      step: 0.01,
      freedom: Freedom.Roll
    },
    // yaw
    {
      id: "magic_square_input_z_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Spread",
      min: -0.3,
      max: 0.3,
      step: 0.1,
      freedom: Freedom.Yaw
    },
    {
      id: "magic_square_input_x_axis_z_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Mouse X",
      min: -1,
      max: 1,
      step: 0.01,
      freedom: Freedom.Yaw
    },
    {
      id: "magic_square_input_y_axis_z_rot_coeff",
      initialValue: 0,
      value: 0,
      label: "Mouse Y",
      min: -1,
      max: 1,
      step: 0.01,
      freedom: Freedom.Yaw
    }
  ]

  $: pitch_sliders = rotationSliders.filter(x => x.freedom === Freedom.Pitch)
  $: roll_sliders = rotationSliders.filter(x => x.freedom === Freedom.Roll)
  $: yaw_sliders = rotationSliders.filter(x => x.freedom === Freedom.Yaw)

  function freedomToString(freedom: Freedom) {
    switch(freedom) {
      case Freedom.Pitch:
        return 'Pitch'
      case Freedom.Roll:
        return 'Roll'
      case Freedom.Yaw:
        return 'Yaw'
    }
  }

  function freedomToSliders(freedom: Freedom) {
    switch(freedom) {
      case Freedom.Pitch:
        return pitch_sliders
      case Freedom.Roll:
        return roll_sliders
      case Freedom.Yaw:
        return yaw_sliders
    }
  }

  const freedoms = [Freedom.Pitch, Freedom.Roll, Freedom.Yaw]

  function handleRotationChange(e: any, id: string) {
    var slider = rotationSliders.find(x => x.id === id)
    slider.value = e.detail.value
    rotationSliders = rotationSliders
    var input = document.getElementById(id)
    input.value = e.detail.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  // zero value on double click
  // TODO: debug updating slider UI
  // this might be solved when moving settings into local storage
  function handleDoubleClick(e: any, id: string) {
    rotationSliders.find(x => x.id === id).value = 0.0
    rotationSliders = JSON.parse(JSON.stringify(rotationSliders))
    var input = document.getElementById(id)
    input.value = 0.0
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<ControlModule title="ROTATION">
  {#each freedoms as freedom}
    <div class="freedom_group flex gap-2">
      <div class="freedom_group_title">
        {freedomToString(freedom)}
      </div>
      <div class="freedom_group_body">
        {#each freedomToSliders(freedom) as  {id, label, min, max, value}}
          <div class="rotation_input flex flex-col"
               on:dblclick={(e) => handleDoubleClick(e,id)}>
            <label class="rotation_input_label flex justify-between" 
                   for={id}>
              <div> {label} </div>
              <div> {value} </div>
            </label>
            <input id={id}
                   type="range"
                   min={min}
                   max={max}
                   bind:value={value}
                   step={.01}/>
          </div>
        {/each}
      </div>
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
      padding-left: 10px
      padding-right: 10px

  .freedom_group
    align-items: center
    padding: 5px
    margin: 5px

    &_title
      font-weight: text.$fw-m
      font-size: text.$fs-l
      transform: rotate(-90deg)
      max-width: 15px
      color: color.$blue-7

    &_body
      border-top: 5px double color.$blue-7
      border-radius: 10px
      flex-grow: 1
    
</style>
