<script lang="ts">
  import { onMount } from 'svelte'
  import Color from './ControlModules/Color.svelte'
  import Radius from './ControlModules/Radius.svelte'
  import Rotation from './ControlModules/Rotation.svelte'
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
</script>

<div id="magic_square_control"
     class="magic_square_control flex">
  <Color />
  <Radius />
  <Rotation />

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
      height: 100%
      width: 100%
      overflow: hidden
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
