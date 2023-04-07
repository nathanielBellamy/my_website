<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"
  

  let buffer = new rust.AppBuffer
 
  let x: number = 0 
  let y: number = 0

  const handleMouseMove = (e: any) => {
    console.log(buffer.set_point(e.clientX, e.clientY))
    x = buffer.x
    y = buffer.y
  }

  let magicSquareWidth: number = 0
  let magicSquareHeight: number = 0
  
  let captureInterval: number

  onMount(async () => {
    const pre = document.getElementById("magic_square_canvas_container")

    const captureLoop = () => {
      buffer.write()
      requestAnimationFrame(renderLoop)
    };
    
    captureInterval = requestAnimationFrame(captureLoop)
   })

// https://rustwasm.github.io/docs/wasm-bindgen/examples/request-animation-frame.html

  onDestroy(() => {
    window.cancelAnimationFrame(captureInterval) 
	});

</script>

<div class="magic_square_container rounded-md flex flex-col justify-start">
  <div class="flex flex-row justify-around">
    <div>
     x:: {x}
    </div>
    <div>
     y:: {y}
    </div>
  </div>
  <div  class="magic_square_canvas_container grow"
        id="magic_square_canvas_container">
    <canvas id="magic_square"
            class="magic_square_canvas"
            width={magicSquareWidth} 
            height={magicSquareHeight}
            on:mousemove={handleMouseMove}/>
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
