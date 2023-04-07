<script lang="ts">
  import { onDestroy } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"

  let x: number = -1 
  let y: number = -1

  let magicSquare = new rust.MagicSquare

  const handleMouseMove = (e: any) => {
    x = e.clientX
    y = e.clientY
  }

  const handleMouseLeave = () => {
    x = -1 // negative indicates off element
    y = -1
  }

  // represents the loop used to capture mouse movement and deliver coordinates to WASM
  // this loop writes to the MagicSquareBuffer
  // an animation loop within a closure within WASM reads from the buffer
  const captureLoop = () => {
    console.log(magicSquare.write_to_buffer(x, y))
    captureInterval = requestAnimationFrame(captureLoop)
  };
    
  let captureInterval = requestAnimationFrame(captureLoop)
    
  onDestroy(async () => {
    cancelAnimationFrame(captureInterval)
  })
</script>

<div class="magic_square_container rounded-md flex flex-col justify-start">
  <div  class="magic_square_canvas_container grow"
        id="magic_square_canvas_container">
    <canvas id="magic_square"
            class="magic_square_canvas"
            on:mousemove={handleMouseMove}
            on:mouseleave={handleMouseLeave}/>
  </div>
</div>

<style lang="sass">
  .magic_square
    background-color: black

    &_container
      width: 100%
      height: 100%
      border: 2px solid black

    &_canvas
      background-color: gold
      width: 100%
      height: 100%
    
      &_container
        width: 100%
        background-color: red

</style>
