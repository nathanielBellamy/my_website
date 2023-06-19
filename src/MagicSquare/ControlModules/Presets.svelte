<script lang="ts">
  import { WasmInputId } from "../WasmInputId"
  import { PresetAction } from "./Preset"

  let bank: number = 0
  let preset: number = 0

  const banks: string[] = ['A', 'B', 'C', 'D']
  
  function toBank(preset: number): string {

    if ( 0< preset && preset < 16) {
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
    var input = document.getElementById(WasmInputId.preset)
    input.value = JSON.stringify({preset, action})
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  function handlePresetClick(idx: number) {
    preset = idx
    presetAction(PresetAction.load)
  }
</script>

<section class="h-full flex flex-col justify-around items-stretch">
  <slot name="preset"/>
  <div class="flex justify-around items-center">
    <div class="flex justify-around">
      CURR
    </div>
    <div class="flex justify-around">
      BANK {toBank(preset)}
    </div>
    <div class="flex justify-around">
      PRESET {preset + 1}
    </div>
  </div>
  <div class="flex justify-around items-center">
    <button on:click={() => presetAction(PresetAction.set)}>
      SET
    </button>
    <button on:click={() => presetAction(PresetAction.show)}>
      SHOW
    </button>
    <button on:click={() => presetAction(PresetAction.save)}>
      SAVE
    </button>
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
    <div class="grow p-5 grid grid-cols-4 grid-rows-4">
      {#if bank === 0}
        {#each {length: 16} as _, idx}
          <button on:click={() => handlePresetClick(idx)}
                  class:selected={preset === idx}>
            {idx + 1}
          </button>
        {/each}
      {:else if bank === 1}
        {#each {length: 16} as _, idx}
          <button on:click={() => handlePresetClick(idx + 16)}
                  class:selected={preset === idx + 16}>
            {idx + 17}
          </button>
        {/each}
      {:else if bank === 2}
        {#each {length: 16} as _, idx}
          <button on:click={() => handlePresetClick(idx + 32)}
                  class:selected={preset === idx + 32}>
            {idx + 33}
          </button>
        {/each}

      {:else if bank === 3}
        {#each {length: 16} as _, idx}
          <button on:click={() => handlePresetClick(idx + 48)}
                  class:selected={preset === idx + 48}>
            {idx + 49}
          </button>
        {/each}
      {/if}
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

  .selected
    background-color: color.$blue-7
</style>
