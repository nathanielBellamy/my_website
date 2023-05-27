<script lang="ts">
  import { onMount } from "svelte"
  import ControlModule from "../ControlModule.svelte"

  const storageKey = 'magic_square_storage'

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

  let formId = 'draw_pattern_form'
  let hiddenInputId = 'magic_square_input_draw_pattern'
  let curr_pattern: string = ''

  function getCurrPatternLocal() {
    return curr_pattern
  }

  function getCurrPatternStorage () {
    const storageData = JSON.parse(localStorage.getItem(storageKey))
    return storageData.settings.draw_pattern
  }

  function setCurrPattern(pattern: string) {
    curr_pattern = pattern
  }

  function handlePatternKeydown(e: any, pattern: string) {
    if (e.keyCode === 13){
      setCurrPattern(pattern)
      let form = document.getElementById(formId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  function handlePatternClick(pattern: string) {
    setCurrPattern(pattern)
    let form = document.getElementById(formId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  } 

  onMount(async () => {
    curr_pattern = getCurrPatternStorage()
    // wasm listens to input events on the forms
    // within the manual call to dispatchEvent we must
    // explicitly set bubbles:true so that wasm can catch the event
    // while listening to the form
    // this way a single wasm closure can handle all ui data updates
    var form = document.getElementById(formId)
    form.addEventListener('submit', () => {
      var input = document.getElementById(hiddenInputId)
      input.value = getCurrPatternLocal()
      input.dispatchEvent(new Event('input', {bubbles: true}))
    })
  })

</script>

<ControlModule title="PATTERN">
  <div id={formId}
       class="draw_pattern_container justify-around">
    <div class="draw_pattern_options flex flex-col">
      {#each drawPatterns as pattern}
        <button class="draw_pattern_option"
                class:selected="{curr_pattern === pattern}"
                on:click={() => handlePatternClick(pattern)}
                on:keydown={(e) => handlePatternKeydown(e, pattern)}>
            {pattern.toUpperCase()}
        </button>
      {/each}
    </div>
    <input id={hiddenInputId}
           class="hidden_input"/>
  </div>
</ControlModule>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .draw_pattern
    &_container
      height: 100%
      width: 100%
      padding: 10px 0 10px 0
    &_options
      width: 100%
      height: 100%
      overflow-y: scroll
    &_option
      flex-grow: 1

  .selected
    background-color: color.$blue-8
  .hidden_input
    display: none
</style>
