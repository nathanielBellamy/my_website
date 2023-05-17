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

  const colors = {
    'magic_square_input_color_origin': "rgba(0.0, 0.0, 0.0, 1.0)",
    'magic_square_input_color_nw': "rgba(0.0, 0.0, 0.0, 1.0)",
    'magic_square_input_color_ne': "rgba(0.0, 0.0, 0.0, 1.0)",
    'magic_square_input_color_se': "rgba(0.0, 0.0, 0.0, 1.0)",
    'magic_square_input_color_sw': "rgba(0.0, 0.0, 0.0, 1.0)",
  }

  $: colors

  const colorPickerIds = Object.keys(colors)

  onMount(() => {
    colorPickerIds.forEach((id:string) => {
      const picker = iro.ColorPicker(`#${id}_picker`, colorPickerOptions)

      picker.on('color:change', (color: any) => {
        colors[id] = `rgba( ${color.rgba.r}, ${color.rgba.g}, ${color.rgba.b}, ${color.rgba.a} )`
        var form = document.getElementById("magic_square_control")
        form.dispatchEvent(new Event('input'))
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
             value={colors[id]}
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
