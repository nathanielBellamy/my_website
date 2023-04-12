<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"

  let x: number = -1 
  let y: number = -1

  const handleMouseMove = (e: any) => {
    x = e.clientX
    y = e.clientY
  }

  const handleMouseLeave = () => {
    x = -1 // negative indicates off element
    y = -1
  }

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

  let magicBanner: any
  let captureInterval: number

  onMount(async () => {
    await timeout(500) // TODO: make this better than a hack
                       // must wait for wasm to load
    magicBanner = new rust.MagicBanner

    // represents the loop used to capture mouse movement and deliver coordinates to WASM
    // this loop writes to the MagicSquareBuffer
    // an animation loop within a closure within WASM reads from the buffer
    const captureLoop = () => {
      magicBanner.write_to_buffer(x, y)
      console.log('foo')
      captureInterval = requestAnimationFrame(captureLoop)
    };
    
    captureInterval = requestAnimationFrame(captureLoop)
  })

    
  onDestroy(async () => {
    cancelAnimationFrame(captureInterval)
  })
</script>

<div class="magic_banner_container rounded-md flex flex-col justify-start">
  <div  class="magic_banner_canvas_container"
        id="magic_banner_canvas_container">
    <canvas id="magic_banner"
            class="magic_banner_canvas"
            on:mousemove={handleMouseMove}
            on:mouseleave={handleMouseLeave}/>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .magic_banner
    background-color: color.$black-4

    &_container
      border: 2px solid color.$black-3
      width: 100%

    &_canvas
      height: 100%
      width: 100%
    
      &_container
        background: color.$black-grad

</style>
