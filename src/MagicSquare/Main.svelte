<script lang="ts" type="module">
  import { onMount } from 'svelte'
  import ControlRack from './ControlRack.svelte'

  let sideLength: number = 0

  var animation_id: any

  onMount(async () => {
    let element = document.getElementById("magic_square_canvas")
    sideLength = element.offsetWidth
    
    await wasm_bindgen() // loaded in index.html from ./pkg/src_rust.js
    const { MagicSquare, init_message } = wasm_bindgen
    console.log(init_message("Wasm Running for Magic Square"))
    
    MagicSquare.run()
  })

</script>

<div id="magic_square"
     class="magic_square grid grid-cols-2">
  <div class="magic_square_canvas_container flex flex-col justify-around display">
    <canvas id="magic_square_canvas"
            class="magic_square_canvas"
            height={sideLength}
            width={sideLength}/>
  </div>
  <div class="control">
    <ControlRack />
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .magic_square
    height: 100%
    width: 100%
    grid-template-areas: "display control"
    grid-template-columns: 40% 60%

    &_canvas
      width: 60%
      border: 5px solid color.$yellow-4
      &_container
        height: 100%
        align-items: center
        flex-grow: 1

  .display
    grid-area: display

  .control
    grid-area: control
</style>
