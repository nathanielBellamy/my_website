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
    'geometry',
    'mouseTracking',
    'radius',
    'rotation',
    'lfos'
  ]

  let curr_mod_left: string = 'color'
  let curr_mod_right: string = 'rotation'  
</script>

<div id="magic_square_control_rack"
     class="magic_square_control_rack flex flex-row-reverse justify-between">
  <div class="mod_select">
    <ControlModule title="MODS">
      <Select modules={modules}
              bind:curr_mod_left={curr_mod_left}
              bind:curr_mod_right={curr_mod_right}/>
    </ControlModule>
  </div>
  <div class="left_right_slots grid grid-cols-2 gap-2">
    <div class="left_slot">
      {#if curr_mod_left == 'color'}
        <ControlModule title="COLOR"
                       side="left">
         <Color />
        </ControlModule>
      {:else if curr_mod_left == 'drawPattern'}
        <ControlModule title="PATTERN"
                       side="left">
          <DrawPattern />
        </ControlModule>
      {:else if curr_mod_left == 'mouseTracking'}
        <ControlModule title="MOUSE"
                       side="left">
          <MouseTracking />
        </ControlModule>
      {:else if curr_mod_left == 'radius'}
        <ControlModule title="MOUSE"
                       side="left">
          <Radius />
        </ControlModule>
      {:else if curr_mod_left == 'rotation'}
        <ControlModule title="ROTATION"
                       side="left">
          <Rotation />
        </ControlModule>
      {:else}
        <ControlModule />
      {/if}
    </div>
    <div class="right_slot">
      {#if curr_mod_right == 'color'}
        <ControlModule title="COLOR"
                       side="right">
         <Color />
        </ControlModule>
      {:else if curr_mod_right == 'drawPattern'}
        <ControlModule title="PATTERN"
                       side="right">
          <DrawPattern />
        </ControlModule>
      {:else if curr_mod_right == 'mouseTracking'}
        <ControlModule title="MOUSE"
                       side="right">
          <MouseTracking />
        </ControlModule>
      {:else if curr_mod_right == 'radius'}
        <ControlModule title="RADIUS"
                       side="right">
          <Radius />
        </ControlModule>
      {:else if curr_mod_right == 'rotation'}
        <ControlModule title="ROTATION"
                       side="right">
          <Rotation />
        </ControlModule>
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
    padding: 5px 40px 5px 40px
    height: 100%
    border-radius: 5px
    background: color.$black-blue-grad
    min-height: 500px

  .mod_select
    height: calc(100% - 10px)
    overflow: hidden

  .left_right_slots
    height: calc(100% - 10px)
    grid-template-areas: "left_slot right_slot"
    grid-template-columns: 45% 45%
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
