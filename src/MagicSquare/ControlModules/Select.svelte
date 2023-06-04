<script lang="ts">
  export let modules: string[] = []
  export let curr_mod_left: string = 'color'
  export let curr_mod_right: string = 'rotation'

  let sideToSet:string = 'left'

  function handleModKeydown(e: any, mod: string) {
    if (e.keyCode === 13){
      setMod(mod)
    }
  }

  function setMod(mod: string) {
    if (mod === curr_mod_left || mod === curr_mod_right) return
    if (sideToSet == 'left') {
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
            class:side_set_left_selected="{sideToSet === 'left'}"
            on:dblclick={() => swap()}
            on:click={() => sideToSet = 'left'}>
      LEFT
    </button>
    <button class="side_set side_set_right"
            class:side_set_right_selected="{sideToSet === 'right'}"
            on:dblclick={() => swap()}
            on:click={() => sideToSet = 'right'}>
      RIGHT
    </button>
  </div>
  {#each modules as mod}
    <button class="module_option"
            class:selected_left="{curr_mod_left === mod}"
            class:selected_right="{curr_mod_right === mod}"
            on:click={() => setMod(mod)}
            on:keydown={(e) => handleModKeydown(e, mod)}>
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
