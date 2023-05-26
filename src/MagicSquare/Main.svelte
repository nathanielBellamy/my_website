<script lang="ts" type="module">
  import { onMount, onDestroy } from 'svelte'
  import ControlRack from './ControlRack.svelte'

  let sideLength: number = 0

  var animation_id: any

  onMount(async () => {
    let element = document.getElementById("magic_square_canvas_container")
    sideLength = Math.floor(Math.min(element.offsetWidth, element.offsetHeight) / 1) - 25

    // clear old ui_buffer from localStorage
    localStorage.clear()
    
    await wasm_bindgen() // loaded in index.html from ./pkg/src_rust.js
    const { MagicSquare, init_message } = wasm_bindgen
    console.log(
      init_message("Magic Square Wasm!")
    )
    
    MagicSquare.run()
  })

  onDestroy(async () => {
    let app = document.getElementById(("app_main"))
    app.dispatchEvent(new Event("destroymswasm", {bubbles: true}))
  })

</script>

<div id="magic_square"
     class="magic_square flex flex-wrap gap-2">
  <div id="magic_square_canvas_container"
       class="magic_square_canvas_container flex flex-col justify-around display">
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
    overflow-y: scroll

    &_canvas
      border-top: 5px double color.$blue-7
      border-bottom: 5px double color.$blue-7
      border-radius: 10px
      &_container
        height: 100%
        background: color.$black-blue-horiz-grad
        border: 10px double color.$blue-7
        border-radius: 5px
        flex-grow: 1

  .display
    align-items: center
  
  .control
    flex-grow: 1
    height: 100%

</style>
