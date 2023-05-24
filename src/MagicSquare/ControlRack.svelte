<script lang="ts">
  import { onMount } from 'svelte'
  import Color from './ControlModules/Color.svelte'
  import DrawPattern from './ControlModules/DrawPattern.svelte'
  import Radius from './ControlModules/Radius.svelte'
  import Rotation from './ControlModules/Rotation.svelte'
  import MouseTracking from './ControlModules/MouseTracking.svelte'
  import Select from './ControlModules/Select.svelte'
  // { [selectId]: hiddehInputId}
  const selects: { [key: string]: string; }= {
    'draw_pattern_select': 'magic_square_input_draw_pattern',
    'mouse_tracking_select': 'magic_square_input_mouse_tracking'
  }

  const modules: string[] = [
    'color',
    'drawPattern',
    'mouseTracking',
    'radius',
    'rotation',
    'none'
  ]

  let curr_mod_left: string = 'color'
  let curr_mod_right: string = 'rotation'

  const getModule = (modName: string) => {
    switch(modName) {
      case 'color':
        return Color
      case 'drawPattern':
        return DrawPattern
      case 'mouseTracking':
        return MouseTracking
      case 'radius':
        return Radius
      case 'rotation':
        return Rotation
      case 'none':
        return null
    }
  }

  // rust sets values on hidden inputs
  // this method reads those values into the iro elmeents
  onMount(() => {
    // wasm listens to input events on the forms
    // within the manual call to dispatchEvent we must
    // explicitly set bubbles:true so that wasm can catch the event
    // while listening to the form
    // this way a single wasm closure can handle all ui data updates


    document.getElementById("magic_square_control_rack").addEventListener("click", (_e:any) => {
      console.dir({curr_mod_left, curr_mod_right})
    })
  })

  const handleSelectChange = (e: Event) => {
    var input = document.getElementById("magic_square_input_draw_pattern")
    input.value = e.target.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<div id="magic_square_control_rack"
     class="magic_square_control_rack flex">
  <Select modules={modules}
          bind:curr_mod_left={curr_mod_left}
          bind:curr_mod_right={curr_mod_right}/>
  
  {#if curr_mod_left == 'color'}
    <Color />
  {:else if curr_mod_left == 'drawPattern'}
    <DrawPattern />
  {:else if curr_mod_left == 'mouseTracking'}
    <MouseTracking />
  {:else if curr_mod_left == 'radius'}
    <Radius />
  {:else if curr_mod_left == 'rotation'}
    <Rotation />
  {/if}

  {#if curr_mod_right == 'color'}
    <Color />
  {:else if curr_mod_right == 'drawPattern'}
    <DrawPattern />
  {:else if curr_mod_right == 'mouseTracking'}
    <MouseTracking />
  {:else if curr_mod_right == 'radius'}
    <Radius />
  {:else if curr_mod_right == 'rotation'}
    <Rotation />
  {/if}
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"
  
  .magic_square_control_rack
      height: 100%
      width: 100%
      overflow: hidden
      padding: 3px 20px 3px 20px
</style>
