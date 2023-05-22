<script lang="ts" type="module">
  import { onMount } from 'svelte'
  import MagicSquareControl from './MagicSquareControl.svelte'

  let height: number = 0
  let width: number = 0

  $: height
  $: width

  onMount(async () => {
    let element = document.getElementById("magic_square")
    height = element.offsetHeight
    width = element.offsetWidth
    
    await wasm_bindgen() // loaded in index.html from ./pkg/src_rust.js
    const { MagicSquare, init_message } = wasm_bindgen
    console.log(init_message("Wasm Running for Magic Square"))
    
    MagicSquare.run()
  })
</script>

<div id="magic_square"
     class="magic_square flex flex justify-start">
  <canvas id="magic_square_canvas"
          class="magic_square_canvas"
          height={height}
          width={width}/>
  <MagicSquareControl />
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
