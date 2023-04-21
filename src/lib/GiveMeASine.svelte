<script lang='ts'>
  import { onMount } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"

  let title = "Give Me a Sine"

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

  onMount(async () => {
    await timeout(200) // await wasm init
    new rust.GmasWasm
  })
</script>

<div class="give_me_a_sine">
  <h3>
    {title}
  </h3>
  <div id="give_me_a_sine_output"
       class="give_me_a_sine_output"/>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .give_me_a_sine
    &_output
      display: flex
      flex-direction: column
      align-items: left
      justify-content: flex-start

  .gmas_graph_row
    text-align: left

</style>
