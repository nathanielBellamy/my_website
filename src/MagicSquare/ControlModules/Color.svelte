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

  onMount(async () => {
    colorPickerIds.forEach((id:string) => {
      var input = document.getElementById(id)
      const picker = iro.ColorPicker(`#${id}_picker`, colorPickerOptions)

      picker.on('color:change', (color: any) => {
        input.value = `${color.rgba.r},${color.rgba.g},${color.rgba.b},${color.rgba.a}`
        input.dispatchEvent(new Event('input', {bubbles: true}))
      })
    })
  })

  let curr_id: string = 'magic_square_input_color_1'
  
  function onClick(id: string) {
    curr_id = id
  }

  function idAsNumber(curr_id: string) {
    return parseInt(curr_id.split("_").slice(-1)[0].toUpperCase())
  }
</script>

<ControlModule title="COLOR">
  <div class="color_container flex flex-col justify-around">
    <div class="curr_picker flex justify-around">
      <div class="flex justify-around items-stretch">
        {#each colorPickerIds as id}
          <div id={`${id}_picker`}
               class="color_picker"
               class:hidden_input={curr_id !== id}/>
        {/each}
      </div>
    </div>
    <div class="color_rows grid grid-rows-2">
      <div class="color_row">
        {#each colorPickerIds.slice(0, 4) as id}
          <button class="color_button"
                  on:click={() => onClick(id)}
                  style:background-color={"green"}>
            {idAsNumber(id)}
            <input id={id}
                   class="hidden_input">
          </button>
        {/each}
      </div>
      <div class="color_row">
        {#each colorPickerIds.slice(4, 8) as id}
          <button class="color_button"
                  on:click={() => onClick(id)}
                  style:background-color={"blue"}>
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




