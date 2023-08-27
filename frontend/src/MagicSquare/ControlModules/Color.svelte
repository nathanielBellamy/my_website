<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import iro from '@jaames/iro'
  import { ColorDirection } from './Color'
  import { WasmInputId } from '../WasmInputId'
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'

  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)
  onDestroy(unsubLang)
  let i18n = new I18n("magicSquare/color")


  import { smallScreen } from '../../stores/smallScreen'
  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean | null) => smallScreenVal = val)

  function rgbaToString(rgba: number[]): string {
    // rgba = !!rgba ? rgba : [0, 255, 0, 1]
    // while we have some infrastructure set up to accept opacity values
    // our WebGl implimentation does not make use of them at the moment
    // so we keep everything rgb in practice
    // console.dir({r: rgba[0], g: rgba[1], b: rgba[2]})
    return `rgba(${rgba[0]}, ${rgba[1]}, ${rgba[2]}, 1)`
  }

  let idx_a: number = 0
  let idx_b: number = 15

  $: idxLeft = idx_a < idx_b ? idx_a : idx_b
  $: idxRight = idx_a > idx_b ? idx_a : idx_b

  export let colorDirection: ColorDirection
  export let colors: number[][]
  // will store iro colorPicker elements here
  let colorPickers: any[] = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]

  function handleColorDirectionClick(cd: ColorDirection) {
    colorDirection = cd
    const input = document.getElementById(WasmInputId.colorDirection)
    input.value = cd
    input.dispatchEvent(new Event('input', { bubbles: true }))
  }

  function colorGradientAtStep(step: number, width: number): number[] {
    const t: number = step / width
    var result: number[] = [0, 0, 0, 0]
    for (let idx = 0; idx < 4; idx++) {
      result[idx] = (1 - t) * colors[idxLeft][idx] + t * colors[idxRight][idx]
    }
    return result
  }

  function setColorGradient() {
    const width: number = idxRight - idxLeft
    if (!!width) {
      var input = document.getElementById(WasmInputId.colors)
      let step: number = 0
      while (idxLeft + step < idxRight + 1) {
        const newColor = colorGradientAtStep(step, width)
        colors[idxLeft + step] = newColor
        colorPickers[idxLeft + step].color.rgba = {r: newColor[0], g: newColor[1], b: newColor[2], a: newColor[3]}
        input.value = JSON.stringify({idx: idxLeft + step, rgba: newColor})
        input.dispatchEvent(new Event('input', {bubbles: true}))
        step += 1 
      }
      colors = colors.map(color => [...color])
    }
  }

  function handleGradientIndexChange(e: any, id: string) {
    e.stopPropagation()
    switch (id) {
      case 'a':
        idx_a = parseInt(e.target.value)
        break
      case 'b':
        idx_b = parseInt(e.target.value)
        break
      default:
        break
    }
  }

  let colorStrings: string[]
  $: colorStrings = colors.map(x => rgbaToString(x))
  $: gradient = `linear-gradient(90deg, ${colorStrings[idx_a]} 0%, ${colorStrings[idx_b]} 100%)`
  let currIdx: number = 0

  const colorPickerOptions = {
    width: 110,
    height: 90,
    borderWidth: 2,
    borderColor: "#EFF0E9", // white in styles/color.sass
    layoutDirection: 'horizontal'
  }

  function onIdxClick(idx: number) {
    currIdx = idx
  }

  function setNewColor(color: number[], idx: number) {
    colors[idx] = [...color]
    colors = colors.map(color => [...color])
  }

  function toIdxString(idx: number): string {
    return `ms_color_picker_picker_${idx}`
  }

  function deriveColorPickerWidth(iw: number, ih: number): number {
    var res: number
    if (smallScreenVal) {
      const shorterLeg: number = Math.min(iw, ih)
      res = Math.floor(shorterLeg / 3.5);
    } else {
      res = Math.floor(iw / 10);
    }
    return Math.max(115, res)
  }

  function handlePickerResize() {
    colorPickers.forEach((p: any) => p.resize(colorPickerWidth))
  }

  let innerWidth: number = 375 // assume small
  let innerHeight: number = 375 // assume small
  $: colorPickerWidth = deriveColorPickerWidth(innerWidth, innerHeight)

  onMount(async () => {
    // get height/width for picker
    const width: number = deriveColorPickerWidth(innerWidth, innerHeight)
    var input = document.getElementById(WasmInputId.colors)

    colors.forEach((color: number[], idx: number) => {
      var picker = iro.ColorPicker(`#${toIdxString(idx)}`, Object.assign(colorPickerOptions, {height: width, width}))
      picker.color.rgba = { r: color[0], g: color[1], b: color[2], a: 1 }

      picker.on('color:change', (newColor: any) => {
        const rgba = [newColor.rgba.r, newColor.rgba.g, newColor.rgba.b, 1]       // TODO: simplify/unwind
        // -> due to color value being bound to input.value
        // -> order matters here
        // -> dispatchEvent synchronyously to send the String to be parsed by Wasm
        // -> sets color1 to be the corresponding number[]
        colors[idx] = rgba
        input.value = JSON.stringify({idx, rgba})
        input.dispatchEvent(new Event('input', {bubbles: true}))
        setNewColor(rgba, idx)
      })
      colorPickers[idx] = picker
    })

    window.addEventListener('resize', handlePickerResize)
  })

  onDestroy(() => {
    window.removeEventListener('resize', handlePickerResize)
  })
