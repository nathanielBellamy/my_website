<script lang="ts" type="module">
  import { onMount, onDestroy } from 'svelte'
  import ControlRack from './ControlRack.svelte'
  // this component will be large
  // the decision was made to optimize for minimal plumbing
  // this component instantiates the wasm module and retrieves the initial UI values from it
  // these values do not go further than this file (in JS)
  // the mantra is
  //   -> Svelte/JS is for layout/display
  //   -> Rust/Wasm is for handling data

  // This pattern has an interesting, happy consequence
  // It was designed to model a modular synthesizer system
  // But the ui input data flow is
  // -> wasm returns initial values
  // -> JS binds these values to their HTML elements here, wrapping them in CSS delivered through Svelte slots
  // -> wasm instantiates event listeners for these inputs and records changes to its internal ui_buffer
  // -> as a result, the actual HTML becomes "the source of truth" for the input value (hooray for avoidding JS State!)
  // -> this models a physical - knob per function - system


  // TODO:
  //  -> On load
  //    -> Js passes the MagicSquareInstanceId to wasm
  //    -> wasm checks localStorage for the key MagicSquareInstanceId
  //      -> if exists it retrieves + deserializes ui_settings from localStorage
  //      -> if none exist, it calls ::new()
  //  -> onDestroy
  //    -> js writes the current ui_settings to localStorage
  //  -> onResize
  //    -> Container.svelte manages MagicSquareInstanceId 
  //    -> this should persist ui_settings
  //    -> while destroying and loading new wasm module instances
  
  export let sideLength: number = 0.0

  let initialUiBuffer: any
  
  // drawPattern vars
  let currDrawPattern: string
  const drawPatternFormId: string = 'draw_pattern_form'
  const drawPatternHiddenInputId: string = 'magic_square_input_draw_pattern'
  const drawPatterns: string[] = [
    'All',
    'One',
    'Two',
    'Three',
    'Four',
    'Five',
    'Six',
    'Seven',
    'Eight',
    'Out1',
    'Out2',
    'Out3',
    'Out4',
    'Out5',
    'Out6',
    'Out7',
    'Out8',
    'In1',
    'In2',
    'In3',
    'In4',
    'In5',
    'In6',
    'In7',
    'In8',
    'Conv',
    'Div',
    'Random'
  ]
  function setCurrDrawPattern(pattern: string) {
    currDrawPattern = pattern
  }
  function handleDrawPatternClick(pattern: string) {
    setCurrDrawPattern(pattern)
    let form = document.getElementById(drawPatternFormId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  }
  function handleDrawPatternKeydown(e: any, pattern: string) {
    if (e.keyCode === 13){
      setCurrDrawPattern(pattern)
      let form = document.getElementById(drawPatternFormId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }

  onMount(async () => {
    // clear old ui_buffer from localStorage
    localStorage.clear()
    
    // load wasm
    await wasm_bindgen() // loaded in index.html from ./pkg/src_rust.js
    const { MagicSquare, init_message } = wasm_bindgen
    console.log(
      init_message("Magic Square Wasm!")
    )
    
    // set initial values
    initialUiBuffer =  MagicSquare.run().then((initialUiBuffer: any) => {
      console.dir(initialUiBuffer)
      currDrawPattern = initialUiBuffer.settings.draw_pattern
    })

    // set up event listeners for the differnt modules

    // set up drawPattern form listener
    var drawPatternForm = document.getElementById(drawPatternFormId)
    drawPatternForm.addEventListener('submit', () => {
      var input = document.getElementById(drawPatternHiddenInputId)
      input.value = currDrawPattern
      input.dispatchEvent(new Event('input', {bubbles: true}))
    })
  })

  onDestroy(async () => {
    let app = document.getElementById(("app_main"))
    app.dispatchEvent(new Event("destroymswasm", {bubbles: true}))
  })
</script>

<div id="magic_square"
     class="magic_square flex flex-wrap gap-2">
  <div id="magic_square_canvas_container"
       class="magic_square_canvas_container flex flex-col justify-around display">
    <canvas id="magic_square_canvas"
            class="magic_square_canvas"
            height={sideLength}
            width={sideLength}/>
  </div>
  <div class="control">
    <ControlRack>
      <div slot="drawPattern"
           id={drawPatternFormId}>
        <div class="draw_pattern_options flex flex-col">
          {#each drawPatterns as pattern}
            <button class="draw_pattern_option"
                    class:selected="{currDrawPattern === pattern}"
                    on:click={() => handleDrawPatternClick(pattern)}
                    on:keydown={(e) => handleDrawPatternKeydown(e, pattern)}>
                {pattern.toUpperCase()}
            </button>
          {/each}
        </div>
        <input id={drawPatternHiddenInputId}
               class="hidden_input"/>
      </div>
    </ControlRack>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  
  .magic_square
    height: 100%
    width: 100%
    overflow-y: scroll

    &_canvas
      border-top: 5px double color.$blue-7
      border-bottom: 5px double color.$blue-7
      border-radius: 10px
      &_container
        height: 100%
        background: color.$black-blue-horiz-grad
        border: 10px double color.$blue-7
        border-radius: 5px
        flex-grow: 1

  .display
    align-items: center
  
  .control
    flex-grow: 1
    height: 100%

  .selected
    background-color: color.$blue-8
</style>
