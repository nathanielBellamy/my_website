<script lang='ts'>
  import { afterUpdate, onMount } from 'svelte'

  let a: number
  let b: number
  let c: number
  let ep: number
  let height: number
  let width: number
  let aboveCharId: number
  let aboveChar: string
  $: aboveChar = colorSquareFromId(aboveCharId)
  // let aboveChar: string
  let belowCharId: number
  let belowChar: string
  $: belowChar = colorSquareFromId(belowCharId)
  let graphCharId: number
  let graphChar: string
  $: graphChar = colorSquareFromId(graphCharId)

  onMount(async () => {
    await wasm_bindgen() // loaded in index.html from pkg/src_rust.js
    const { GmasWasm, init_message } = wasm_bindgen
    console.log(init_message("Wasm Running for GMAS"))

    const initialData = GmasWasm.run()
    a = initialData.function[0]
    b = initialData.function[1]
    c = initialData.function[2]
    ep = initialData.settings.ep
    height = initialData.settings.height
    width = initialData.settings.width
    aboveCharId = initialData.settings.above_char_id
    belowCharId = initialData.settings.below_char_id
    graphCharId = initialData.settings.graph_char_id
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

  function colorSquareFromId (id: number): string {
    switch (id) {
      case 0: 
        return '⬛'
      case 1:
        return '⬜'
      case 2:
        return '🟪'
      case 3:
        return '🟦'
      case 4:
        return '🟩'
      case 5: 
        return '🟨'
      case 6:
        return '🟧'
      case 7:
        return '🟥'
      case 8: 
       return '🟫'
      default: 
        return ''
    }
  }

  function idFromColorSquare(colorSquare: string): number {
    switch (colorSquare) {
      case '⬛': 
        return 0
      case '⬜':
        return 1
      case '🟪':
        return 2
      case '🟦':
        return 3
      case '🟩':
        return 4
      case '🟨':
        return 5
      case '🟧':
        return 6
      case '🟥':
        return 7
      case '🟫':
        return 8
      default: 
        return -1
    }
  }

  // the decision below to keep all inputs as raw html
  // as opposed to excising into a Range.svelte component
  // was the decision to avoid JS value bindings as much as possible
  // 
  // for similar "avoid JS data handling" reasons
  // we do not "dry out" our code by storing data in an object/array
  // and then using some form of an {#each} block
  // doing so would require some getting/setting of the 
  // component values we are binding to the inputs
  // explicit layout means bound values appear directly in <script> and <body>
  //
  // from anecdotal testing, this style cooperates better with Svelte's compilation process
  // particularly when hot-updating in dev
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
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_a'}>
            <p>a</p> 
            <p>{a}</p>
          </label>
          <input id={'gmas_form_input_a'}
                 type="range"
                 min={-6.28}
                 max={6.28}
                 bind:value={a}
                 step={0.1}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_b'}>
            <p>b</p> 
            <p>{b}</p>
          </label>
          <input id={'gmas_form_input_b'}
                 type="range"
                 min={-6.28}
                 max={6.28}
                 bind:value={b}
                 step={0.1}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_c'}>
            <p>c</p> 
            <p>{c}</p>
          </label>
          <input id={'gmas_form_input_c'}
                 type="range"
                 min={-6.28}
                 max={6.28}
                 bind:value={c}
                 step={0.01}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_ep'}>
            <p>ep</p> 
            <p>{ep}</p>
          </label>
          <input id={'gmas_form_input_ep'}
                 type="range"
                 min={0}
                 max={1.26}
                 bind:value={ep}
                 step={0.01}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_height'}>
            <p>Height</p> 
            <p>{height}</p>
          </label>
          <input id={'gmas_form_input_height'}
                 type="range"
                 min={1}
                 max={50}
                 bind:value={height}
                 step={1}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_width'}>
            <p>Width</p> 
            <p>{width}</p>
          </label>
          <input id={'gmas_form_input_width'}
                 type="range"
                 min={1}
                 max={50}
                 bind:value={width}
                 step={1}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_above_char'}>
            <p>Above Color</p> 
            <p>{aboveChar}</p>
          </label>
          <input id={'gmas_form_input_above_char'}
                 type="range"
                 min={0}
                 max={8}
                 bind:value={aboveCharId}
                 step={1}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_below_char'}>
            <p>Below Color</p> 
            <p>{belowChar}</p>
          </label>
          <input id={'gmas_form_input_below_char'}
                 type="range"
                 min={0}
                 max={8}
                 bind:value={belowCharId}
                 step={1}/>
        </div>
        <div class="give_me_a_sine_form_cell">
          <label  class="give_me_a_sine_form_cell_label"
                  for={'gmas_form_input_graph_char'}>
            <p>Graph Color</p> 
            <p>{graphChar}</p>
          </label>
          <input id={'gmas_form_input_graph_char'}
                 type="range"
                 min={0}
                 max={8}
                 bind:value={graphCharId}
                 step={1}/>
        </div>

      </div>
    </div>
    <div id="give_me_a_sine_output"
         class="give_me_a_sine_output device_graph_font"/>
  </div>
</body>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

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
        color: color.$blue-4
      &_body
        padding: 0 30px 0 30px
        flex-grow: 1
  
      &_cell
        display: flex
        flex-direction: column
        justify-content: flex-start
        border-radius: 5px
        color: color.$blue-4
        &_label
          width: 100%
          display: flex
          justify-content: space-between
          align-items: center
          padding-left: 15px
          font-weight: text.$fw-l
          font-size: text.$fs-m
          padding-left: 10px
          padding-right: 10px
</style>

