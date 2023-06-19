<script lang="ts">
  import { WasmInputId } from "../WasmInputId"
  import { PresetAction } from "./Preset"

  let bank: number = 0
  export let preset: number
  export let updateUiSettings: Function
  let presetNext: number = preset

  const banks: string[] = ['A', 'B', 'C', 'D']
  
  function toBank(preset: number): string {

    if (-1 < preset && preset < 16) {
      return banks[0]
    } else if (15 < preset && preset < 32) {
      return banks[1]
    } else if (31 < preset && preset < 48) {
      return banks[2]
    } else if (47 < preset && preset < 64) {
      return banks[3]
    }

    return ''
  }

  function presetAction(action: PresetAction) {
    preset = presetNext
    var input = document.getElementById(WasmInputId.preset)
    input.value = JSON.stringify({preset, action})
    input.dispatchEvent(new Event('input', {bubbles: true}))

    updateUiSettings()
  }

  function handlePresetClick(idx: number) {
    presetNext = idx
  }
</script>

<section class="h-full flex flex-col justify-around items-stretch">
  <slot name="preset"/>

  <div class="grid grid-cols-4 grid-rows-1">
    <div class="title col-span-2">
      ACTIVE:
    </div>
    <div class="title flex justify-around">
      {toBank(preset)}
    </div>
    <div class="title flex justify-around">
      {preset + 1}
    </div>
  </div>
  <div class="flex flex-col justify-between items-stretch">
    <div class="title pl-5 text-left">
      Bank
    </div>
    <div class="pl-5 pr-5 grid grid-cols-4">
      {#each {length: 4} as _, i}
        <button on:click={() => bank = i}
                class:selected={bank === i}>
        {banks[i]}
        </button>
      {/each}
    </div>
  </div>
  <div class="grow pt-5 flex flex-col justify-around items-stretch">
    <div class="title pl-5 text-left">
      Preset
    </div>
    <div class="grow p-5 grid grid-cols-4 grid-rows-5 gap-2">
      <button class="col-span-2"
              on:click={() => presetAction(PresetAction.set)}>
        LOAD
      </button>
      <button class="col-span-2"
              on:click={() => presetAction(PresetAction.save)}>
        SAVE
      </button>
      {#if bank === 0}
        {#each {length: 16} as _, idx}
          <button class="no_shadow"
                  on:click={() => handlePresetClick(idx)}
                  class:active={preset === idx}
                  class:selected={presetNext === idx}>
            {idx + 1}
          </button>
        {/each}
      {:else if bank === 1}
        {#each {length: 16} as _, idx}
          <button class="no_shadow"
                  on:click={() => handlePresetClick(idx + 16)}
                  class:active={preset === idx + 16}
                  class:selected={presetNext === idx + 16}>
            {idx + 17}
          </button>
        {/each}
      {:else if bank === 2}
        {#each {length: 16} as _, idx}
          <button class="no_shadow"
                  on:click={() => handlePresetClick(idx + 32)}
                  class:active={preset === idx + 32}
                  class:selected={presetNext === idx + 32}>
            {idx + 33}
          </button>
        {/each}
      {:else if bank === 3}
        {#each {length: 16} as _, idx}
          <button class="no_shadow"
                  on:click={() => handlePresetClick(idx + 48)}
                  class:active={preset === idx + 48}
                  class:selected={presetNext === idx + 48}>
            {idx + 49}
          </button>
        {/each}
      {/if}
    </div>
    <div class="pl-5 pr-5 flex justify-around items-center">

    </div>
  </div>
</section>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .title
    color: color.$blue-7
    font-size: text.$fs-ml
    font-weight: text.$fw-l

  .active
    border: solid
    border-width: 5px
    border-color: color.$red-5

  .selected
    background-color: color.$blue-7

  .no_shadow
    box-shadow: 0 0 
</style>
