<script lang="ts">
  import { onMount } from 'svelte'
  import Color from './ControlModules/Color.svelte'
  import ControlModule from './ControlModule.svelte'
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
</script>

<div id="magic_square_control_rack"
     class="magic_square_control_rack flex flex-row-reverse justify-between flex-wrap-reverse">
  <Select modules={modules}
          bind:curr_mod_left={curr_mod_left}
          bind:curr_mod_right={curr_mod_right}/>
  <div class="left_right_slots grid grid-cols-2 gap-2">
    <div class="left_slot">
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
      {:else}
        <ControlModule />
      {/if}
    </div>
    <div class="right_slot">
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
      {:else}
        <ControlModule />
      {/if}
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"
  
  .magic_square_control_rack
    flex-grow: 1
    overflow-y: scroll
    padding: 5px 40px 5px 40px
    height: 100%
    border: 10px double color.$blue-4
    border-radius: 5px
    background: color.$black-blue-grad

  .left_right_slots
    grid-template-areas: "left_slot right_slot"
    grid-template-columns: 45% 45%
    height: 100%
    min-width: 500px
    flex-grow: 1

  .left_slot
    min-width: 200px
    grid-area: left_slot
    overflow: hidden
  
  .right_slot
    min-width: 200px
    grid-area: right_slot
    overflow: hidden
  

</style>
