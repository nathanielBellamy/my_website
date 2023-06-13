<script lang="ts" type="module">
  import { afterUpdate, onMount, onDestroy } from 'svelte'
  import Loading from '../lib/Loading.svelte'
  import DrawPatternContainer from './ControlModules/DrawPattern.svelte'
  import type { DrawPattern } from './ControlModules/DrawPattern'
  import { intoDrawPattern } from './ControlModules/DrawPattern'
  import Color from './ControlModules/Color.svelte'
  import ControlRack from './ControlRack.svelte'
  import Geometry from './ControlModules/Geometry.svelte'
  import LfoContainer from './ControlModules/Lfo.svelte'
  import { Lfo } from './ControlModules/Lfo'
  import { intoLfoDestination } from './ControlModules/LfoDestination'
  import type { LfoDestination } from './ControlModules/LfoDestination'
  import { intoLfoShape } from './ControlModules/LfoShape'
  import type { LfoShape } from './ControlModules/LfoShape'
  import type { MouseTracking } from './ControlModules/MouseTracking'
  import MouseTrackingContainer from './ControlModules/MouseTracking.svelte'
  import Rotation from  './ControlModules/Rotation.svelte'
  import type { StorageSettings } from './StorageSettings'
  import Translation from './ControlModules/Translation.svelte'
  import { WasmInputId } from './WasmInputId'
  import { prevSettingsStore } from './PrevSettingsStore'
  // INIT LANG BOILER PLATE
  import { I18n, Lang } from '../I18n'
  import { lang } from '../stores/lang'

  const i18n = new I18n('magicSquare/main')
  let langVal: Lang
  lang.subscribe(val => langVal = val)

  // this component will be large
  // the decision was made to optimize for minimal plumbing
  // this component instantiates the wasm module and retrieves the initial UI values from it
  // the mantra is
  //   -> Svelte/JS is for layout + display logic
  //   -> Rust/Wasm is for handling data

  // update: we have reworked this pattern
  // -> ui-related values are stored in JS and bound to pass to ui-logic components
  // -> but input values are bound here
  // -> ui-logic component can retrieve and edit input value
  // -> some ui logic components are forms
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

  // DRAW PATTERN
  let drawPattern: DrawPattern
  enum DrawPatternDirection {
    Fix = "Fix",
    In = "In",
    Out = "Out"
  }
  let initialDrawPatternDirection: DrawPatternDirection = DrawPatternDirection.Fix
  let initialDrawPatternCount: number
  function setInitialDrawPatternVars(initialUiBuffer: any) {
    drawPattern = intoDrawPattern(initialUiBuffer.settings.draw_pattern)
    setInitialDrawPatternCount(parseInt(drawPattern.slice(-1)[0]))
    let first_letter = drawPattern[0]
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

  // COLOR
  let color1: number[]
  let color2: number[]
  let color3: number[]
  let color4: number[]
  let color5: number[]
  let color6: number[]
  let color7: number[]
  let color8: number[]

  // CSS (inline in Color.svelte) uses Int:0-255, WebGL uses Float:0.0-1.0
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

  // LFO
  let lfo1Active: boolean
  let lfo1Amp: number
  let lfo1Dest: LfoDestination
  let lfo1Freq: number
  let lfo1Phase: number
  let lfo1Shape: LfoShape

  let lfo2Active: boolean
  let lfo2Amp: number
  let lfo2Dest: LfoDestination
  let lfo2Freq: number
  let lfo2Phase: number
  let lfo2Shape: LfoShape

  let lfo3Active: boolean
  let lfo3Amp: number
  let lfo3Dest: LfoDestination
  let lfo3Freq: number
  let lfo3Phase: number
  let lfo3Shape: LfoShape

  let lfo4Active: boolean
  let lfo4Amp: number
  let lfo4Dest: LfoDestination
  let lfo4Freq: number
  let lfo4Phase: number
  let lfo4Shape: LfoShape
  
  function setInitialLfoVars(initialUiBuffer: any) {
    lfo1Active = initialUiBuffer.settings.lfo_1_active
    lfo1Amp = initialUiBuffer.settings.lfo_1_amp
    lfo1Dest = intoLfoDestination(initialUiBuffer.settings.lfo_1_dest)
    lfo1Freq = initialUiBuffer.settings.lfo_1_freq
    lfo1Phase = initialUiBuffer.settings.lfo_1_phase
    lfo1Shape = intoLfoShape(initialUiBuffer.settings.lfo_1_shape)

    lfo2Active = initialUiBuffer.settings.lfo_2_active
    lfo2Amp = initialUiBuffer.settings.lfo_2_amp
    lfo2Dest = intoLfoDestination(initialUiBuffer.settings.lfo_2_dest)
    lfo2Freq = initialUiBuffer.settings.lfo_2_freq
    lfo2Phase = initialUiBuffer.settings.lfo_2_phase
    lfo2Shape = intoLfoShape(initialUiBuffer.settings.lfo_2_shape)

    lfo3Active = initialUiBuffer.settings.lfo_3_active
    lfo3Amp = initialUiBuffer.settings.lfo_3_amp
    lfo3Dest = intoLfoDestination(initialUiBuffer.settings.lfo_3_dest)
    lfo3Freq = initialUiBuffer.settings.lfo_3_freq
    lfo3Phase = initialUiBuffer.settings.lfo_3_phase
    lfo3Shape = intoLfoShape(initialUiBuffer.settings.lfo_3_shape)

    lfo4Active = initialUiBuffer.settings.lfo_4_active
    lfo4Amp = initialUiBuffer.settings.lfo_4_amp
    lfo4Dest = intoLfoDestination(initialUiBuffer.settings.lfo_4_dest)
    lfo4Freq = initialUiBuffer.settings.lfo_4_freq
    lfo4Phase = initialUiBuffer.settings.lfo_4_phase
    lfo4Shape = intoLfoShape(initialUiBuffer.settings.lfo_4_shape)
  }

  function handleLfoActiveToggle(lfo: Lfo) {
    var val: boolean
    switch (lfo) {
      case Lfo.one:
        lfo1Active = !lfo1Active
        val = lfo1Active
        break
      case Lfo.two:
        lfo2Active = !lfo2Active
        val = lfo2Active
        break
      case Lfo.three:
        lfo3Active = !lfo3Active
        val = lfo3Active
        break
      case Lfo.four:
        lfo4Active = !lfo4Active
        val = lfo4Active
        break
    }
    const inputId: WasmInputId = intoLfoActiveInputId(lfo)
    var input = document.getElementById(inputId)
    if (!!input) {
      input.value = val
      input.dispatchEvent(new Event('input', {bubbles: true}))
    }
  }

  function intoLfoActiveInputId(lfo: Lfo): WasmInputId {
    switch (lfo) {
      case Lfo.two:
        return WasmInputId.lfo2Active
      case Lfo.three:
        return WasmInputId.lfo3Active
      case Lfo.four:
        return WasmInputId.lfo4Active
      case Lfo.one:
      default:
        return WasmInputId.lfo1Active
    }
  }

  // TRANSLATION
  let translationXBase: number
  let translationXSpread: number
  let translationYBase: number
  let translationYSpread: number
  // let translationZ: number

  function setInitialTranslationVars(initialUiBuffer: any) {
    translationXBase = initialUiBuffer.settings.translation_x_base
    translationXSpread = initialUiBuffer.settings.translation_x_spread
    translationYBase = initialUiBuffer.settings.translation_y_base
    translationYSpread = initialUiBuffer.settings.translation_y_spread
    // translationZ = initialUiBuffer.settings.translation_z
  }

  let mouseTracking: MouseTracking

  function setInitialMouseTracking(initialUiBuffer: any) {
    mouseTracking = initialUiBuffer.settings.mouse_tracking
  }

  // RADIUS
  let radiusBase: number
  let radiusStep: number

  function setInitialRadiusVars(initialUiBuffer: any) {
    radiusBase = Math.floor(initialUiBuffer.settings.radius_base * 100) / 100
    radiusStep = Math.floor(initialUiBuffer.settings.radius_step * 100) / 100
  }

  function round2(val: number){
    return Math.floor(val * 100) / 100
  }

  // ROTATION
  let pitchBase: number
  let pitchSpread: number
  let pitchX: number
  let pitchY: number
  let rollBase: number
  let rollSpread: number  
  let rollX: number  
  let rollY: number
  let yawBase: number
  let yawSpread: number  
  let yawX: number  
  let yawY: number

  function setInitialRotationVars(initialUiBuffer: any) {
    pitchBase = round2(initialUiBuffer.settings.y_rot_base)
    pitchSpread = round2(initialUiBuffer.settings.y_rot_spread)
    pitchX = round2(initialUiBuffer.settings.x_axis_y_rot_coeff)
    pitchY = round2(initialUiBuffer.settings.y_axis_y_rot_coeff)
    rollBase = round2(initialUiBuffer.settings.x_rot_base)
    rollSpread = round2(initialUiBuffer.settings.x_rot_spread)
    rollX = round2(initialUiBuffer.settings.x_axis_x_rot_coeff)
    rollY = round2(initialUiBuffer.settings.y_axis_x_rot_coeff)
    yawBase = round2(initialUiBuffer.settings.z_rot_base)
    yawSpread = round2(initialUiBuffer.settings.z_rot_spread)
    yawX = round2(initialUiBuffer.settings.x_axis_z_rot_coeff)
    yawY = round2(initialUiBuffer.settings.y_axis_z_rot_coeff)
  }

  function handleRotationSliderDoubleClick(inputId: string) {
    var input = document.getElementById(inputId)
    input.value = 0.0
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  export let instance: number
  let prevSettings: StorageSettings | null
  let renderDataReady = false
  let hasBeenDestroyed = false

  onMount(async () => {
    // console.dir({instance, prevSettings})


    // load wasm
    await wasm_bindgen() // loaded in index.html from ./pkg/src_rust.js
    
    let ses = localStorage.getItem("magic_square_settings")
    if (ses) {
      const res = JSON.parse(ses)
      console.dir({res})
      prevSettingsStore.update((_: StorageSettings): StorageSettings => {
        prevSettings = res
        return res
      })
    }

    if (!hasBeenDestroyed) { 
      // resize + key block in Container.svelte may destroy component before wasm_bindgen can load
      // without this check, it is possible to load two wasm instances
      // since wasm retrieves the elements using .get_element_by_id
      // and since a new instance of the component will havee been mounted by the time wasm_bindgen loads
      // the result is two identical wasm instances listening to the same ui elements and drawing to the same context
      const { MagicSquare, init_message } = wasm_bindgen
      console.log(
        init_message("Magic Square Wasm!")
      )
      
      // init wasm process and set initial values
      const initialUiBuffer = await MagicSquare.run(prevSettings)
      console.dir({prevSettings, uibuff: initialUiBuffer.settings})
      setInitialDrawPatternVars(initialUiBuffer)
      setInitialColorVars(initialUiBuffer)
      setInitialLfoVars(initialUiBuffer)
      setInitialMouseTracking(initialUiBuffer)
      setInitialRadiusVars(initialUiBuffer)
      setInitialRotationVars(initialUiBuffer)
      setInitialTranslationVars(initialUiBuffer)
      renderDataReady = true
    }
  })

  function deriveStorageSettings(): StorageSettings {
    return {
      color_1: color1,
      color_2: color2,
      color_3: color3,
      color_4: color4,
      color_5: color5,
      color_6: color6,
      color_7: color7,
      color_8: color8,

      // lfo_1
      lfo_1_active: lfo1Active,
      lfo_1_amp: lfo1Amp,
      lfo_1_dest: lfo1Dest,
      lfo_1_freq: lfo1Freq,
      lfo_1_phase: lfo1Phase,
      lfo_1_shape: lfo1Shape,

      // lfo_2
      lfo_2_active: lfo2Active,
      lfo_2_amp: lfo2Amp,
      lfo_2_dest: lfo2Dest,
      lfo_2_freq: lfo2Freq,
      lfo_2_phase: lfo2Phase,
      lfo_2_shape: lfo2Shape,

      // lfo_3
      lfo_3_active: lfo3Active,
      lfo_3_amp: lfo3Amp,
      lfo_3_dest: lfo3Dest,
      lfo_3_freq: lfo3Freq,
      lfo_3_phase: lfo3Phase,
      lfo_3_shape: lfo3Shape,

      // lfo_4
      lfo_4_active: lfo4Active,
      lfo_4_amp: lfo4Amp,
      lfo_4_dest: lfo4Dest,
      lfo_4_freq: lfo4Freq,
      lfo_4_phase: lfo4Phase,
      lfo_4_shape: lfo4Shape,

      // PATTERN
      draw_pattern: drawPattern,

      // RADIUS
      radius_base: radiusBase,
      radius_step: radiusStep,

      // ROTATION
      x_rot_base: rollBase,
      y_rot_base: pitchBase,
      z_rot_base: yawBase,

      x_rot_spread: rollSpread,
      y_rot_spread: pitchSpread,
      z_rot_spread: yawSpread,

      // rotation sensitivity to mouse movement
      x_axis_x_rot_coeff: rollX,
      x_axis_y_rot_coeff: pitchX,
      x_axis_z_rot_coeff: yawX,

      y_axis_x_rot_coeff: rollY,
      y_axis_y_rot_coeff: pitchY,
      y_axis_z_rot_coeff: yawY,

      // TRANSLATION
      translation_x_base: translationXBase,
      translation_x_spread: translationXSpread,
      translation_y_base: translationYBase,
      translation_y_spread: translationYSpread,
      translation_z_base: 0.0,
      translation_z_spread: 0.0,
      mouse_tracking: mouseTracking,
    }
  }

  const unsubscribe = prevSettingsStore.subscribe(val => prevSettings = val)

  onDestroy(() => {
    hasBeenDestroyed = true
    prevSettingsStore.update((_: StorageSettings) => {
      return deriveStorageSettings()
    })
    unsubscribe()
    let app = document.getElementById(("app_main"))
    app.dispatchEvent(new Event("destroymswasm", {bubbles: true}))
  })

  afterUpdate(() => {
    const res = deriveStorageSettings()
    if (Object.values(res).every(x => typeof x !== 'undefined')){
      window.localStorage.setItem(
        "magic_square_settings",
        JSON.stringify(res)
      )
    }
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
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Color bind:color1={color1}
                 bind:color2={color2}
                 bind:color3={color3}
                 bind:color4={color4}
                 bind:color5={color5}
                 bind:color6={color6}
                 bind:color7={color7}
                 bind:color8={color8}>
            <div slot="hiddenInputs">
              <input id={WasmInputId.color1}
                     bind:value={color1}
                     class="hidden_input">
              <input id={WasmInputId.color2}
                     bind:value={color2}
                     class="hidden_input">
              <input id={WasmInputId.color3}
                     bind:value={color3}
                     class="hidden_input">
              <input id={WasmInputId.color4}
                     bind:value={color4}
                     class="hidden_input">
              <input id={WasmInputId.color5}
                     bind:value={color5}
                     class="hidden_input">
              <input id={WasmInputId.color6}
                     bind:value={color6}
                     class="hidden_input">
              <input id={WasmInputId.color7}
                     bind:value={color7}
                     class="hidden_input">
              <input id={WasmInputId.color8}
                     bind:value={color8}
                     class="hidden_input">
            </div>
          </Color>
        {/if}
      </div>
      <div slot="drawPattern"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <DrawPatternContainer bind:drawPatternDirection={initialDrawPatternDirection}
                                bind:drawPatternCount={initialDrawPatternCount}>
            <div slot="hiddenInput">
              <input id={WasmInputId.drawPattern}
                     bind:value={drawPattern}
                     class="hidden_input"/>
            </div>
          </DrawPatternContainer>
        {/if}
      </div>
      <div slot="lfo"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <LfoContainer   lfo1Active={lfo1Active}
                          lfo1Dest={lfo1Dest}
                          lfo1Shape={lfo1Shape}
                          lfo2Active={lfo2Active}
                          lfo2Dest={lfo2Dest}
                          lfo2Shape={lfo2Shape}
                          lfo3Active={lfo3Active}
                          lfo3Dest={lfo3Dest}
                          lfo3Shape={lfo3Shape}
                          lfo4Active={lfo4Active}
                          lfo4Dest={lfo4Dest}
                          lfo4Shape={lfo4Shape}>

            <!-- LFO1 START-->
            <div class="w-full h-full p-5 grow flex flex-col justify-around items-stretch"
                 slot="lfo1">
              <!-- hidden input for destination select  -->
              <input id={WasmInputId.lfo1Dest}
                     bind:value={lfo1Dest}
                     class="hidden_input"/>
              <input id={WasmInputId.lfo1Shape}
                     bind:value={lfo1Shape}
                     class="hidden_input"/>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <!-- TODO: lfo active/selected colors for buttons  -->
                <button class="mb-5"
                        class:active={lfo1Active}
                        on:click={() => handleLfoActiveToggle(Lfo.one)}>
                  <input id={WasmInputId.lfo1Active}
                         value={lfo1Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo1Freq}>
                  <div> {i18n.t("frequency", langVal)} </div>
                  <div> {lfo1Freq} </div>
                </label>
                <input id={WasmInputId.lfo1Freq}
                       type="range"
                       min={1}
                       max={255}
                       bind:value={lfo1Freq}
                       step={1}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo1Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo1Amp)} </div>
                </label>
                <input id={WasmInputId.lfo1Amp}
                       type="range"
                       min={0}
                       max={3}
                       bind:value={lfo1Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo1Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {lfo1Phase} </div>
                </label>
                <input id={WasmInputId.lfo1Phase}
                       type="range"
                       min={-3.14159}
                       max={3.13159}
                       bind:value={lfo1Phase}
                       step={.01}/>
              </div>
            </div>
            <!-- LFO1 END-->

            <!-- lfo2 START-->
            <div class="w-full h-full p-5 grow flex flex-col justify-around items-stretch"
                 slot="lfo2">
              <!-- hidden input for destination select  -->
              <input id={WasmInputId.lfo2Dest}
                     bind:value={lfo2Dest}
                     class="hidden_input"/>
              <input id={WasmInputId.lfo2Shape}
                     bind:value={lfo2Shape}
                     class="hidden_input"/>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <!-- TODO: lfo active/selected colors for buttons  -->
                <button class="mb-5"
                        class:active={lfo2Active}
                        on:click={() => handleLfoActiveToggle(Lfo.two)}>
                  <input id={WasmInputId.lfo2Active}
                         value={lfo2Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo2Freq}>
                  <div> {i18n.t("frequency", langVal)} </div>
                  <div> {lfo2Freq} </div>
                </label>
                <input id={WasmInputId.lfo2Freq}
                       type="range"
                       min={1}
                       max={255}
                       bind:value={lfo2Freq}
                       step={1}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo2Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo2Amp)} </div>
                </label>
                <input id={WasmInputId.lfo2Amp}
                       type="range"
                       min={0}
                       max={3}
                       bind:value={lfo2Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo2Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {lfo2Phase} </div>
                </label>
                <input id={WasmInputId.lfo2Phase}
                       type="range"
                       min={-3.14159}
                       max={3.13159}
                       bind:value={lfo2Phase}
                       step={.01}/>
              </div>
            </div>
            <!-- lfo2 END-->

            <!-- lfo3 START-->
            <div class="w-full h-full p-5 grow flex flex-col justify-around items-stretch"
                 slot="lfo3">
              <!-- hidden input for destination select  -->
              <input id={WasmInputId.lfo3Dest}
                     bind:value={lfo3Dest}
                     class="hidden_input"/>
              <input id={WasmInputId.lfo3Shape}
                     bind:value={lfo3Shape}
                     class="hidden_input"/>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <!-- TODO: lfo active/selected colors for buttons  -->
                <button class="mb-5"
                        class:active={lfo3Active}
                        on:click={() => handleLfoActiveToggle(Lfo.three)}>
                  <input id={WasmInputId.lfo3Active}
                         value={lfo3Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo3Freq}>
                  <div> {i18n.t("frequency", langVal)} </div>
                  <div> {lfo3Freq} </div>
                </label>
                <input id={WasmInputId.lfo3Freq}
                       type="range"
                       min={1}
                       max={255}
                       bind:value={lfo3Freq}
                       step={1}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo3Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo3Amp)} </div>
                </label>
                <input id={WasmInputId.lfo3Amp}
                       type="range"
                       min={0}
                       max={3}
                       bind:value={lfo3Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo3Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {lfo3Phase} </div>
                </label>
                <input id={WasmInputId.lfo3Phase}
                       type="range"
                       min={-3.14159}
                       max={3.13159}
                       bind:value={lfo3Phase}
                       step={.01}/>
              </div>
            </div>
            <!-- lfo3 END-->

            <!-- lfo4 START-->
            <div class="w-full h-full p-5 grow flex flex-col justify-around items-stretch"
                 slot="lfo4">
              <!-- hidden input for destination select  -->
              <input id={WasmInputId.lfo4Dest}
                     bind:value={lfo4Dest}
                     class="hidden_input"/>
              <input id={WasmInputId.lfo4Shape}
                     bind:value={lfo4Shape}
                     class="hidden_input"/>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <!-- TODO: lfo active/selected colors for buttons  -->
                <button class="mb-5"
                        class:active={lfo4Active}
                        on:click={() => handleLfoActiveToggle(Lfo.four)}>
                  <input id={WasmInputId.lfo4Active}
                         value={lfo4Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo4Freq}>
                  <div> {i18n.t("frequency", langVal)} </div>
                  <div> {lfo4Freq} </div>
                </label>
                <input id={WasmInputId.lfo4Freq}
                       type="range"
                       min={1}
                       max={255}
                       bind:value={lfo4Freq}
                       step={1}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo4Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo4Amp)} </div>
                </label>
                <input id={WasmInputId.lfo4Amp}
                       type="range"
                       min={0}
                       max={3}
                       bind:value={lfo4Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo4Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {lfo4Phase} </div>
                </label>
                <input id={WasmInputId.lfo4Phase}
                       type="range"
                       min={-3.14159}
                       max={3.13159}
                       bind:value={lfo4Phase}
                       step={.01}/>
              </div>
            </div>
            <!-- lfo4 END-->

          </LfoContainer>
        {/if}
      </div>
      <div slot="translation"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Translation>
            <div  class="p-5 grow flex flex-col justify-around items-stretch"
                  slot="xSliders">
              <div class="w-full flex flex-col justify-between items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationXBase}>
                  <div> base </div>
                  <div> {translationXBase} </div>
                </label>
                <input id={WasmInputId.translationXBase}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={translationXBase}
                       step={.01}/>
              </div>
              <div class="w-full flex flex-col justify-between items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationXSpread}>
                  <div> spread </div>
                  <div> {translationXSpread} </div>
                </label>
                <input id={WasmInputId.translationXSpread}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={translationXSpread}
                       step={.01}/>
              </div>
            </div>
            <div  class="p-5 grow flex flex-col justify-around items-stretch"
                  slot="ySliders">
              <div class="w-full flex flex-col justify-between items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationYBase}>
                  <div> base </div>
                  <div> {translationYBase} </div>
                </label>
                <input id={WasmInputId.translationYBase}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={translationYBase}
                       step={.01}/>
              </div>
              <div class="w-full flex flex-col justify-between items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationYSpread}>
                  <div> spread </div>
                  <div> {translationYSpread} </div>
                </label>
                <input id={WasmInputId.translationYSpread}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={translationYSpread}
                       step={.01}/>
              </div>
              <!-- TODO: impliment depth perspective shifting in WebGl -->
              <!-- <div class="w-full flex flex-col justify-between items-stretch"> -->
              <!--   <label class="slider_label flex justify-between"  -->
              <!--          for={WasmInputId.translationZ}> -->
              <!--     <div> Z </div> -->
              <!--     <div> {translationZ} </div> -->
              <!--   </label> -->
              <!--   <input id={WasmInputId.translationZ} -->
              <!--          type="range" -->
              <!--          min={-2} -->
              <!--          max={2} -->
              <!--          bind:value={translationZ} -->
              <!--          step={.01}/> -->
              <!-- </div> -->
            </div>
            <div slot="mouseTracking">
              <MouseTrackingContainer currOption={mouseTracking}>
                <div slot="hiddenInput">
                  <input id={WasmInputId.mouseTracking}
                         bind:value={mouseTracking}
                         class="hidden_input"/>
                </div>
              </MouseTrackingContainer>
            </div>
          </Translation>
        {/if}
      </div>
      <div slot="geometry"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Geometry>
            <div  class="p-5 grow flex flex-col justify-around items-stretch"
                  slot="radiusSliders">
              <div class="w-full flex flex-col justify-between items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.radiusBase}>
                  <div> {i18n.t("base", langVal)} </div>
                  <div> {radiusBase} </div>
                </label>
                <input id={WasmInputId.radiusBase}
                       type="range"
                       min={0.1}
                       max={1.1}
                       bind:value={radiusBase}
                       step={.01}/>
              </div>
              <div class="w-full flex flex-col justify-between items-stretch">
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.radiusStep}>
                  <div> {i18n.t("step", langVal)} </div>
                  <div> {radiusStep} </div>
                </label>
                <input id={WasmInputId.radiusStep}
                       type="range"
                       min={-0.5}
                       max={0.5}
                       bind:value={radiusStep}
                       step={.01}/>
              </div>
            </div>
          </Geometry>
        {/if}
      </div>
      <div slot="rotation"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Rotation>
            <div slot="pitch"
                 class="grow flex flex-col justify-around items-stretch p-2">
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.pitchBase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.pitchBase}>
                  <div> {i18n.t("base", langVal)} </div>
                  <div> {pitchBase} </div>
                </label>
                <input id={WasmInputId.pitchBase}
                       type="range"
                       min={-6.33}
                       max={6.33}
                       bind:value={pitchBase}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.pitchSpread)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.pitchSpread}>
                  <div> {i18n.t("spread", langVal)} </div>
                  <div> {pitchSpread} </div>
                </label>
                <input id={WasmInputId.pitchSpread}
                       type="range"
                       min={-.33}
                       max={.33}
                       bind:value={pitchSpread}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.pitchX)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.pitchX}>
                  <div> {"X"} </div>
                  <div> {pitchX} </div>
                </label>
                <input id={WasmInputId.pitchX}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={pitchX}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.pitchY)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.pitchY}>
                  <div> {"Y"} </div>
                  <div> {pitchY} </div>
                </label>
                <input id={WasmInputId.pitchY}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={pitchY}
                       step={.01}/>
              </div>
            </div>
            <div slot="roll"
                 class="grow flex flex-col justify-around items-stretch p-2">
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.rollBase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.rollBase}>
                  <div> {i18n.t("base", langVal)} </div>
                  <div> {rollBase} </div>
                </label>
                <input id={WasmInputId.rollBase}
                       type="range"
                       min={-6.33}
                       max={6.33}
                       bind:value={rollBase}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.rollSpread)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.rollSpread}>
                  <div> {i18n.t("spread", langVal)} </div>
                  <div> {rollSpread} </div>
                </label>
                <input id={WasmInputId.rollSpread}
                       type="range"
                       min={-.33}
                       max={.33}
                       bind:value={rollSpread}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.rollX)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.rollX}>
                  <div> {"X"} </div>
                  <div> {rollX} </div>
                </label>
                <input id={WasmInputId.rollX}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={rollX}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.rollY)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.rollY}>
                  <div> {"Y"} </div>
                  <div> {rollY} </div>
                </label>
                <input id={WasmInputId.rollY}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={rollY}
                       step={.01}/>
              </div>
            </div>
            <div slot="yaw"
                 class="grow flex flex-col justify-around items-stretch p-2">
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.yawBase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.yawBase}>
                  <div> {i18n.t("base", langVal)} </div>
                  <div> {yawBase} </div>
                </label>
                <input id={WasmInputId.yawBase}
                       type="range"
                       min={-6.33}
                       max={6.33}
                       bind:value={yawBase}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.yawSpread)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.yawSpread}>
                  <div> {i18n.t("spread", langVal)} </div>
                  <div> {yawSpread} </div>
                </label>
                <input id={WasmInputId.yawSpread}
                       type="range"
                       min={-.33}
                       max={.33}
                       bind:value={yawSpread}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.yawX)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.yawX}>
                  <div> {"X"} </div>
                  <div> {yawX} </div>
                </label>
                <input id={WasmInputId.yawX}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={yawX}
                       step={.01}/>
              </div>
              <div class="flex flex-col"
                   on:dblclick={() => handleRotationSliderDoubleClick(WasmInputId.yawY)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.yawY}>
                  <div> {"Y"} </div>
                  <div> {yawY} </div>
                </label>
                <input id={WasmInputId.yawY}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={yawY}
                       step={.01}/>
              </div>
            </div>
          </Rotation>
        {/if}
      </div>
    </ControlRack>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"
  
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
        cursor: none
        background: color.$black-blue-horiz-grad
        border: 10px double color.$blue-7
        border-radius: 5px
        flex-grow: 1

  .display
    align-items: center
  
  .control
    flex-grow: 1
    height: 100%

  .slider_label
      width: 100%
      font-weight: text.$fw-l
      font-size: text.$fs-m
      padding-right: 5%

  .hidden_input
    display: none

  .active
    background-color: color.$red-5
</style>
