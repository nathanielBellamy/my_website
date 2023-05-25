<script lang="ts">
  import { onMount } from 'svelte'
  import ControlModule from "../ControlModule.svelte"
  
  const mouseTrackingOptions: string[] = [
    'On',
    'Off',
    'Inv X',
    'Inv Y',
    'Inv XY'
  ]

  let selectId = 'mouse_tracking_select'
  let hiddenInputId = 'magic_square_input_mouse_tracking'

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

<ControlModule title="MOUSE">
  <select id={selectId}
          value="Off">
    {#each mouseTrackingOptions as mto}
      <option value={mto}>
        {mto.toUpperCase()}
      </option>
    {/each}
    <input id={hiddenInputId}
           class="hidden_input">
  </select>
</ControlModule>

<style lang="sass">
  .hidden_input
    display: none
</style>
