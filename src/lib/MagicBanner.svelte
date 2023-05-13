<script lang="ts">
  import { onMount } from 'svelte'
  // import * as rust from "../../src-rust/pkg/src_rust.js"

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

  onMount(async () => {
    await timeout(200) // await wasm init
    // rust.MagicBanner.run()
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
      height: 100px

</style>
