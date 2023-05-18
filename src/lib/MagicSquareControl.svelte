<script lang="ts">
  import { onMount } from 'svelte'
  import iro from '@jaames/iro'

  const colorPickerOptions = {
    width: 250,
    height: 250,
    borderWidth: 5,
    borderColor: "#EFF0E9", // white in styles/color.sass
    layoutDirection: 'horizontal'
  }

  const colorPickerIds: string[] = [
    'magic_square_input_color_origin',
    'magic_square_input_color_nw',
    'magic_square_input_color_ne',
    'magic_square_input_color_se',
    'magic_square_input_color_sw'
  ]

  onMount(() => {
    var form = document.getElementById("magic_square_control")
    colorPickerIds.forEach((id:string) => {
      var input = document.getElementById(id)
      const picker = iro.ColorPicker(`#${id}_picker`, colorPickerOptions)

      picker.on('color:change', (color: any) => {
        input.value = `${color.rgba.r}, ${color.rgba.g}, ${color.rgba.b}, ${color.rgba.a}`
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })
    })
  })
</script>

<div id="magic_square_control"
     class="magic_square_control flex">
  {#each colorPickerIds as id}
    <div class="magic_square_input_color flex flex-col space-between">
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

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"
  
  .magic_square
    &_input_color
      max-width: 300px
      align-items: stretch
      
  .title
    font-size: text.$fs_l
    font-weight: text.$fw_l
  
  .color_picker
    display: flex
    justify-content: space-around

  .hidden_input
    display: none
</style>
