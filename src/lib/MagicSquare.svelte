<script lang="ts">
  import { onMount } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

  onMount(async () => {
    await timeout(200) // await wasm init
    rust.MagicSquare.run()
  })
</script>

<div class="magic_square flex flex-col justify-start">
  <canvas id="magic_square_canvas"
          class="magic_square_canvas"/>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .magic_square
    &_canvas
      width: 100%
      height: calc(100vh - 200px)
      border: 5px solid color.$yellow-4

</style>
