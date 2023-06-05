<script lang="ts">
  import { Module } from './Module'
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'

  let langVal: Lang 
  lang.subscribe(val => langVal = val)
  let i18n = new I18n("magicSquare/select")

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
</script>

<div class="module_selector flex flex-col">
  <div class="module_selector_side_set flex">
    <button class="side_set side_set_left"
            class:side_set_left_selected="{sideToSet === Side.left}"
            on:dblclick={() => swap()}
            on:click={() => sideToSet = Side.left}>
      {i18n.t("left", langVal)}
    </button>
    <button class="side_set side_set_right"
            class:side_set_right_selected="{sideToSet === Side.right}"
            on:dblclick={() => swap()}
            on:click={() => sideToSet = Side.right}>
      {i18n.t("right", langVal)}
    </button>
  </div>
  {#each Object.values(Module) as mod}
    <button class="module_option"
            class:selected_left="{curr_mod_left === mod}"
            class:selected_right="{curr_mod_right === mod}"
            on:click={() => setMod(mod)}
            on:keydown={(e) => handleModKeydown(e, mod)}>
        {i18n.t(mod, langVal)}
      <input id={`mod_radio_${mod}`}
             value={mod}
             type="radio"
             name="wow"
             checked={curr_mod_left === mod}
             class="hidden_input"/>
    </button>
  {/each}
</div>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .side_set
    flex-grow: 1
    margin: 5px
    padding: 5px
    border-radius: 5px
    font-size: text.$fs-m
    font-weight: text.$fw-l
    color: color.$cream
    &_left
      &_selected
        background-color: color.$green-4
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


  .selected_left
    background-color: color.$green-4
  .selected_right
    background-color: color.$purple-7

  .module_selector
    justify-content: space-between
    border-radius: 5px
    height: 100%

  .hidden_input
    display: none
</style>
