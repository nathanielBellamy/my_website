<script lang="ts">
  import { onMount } from 'svelte'
  
  enum MouseTrackingOption {
    on = 'On',
    off = 'Off',
    invX = 'Inv X',
    invY = 'Inv Y',
    invXY = 'Inv XY'
  }

  export let currOption: MouseTrackingOption
  const hiddenInputId = 'magic_square_input_mouse_tracking'
  const formId = 'mouse_tracking_form'
  
  function setCurrOpt(opt: MouseTrackingOption) {
    currOption = opt
  }
  function handleFormSubmit(e: any){
    e.preventDefault()
    var input = document.getElementById(hiddenInputId)
    input.value = currOption
    input.dispatchEvent(new Event('input', {bubbles: true}))
    return false // do not refresh page on submit
  }

  onMount(async () => {
    // wasm listens to input events on the forms
    // within the manual call to dispatchEvent we must
    // explicitly set bubbles:true so that wasm can catch the event
    // while listening to the form
    // this way a single wasm closure can handle all ui data updates
    var form = document.getElementById(formId)
    form.addEventListener('submit', handleFormSubmit)
  })

  function handleOptKeydown(e: any, opt: MouseTrackingOption) {
    if (e.keyCode === 13){
      setCurrOpt(opt)
      let form = document.getElementById(formId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  function handleOptClick(opt: MouseTrackingOption) {
    setCurrOpt(opt)
    let form = document.getElementById(formId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  } 
</script>

<form id={formId}
     class="mouse_tracking_container flex flex-col justify-around">
  {#each Object.values(MouseTrackingOption) as opt}
    <button class="mouse_tracking_option"
            class:selected="{currOption === opt}"
            on:click={() => handleOptClick(opt)}
            on:keydown={(e) => handleOptKeydown(e, opt)}>
        {opt.toUpperCase()}
    </button>
  {/each}
  <slot name="hiddenInput" />
</form>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .mouse_tracking
    &_container
      height: 100%
    &_option
      flex-grow: 1
  .selected
    background-color: color.$blue-8
</style>
