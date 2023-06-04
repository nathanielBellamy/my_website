<script lang="ts">
  import { onMount } from 'svelte'
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'

  let langVal: Lang 
  lang.subscribe(val => langVal = val)
  let i18n = new I18n

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

  enum Toggle {
    on = 'on',
    off = 'off'
  }

  enum Inv {
    x = 'x',
    y = 'y',
    xy = 'xy',
    none = 'none'
  }

  let toggle: Toggle
  let inv: Inv
  let option: MouseTrackingOption

  $: option = deriveMouseTrackingOption(toggle, inv)

  function parseVars(opt: MouseTrackingOption) {
    switch(opt) {
      case MouseTrackingOption.on:
        toggle = Toggle.on
        inv = Inv.none
        break
      case MouseTrackingOption.off:
        toggle = Toggle.off
        inv = Inv.none
        break
      case MouseTrackingOption.invX:
        toggle = Toggle.on
        inv = Inv.x
        break
      case MouseTrackingOption.invY:
        toggle = Toggle.on
        inv = Inv.y
        break
      case MouseTrackingOption.invXY:
        toggle = Toggle.on
        inv = Inv.xy
        break
    }
  }

  function deriveMouseTrackingOption(t: Toggle, i: Inv) {
    switch(t) {
      case Toggle.off :
        return MouseTrackingOption.off
      case Toggle.on:
        switch(i) {
          case Inv.none:
            return MouseTrackingOption.on
          case Inv.x:
            return MouseTrackingOption.invX
          case Inv.y:
            return MouseTrackingOption.invY
          case Inv.xy:
            return MouseTrackingOption.invXY
        }
      default:
        return MouseTrackingOption.off
    }
  }

  function handleFormSubmit(e: any){
    e.preventDefault()
    var input = document.getElementById(hiddenInputId)
    console.log(option)
    input.value = option
    input.dispatchEvent(new Event('input', {bubbles: true}))
    return false // do not refresh page on submit
  }

  onMount(async () => {
    // wasm listens to input events on the forms
    // within the manual call to dispatchEvent we must
    // explicitly set bubbles:true so that wasm can catch the event
    // while listening to the form
    // this way a single wasm closure can handle all ui data updates
    parseVars(currOption)
    var form = document.getElementById(formId)
    form.addEventListener('submit', handleFormSubmit)
  })

  function handleToggleKeydown(e: any, newToggle: Toggle) {
    if (e.keyCode === 13){
      toggle = newToggle
      let form = document.getElementById(formId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  function handleToggleClick(newToggle: Toggle) {
    toggle = newToggle
    let form = document.getElementById(formId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  } 

  function handleInvKeydown(e: any, newInv: Inv) {
    if (e.keyCode === 13){
      inv = newInv
      let form = document.getElementById(formId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  function handleInvClick(newInv: Inv) {
    inv = newInv
    let form = document.getElementById(formId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  }

  const invGroup1: Inv[] = [Inv.x, Inv.y, Inv.xy]
  const invGroup2: Inv[] = [Inv.none]
</script>

<form id={formId}
      class="h-full flex flex-col justify-around items-stretch">
  <h2 class="mouse_tracking_title">
    {i18n.t("magicSquare/mouseTracking/mouse", langVal)}
  </h2>
  <div id="mouse_tracking_toggle"
       class="grow flex justify-around items-stretch">
    {#each Object.keys(Toggle) as t}
      <button class="grow"
              class:selected="{toggle === t}"
              on:click={() => handleToggleClick(Toggle[t])}
              on:keydown={(e) => handleToggleKeydown(e, Toggle[t])}>
          {i18n.t(`misc/${t}`, langVal)}
      </button>
    {/each}
  </div>
  <div id="mouse_tracking_inv"
       class="mouse_tracking_inv grow w-full flex justify-between items-stretch">
    <div class="mouse_tracking_inv_title">
      {i18n.t("magicSquare/mouseTracking/invert", langVal)}
    </div>
    <div class="grow flex flex-col justify-between items stretch">
      <div class="grow flex justify-evenly items-center">
        {#each invGroup1 as i}
          <button class="grow"
                  class:selected="{inv === i}"
                  disabled={toggle === Toggle.off}
                  on:click={() => handleInvClick(Inv[i])}
                  on:keydown={(e) => handleInvKeydown(e, Inv[i])}>
              {i.toUpperCase()}
          </button>
        {/each}
      </div>
      <div class="grow flex justify-stretch items-center">
        {#each invGroup2 as i}
          <button class="grow"
                  class:selected="{inv === i}"
                  disabled={toggle === Toggle.off}
                  on:click={() => handleInvClick(Inv[i])}
                  on:keydown={(e) => handleInvKeydown(e, Inv[i])}>
              {i18n.t(`magicSquare/mouseTracking/${i}`, langVal)}
          </button>
        {/each}
      </div>
    </div>
  </div>
  <slot name="hiddenInput" />
</form>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .mouse_tracking
    &_title
      color: color.$blue-4
      font-weight: text.$fw-l
      font-size: text.$fs-l

    &_inv
      border-top: 5px double color.$blue-7
      border-bottom: 5px double color.$blue-7
      &_title
        color: color.$blue-4
        font-weight: text.$fw-l
        transform: rotate(-90deg)
        height: 100%
        min-width: 25px
        max-width: 25px
        margin-top: 70px

  .selected
    background-color: color.$blue-8
</style>
