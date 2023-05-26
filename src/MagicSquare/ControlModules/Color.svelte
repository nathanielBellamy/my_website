<script lang="ts">
  import { onMount } from 'svelte'
  import iro from '@jaames/iro'
  import ControlModule from '../ControlModule.svelte'
  
  const colorPickerOptions = {
    width: 175,
    height: 150,
    borderWidth: 5,
    borderColor: "#EFF0E9", // white in styles/color.sass
    layoutDirection: 'vertical'
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

  let colorData: ColorData = {
    color1: {
      id: 'magic_square_input_color_1',
      idx: 1,
      rgba: [],
      picker: null
    },
    color2: {
      id: 'magic_square_input_color_2',
      idx: 2,
      rgba: [],
      picker: null
    },   
    color3: {
      id: 'magic_square_input_color_3',
      idx: 3,
      rgba: [],
      picker: null
    },
    color4: {
      id: 'magic_square_input_color_4',
      idx: 4,
      rgba: [],
      picker: null
    },
    color5: {
      id: 'magic_square_input_color_5',
      idx: 5,
      rgba: [],
      picker: null
    },
   color6: {
      id: 'magic_square_input_color_6',
      idx: 6,
      rgba: [],
      picker: null
    },
   color7: {
      id: 'magic_square_input_color_7',
      idx: 7,
      rgba: [],
      picker: null
    },
   color8: {
      id: 'magic_square_input_color_8',
      idx: 8,
      rgba: [],
      picker: null
    },
  }

  function setColorData(source: any) {
    // source is JSON object serialized from rust
    colorData.color1.rgba = source.settings.color_1.map((x: number) => 255 * x)
    colorData.color2.rgba = source.settings.color_2.map((x: number) => 255 * x)
    colorData.color3.rgba = source.settings.color_3.map((x: number) => 255 * x)
    colorData.color4.rgba = source.settings.color_4.map((x: number) => 255 * x)
    colorData.color5.rgba = source.settings.color_5.map((x: number) => 255 * x)
    colorData.color6.rgba = source.settings.color_6.map((x: number) => 255 * x)
    colorData.color7.rgba = source.settings.color_7.map((x: number) => 255 * x)
    colorData.color8.rgba = source.settings.color_8.map((x: number) => 255 * x)

  }
  
  const storageKey = 'magic_square_storage'

  function getStorageData () {
    return JSON.parse(localStorage.getItem(storageKey))
  }

  onMount(async () => {
    const storageData = getStorageData()
    setColorData(storageData)

    Object.values(colorData).forEach((color:Color) => {
      var input = document.getElementById(color.id)
      var picker = iro.ColorPicker(`#${color.id}_picker`, colorPickerOptions)
      picker.color.rgb = { r: color.rgba[0], g: color.rgba[1], b: color.rgba[2] }

      picker.on('color:change', (color: any) => {
        input.value = `${color.rgba.r},${color.rgba.g},${color.rgba.b},${color.rgba.a}`
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })

      colorData[`color${color.idx}`].picker = picker 
    })
  })

  function handleStorageEvent (e:Event) {
    if (e.key === storageKey){
      const storageData = getStorageData()
      setColorData(storageData)
      colorData = colorData
      console.dir({
        storage: storageData.settings.color_1,
        local: colorData.color1.rgba
      })
    }
  }

  window.addEventListener("storage", (e:any) => handleStorageEvent(e))

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

<ControlModule title="COLOR">
  <div class="color_container flex flex-col justify-around">
    <div class="curr_picker flex justify-around">
      <div class="flex justify-around items-stretch">
        {#each Object.values(colorData) as { id }}
          <div id={`${id}_picker`}
               class="color_picker"
               class:hidden_input={curr_id !== id}/>
        {/each}
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
</ControlModule>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .color_picker
    display: flex
    justify-content: space-between
    width: 100%
    flex-grow: 1
  
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
</style>




