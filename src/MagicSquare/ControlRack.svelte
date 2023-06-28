<script lang="ts">
  import { onMount } from 'svelte'
  import { into_module, Module } from './ControlModules/Module'
  import ControlModule from './ControlModule.svelte'
  import Select from './ControlModules/Select.svelte'

  // INIT LANG BOILER PLATE
  import { I18n, Lang } from '../I18n'
  import { lang } from '../stores/lang'

  const i18n = new I18n('magicSquare/controlRack')
  let langVal: Lang
  lang.subscribe(val => langVal = val)

  $: translationTitle = i18n.t(Module.translation, langVal)

  export let small: boolean = false

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
      localStorage.setItem('magic_square_curr_mods', JSON.stringify(res))
    }
    return res
  }

  onMount(() => {
    const curr_mods: any = JSON.parse(localStorage.getItem('magic_square_curr_mods'))
    if (curr_mods){
      curr_mod_left = into_module(curr_mods.curr_mod_left)
      curr_mod_right = into_module(curr_mods.curr_mod_right)
    } else {
      curr_mod_left = Module.presets
      curr_mod_right = Module.color
    }
  })
</script>




<div id="magic_square_control_rack"
     class="magic_square_control_rack"
     class:grid_col={small}
     class:grid_row={!small}
    >
  <div class="hidden">
    {storage_mods}
  </div>
  <div class:slot_grid_1={small}
       class:slot_grid_2={!small}>
    <div class="left_slot">
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
      {:else if curr_mod_left === Module.presets}
        <ControlModule title={i18n.t(Module.presets, langVal)}
                       side={Side.left}>
          <slot name="presets" />
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
    {#if !small}
      <div class="right_slot">
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
        {:else if curr_mod_right === Module.presets}
          <ControlModule title={i18n.t(Module.presets, langVal)}
                         side={Side.right}>
            <slot name="presets" />
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
  <div class="mod_select">
    <ControlModule title={i18n.t("modules", langVal)}>
      <Select bind:curr_mod_left={curr_mod_left}
              bind:curr_mod_right={curr_mod_right}
              small={small}/>
    </ControlModule>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .grid_col
    display: grid
    grid-template-columns: 1fr
    grid-template-rows: 1fr 1fr
    gap: 5px

  .grid_row
    display: grid
    grid-template-columns: 1fr 0.5fr
    grid-template-rows: 100%
    gap: 5px


  .slot_grid_1
    height: 100%
    display: grid
    grid-template-columns: 1fr
    grid-template-rows: 1fr
    gap: 5px

  .slot_grid_2
    height: 100%
    display: grid
    grid-template-columns: 1fr 1fr
    grid-template-rows: 1fr
    gap: 5px

  .magic_square_control_rack
    flex-grow: 1
    padding: 5px 40px 5px 40px
    border-radius: 5px
    background: color.$black-blue-grad
    min-height: 100%
    
  .left_slot
    min-width: 226px
    overflow: hidden
  
  .right_slot
    min-width: 226px
    overflow: hidden

  .hidden
    display: none
</style>
