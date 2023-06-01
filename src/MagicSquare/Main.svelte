<script lang="ts" type="module">
  import { onMount, onDestroy } from 'svelte'
  import Loading from '../lib/Loading.svelte'
  import DrawPattern from './ControlModules/DrawPattern.svelte'
  import Color from './ControlModules/Color.svelte'
  import ControlRack from './ControlRack.svelte'
  import MouseTracking from './ControlModules/MouseTracking.svelte'
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
  
  enum HiddenInputId {
    drawPattern = "magic_square_input_draw_pattern",
    color1 = "magic_square_input_color_1",
    color2 = "magic_square_input_color_2",
    color3 = "magic_square_input_color_3",
    color4 = "magic_square_input_color_4",
    color5 = "magic_square_input_color_5",
    color6 = "magic_square_input_color_6",
    color7 = "magic_square_input_color_7",
    color8 = "magic_square_input_color_8",
    mouseTracking = "magic_square_input_mouse_tracking",
    radiusMin="magic_square_input_radius_min",
    radiusStep="magic_square_input_radius_step"
  }

  export let sideLength: number = 0.0

  // DRAW PATTERN
  let currDrawPattern: string
  enum DrawPatternDirection {
    Fix = "Fix",
    In = "In",
    Out = "Out"
  }
  let initialDrawPatternDirection: DrawPatternDirection = DrawPatternDirection.Fix
  let initialDrawPatternCount: number
  function setInitialDrawPatternVars(initialUiBuffer: any) {
    const pattern: string = initialUiBuffer.settings.draw_pattern
    setInitialDrawPatternCount(parseInt(pattern.slice(-1)[0]))
    let first_letter = pattern[0]
    switch (first_letter) {
      case 'F':
        setInitialDrawPatternDirection(DrawPatternDirection.Fix)
        break
      case 'I': 
        setInitialDrawPatternDirection(DrawPatternDirection.In)
        break
      case 'O':
        setInitialDrawPatternDirection(DrawPatternDirection.Out)
        break
    }
  }
  function setInitialDrawPatternDirection(direction: DrawPatternDirection) {
    initialDrawPatternDirection = direction
  }
  function setInitialDrawPatternCount(count: number) {
    initialDrawPatternCount = count
  }

  function set(x: any, y:any) {
    x = y
  }

  // COLOR
  let renderColor: boolean = false
  let color1: number[]
  let color2: number[]
  let color3: number[]
  let color4: number[]
  let color5: number[]
  let color6: number[]
  let color7: number[]
  let color8: number[]

  function convertRgbaValue(val: number, idx: number): number {
    if (idx < 3) { // do for r, g, b, but not a
      val = val * 255
    }
    return val
  }

  function setInitialColorVars(initialUiBuffer: any) {
    color1 = [...initialUiBuffer.settings.color_1].map((x,idx) => convertRgbaValue(x, idx))
    color2 = [...initialUiBuffer.settings.color_2].map((x,idx) => convertRgbaValue(x, idx))
    color3 = [...initialUiBuffer.settings.color_3].map((x,idx) => convertRgbaValue(x, idx))
    color4 = [...initialUiBuffer.settings.color_4].map((x,idx) => convertRgbaValue(x, idx))
    color5 = [...initialUiBuffer.settings.color_5].map((x,idx) => convertRgbaValue(x, idx))
    color6 = [...initialUiBuffer.settings.color_6].map((x,idx) => convertRgbaValue(x, idx))
    color7 = [...initialUiBuffer.settings.color_7].map((x,idx) => convertRgbaValue(x, idx))
    color8 = [...initialUiBuffer.settings.color_8].map((x,idx) => convertRgbaValue(x, idx))
  }

  // MOUSE TRACKING
  enum MouseTrackingOption {
    on = 'On',
    off = 'Off',
    invX = 'Inv X',
    invY = 'Inv Y',
    invXY = 'Inv XY'
  }
  let currMouseTrackingOption: MouseTrackingOption

  function setInitialMouseTrackingOption(initialUiBuffer: any) {
    currMouseTrackingOption = initialUiBuffer.settings.mouse_tracking
  }

  let renderDataReady = false
  onMount(async () => {
    // clear old ui_buffer from localStorage
    localStorage.clear()
    
    // load wasm
    await wasm_bindgen() // loaded in index.html from ./pkg/src_rust.js
    const { MagicSquare, init_message } = wasm_bindgen
    console.log(
      init_message("Magic Square Wasm!")
    )
    
    // init wasm process and set initial values
    const initialUiBuffer = await MagicSquare.run()
    setInitialDrawPatternVars(initialUiBuffer)
    setInitialColorVars(initialUiBuffer)
    setInitialMouseTrackingOption(initialUiBuffer)
    renderDataReady = true
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
      <div slot="color"
           class="h-full">
        {#if renderDataReady}
          <Color bind:color1={color1}
                 bind:color2={color2}
                 bind:color3={color3}
                 bind:color4={color4}
                 bind:color5={color5}
                 bind:color6={color6}
                 bind:color7={color7}
                 bind:color8={color8}/>
        {:else}
          <Loading />
        {/if}
        <input id={HiddenInputId.color1}
               bind:value={color1}
               class="hidden_input">
        <input id={HiddenInputId.color2}
               bind:value={color2}
               class="hidden_input">
        <input id={HiddenInputId.color3}
               bind:value={color3}
               class="hidden_input">
        <input id={HiddenInputId.color4}
               bind:value={color4}
               class="hidden_input">
        <input id={HiddenInputId.color5}
               bind:value={color5}
               class="hidden_input">
        <input id={HiddenInputId.color6}
               bind:value={color6}
               class="hidden_input">
        <input id={HiddenInputId.color7}
               bind:value={color7}
               class="hidden_input">
        <input id={HiddenInputId.color8}
               bind:value={color8}
               class="hidden_input">
      </div>
      <div slot="drawPattern"
           class="h-full">
        {#if renderDataReady}
          <DrawPattern bind:currDrawPatternDirection={initialDrawPatternDirection}
                       bind:currDrawPatternCount={initialDrawPatternCount}/>
        {:else}
          <Loading />
        {/if}
        <input id={HiddenInputId.drawPattern}
               bind:value={currDrawPattern}
               class="hidden_input"/>
      </div>
      <div slot="mouseTracking"
           class="h-full">
        {#if renderDataReady}
          <MouseTracking currOption={currMouseTrackingOption}/>
        {:else}
          <Loading />
        {/if}
        <input id={HiddenInputId.mouseTracking}
               bind:value={currMouseTrackingOption}
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

  .hidden_input
    display: none
</style>
