<script lang='ts'>
  import { onMount } from 'svelte'
  import * as rust from "../../src-rust/pkg/src_rust.js"

  let title = "Give Me a Sine"

  const timeout = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

  onMount(async () => {
    await timeout(50) // await wasm init
    new rust.GmasWasm
  })

  interface GmasRangeInput {
    id: string,
    label: string,
    min: number,
    max: number,
    step: number
  }

  const range_inputs: GmasRangeInput[] = [
    {
      id: "gmas_form_input_a",
      label: "a",
      min: -3,
      max: 3,
      step: 0.1
    },
    {
      id: "gmas_form_input_b",
      label: "b",
      min: -12,
      max: 12,
      step: 0.1
    },
    {
      id: "gmas_form_input_c",
      label: "c",
      min: -3,
      max: 3,
      step: 0.1
    },
    {
      id: "gmas_form_input_ep",
      label: "ep",
      min: 0.01,
      max: 1,
      step: 0.01
    },
    {
      id: "gmas_form_input_height",
      label: "height",
      min: 5,
      max: 50,
      step: 1
    },
    {
      id: "gmas_form_input_width",
      label: "width",
      min: 10,
      max: 255,
      step: 1
    },
    {
      id: "gmas_form_input_graph_char",
      label: "graph color",
      min: 0,
      max: 8,
      step: 1
    },
    {
      id: "gmas_form_input_above_char",
      label: "above color",
      min: 0,
      max: 8,
      step: 1
    },
    {
      id: "gmas_form_input_below_char",
      label: "below color",
      min: 0,
      max: 8,
      step: 1
    }
  ]
</script>

<div class="give_me_a_sine grid grid-rows-10 md:grid-cols-4">
  <div class="give_me_a_sine_form_container row-span-4 device_form">
    <div id="give_me_a_sine_form"
         class="give_me_a_sine_form grid grid-cols-2">
      <div class="give_me_a_sine_form_cell font-bold title">
        Give Me A Sine
      </div>
      <div class="give_me_a_sine_form_cell font-bold sub_title">
        f(x) = a * sin(b*x + c)
      </div>
      {#each range_inputs as {id, label, min, max, step}}
        <div class="give_me_a_sine_form_cell">
          <label for={id}>
            {label}
          </label>
          <input id={id}
                 type="range"
                 min={min}
                 max={max}
                 step={step}/>
        </div>
      {/each}
    </div>
  </div>
  <div id="give_me_a_sine_output"
       class="give_me_a_sine_output device_graph_font row-span-8 md:col-span-3"/>
</div>

<style lang="sass">
  @use "./../styles/color"

  @media (max-width : 700px) 
      .device_graph_font
        font-size: 3px
    
      .device_form
        max-height: 100px

  .gmas_graph_row
    margin: 0px

  .title
    color: color.$blue-2
    font-weight: 900
    display: flex
    flex-direction: column
    justify-content: space-around
    align-items: stretch
    
  .sub_title
    color: color.$blue-2
    font-weight: 700
  
  .give_me_a_sine
    &_output
      display: flex
      flex-direction: column
      align-items: left
      justify-content: flex-start
      overflow-x: scroll
      border: 5px solid color.$blue-2
      border-radius: 5px

    &_form
      border: 5px solid color.$blue-4
      border-radius: 5px
      margin: 0 5px 0 5px

      &_container
        display: flex
        flex-direction: column
        justify-content: flex-start
        align-items: stretch
  
      &_cell
        display: flex
        flex-direction: column
        justify-content: flex-start
        margin: 5px
        padding: 5px
        border: 5px solid color.$blue-4
        border-radius: 5px
        max-height: 100px
</style>

