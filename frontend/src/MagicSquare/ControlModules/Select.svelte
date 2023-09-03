<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Module } from './Module'
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'
  import { smallScreen } from '../../stores/smallScreen'
  import Icon from '../../lib/Icon.svelte'
  import { Icons } from '../../lib/Icons'
  import { ToastColor } from '../../lib/Toasty';
  import DrawPattern from './DrawPattern.svelte';
    import { MessagesSolid } from 'flowbite-svelte-icons';

  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)
  let i18n = new I18n("magicSquare/select")

  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean) => smallScreenVal = val)

  export let pub: boolean = false
  export let curr_mod_left: Module = Module.color
  export let curr_mod_right: Module = Module.rotation

  enum Side {
    left = 'left',
    right = 'right'
  }

  let sideToSet:Side = Side.left

  function handleModKeydown(e: any, mod: Module) {
    if (e.keyCode === 13){
      setMod(mod)
    }
  }

  function setMod(mod: Module) {
    // this condition prevents a visual bug that occurs
    // when two instances of the Color module are loaded simultaneously
    // the issue results from element confusion in document.getElementById
    if (smallScreenVal && curr_mod_right === Module.color && mod === Module.color) {
      // destroy color instance in curr_mod_right
      curr_mod_right = Module.feed
      curr_mod_left = mod 
      return
    }

    if (smallScreenVal) { // only one module is displayed at a time
      curr_mod_left = mod 
      return
    }
    if (mod === curr_mod_left || mod === curr_mod_right) return
    if (sideToSet === Side.left) {
      curr_mod_left = mod
    } else {
      curr_mod_right = mod
    }
  }

  function swap() {
    const old_left = JSON.parse(JSON.stringify(curr_mod_left))
    curr_mod_left = curr_mod_right
    curr_mod_right = old_left
  }

  let modules = Object.values(Module).filter(x => {
    // feed for public
    // presets for private
    var res: boolean
    if (pub) {
      res = x !== "presets"
    } else {
      res = x !== "feed"
    }
    return res
  })

  onDestroy(() => {
    unsubLang()
    unsubSmallScreen()
  })
</script>

<div class="select_container rounded-md h-10 w-11/12 pl-2 pr-2 overflow-x-scroll"
     class:module_selector_grid={!smallScreenVal}
     class:module_selector_flex={smallScreenVal}
     class:text-xs={smallScreenVal}>
  {#if !smallScreenVal}
    <div class="left_right_buttons pr-2 h-full w-fit flex justify-between items-center">
      <button class="side_set h-5/6 text-cyan-500 flex justify-around items-center pl-2 pr-2"
              class:side_set_left_selected={sideToSet === Side.left}
              on:dblclick={() => swap()}
              on:click={() => sideToSet = Side.left}>
        <Icon icon={Icons.ChevronLeftSolid} />
      </button>
      <button class="side_set h-5/6 text-cyan-500 flex justify-around items-center pl-2 pr-2"
              class:side_set_right_selected={sideToSet === Side.right}
              on:dblclick={() => swap()}
              on:click={() => sideToSet = Side.right}>
        <Icon icon={Icons.ChevronRightSolid} />
      </button>
    </div>
  {/if}
  <div class="h-full w-full pl-2 pr-2 flex justify-between items-center overflow-x-scroll">
    {#each modules as mod}
      <button class="module_option text-cyan-700 h-5/6 w-fit pr-2 pl-2"
              title={i18n.t(mod, langVal)}
              class:selected_left={curr_mod_left === mod}
              class:selected_right={curr_mod_right === mod && !smallScreenVal}
              on:click={() => setMod(mod)}
              on:keydown={(e) => handleModKeydown(e, mod)}>
        {#if mod == Module.drawPattern}
          <span class="text-cyan-500">
            <Icon icon={Icons.AtomSolid}/>
          </span>
        {:else if mod === Module.feed}
          <span class="text-cyan-500">
            <Icon icon={Icons.MessagesOutline} />
          </span>
        {:else if mod == Module.translation}
          <span class="text-cyan-500">
            <Icon icon={Icons.ArrowUpDownSolid}/>
          </span>
        {:else if mod == Module.presets}
          <span class="text-cyan-500">
            <Icon icon={Icons.BriefcaseSolid} />
          </span>
        {:else if mod == Module.rotation}
          <span class="text-cyan-500">
            <Icon icon={Icons.RotateOutline} />
          </span>
        {:else}
          <span class="font-extrabold text-xl">
            {i18n.t(mod + "_emoji", langVal)} 
          </span>
        {/if}
      </button>
      <input id={`mod_radio_${mod}`}
             value={mod}
             type="radio"
             checked={curr_mod_left === mod}
             class="hidden_input"/>
    {/each}
  </div>
</div>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .left_right_buttons
    border-right: 5px double color.$blue-7

  .select_container
    border: 5px double color.$blue-7

  .side_set
    box-shadow: none

    &_left
      &_selected
        background-color: color.$green-7
    &_right
      &_selected
        background-color: color.$purple-7

  .module_option
    color: color.$cream
    display: flex
    justify-content: space-around
    align-items: center
    cursor: pointer
    font-size: text.$fs-s
    font-weight: text.$fw-xl
    flex-grow: 1
    cursor: pointer
    box-shadow: none
    min-width: 30px

  .selected_left
    background-color: color.$green-7
  .selected_right
    background-color: color.$purple-7

  .module_selector
    &_grid
      display: grid
      grid-template-columns: .1fr 1fr
      grid-template-rows: 100%
      gap: 10px
    &_flex
      display: flex
      justify-content: space-between
      align-items: center
      padding: 0 10px 0 10px

  .hidden_input
    display: none
</style>
