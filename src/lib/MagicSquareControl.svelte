<script lang="ts">
  import { onMount } from 'svelte'
  import iro from '@jaames/iro'
  import Range from './Range.svelte'

  const colorPickerOptions = {
    width: 150,
    height: 150,
    borderWidth: 5,
    borderColor: "#EFF0E9", // white in styles/color.sass
    layoutDirection: 'horizontal'
  }

  const colorPickerIds: string[] = [
    'magic_square_input_color_1',
    'magic_square_input_color_2',
    'magic_square_input_color_3',
    'magic_square_input_color_4',
    'magic_square_input_color_5',
    'magic_square_input_color_6',
    'magic_square_input_color_7',
    'magic_square_input_color_8',
  ]
  
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

  var rotationSliders: RangeInput[] = [
    {
      id: "magic_square_input_x_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Roll Spread",
      min: -3.14,
      max: 3.14,
      step: 0.1
    },
    {
      id: "magic_square_input_y_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Pitch Spread",
      min: -3.14,
      max: 3.14,
      step: 0.1
    },
    {
      id: "magic_square_input_z_rot_spread",
      initialValue: 0.0,
      value: 0.0,
      label: "Yaw Spread",
      min: -3.14,
      max: 3.14,
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

  // { [selectId]: hiddehInputId}
  const selects: { [key: string]: string; }= {
    'draw_pattern_select': 'magic_square_input_draw_pattern',
    'mouse_tracking_select': 'magic_square_input_mouse_tracking'
  }

  const drawPatterns: string[] = [
    'All',
    'One',
    'Two',
    'Three',
    'Four',
    'Five',
    'Six',
    'Seven',
    'Eight',
    'Out1',
    'Out2',
    'Out3',
    'Out4',
    'Out5',
    'Out6',
    'Out7',
    'Out8',
    'In1',
    'In2',
    'In3',
    'In4',
    'In5',
    'In6',
    'In7',
    'In8',
    'Conv',
    'Div',
    'Random'
  ]

  const mouseTrackingOptions: string[] = [
    'On',
    'Off',
    'Inv X',
    'Inv Y',
    'Inv XY'
  ]
  
  // rust sets values on hidden inputs
  // this method reads those values into the iro elmeents

  onMount(() => {
    colorPickerIds.forEach((id:string) => {
      var input = document.getElementById(id)
      const picker = iro.ColorPicker(`#${id}_picker`, colorPickerOptions)

      picker.on('color:change', (color: any) => {
        input.value = `${color.rgba.r},${color.rgba.g},${color.rgba.b},${color.rgba.a}`
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })
    })

    // wasm listens to input events on the forms
    // within the manual call to dispatchEvent we must
    // explicitly set bubbles:true so that wasm can catch the event
    // while listening to the form
    // this way a single wasm closure can handle all ui data updates

    for (const [selectId, hiddenInputId] of Object.entries(selects)) {
      var select = document.getElementById(selectId)
      select.addEventListener('change', (e: Event) => {
        var input = document.getElementById(hiddenInputId)
        input.value = e.target.value
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })
    }
  })

  const handleSelectChange = (e: Event) => {
    var input = document.getElementById("magic_square_input_draw_pattern")
    input.value = e.target.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  } 

  function handleRadiusChange(e: any, id: string, idx: number) {
    radiusSliders[idx].value = e.detail.value
    radiusSliders = radiusSliders
    var input = document.getElementById(id)
    input.value = e.detail.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  function handleRotationChange(e: any, id: string, idx: number) {
    rotationSliders[idx].value = e.detail.value
    rotationSliders = rotationSliders
    var input = document.getElementById(id)
    input.value = e.detail.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<div id="magic_square_control"
     class="magic_square_control flex">
  {#each colorPickerIds as id}
    <div class="magic_square_input flex flex-col space-between">
      <label for={id}
             class="title">
        {id.split("_").slice(-1)[0].toUpperCase()}
      </label>
      <div id={`${id}_picker`} 
           class="color_picker"/>
      <input id={id}
             class="hidden_input">
    </div>
  {/each}
  <div class="magic_square_input flex flex-col space-between">
    <label for="radius_sliders"
           class="title">
      RADIUS
    </label>
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
  <div class="magic_square_input flex flex-col space-between">
    <label for="radius_sliders"
           class="title">
      ROTATION
    </label>
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
  <div class="magic_square_input flex flex-col space-between">
    <label for="draw_pattern_select"
           class="title">
      DRAW PATTERN
    </label>
    <select id="draw_pattern_select" 
            value="Three">
      {#each drawPatterns as pattern}
        <option value={pattern}>
          {pattern.toUpperCase()}
        </option>
      {/each}
      <input id="magic_square_input_draw_pattern"
             class="hidden_input">
    </select>
  </div>
  <div class="magic_square_input flex flex-col space-between">
    <label for="draw_pattern_select"
           class="title">
      MOUSE TRACKING
    </label>
    <select id="mouse_tracking_select" 
            value="Off">
      {#each mouseTrackingOptions as mto}
        <option value={mto}>
          {mto.toUpperCase()}
        </option>
      {/each}
      <input id="magic_square_input_mouse_tracking"
             class="hidden_input">
    </select>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"
  
  .magic_square
    &_control
      width: 100%
      overflow-x: scroll
      padding: 3px 20px 3px 20px

    &_input
      max-width: 300px
      align-items: stretch
      border: 5px solid color.$blue-4
      border-width: 3px
      border-radius: 5px
      margin: 5px
      padding: 5px
      overlfow-y: scroll
      
  .title
    font-size: text.$fs_l
    font-weight: text.$fw_l
    text-align: left
  
  .color_picker
    display: flex
    justify-content: space-around

  .hidden_input
    display: none
</style>
