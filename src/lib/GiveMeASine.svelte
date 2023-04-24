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
             type="range"
             min="-3"
             max="3"
             step="0.1"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_b">
        b
      </label>
      <input id="gmas_form_input_b"
             type="range"
             min="-3"
             max="3"
             step="0.1"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_c">
        c
      </label>
      <input id="gmas_form_input_c"
             type="range"
             min="-3"
             max="3"
             step="0.1"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_ep">
        ep
      </label>
      <input id="gmas_form_input_ep"
             type="range"
             min="-1"
             max="1"
             step="0.1"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_x_min">
        x_min
      </label>
      <input id="gmas_form_input_x_min"
             type="range"
             min="-3.14"
             max="3.14"
             step="0.1"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_height">
        height
      </label>
      <input id="gmas_form_input_height"
             type="range"
             min="10"
             max="50"
             step="1"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_width">
        width
      </label>
      <input id="gmas_form_input_width"
             type="range"
             min="3"
             max="255"
             step="1"/>
    </div>
    <div class="give_me_a_sine_form_cell">
      <label for="gmas_form_input_c">
        c
      </label>
      <input id="gmas_form_input_c"
             type="range"
             min="-3"
             max="3"
             step="0.5"/>
    </div>
  </div>
  <div id="give_me_a_sine_output"
       class="give_me_a_sine_output font-xs"/>
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
      flex-wrap: wrap
  
      &_cell
        display: flex
        flex-direction: column
        justify-content: space-around
        align-items: stretch
        margin: 15px
        padding: 5px
        border: 5px solid color.$blue-4
        border-radius: 5px
</style>

