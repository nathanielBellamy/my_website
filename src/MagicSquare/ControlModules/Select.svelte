<script lang="ts">
  import ControlModule from "../ControlModule.svelte"

  export let modules: string[] = []
  export let curr_mod_left: string = 'color'
  export let curr_mod_right: string = 'rotation'

  let sideToSet:string = 'left'

  function onChange(mod:string, side: string) {
    if (side == 'left') {
      curr_mod_left = mod
    } else {
      curr_mod_right = mod
    }
  }
</script>

<ControlModule title="MODS">
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
</ControlModule>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .module_option
    border: 5px solid color.$red-4
    border-radius: 5px
    cursor: pointer
    font-size: text.$fs-s
    font-weight: text.$fw-l

  .selected
    background-color: color.$red-2

  .module_selector
    align-items: stretch
    justify-content: space-around
    border: 3px solid color.$red-2
    border-radius: 5px
    padding-top: 25px

  .hidden_input
    display: none
</style>
