<script lang="ts" type="module">
  import { onMount } from 'svelte'
  import { watchResize } from "svelte-watch-resize"
  import * as rust from "../../src-rust/pkg/src_rust.js"

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))
  let height: number = 0
  let width: number = 0

  let worker: any = null

  $: height
  $: width

  function handleResize(node: HTMLCanvasElement) {
    height = node.offsetHeight
    width = node.offsetWidth
  }

  onMount(async () => {
    let element = document.getElementById("magic_square")
    height = element.offsetHeight
    width = element.offsetWidth

    await timeout(150) // await wasm init
    // rust.MagicSquare.run()

    if ('serviceWorker' in navigator) {
      navigator.serviceWorker.register('../src/magicSquareWasm.js')
        .then((reg) =>  {
          worker = reg
          rust.MagicSquare.run()
        })
    }
    // const myWorker = new Worker("../src/magicSquareWasm.js", { type: 'module' })
  })
</script>

<div id="magic_square"
     class="magic_square flex flex-col justify-start">
  <canvas id="magic_square_canvas"
          class="magic_square_canvas"
          use:watchResize={handleResize}
          height={height}
          width={width}/>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .magic_square
    height: 100%
    width: 100%
    &_canvas
      height: 100%
      width: 100%
      border: 5px solid color.$yellow-4

</style>
