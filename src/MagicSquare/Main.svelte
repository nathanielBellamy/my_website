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
  enum DrawPatternDirection {
    Fix = "Fix",
    In = "In",
    Out = "Out"
  }
  interface Drawpattern {
    dir: DrawPatternDirection,
    count: number // 0 < x < 9
  }
  let currDrawPatternDirection: DrawPatternDirection = DrawPatternDirection.Fix
  let currDrawPatternCount: number
  const drawPatternFormId: string = 'draw_pattern_form'
  const drawPatternHiddenInputId: string = 'magic_square_input_draw_pattern'
  // const drawPatterns: string[] = [
  //   'Fix1',
  //   'Fix2',
  //   'Fix3',
  //   'Fix4',
  //   'Fix5',
  //   'Fix6',
  //   'Fix7',
  //   'Fix8',
  //   'Out1',
  //   'Out2',
  //   'Out3',
  //   'Out4',
  //   'Out5',
  //   'Out6',
  //   'Out7',
  //   'Out8',
  //   'In1',
  //   'In2',
  //   'In3',
  //   'In4',
  //   'In5',
  //   'In6',
  //   'In7',
  //   'In8',
  // ]
  function setInitialDrawPatternVars(pattern: string) {
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
  function deriveCurrDrawPattern(): string {
    var result: string
    switch (currDrawPatternDirection) {
      case DrawPatternDirection.Fix:
        result = DrawPatternDirection.Fix
        break
      case DrawPatternDirection.In:
        result = DrawPatternDirection.In
        break
      case DrawPatternDirection.Out:
        result = DrawPatternDirection.Out
        break
    }

    return `${result}${currDrawPatternCount}`
  }
  function setCurrDrawPattern() {
    currDrawPattern = deriveCurrDrawPattern()
  }
  function setCurrDrawPatternDirection(direction: DrawPatternDirection) {
    currDrawPatternDirection = direction
  }
  function handleDrawPatternDirectionClick(direction: DrawPatternDirection) {
    setCurrDrawPatternDirection(direction)
    setCurrDrawPattern()
    let form = document.getElementById(drawPatternFormId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  }
  function handleDrawPatternDirectionKeydown(e: any, direction: DrawPatternDirection) {
    if (e.keyCode === 13){
      currDrawPatternDirection = direction
      setCurrDrawPattern()
      let form = document.getElementById(drawPatternFormId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }
  function setCurrDrawPatternCount(count: number) {
    currDrawPatternCount = count
  }
  function handleDrawPatternCountClick(count: number) {
    setCurrDrawPatternCount(count)
    let form = document.getElementById(drawPatternFormId)
    form.dispatchEvent(new Event('submit', {bubbles: true}))
  }
  function handleDrawPatternCountKeydown(e: any, count: number) {
    if (e.keyCode === 13){
      setCurrDrawPatternCount(count)
      let form = document.getElementById(drawPatternFormId)
      form.dispatchEvent(new Event('submit', {bubbles: true}))
    }
  }
  function handleDrawPatternFormSubmit() {
    var input = document.getElementById(drawPatternHiddenInputId)
    input.value = deriveCurrDrawPattern()
    input.dispatchEvent(new Event('input', {bubbles: true}))
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
    })

    // set up event listeners for the differnt modules

    // set up drawPattern form listener
    var drawPatternForm = document.getElementById(drawPatternFormId)
    drawPatternForm.addEventListener('submit', handleDrawPatternFormSubmit)
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
           id="draw_pattern_form"
           class="flex flex-col justify-around items-stretch h-full w-full">
        <div id="draw_pattern_buttons"
             class="h-full flex flex-col justify-around">
          <div id="draw_pattern_directions_outer"
               class="grow flex flex-col justify-around items-streth">
            <div id="draw_pattern_directions_inner"
                 class="grow max-h-20 flex justify-around items-stretch">
              {#each Object.values(DrawPatternDirection) as dir}
                <button class="grow max-h-26 pr-3 pl-3"
                        on:click={() => handleDrawPatternDirectionClick(dir)}
                        on:keydown={(e) => handleDrawPatternDirectionKeydown(e, dir)}
                        class:selected={currDrawPatternDirection === dir}>
                  {dir}
                </button>
              {/each}
            </div>
          </div>
          <div id="draw_pattern_counts"
               class="grow flex flex-col justify-around items-stretch">
            {#each [0,4] as countShifter}
              <div id="draw_pattern_counts_row"
                   class="grow flex justify-evenly items-stretch gap-0">
                {#each [1,2,3,4].map(x => x + countShifter) as count}
                  <button class="grow max-h-20"
                          on:click={() => handleDrawPatternCountClick(count)}
                          on:keydown={(e) => handleDrawPatternCountKeydown(e, count)}
                          class:selected={currDrawPatternCount === count}>
                    {count}
                  </button>
                {/each}
              </div>
            {/each}
          </div>
        </div>
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
