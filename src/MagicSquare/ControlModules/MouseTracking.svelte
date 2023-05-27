<script lang="ts">
  import { onMount } from 'svelte'

  const storageKey = 'magic_square_storage'
  
  const mouseTrackingOptions: string[] = [
    'On',
    'Off',
    'Inv X',
    'Inv Y',
    'Inv XY'
  ]

  let formId = 'mouse_tracking_form'
  let hiddenInputId = 'magic_square_input_mouse_tracking'
  let curr_option: string = ''

  function getCurrOptionLocal() {
    return curr_option
  }

  function getCurrOptionStorage () {
    const storageData = JSON.parse(localStorage.getItem(storageKey))
    return storageData.settings.mouse_tracking
  }

  function setCurrOpt(opt: string) {
    curr_option = opt
  }

  onMount(async () => {
    curr_option = getCurrOptionStorage()
    // wasm listens to input events on the forms
    // within the manual call to dispatchEvent we must
    // explicitly set bubbles:true so that wasm can catch the event
    // while listening to the form
    // this way a single wasm closure can handle all ui data updates
    var form = document.getElementById(formId)
    form.addEventListener('submit', () => {
      var input = document.getElementById(hiddenInputId)
      input.value = getCurrOptionLocal()
      input.dispatchEvent(new Event('input', {bubbles: true}))
    })
  })

  function handleOptKeydown(e: any, opt: string) {
    if (e.keyCode === 13){
      setCurrOpt(opt)
      let form = document.getElementById(formId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  function handleOptClick(opt: string) {
    setCurrOpt(opt)
    let form = document.getElementById(formId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  } 
</script>

<div id={formId}
     class="mouse_tracking_container flex flex-col justify-around">
  {#each mouseTrackingOptions as opt}
    <button class="mouse_tracking_option"
            class:selected="{curr_option === opt}"
            on:click={() => handleOptClick(opt)}
            on:keydown={(e) => handleOptKeydown(e, opt)}>
        {opt.toUpperCase()}
    </button>
  {/each}
  <input id={hiddenInputId}
         class="hidden_input"/>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .mouse_tracking
    &_container
      height: 100%
    &_option
      flex-grow: 1
  .hidden_input
    display: none

  .selected
    background-color: color.$blue-8
</style>
