<script lang="ts">
  import ControlModule from "../ControlModule.svelte"

  export let modules: string[] = []
  export let curr_mod_left: string = 'color'
  export let curr_mod_right: string = 'rotation'

  function onChange(mod:string, side: string) {
    if (side == 'left') {
      curr_mod_left = mod
    } else {
      curr_mod_right = mod
    }
  }
</script>

<ControlModule title="MODS">
  <form class="module_selectors flex flex-col gap-2">
    <div class="module_selector flex flex-col">
      {#each modules as mod}
        <label for={`mod_radio_${mod}_left`}
               class="module_option left"
               class:selected="{curr_mod_left === mod}"
               on:click={() => {
                 const input = document.getElementById(`mod_radio_${mod}_left`)
                 input.dispatchEvent(new Event("input", {bubbles: true}))
               }}
               on:keydown={() => {
                 const input = document.getElementById(`mod_radio_${mod}_left`)
                 input.dispatchEvent(new Event("input", {bubbles: true}))
               }}>
            {mod.toUpperCase()}
          <input id={`mod_radio_${mod}_left`}
                 value={mod}
                 type="radio"
                 name="wow"
                 checked={curr_mod_left === mod}
                 on:change={() => onChange(mod, 'left')}
                 class="hidden_input"/>
        </label>
      {/each}
    </div>
    <div class="module_selector flex flex-col">
      {#each modules as mod}
        <label for={`mod_radio_${mod}_right`}
               class="module_option right"
               class:selected="{curr_mod_right === mod}"
               on:click={() => {
                 const input = document.getElementById(`mod_radio_${mod}_right`)
                 input.dispatchEvent(new Event("input", {bubbles: true}))
               }}
               on:keydown={() => {
                 const input = document.getElementById(`mod_radio_${mod}_right`)
                 input.dispatchEvent(new Event("input", {bubbles: true}))
               }}>
            {mod.toUpperCase()}
          <input id={`mod_radio_${mod}_right`}
                 value={mod}
                 type="radio"
                 name="wow"
                 checked={curr_mod_right === mod}
                 on:change={() => onChange(mod, 'right')}
                 class="hidden_input"/>
        </label>
      {/each}
    </div>
  </form>
</ControlModule>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .module_option
    border: 5px solid color.$red-4
    border-radius: 5px
    margin: 5px
    cursor: pointer
    padding: 7px

  .selected
    background-color: color.$red-2

  .module_selectors
    align-items: stretch
    height: 100%
    min-width: 200px
    grid-template-areas: "left", "right"
    overflow-y: scroll

  .left
    grid-area: "left"

  .right
    grid-area: "right"
    
  .module_selector
    align-items: stretch
    justify-content: space-around
    border: 3px solid color.$red-2
    border-radius: 5px

  .hidden_input
    display: none
</style>
