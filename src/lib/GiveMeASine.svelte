<script lang='ts'>
  import { onMount } from 'svelte'

  onMount(async () => {
    await wasm_bindgen() // loaded in index.html from pkg/src_rust.js
    const { GmasWasm, init_message } = wasm_bindgen
    console.log(init_message("Wasm Running for GMAS"))
    GmasWasm.run()
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

<body>
  <div class="give_me_a_sine flex flex-col md:flex-row md:justify-stretch">
    <div id="give_me_a_sine_form"
         class="give_me_a_sine_form flex flex-col">
      <div class="give_me_a_sine_form_header font-bold">
        <p>Give Me A Sine</p>
        <p>f(x) = a * sin(b*x + c)</p>
      </div>
      <div class="give_me_a_sine_form_body flex flex-col justify-between items-stretch">
        {#each range_inputs as {id, label, min, max, step}}
          <div class="give_me_a_sine_form_cell">
            <label  class="give_me_a_sine_form_cell_label"
                    for={id}>
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
         class="give_me_a_sine_output device_graph_font"/>
  </div>
</body>

<style lang="sass">
  @use "./../styles/color"

  @media (max-width : 700px) 
      .device_graph_font
        font-size: 3px
    
  .give_me_a_sine
    flex-grow: 1
    overflow-y: scroll
    margin: 5px

    &_output
      display: flex
      flex-direction: column
      align-items: left
      justify-content: flex-start
      overflow-x: scroll
      border-top: 10px double color.$blue-7
      border-bottom: 10px double color.$blue-7
      border-radius: 5px
      maargin: 10px
      height: 100%
      flex-grow: 1

    &_form
      border-radius: 5px
      margin: 0 5px 0 5px
      padding: 0 0 10px 0
      height: 100%
      background: color.$black-blue-grad
      min-height: 500px
      overflow-y: scroll
      min-width: 200px

      &_header
        color: color.$blue-7
      &_body
        padding: 0 30px 0 30px
        flex-grow: 1
  
      &_cell
        display: flex
        flex-direction: column
        justify-content: flex-start
        border-radius: 5px

        &_label
          width: 100%
          text-align: left
          padding-left: 15px
</style>

