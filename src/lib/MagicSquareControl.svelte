<script lang="ts">
  import { onMount } from 'svelte'
  import iro from '@jaames/iro'

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
    var select = document.getElementById("draw_pattern_select")
    select.addEventListener('change', (e: Event) => {
      var input = document.getElementById("magic_square_input_draw_pattern")
      input.value = e.target.value
      input.dispatchEvent(new Event('input', {bubbles: true}))
    })
  })

  const handleSelectChange = (e: Event) => {
    var input = document.getElementById("magic_square_input_draw_pattern")
    input.value = e.target.value
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
      
  .title
    font-size: text.$fs_l
    font-weight: text.$fw_l
  
  .color_picker
    display: flex
    justify-content: space-around

  .hidden_input
    display: none
</style>
