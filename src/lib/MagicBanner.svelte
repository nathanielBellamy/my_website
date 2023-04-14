<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

  let magicBanner: any
  let captureInterval: number

  onMount(async () => {
    await timeout(100) // await wasm init
    magicBanner = new rust.MagicBanner
    magicBanner.run()
  })

    
  onDestroy(async () => {
    cancelAnimationFrame(captureInterval)
  })
</script>

<div class="magic_banner flex flex-col justify-start">
  <canvas id="magic_banner_canvas"
          class="magic_banner_canvas"/>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .magic_banner
    &_canvas
      width: 100%
      height: 100%

</style>
