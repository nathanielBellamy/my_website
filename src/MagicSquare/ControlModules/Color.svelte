<script lang="ts">
  import { afterUpdate, onMount } from 'svelte'
  import iro from '@jaames/iro'
  import App from '../../App.svelte'

  function rgbaToString(rgba: number[]): string {
    // rgba = !!rgba ? rgba : [0, 255, 0, 1]
    // while we have some infrastructure set up to accept opacity values
    // our WebGl implimentation does not make use of them at the moment
    // so we keep everything rgb in practice
    // console.dir({r: rgba[0], g: rgba[1]})
    return `rgba(${rgba[0]}, ${rgba[1]}, ${rgba[2]}, 1)`
  }
  
  export let color1: number[]
  export let color2: number[]
  export let color3: number[]
  export let color4: number[]
  export let color5: number[]
  export let color6: number[]
  export let color7: number[]
  export let color8: number[]

  let color1Str: string
  let color2Str: string
  let color3Str: string
  let color4Str: string
  let color5Str: string
  let color6Str: string
  let color7Str: string
  let color8Str: string

  $: color1Str = rgbaToString(color1)
  $: color2Str = rgbaToString(color2)
  $: color3Str = rgbaToString(color3)
  $: color4Str = rgbaToString(color4)
  $: color5Str = rgbaToString(color5)
  $: color6Str = rgbaToString(color6)
  $: color7Str = rgbaToString(color7)
  $: color8Str = rgbaToString(color8)
  let currId: string = 'magic_square_input_color_1'
  

  enum HiddenInputIds {
    color1 = "magic_square_input_color_1",
    color2 = "magic_square_input_color_2",
    color3 = "magic_square_input_color_3",
    color4 = "magic_square_input_color_4",
    color5 = "magic_square_input_color_5",
    color6 = "magic_square_input_color_6",
    color7 = "magic_square_input_color_7",
    color8 = "magic_square_input_color_8",
  }

  const colorPickerOptions = {
    width: 110,
    height: 90,
    borderWidth: 2,
    borderColor: "#EFF0E9", // white in styles/color.sass
    layoutDirection: 'horizontal'
  }

  function onIdClick(id: number) {
    currId = toIdString(id)
  }

  function toIdString(id: number) {
    return `magic_square_input_color_${id}`
  }

  function setNewColor(color: number[], id: number) {
    switch (id) {
      case 1:
        color1 = [...color]
        break
      case 2:
        color2 = [...color]
        break
      case 3:
        color3 = [...color]
        break
      case 4:
        color4 = [...color]
        break
      case 5:
        color5 = [...color]
        break
      case 6:
        color6 = [...color]
        break
      case 7:
        color7 = [...color]
        break
      case 8:
        color8 = [...color]
        break
    }
  }

  onMount(async () => {
    // get height/width for picker
    var colorPickerDiv: any = document.getElementById('color_mode_and_curr')
    const width: number = Math.floor(colorPickerDiv.offsetWidth / 1.7);

    [color1, color2, color3, color4, color5, color6, color7, color8].forEach((color: number[], idx: number) => {
      const id: number = idx + 1
      const idStr: string = toIdString(id)
      var picker = iro.ColorPicker(`#${idStr}_picker`, Object.assign(colorPickerOptions, {height: width, width}))
      picker.color.rgba = { r: color[0], g: color[1], b: color[2], a: 1 }
      var input = document.getElementById(idStr)

      picker.on('color:change', (newColor: any) => {
        const arr = [newColor.rgba.r, newColor.rgba.g, newColor.rgba.b, 1]
        // TODO: simplify/unwind
        // -> due to color value being bound to input.value
        // -> order matters here
        // -> dispatchEvent synchronyously to send the String to be parsed by Wasm
        // -> sets color1 to be the corresponding number[]
        input.value = `${arr[0]},${arr[1]},${arr[2]},1`
        input.dispatchEvent(new Event('input', {bubbles: true}))
        setNewColor(arr, id)
      })
    })
  })

  
</script>

<div class="color_container flex flex-col justify-between items-stretch">
  <div id="color_mode_and_curr"
       class="color_mode_and_curr flex flex-col justify-around items-center">
    <div class="color_modes flex flex-col justify-around items-stretch">
      <div class="color_mode flex justify-evenly">
        <button class="color_mode_option">
          Out
        </button>
        <button class="color_mode_option">
          In
        </button>
        <button class="color_mode_option">
          Fix
        </button>
      </div>
      <div class="color_mode flex flex-col justify-between">
        <button class="color_mode_option">
          Eight
        </button>
        <button class="color_mode_option">
          Grad
        </button>
      </div>
    </div>
    <div  id="magic_square_color_curr_picker"
          class="curr_picker flex justify-around">
      <div class="curr_picker_id">
        {currId.split("_").slice(-1)[0]}
      </div>
      <div class="mt-5 flex justify-around items-stretch">
        {#each [1,2,3,4,5,6,7,8] as id }
          <div id={`${toIdString(id)}_picker`}
               class:hidden_input={currId !== toIdString(id)}/>
        {/each}
      </div>
    </div>
  </div>
  <div class="color_rows grid grid-rows-2">
    <div class="color_row">
      {#each [color1Str, color2Str, color3Str, color4Str] as rgbaStr, idx}
        <button class="color_button"
                on:click={() => onIdClick(idx + 1)}
                style:background-color={rgbaStr}>
          {idx + 1}
        </button>
      {/each}
    </div>
    <div class="color_row">
      {#each [color5Str, color6Str, color7Str, color8Str] as rgbaStr, idx}
        <button class="color_button"
                on:click={() => onIdClick(idx + 5)}
                style:background-color={rgbaStr}>
          {idx+5}
        </button>
      {/each}
    </div>
  </div>
  <slot />
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .color_mode_and_curr
    flex-grow: 1
    width: 90%
    margin-left: 5%
  .color_mode
    &s
      flex-grow: .1
      border-bottom: 5px double color.$blue-7
      padding-bottom: 10px
    &_option
      flex: 1 1 0px
      font-weight: text.$fw-l
      font-size: text.$fs-m
      color: color.$blue-7

  .color_picker
    display: flex
    justify-content: space-between
    width: 100%
    flex-grow: 1
    margin: 0 5px 0 5px
  
  .color_rows
    flex-grow: .75

  .color_row
    width: 100%
    display: flex
    justify-content: stretch
    align-items: stretch

  .color_button
    flex-grow: 1

  .hidden_input
    display: none

  .color_container
    height: 100%
    overflow-y: scroll
    overflow-x: hidden
    padding: 10px 0 10px 0

  .curr_picker
    position: relative
    &_id
      position: absolute
      margin-top: 30%
      margin-right: 40px
      z-index: 100
      font-weight: text.$fw-m
      font-size: text.$fs-xl
      color: color.$black-7
      pointer-events: none
</style>