</script>

<svelte:window bind:innerWidth />

<div class="color_container h-full pb-10 flex flex-col justify-between items-stretch">
  <div id="color_mode_and_curr"
       class="grow w-full flex flex-col justify-around items-center">
    <div class="color_modes w-full flex flex-col justify-around items-stretch">
      <div class="title">
        {i18n.t("animation", langVal)}
      </div>
      <div class="grow w-full pl-2 pr-2 flex justify-around">
        <button class="grow color_mode_option"
                class:selected={colorDirection === ColorDirection.out}
                on:click={() => handleColorDirectionClick(ColorDirection.out)}>
          {i18n.t("out", langVal)}
        </button>
        <button class="grow color_mode_option"
                class:selected={colorDirection === ColorDirection.in}
                on:click={() => handleColorDirectionClick(ColorDirection.in)}>
          {i18n.t("in", langVal)}
        </button>
        <button class="grow color_mode_option"
                class:selected={colorDirection === ColorDirection.fix}
                on:click={() => handleColorDirectionClick(ColorDirection.fix)}>
          {i18n.t("fix", langVal)}
        </button>
      </div>
      <div class="grow pl-2 pr-2 m-2">
        <slot name="speed"/>
      </div>
    </div>
    <div  id="magic_square_color_curr_picker"
          class="curr_picker flex justify-around">
      <div class="curr_picker_id">
        {currIdx + 1}
      </div>
      <div class="mt-5 flex justify-around items-stretch">
        {#each {length: 16} as _, idx }
          <div id={toIdxString(idx)}
               class:hidden_input={currIdx !== idx}/>
        {/each}
      </div>
    </div>
  </div>
  <div class="color_rows pl-2 pr-2 grid grid-cols-4 grid-rows-4">
      {#each colorStrings as rgbaStr, idx}
        <button class="color_button"
                on:click={() => onIdxClick(idx)}
                style:background-color={rgbaStr}>
          {idx + 1}
        </button>
      {/each}
  </div>
  <div class="color_gradient h-fit rounded-md pt-4 p-2 m-2 grid grid-cols-2 grid-rows-2 gap-y-2 gap-x-1">
    <select bind:value={idx_a}
            class="h-fit"
            on:input={(e) => e.stopPropagation()}
            on:change={(e) => handleGradientIndexChange(e, 'a')}>
      {#each {length: 16} as _, idx}
        <option selected={idx_a === idx}
                value={idx}>
          {idx + 1}
        </option>
      {/each}
    </select>
    <select bind:value={idx_b}
            class="h-fit"
            on:input={(e) => e.stopPropagation()}
            on:change={(e) => handleGradientIndexChange(e, 'b')}>
      {#each {length: 16} as _, idx}
        <option selected={idx_b === idx}
                value={idx}>
          {idx + 1}
        </option>
      {/each}
    </select>
    <button class="h-6 col-span-2"
            style:background="{gradient}"
            on:click={setColorGradient}/>
  </div>
  <slot name="hiddenInputs"/>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"
  @import "../styles/control_module_title.sass"

  .title
    @include control_module_title


  .selected
    background-color: color.$blue-4

  .color_gradient
    border: 3px solid $blue-7

  .color_mode
    &s
      flex-grow: .75
      border-bottom: 5px double color.$blue-7
      padding-bottom: 10px
    &_option
      font-weight: text.$fw-l
      font-size: text.$fs-m
      color: color.$blue-7
  
  .color_rows
    flex-grow: .75

  .color_button
    flex-grow: 1

  .hidden_input
    display: none

  .color_container
    overflow-y: scroll
    overflow-x: hidden

  .curr_picker
    position: relative
    &_id
      position: absolute
      margin-top: 26%
      margin-right: 40px
      z-index: 100
      font-weight: text.$fw-m
      font-size: text.$fs-xl
      color: color.$black
      pointer-events: none
</style>




