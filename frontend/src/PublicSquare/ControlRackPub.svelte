<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { into_module, Module } from '../MagicSquare/ControlModules/Module'
  import ControlModule from '../MagicSquare/ControlModule.svelte'
  import Select from '../MagicSquare/ControlModules/Select.svelte'

  // INIT LANG BOILER PLATE
  import { I18n, Lang } from '../I18n'
  import { lang } from '../stores/lang'

  const i18n = new I18n('magicSquare/controlRack')
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)

  import { smallScreen } from '../stores/smallScreen'

  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean) => smallScreenVal = val)

  $: translationTitle = i18n.t(Module.translation, langVal)

  enum Side {
    left = 'left',
    right = 'right'
  }

  let curr_mod_left: Module
  let curr_mod_right: Module

  $: storage_mods = set_curr_mods(curr_mod_left, curr_mod_right)

  function set_curr_mods(left: Module, right: Module): {[key: string]: Module} {
    const res: {[key: string]: Module} = {curr_mod_left: left, curr_mod_right: right}
    if (!!left && !!right){
      localStorage.setItem('public_square_curr_mods', JSON.stringify(res))
    }
    return res
  }

  onMount(() => {
    const curr_mods: any = JSON.parse(localStorage.getItem('public_square_curr_mods'))
    if (curr_mods){
      curr_mod_left = into_module(curr_mods.curr_mod_left)
      curr_mod_right = into_module(curr_mods.curr_mod_right)
    } else {
      curr_mod_left = Module.color
      curr_mod_right = Module.feed
    }
  })

  onDestroy(() => {
    unsubLang()
    unsubSmallScreen()
  })
</script>

<div id="magic_square_control_rack"
     class="magic_square_control_rack grid_col h-full">
  <div class="hidden">{storage_mods}</div>
  <div class="h-full w-full"
       class:slot_grid={!smallScreenVal}
       class:slot_flex={smallScreenVal}>
    <div class="left_slot h-full">
      {#if curr_mod_left === Module.color}
        <ControlModule title={i18n.t(Module.color, langVal)}
                       side={Side.left}>
          <slot name="color"/>
        </ControlModule>
      {:else if curr_mod_left === Module.drawPattern}
        <ControlModule title={i18n.t(Module.drawPattern, langVal)}
                       side={Side.left}>
          <slot name="drawPattern"/>
        </ControlModule>
      {:else if curr_mod_left === Module.lfo}
        <ControlModule title={i18n.t(Module.lfo, langVal)}
                       side={Side.left}>
          <slot name="lfo"/>
        </ControlModule>
      {:else if curr_mod_left === Module.geometry}
        <ControlModule title={i18n.t(Module.geometry, langVal)}
                       side={Side.left}>
          <slot name="geometry" />
        </ControlModule>
      {:else if curr_mod_left === Module.feed}
        <ControlModule title={i18n.t(Module.feed, langVal)}
                       side={Side.left}>
          <slot name="feed" />
        </ControlModule>
      {:else if curr_mod_left === Module.rotation}
        <ControlModule title={i18n.t(Module.rotation, langVal)}
                       side={Side.left}>
          <slot name="rotation" />
        </ControlModule>
      {:else if curr_mod_left === Module.translation}
        <ControlModule bind:title={translationTitle}
                       side={Side.left}>
          <slot name="translation"/>
        </ControlModule>
      {:else}
        <ControlModule side="left"/>
      {/if}
    </div>
    {#if !smallScreenVal}
      <div class="right_slot h-full">
        {#if curr_mod_right === Module.color}
          <ControlModule  title={i18n.t(Module.color, langVal)}
                          side={Side.right}>
            <slot name="color"/>
          </ControlModule>
        {:else if curr_mod_right === Module.drawPattern}
          <ControlModule title={i18n.t(Module.drawPattern, langVal)}
                         side={Side.right}>
            <slot name="drawPattern" />
          </ControlModule>
        {:else if curr_mod_right === Module.lfo}
          <ControlModule title={i18n.t(Module.lfo, langVal)}
                         side={Side.right}>
            <slot name="lfo"/>
          </ControlModule>
        {:else if curr_mod_right === Module.geometry}
          <ControlModule title={i18n.t(Module.geometry, langVal)}
                         side={Side.right}>
            <slot name="geometry" />
          </ControlModule>
        {:else if curr_mod_right === Module.feed}
          <ControlModule title={i18n.t(Module.feed, langVal)}
                         side={Side.right}>
            <slot name="feed" />
          </ControlModule>
        {:else if curr_mod_right === Module.rotation}
          <ControlModule title={i18n.t(Module.rotation, langVal)}
                         side={Side.right}>
            <slot name="rotation" />
          </ControlModule>
        {:else if curr_mod_right == Module.translation}
          <ControlModule title={i18n.t(Module.translation, langVal)}
                         side={Side.right}>
            <slot name="translation"/>
          </ControlModule>
        {:else}
          <ControlModule side="right"/>
        {/if}
      </div>
    {/if}
  </div>
  <div class="h-full w-full flex justify-around items-center">
    <Select bind:curr_mod_left={curr_mod_left}
            bind:curr_mod_right={curr_mod_right}
            pub={true}/>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .grid_col
    display: grid
    grid-template-columns: 1fr
    grid-template-rows: 85% 15%
    gap: 5px
 
  .slot
    &_flex
      display: flex
      justify-content: space-around
      align-items: center
    &_grid
      display: grid
      grid-template-columns: 1fr 1fr
      grid-template-rows: 1fr
      gap: 5px

  .magic_square_control_rack
    padding: 5px 40px 5px 40px
    border-radius: 5px
    background: color.$black-blue-grad
    
  .left_slot
    /* min-width: 226px */
    overflow: hidden
  
  .right_slot
    /* min-width: 226px */
    overflow: hidden

  .hidden
    display: none
</style>
