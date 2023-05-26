<script lang="ts">
  import ControlModule from "../ControlModule.svelte"

  export let modules: string[] = []
  export let curr_mod_left: string = 'color'
  export let curr_mod_right: string = 'rotation'

  let sideToSet:string = 'left'

  function handleModClick(mod: any) {
    {
      if (mod === curr_mod_left || mod === curr_mod_right) return
      if (sideToSet == 'left') {
        curr_mod_left = mod
      } else {
        curr_mod_right = mod
      }
    }
  }
</script>

<ControlModule title="MODS">
  <div class="module_selector flex flex-col">
    <div class="module_selector_side_set flex">
      <button class="side_set side_set_left"
              class:side_set_left_selected="{sideToSet === 'left'}"
              on:click={() => sideToSet = 'left'}>
        LEFT
      </button>
      <button class="side_set side_set_right"
              class:side_set_right_selected="{sideToSet === 'right'}"
              on:click={() => sideToSet = 'right'}>
        RIGHT
      </button>
    </div>
    {#each modules as mod}
      <button class="module_option"
              class:selected_left="{curr_mod_left === mod}"
              class:selected_right="{curr_mod_right === mod}"
              on:click={() => handleModClick(mod)}
              on:keydown={() => handleModClick(mod)}>
          {mod.toUpperCase()}
        <input id={`mod_radio_${mod}`}
               value={mod}
               type="radio"
               name="wow"
               checked={curr_mod_left === mod}
               class="hidden_input"/>
      </button>
    {/each}
  </div>
</ControlModule>

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
      border: 5px solid color.$green-4
      &_selected
        background-color: color.$green-4
    &_right
      border: 5px solid color.$red-4
      &_selected
        background-color: color.$red-4

  .module_option
    color: color.$cream
    display: flex
    justify-content: space-around
    align-items: center
    cursor: pointer
    font-size: text.$fs-s
    font-weight: text.$fw-l
    flex-grow: 1
    cursor: pointer


  .selected_left
    background-color: color.$green-4
  .selected_right
    background-color: color.$red-4

  .module_selector
    justify-content: space-between
    border-radius: 5px
    height: 100%
    &_module_selector_side_to_set
      background-color: blue

  .hidden_input
    display: none
</style>
