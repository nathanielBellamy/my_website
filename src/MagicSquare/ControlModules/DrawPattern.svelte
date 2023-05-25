<script lang="ts">
  import { onMount } from "svelte"
  import ControlModule from "../ControlModule.svelte"

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

  let selectId = 'draw_pattern_select'
  let hiddenInputId = 'magic_square_input_draw_pattern'

  onMount(async () => {
    // wasm listens to input events on the forms
    // within the manual call to dispatchEvent we must
    // explicitly set bubbles:true so that wasm can catch the event
    // while listening to the form
    // this way a single wasm closure can handle all ui data updates
    var select = document.getElementById(selectId)
    select.addEventListener('change', (e: Event) => {
      var input = document.getElementById(hiddenInputId)
      input.value = e.target.value
      input.dispatchEvent(new Event('input', {bubbles: true}))
    })
  })

</script>

<ControlModule title="PATTERN">
  <select id={selectId} 
          value="Three">
    {#each drawPatterns as pattern}
      <option value={pattern}>
        {pattern.toUpperCase()}
      </option>
    {/each}
    <input id="magic_square_input_draw_pattern"
           class="hidden_input">
  </select>
</ControlModule>

<style lang="sass">
  .hidden_input
    display: none
</style>
