<script lang="ts" type="module">
  import { onMount, onDestroy } from 'svelte'
    import DrawPattern from './ControlModules/DrawPattern.svelte';
  import ControlRack from './ControlRack.svelte'
  // this component will be large
  // the decision was made to optimize for minimal plumbing
  // this component instantiates the wasm module and retrieves the initial UI values from it
  // these values do not go further than this file (in JS)
  // the mantra is
  //   -> Svelte/JS is for layout + display logic
  //   -> Rust/Wasm is for handling data

  // update: we have reworked this pattern
  // -> ui-related values are stored in JS and bound to pass to ui-logic components
  // -> but input values are bound here
  // -> ui-logic component can retrieve and edit input value
  // -> ui logic components are forms
  //    -> on submit they set the value in the hidden input and trigger an input event
  //    -> the input event bubbles up triggering an input event on the ui_buffer_form
  //    -> wasm is listening to input events on the ui_buffer_form
  //    -> wasm receives the input event and updates the buffer with the desired value

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
  
  // DRAW PATTERN
  let currDrawPattern: string
  enum DrawPatternDirection {
    Fix = "Fix",
    In = "In",
    Out = "Out"
  }
  let currDrawPatternDirection: DrawPatternDirection = DrawPatternDirection.Fix
  let currDrawPatternCount: number
  let initialDrawPattern: string
  const drawPatternHiddenInputId: string = 'magic_square_input_draw_pattern'
  function setInitialDrawPatternVars(pattern: string) {
    console.log("wooooo")
    console.log(pattern)
    setCurrDrawPatternCount(parseInt(pattern.slice(-1)[0]))
    let first_letter = pattern[0]
    switch (first_letter) {
      case 'F':
        setCurrDrawPatternDirection(DrawPatternDirection.Fix)
        break
      case 'I': 
        setCurrDrawPatternDirection(DrawPatternDirection.In)
        break
      case 'O':
        setCurrDrawPatternDirection(DrawPatternDirection.Out)
        break
    }
  }
  function setCurrDrawPatternDirection(direction: DrawPatternDirection) {
    currDrawPatternDirection = direction
  }
  function setCurrDrawPatternCount(count: number) {
    currDrawPatternCount = count
  }

  // COLOR
  let color1: number[]
  let color2: number[]
  let color3: number[]
  let color4: number[]
  let color5: number[]
  let color6: number[]
  let color7: number[]
  let color8: number[]

  function setInitialColors(initialUiBuffer: any) {
    color1 = initialUiBuffer.settings.color_1
    color2 = initialUiBuffer.settings.color_2
    color3 = initialUiBuffer.settings.color_3
    color4 = initialUiBuffer.settings.color_4
    color5 = initialUiBuffer.settings.color_5
    color6 = initialUiBuffer.settings.color_6
    color7 = initialUiBuffer.settings.color_7
    color8 = initialUiBuffer.settings.color_8

    console.dir({
      color1,
      color2,
      color3,
      color4,
      color5,
      color6,
      color7,
      color8
    })
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
      // console.dir(initialUiBuffer)
      setInitialDrawPatternVars(initialUiBuffer.settings.draw_pattern)
      setInitialColors(initialUiBuffer)
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
           class="h-full">
        <DrawPattern bind:currDrawPatternDirection={currDrawPatternDirection}
                     bind:currDrawPatternCount={currDrawPatternCount}/>
        <input id={drawPatternHiddenInputId}
               bind:value={currDrawPattern}
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

  .hidden_input
    display: none
</style>
