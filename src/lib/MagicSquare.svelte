<script lang="ts" type="module">
  import { onMount } from 'svelte'

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))
  let height: number = 0
  let width: number = 0

  $: height
  $: width

  onMount(async () => {
    let element = document.getElementById("magic_square")
    height = element.offsetHeight
    width = element.offsetWidth

    await timeout(150) // await wasm init

    const { MagicSquare, init_message } = wasm_bindgen
    wasm_bindgen('./src-rust/pkg/src_rust_bg.wasm').then(() => {
      console.log(init_message("Wasm Running for Magic Square"))
      MagicSquare.run()
    }).catch(console.error)
  })
</script>

<div id="magic_square"
     class="magic_square flex flex-col justify-start">
  <canvas id="magic_square_canvas"
          class="magic_square_canvas"
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
