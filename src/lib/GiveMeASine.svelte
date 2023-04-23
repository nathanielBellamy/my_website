<script lang='ts'>
  import { onMount } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"

  let title = "Give Me a Sine"


  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

  onMount(async () => {
    await timeout(50) // await wasm init
    new rust.GmasWasm
  })
</script>

<div class="give_me_a_sine">
  <h3>
    {title}
  </h3>
  <div id="give_me_a_sine_output"
       class="give_me_a_sine_output"/>
  <div id="give_me_a_sine_form"
       class="give_me_a_sine_form">
    <div class="give_me_a_sine_form_cell font-bold">
      f(x) = a * sin(b*x + c)
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_a">
        a
      </label>
      <input id="gmas_form_input_a"
             type="number"
             step="0.01"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_b">
        b
      </label>
      <input id="gmas_form_input_b"
             type="number"
             step="0.01"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_c">
        c
      </label>
      <input id="gmas_form_input_c"
             type="number"
             step="0.01"/>
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .give_me_a_sine
    &_output
      display: flex
      flex-direction: column
      align-items: left
      justify-content: flex-start
      overflow-x: scroll

    &_form
      display: flex
      justify-content: space-around
  
      &_cell
        display: flex
        flex-direction: column
        justify-content: space-around
        align-items: stretch
</style>
