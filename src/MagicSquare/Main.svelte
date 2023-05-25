<script lang="ts" type="module">
  import { onMount, onDestroy } from 'svelte'
  import ControlRack from './ControlRack.svelte'

  let sideLength: number = 0

  var animation_id: any

  onMount(async () => {
    let element = document.getElementById("magic_square_canvas_container")
    sideLength = Math.min(Math.floor(element.offsetWidth / 1), Math.floor(element.offsetHeight /1)) - 25
    
    await wasm_bindgen() // loaded in index.html from ./pkg/src_rust.js
    const { MagicSquare, init_message } = wasm_bindgen
    console.log(
      init_message("Wasm Running for Magic Square")
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
    grid-template-areas: "display control"
    grid-template-columns: 27% 73%
    overflow: hidden

    &_canvas
      border: 5px solid color.$yellow-4
      border-radius: 5px
      margin: 0 -500px 0 -500px
      &_container
        height: 100%
        margin: 5px 0 10px 0
        flex-grow: 1

  .display
    grid-area: display
    align-items: center
    height: 100%

  .control
    grid-area: control
    height: 100%
</style>
