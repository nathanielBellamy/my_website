<script lang="ts" type="module">
  import { onMount, onDestroy } from 'svelte'
  import ControlRack from './ControlRack.svelte'

  let sideLength: number = 0

  var animation_id: any

  onMount(async () => {
    let element = document.getElementById("magic_square_canvas_container")
    sideLength = Math.floor(Math.min(element.offsetWidth, element.offsetHeight) / 1) - 25
    
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
     class="magic_square flex">
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
    display: grid
    grid-template-areas: "display control"
    grid-template-columns: 30% 70%
    grid-template-rows: 100%
    overflow: hidden

    &_canvas
      border-top: 5px double color.$blue-4
      border-bottom: 5px double color.$blue-4
      border-radius: 10px
      &_container
        height: 100%
        background: color.$black-blue-horiz-grad
        border: 10px double color.$blue-4
        border-radius: 5px

  .display
    grid-area: display
    align-items: center

  .control
    grid-area: control
</style>
