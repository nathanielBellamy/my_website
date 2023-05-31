<script lang="ts">
  import { afterUpdate, onMount } from 'svelte'
  import iro from '@jaames/iro'
  
  export let color1: number[]
  export let color2: number[]
  export let color3: number[]
  export let color4: number[]
  export let color5: number[]
  export let color6: number[]
  export let color7: number[]
  export let color8: number[]


  let curr_id: string = 'magic_square_input_color_1'

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
    curr_id = toIdString(id)
  }

  function toIdString(id: number) {
    return  `magic_square_input_color_${id}`
  }

  function rgbaToString(rgba: number[]) {
    rgba = !!rgba ? rgba : [1, 0, 1]
    // while we have some infrastructure set up to accept opacity values
    // our WebGl implimentation does not make use of them at the moment
    // so we keep everything rgb in practice
    return `rgba(${rgba[0]}, ${rgba[1]}, ${rgba[2]}, 1)`
  }

  onMount(() => {
    [color1, color2, color3, color4, color5, color6, color7, color8].forEach((color: number[], idx: number) => {
      console.log("OOOOOOOO")
      console.log(color)
      const id: number = idx + 1
      const idStr: string = toIdString(id)
      var input = document.getElementById(idStr)
      var picker = iro.ColorPicker(`#${idStr}_picker`, colorPickerOptions)
      // picker.color.rgba = { r: color[0], g: color[1], b: color[2], a: 1 }
      picker.color.rgba = { r: 255, g: 0, b: 255, a: 1 }

      picker.on('color:change', (color: any) => {
        input.value = `${color.rgba.r},${color.rgba.g},${color.rgba.b},1`
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })
    })
  })
</script>

<div class="color_container flex flex-col justify-around">
  <div class="color_mode_and_curr flex flex-col justify-around items-center">
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
    <div class="curr_picker flex justify-around">
      <div class="curr_picker_id">
        {curr_id.split("_").slice(-1)[0]}
      </div>
      <div class="flex justify-around items-stretch">
        {#each [1,2,3,4,5,6,7,8] as id }
          <div id={`${toIdString(id)}_picker`}
               class="color_picker"
               class:hidden_input={curr_id !== toIdString(id)}/>
        {/each}
      </div>
    </div>
  </div>
  <div class="color_rows grid grid-rows-2">
    <div class="color_row">
      {#each [color1, color2, color3, color4] as rgba, idx}
        <button class="color_button"
                on:click={() => onIdClick(idx + 1)}
                style:background-color={rgbaToString(rgba)}>
          {idx + 1}
        </button>
      {/each}
    </div>
    <div class="color_row">
      {#each [color5, color6, color7, color8] as rgba, idx}
        <button class="color_button"
                on:click={() => onIdClick(idx +1)}
                style:background-color={rgbaToString(rgba)}>
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
    overflow: hidden
    padding: 10px 0 10px 0

  .curr_picker
    position: relative
    &_id
      position: absolute
      margin-top: 20px
      margin-right: 40px
      z-index: 100
      font-weight: text.$fw-m
      font-size: text.$fs-xl
      color: color.$black-7
      pointer-events: none
</style>




