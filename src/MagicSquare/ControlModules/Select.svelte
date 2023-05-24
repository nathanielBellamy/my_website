<script lang="ts">
  import ControlModule from "../ControlModule.svelte"

  export let modules: string[] = []
  export let curr_mod_left: string = 'color'
  export let curr_mod_right: string = 'rotation'
</script>

<ControlModule>
  <div class="flex flex-col">
    <div class="module_selector flex flex-col space-around">
      <div class="module_selector_title">
        LEFT
      </div>
      <div class="module_selector_current_value">
        {curr_mod_left.toUpperCase()}
      </div>
      <select bind:value={curr_mod_left}
              class="control_module_select">
        {#each modules as mod}
          <option value={mod}>
            {mod.toUpperCase()}
          </option>
        {/each}
      </select>
    </div>
    <div class="module_selector flex flex-col">
      <div class="module_selector_title">
        RIGHT
      </div>
      <div class="module_selector_current_value">
        {curr_mod_right.toUpperCase()}
      </div>
      <select bind:value={curr_mod_right}
              class="control_module_select">
        {#each modules as mod}
          <option value={mod}>
            {mod.toUpperCase()}
          </option>
        {/each}
      </select>
    </div>
  </div>
</ControlModule>

<style lang="sass">
  .module_selector
    height: 100%
    max-width: 75px
    align-items: stretch
    
    &_title
      text-align: center

    &_current_value
      text-align: left
  
</style>
