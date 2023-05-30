<script lang="ts">
  import { onMount, afterUpdate } from 'svelte'
  import iro from '@jaames/iro'
  
  interface Props {// {[key: string]: rgba[]}
    color1: number[],
    color2: number[],
    color3: number[],
    color4: number[],
    color5: number[],
    color6: number[],
    color7: number[],
    color8: number[],
  }

  export let props: Props

  const colorPickerOptions = {
    width: 110,
    height: 90,
    borderWidth: 2,
    borderColor: "#EFF0E9", // white in styles/color.sass
    layoutDirection: 'horizontal'
  }

  interface Color {
    id: string,
    idx: number,
    rgba: number[],
    picker: any
  }

  interface ColorData {
    color1: Color,
    color2: Color,
    color3: Color,
    color4: Color,
    color5: Color,
    color6: Color,
    color7: Color,
    color8: Color,
  }

  afterUpdate(() => console.log(props))

  let colorData: ColorData
  $: colorData = {
    color1: {
      id: 'magic_square_input_color_1',
      idx: 1,
      rgba: props.color1,
      picker: null
    },
    color2: {
      id: 'magic_square_input_color_2',
      idx: 2,
      rgba: props.color2,
      picker: null
    },   
    color3: {
      id: 'magic_square_input_color_3',
      idx: 3,
      rgba: props.color3,
      picker: null
    },
    color4: {
      id: 'magic_square_input_color_4',
      idx: 4,
      rgba: props.color4,
      picker: null
    },
    color5: {
      id: 'magic_square_input_color_5',
      idx: 5,
      rgba: props.color5,
      picker: null
    },
   color6: {
      id: 'magic_square_input_color_6',
      idx: 6,
      rgba: props.color6,
      picker: null
    },
   color7: {
      id: 'magic_square_input_color_7',
      idx: 7,
      rgba: props.color7,
      picker: null
    },
   color8: {
      id: 'magic_square_input_color_8',
      idx: 8,
      rgba: props.color8,
      picker: null
    },
  }
  
  onMount(async () => {

    Object.values(colorData).forEach((color:Color) => {
      var input = document.getElementById(color.id)
      var picker = iro.ColorPicker(`#${color.id}_picker`, colorPickerOptions)
      picker.color.rgb = { r: color.rgba[0], g: color.rgba[1], b: color.rgba[2] }

      picker.on('color:change', (color: any) => {
        input.value = `${color.rgba.r},${color.rgba.g},${color.rgba.b},1`
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })

      colorData[`color${color.idx}`].picker = picker 
    })
  })

  let curr_id: string = 'magic_square_input_color_1'
  
  function onClick(id: string) {
    curr_id = id
  }

  function idAsNumber(curr_id: string) {
    return parseInt(curr_id.split("_").slice(-1)[0].toUpperCase())
  }

  function rgbaToString(rgba: number[]) {
    // while we have some infrastructure set up to accept opacity values
    // our WebGl implimentation does not make use of them at the moment
    // so we keep everything rgb in practice
    return `rgba(${rgba[0]}, ${rgba[1]}, ${rgba[2]}, 1)`
  }
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
        {#each Object.values(colorData) as { id }}
          <div id={`${id}_picker`}
               class="color_picker"
               class:hidden_input={curr_id !== id}/>
        {/each}
      </div>
    </div>
  </div>
  <div class="color_rows grid grid-rows-2">
    <div class="color_row">
      {#each [colorData.color1, colorData.color2, colorData.color3, colorData.color4] as { id, rgba }}
        <button class="color_button"
                on:click={() => onClick(id)}
                style:background-color={rgbaToString(rgba)}>
          {idAsNumber(id)}
          <input id={id}
                 class="hidden_input">
        </button>
      {/each}
    </div>
    <div class="color_row">
      {#each [colorData.color5, colorData.color6, colorData.color7, colorData.color8] as { id, rgba }}
        <button class="color_button"
                on:click={() => onClick(id)}
                style:background-color={rgbaToString(rgba)}>
          {idAsNumber(id)}
          <input id={id}
                 class="hidden_input">
        </button>
      {/each}
    </div>
  </div>
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




