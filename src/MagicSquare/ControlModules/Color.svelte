<script lang="ts">
  import { onMount } from 'svelte'
  import iro from '@jaames/iro'
  import ControlModule from '../ControlModule.svelte'
  
  const colorPickerOptions = {
    width: 50,
    height: 50,
    borderWidth: 5,
    borderColor: "#EFF0E9", // white in styles/color.sass
    layoutDirection: 'horizontal'
  }

  const colorPickerIds: string[] = [
    'magic_square_input_color_1',
    'magic_square_input_color_2',
    'magic_square_input_color_3',
    'magic_square_input_color_4',
    'magic_square_input_color_5',
    'magic_square_input_color_6',
    'magic_square_input_color_7',
    'magic_square_input_color_8',
  ]

  onMount(() => {
    colorPickerIds.forEach((id:string) => {
      var input = document.getElementById(id)
      const picker = iro.ColorPicker(`#${id}_picker`, colorPickerOptions)

      picker.on('color:change', (color: any) => {
        input.value = `${color.rgba.r},${color.rgba.g},${color.rgba.b},${color.rgba.a}`
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })
    })
  })
</script>

<ControlModule title="COLORS">
  <div class="flex flex-col">
    {#each colorPickerIds as id}
      <div class="flex flex-col space-between">
        <label for={id}
               class="title">
          {id.split("_").slice(-1)[0].toUpperCase()}
        </label>
        <div id={`${id}_picker`} 
             class="color_picker"/>
        <input id={id}
               class="hidden_input">
      </div>
    {/each}
  </div>
</ControlModule>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"
  
  .title
    font-size: text.$fs_l
    font-weight: text.$fw_l
    text-align: center
  
  .color_picker
    display: flex
    justify-content: space-around

  .hidden_input
    display: none
</style>




