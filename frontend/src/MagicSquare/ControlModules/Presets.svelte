<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { WasmInputId } from "../WasmInputId"
  import { PresetAction } from "./Preset"

  import { msStoreSettings } from '../../stores/msStoreSettings'
  import type { MsStoreSettings } from '../../stores/msStoreSettings'
  let msStoreSettingsVal: MsStoreSettings
  const unsubMsStoreSettings = msStoreSettings.subscribe((val: MsStoreSettings) => msStoreSettingsVal = val)

  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'
  let i18n = new I18n("magicSquare/presets")
  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)

  // TODO:
  // bring presets to PublicSquare
  import { currSquare, SquareType } from '../../stores/currSquare'
  let currSquareVal: SquareType
  const unsubCurrSquare = currSquare.subscribe((val: SquareType) => currSquareVal = val)

  // bank is non-input setting
  let bank: number
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

    // triggers update in Main that will read from localStorage
    updateUiSettings()
  }

  function handleBankClick(id: number) {
    bank = id
    msStoreSettings.update((prevSettings: MsStoreSettings) => {
      prevSettings.msPresetBank = id
      return prevSettings
    })
  }

  function handlePresetClick(idx: number) {
    presetNext = idx
  }

  onMount(() => {
    switch(currSquareVal){
      case SquareType.magic:
        bank = msStoreSettingsVal.msPresetBank
        break
      case SquareType.public:
        bank = msStoreSettingsVal.psPresetBank
        break
      case SquareType.none:
        bank = 0
        break
    }
  })

  onDestroy(() => {
    unsubCurrSquare()
    unsubLang()
    unsubMsStoreSettings()
  })
</script>

<section class="h-full flex flex-col justify-between items-stretch gap-2">
  <slot name="preset"/>
  <div class="grid grid-cols-4 grid-rows-1">
    <div class="title_m pl-5 col-span-2 w-full flex justify-between items-center">
      {i18n.t("curr", langVal)}
      <div>:</div>
    </div>
    <div class="title_m flex justify-around">
      {toBank(preset)}
    </div>
    <div class="title_m flex justify-around items-center">
      {preset + 1}
    </div>
  </div>
  <div class="grid grid-cols-4 grid-rows-1">
    <div class="title_m pl-5 col-span-2 w-full flex justify-between items-center">
      {i18n.t("next", langVal)}
      <div>:</div>
    </div>
    <div class="title_m flex justify-around">
      {toBank(presetNext)}
    </div>
    <div class="title_m flex justify-around items-center">
      {presetNext + 1}
    </div>
  </div>
  <div class="title pl-5 text-left">
    {i18n.t("bank", langVal)}
  </div>
  <div class="flex flex-col justify-between items-stretch">
    <div class="pl-5 pr-5 grid grid-cols-4">
      {#each {length: 4} as _, i}
        <button class="bank_button flex justify-around items-center"
                on:click={() => handleBankClick(i)}
                class:selected={bank === i}>
        {banks[i]}
        </button>
      {/each}
    </div>
  </div>
  <div class="title pl-5 text-left">
    {i18n.t("preset", langVal)}
  </div>
  <div class="grow flex flex-col justify-around items-stretch">
    <div class="preset_buttons grow pl-5 pr-5 pb-5 grid grid-cols-4 grid-rows-5">
      <div class="col-span-2 flex justify-around items-stretch">
        <button class="grow pl-2 pr-2 flex justify-around items-center"
                on:click={() => presetAction(PresetAction.set)}>
          {i18n.t("load", langVal)}
        </button>
      </div>
      <div class="col-span-2 flex justify-around items-stretch">
        <button class="grow flex justify-around items-center"
                on:click={() => presetAction(PresetAction.save)}>
          {i18n.t("save", langVal)}
        </button>
      </div>
      {#if bank === 0}
        {#each {length: 16} as _, idx}
          <button class="no_shadow flex justify-around items-center p-2"
                  on:click={() => handlePresetClick(idx)}
                  class:active={preset === idx}
                  class:selected={presetNext === idx}>
            {idx + 1}
          </button>
        {/each}
      {:else if bank === 1}
        {#each {length: 16} as _, idx}
          <button class="no_shadow flex justify-around items-center p-2"
                  on:click={() => handlePresetClick(idx + 16)}
                  class:active={preset === idx + 16}
                  class:selected={presetNext === idx + 16}>
            {idx + 17}
          </button>
        {/each}
      {:else if bank === 2}
        {#each {length: 16} as _, idx}
          <button class="no_shadow flex justify-around items-center p-2"
                  on:click={() => handlePresetClick(idx + 32)}
                  class:active={preset === idx + 32}
                  class:selected={presetNext === idx + 32}>
            {idx + 33}
          </button>
        {/each}
      {:else if bank === 3}
        {#each {length: 16} as _, idx}
          <button class="no_shadow flex justify-around items-center p-2"
                  on:click={() => handlePresetClick(idx + 48)}
                  class:active={preset === idx + 48}
                  class:selected={presetNext === idx + 48}>
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
  
  @import "../styles/control_module_title.sass"

  .bank_button
    min-width: 30px  

  .preset_buttons
    grid-template-rows: 0.5fr 1fr 1fr 1fr 1fr

  .title
    @include control_module_title

    &_m
      color: color.$blue-7
      font-size: text.$fs-m
      font-weight: text.$fw-l
      text-align: left

  .active
    border: solid
    border-width: 5px
    border-color: color.$red-7

  .selected
    background-color: color.$blue-7

  .no_shadow
    box-shadow: 0 0 
    min-width: 30px
</style>
