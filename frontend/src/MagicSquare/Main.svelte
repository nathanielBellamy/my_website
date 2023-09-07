<script lang="ts" type="module">
  import init, { MagicSquare, rust_init_message } from '../../pkg/src_rust.js'
  import { afterUpdate, onDestroy, onMount } from 'svelte'
  import Loading from '../lib/Loading.svelte'
  import DrawPatternContainer from './ControlModules/DrawPattern.svelte'
  import { DrawPatternType } from './ControlModules/DrawPattern'
  import { intoDrawPatternType } from './ControlModules/DrawPattern'
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
  import { intoTransformOrder, TransformOrder } from './ControlModules/TransformOrder'
  import { ColorDirection, intoColorDirection } from './ControlModules/Color'
  import { intoShape } from './ControlModules/Shape'
  import type { Shape } from './ControlModules/Shape'
  import Presets from './ControlModules/Presets.svelte'
  // INIT LANG BOILER PLATE
  import { I18n, Lang } from '../I18n'
  import { lang } from '../stores/lang'
  import { touchScreen } from '../stores/touchScreen'
  import { smallScreen } from '../stores/smallScreen'
  import Icon from '../lib/Icon.svelte'
  import { Icons } from '../lib/Icons.js'

  // TODO:
  // this combination of touchSreen store and value updates works 
  // to ensure that the touchScreen value is updated by the time it is passed to RustWasm
  // this includes the hidden element using touchScreenVal in the html
  // not sure why this magic combo gets it done
  // but we won't worry about it right now
  const id = (x: any): any => x
  let touchScreenVal: boolean
  const unsubTouchScreen = touchScreen.subscribe((val: boolean) => touchScreenVal = val)
  $: isTouchScreen = id(touchScreenVal)

  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean) => smallScreenVal = val)
  
  const i18n = new I18n('magicSquare/main')
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)

  enum MagicSquareView {
    square = "square",
    controls = "controls"
  }

  let magicSquareView: MagicSquareView = MagicSquareView.square
  function setMagicSquareView(msv: MagicSquareView) {
    magicSquareView = msv
  }

  // this component will be large
  // but it is meant to stay flat
  // it inits wasm and sets a lot of values
  // but there should not be a lot of logic in here
  // the mantra is
  //   -> Svelte/JS is for layout + display logic
  //   -> Rust/Wasm is for handling data
  
  // DRAW PATTERN
  let drawPatternType: DrawPatternType
  let drawPatternCount: number
  let drawPatternOffset: number
  let drawPatternSpeed: number

  function setInitialDrawPatternVars(initialSettings: any) {
    drawPatternType = intoDrawPatternType(initialSettings.draw_pattern_type)
    drawPatternCount = initialSettings.draw_pattern_count
    drawPatternOffset = initialSettings.draw_pattern_offset
    drawPatternSpeed = initialSettings.draw_pattern_speed
  }

  // COLOR
  let colorDirection: ColorDirection
  let colorSpeed: number
  let colors: number[][]

  // CSS (inline in Color.svelte) uses Int:0-255, WebGL uses Float:0.0-1.0
  function convertRgba(rgba: number[], dir: string): number[] {
    return rgba.map((x:number, idx: number) => {
      if (idx < 3 && dir == 'up') {
        return round2(x * 255)
      } else if (idx < 3 && dir == 'down') {
        return round2(x / 255)
      } else {
        return round2(x)
      }
    })
  }

  function handleRangeDoubleClick(id: WasmInputId) {
    var input = document.getElementById(id)
    input.value = 0
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }

  function setInitialColorVars(initialSettings: any) {
    colors = initialSettings.colors.map((x: number[]) => convertRgba(x, 'up'))
    colorDirection = intoColorDirection(initialSettings.color_direction)
    colorSpeed = initialSettings.color_speed
  }

  // GEOMETRY
  let radiusBase: number
  let radiusStep: number
  let shapes: Shape[]
  let transformOrder: TransformOrder

  function setInitialGeometryVars(initialSettings: any) {
    radiusBase = round2(initialSettings.radius_base)
    radiusStep = round2(initialSettings.radius_step)
    shapes = initialSettings.shapes.map((x:string) => intoShape(x))
    transformOrder = intoTransformOrder(initialSettings.transform_order)
  }

  function round2(val: number){
    return Math.floor(val * 100) / 100
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
  
  function setInitialLfoVars(initialSettings: any) {
    lfo1Active = initialSettings.lfo_1_active
    lfo1Amp = round2(initialSettings.lfo_1_amp)
    lfo1Dest = intoLfoDestination(initialSettings.lfo_1_dest)
    lfo1Freq = round2(initialSettings.lfo_1_freq)
    lfo1Phase = round2(initialSettings.lfo_1_phase)
    lfo1Shape = intoLfoShape(initialSettings.lfo_1_shape)

    lfo2Active = initialSettings.lfo_2_active
    lfo2Amp = round2(initialSettings.lfo_2_amp)
    lfo2Dest = intoLfoDestination(initialSettings.lfo_2_dest)
    lfo2Freq = round2(initialSettings.lfo_2_freq)
    lfo2Phase = round2(initialSettings.lfo_2_phase)
    lfo2Shape = intoLfoShape(initialSettings.lfo_2_shape)

    lfo3Active = initialSettings.lfo_3_active
    lfo3Amp = round2(initialSettings.lfo_3_amp)
    lfo3Dest = intoLfoDestination(initialSettings.lfo_3_dest)
    lfo3Freq = round2(initialSettings.lfo_3_freq)
    lfo3Phase = round2(initialSettings.lfo_3_phase)
    lfo3Shape = intoLfoShape(initialSettings.lfo_3_shape)

    lfo4Active = initialSettings.lfo_4_active
    lfo4Amp = round2(initialSettings.lfo_4_amp)
    lfo4Dest = intoLfoDestination(initialSettings.lfo_4_dest)
    lfo4Freq = round2(initialSettings.lfo_4_freq)
    lfo4Phase = round2(initialSettings.lfo_4_phase)
    lfo4Shape = intoLfoShape(initialSettings.lfo_4_shape)
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

  // PRESET
  let preset: number

  function setInitialPreset(initialSettings: any) {
    preset = initialSettings.preset
  }

  // TRANSLATION
  let translationXBase: number
  let translationXSpread: number
  let translationYBase: number
  let translationYSpread: number
  // let translationZ: number

  function setInitialTranslationVars(initialSettings: any) {
    translationXBase = round2(initialSettings.translation_x_base)
    translationXSpread = round2(initialSettings.translation_x_spread)
    translationYBase = round2(initialSettings.translation_y_base)
    translationYSpread = round2(initialSettings.translation_y_spread)
    // translationZ = initialSettings.translation_z
  }

  let mouseTracking: MouseTracking

  function setInitialMouseTracking(initialSettings: any) {
    mouseTracking = initialSettings.mouse_tracking
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

  function setInitialRotationVars(initialSettings: any) {
    pitchBase = round2(initialSettings.y_rot_base)
    pitchSpread = round2(initialSettings.y_rot_spread)
    pitchX = round2(initialSettings.x_axis_y_rot_coeff)
    pitchY = round2(initialSettings.y_axis_y_rot_coeff)
    rollBase = round2(initialSettings.x_rot_base)
    rollSpread = round2(initialSettings.x_rot_spread)
    rollX = round2(initialSettings.x_axis_x_rot_coeff)
    rollY = round2(initialSettings.y_axis_x_rot_coeff)
    yawBase = round2(initialSettings.z_rot_base)
    yawSpread = round2(initialSettings.z_rot_spread)
    yawX = round2(initialSettings.x_axis_z_rot_coeff)
    yawY = round2(initialSettings.y_axis_z_rot_coeff)
  }

  let prevSettings: StorageSettings | null
  let renderDataReady = false
  let hasBeenDestroyed = false

  function setAllSettings(settings: StorageSettings) {
    setInitialDrawPatternVars(settings)
    setInitialColorVars(settings)
    setInitialLfoVars(settings)
    setInitialGeometryVars(settings)
    setInitialMouseTracking(settings)
    setInitialPreset(settings)
    setInitialRotationVars(settings)
    setInitialTranslationVars(settings)
  }

  function setAllSettingsFromPreset() {
    // if not already present
    // presets set in local_storage by RustWasm during init
    let presets = JSON.parse(localStorage.getItem("magic_square_presets"))
    if (!!presets && !!presets[preset]) {
      setAllSettings(presets[preset])
    }
  }

  function deriveStorageSettings(): StorageSettings {
    return {
      // Color
      colors: (colors || []).map(color => convertRgba(color, 'down')),
      color_direction: colorDirection,
      color_speed: colorSpeed,

      // DRAW PATTERN
      draw_pattern_type: drawPatternType,
      draw_pattern_count: drawPatternCount,
      draw_pattern_offset: drawPatternOffset,
      draw_pattern_speed: drawPatternSpeed,

      // Geometry
      radius_base: radiusBase,
      radius_step: radiusStep,
      shapes: shapes,
      transform_order: transformOrder,

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

      preset: preset,

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
      translation_z_base: translationXBase, // unused
      translation_z_spread: translationXSpread, // unused
      mouse_tracking: mouseTracking,
    }
  }

  const unsubPrevSettings = prevSettingsStore.subscribe(val => prevSettings = val)

  onDestroy(() => {
    hasBeenDestroyed = true
    const storageSettings: StorageSettings = deriveStorageSettings()
    if (!Object.values(storageSettings).some(x => typeof x === 'undefined')) {
      prevSettingsStore.update((_: StorageSettings) => storageSettings)
      window.localStorage.setItem("magic_square_settings", JSON.stringify(storageSettings))
    }
    unsubLang()
    unsubPrevSettings()
    unsubSmallScreen()
    unsubTouchScreen()
    let app = document.getElementById(("app_main"))
    app.dispatchEvent(new Event("destroymswasm", {bubbles: true}))
  })

  afterUpdate(() => {
    const res: StorageSettings = deriveStorageSettings()
    if (Object.values(res).every(x => typeof x !== 'undefined')){
      window.localStorage.setItem(
        "magic_square_settings",
        JSON.stringify(res)
      )
    }
  })

  async function run() {
    let ses = localStorage.getItem("magic_square_settings")
    if (ses) {
      const res = JSON.parse(ses)
      // console.dir({res})
      prevSettingsStore.update((_: StorageSettings): StorageSettings => {
        prevSettings = res
        return res
      })
    }

    let presets = JSON.parse(localStorage.getItem("magic_square_presets"))

    if (!hasBeenDestroyed) { 
      // resize + key block in Container.svelte may destroy component before wasm_bindgen can load
      // without this check, it is possible to load two wasm instances
      // since wasm retrieves the elements using .get_element_by_id
      // and since a new instance of the component will havee been mounted by the time wasm_bindgen loads
      // the result is two identical wasm instances listening to the same ui elements and drawing to the same context
      await init()
      rust_init_message("Magic Square Wasm!")
      
      // init wasm process and set initial values
      const initialSettings = await MagicSquare.run(prevSettings, presets, touchScreenVal)
      setAllSettings(initialSettings)
      renderDataReady = true
    }
  }

  onMount(async () => {
    await run()
  })
</script>

<!-- we use touchSceenVal here to ensure Svelte has it updated by the time it reaches RustWasm -->
<div style="display: none"> {touchScreenVal}  </div>

<div id="magic_square"
     class="magic_square overscroll-none"
     class:grid_col={smallScreenVal}
     class:grid_row={!smallScreenVal}>
  {#if smallScreenVal}
     <div class="text-sm grid grid-cols-2 grid-rows-1">
        <button on:click={() => setMagicSquareView(MagicSquareView.square)}
                class="view_select_button pt-2 pb-2 flex justify-around items-center"
                class:selected={magicSquareView === MagicSquareView.square}>
          <span class="text-cyan-500">
            <Icon icon={Icons.EyeSolid} />
          </span>
        </button>
        <button on:click={() => setMagicSquareView(MagicSquareView.controls)}
                class="view_select_button text-sm pt-2 pb-2 flex justify-around items-center"
                class:selected={magicSquareView === MagicSquareView.controls}>
          <span class="text-cyan-500">
            <Icon icon={Icons.GearSolid} />
          </span>
        </button>
     </div>
  {/if}
  <div id="magic_square_canvas_container"
       class="magic_square_canvas_container flex flex-col justify-around display"
       class:hidden={smallScreenVal && magicSquareView !== MagicSquareView.square}>
    <canvas id="magic_square_canvas"
            class="magic_square_canvas"/>
  </div>
  <div class="w-full"
       class:hidden={smallScreenVal && magicSquareView !== MagicSquareView.controls}
       class:overflow-hidden={!smallScreenVal}>
    <ControlRack>
      <div slot="color"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Color bind:colorDirection={colorDirection}
                 bind:colors={colors}>
            <div slot="speed">
              <div class="grow w-full flex flex-col justify-center items-stretch">
                <label class="slider_label flex justify-between" 
                       class:disabled={colorDirection === ColorDirection.fix} 
                       for={WasmInputId.colorSpeed}>
                  <div> {i18n.t("speed", langVal)} </div>
                  <div> {colorSpeed} </div>
                </label>
                <input id={WasmInputId.colorSpeed}
                       on:dblclick={() => handleRangeDoubleClick(WasmInputId.colorSpeed)}
                       disabled={colorDirection === ColorDirection.fix}
                       type="range"
                       min={1}
                       max={20}
                       bind:value={colorSpeed}
                       step={1}/>
              </div>
            </div>
            <div slot="hiddenInputs">
              <input id={WasmInputId.colorDirection}
                     bind:value={colorDirection}
                     class="hidden_input">
              <input id={WasmInputId.colors}
                     bind:value={colors}
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
          <DrawPatternContainer bind:transformOrder={transformOrder}
                                drawPatternType={drawPatternType}>
            <div slot="transformOrder">
              <input id={WasmInputId.transformOrder}
                     bind:value={transformOrder}
                     class="hidden_input"/>
            </div>
            <div slot="countAndSpeed"
                 class="grow flex flex-col justify-between items-stretch">
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.drawPatternCount)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.drawPatternCount}>
                  <div> {i18n.t("count", langVal)} </div>
                  <div> {drawPatternCount} </div>
                </label>
                <input id={WasmInputId.drawPatternCount}
                       type="range"
                       min={1}
                       max={16}
                       bind:value={drawPatternCount}
                       step={1}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.drawPatternSpeed)}>
                <label class="slider_label flex justify-between" 
                       class:disabled={drawPatternType == DrawPatternType.fix}
                       for={WasmInputId.drawPatternSpeed}>
                  <div> {i18n.t("speed", langVal)} </div>
                  <div> {drawPatternSpeed} </div>
                </label>
                <input id={WasmInputId.drawPatternSpeed}
                       type="range"
                       min={1}
                       max={20}
                       disabled={drawPatternType == DrawPatternType.fix}
                       bind:value={drawPatternSpeed}
                       step={1}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.drawPatternOffset)}>
                <label class="slider_label flex justify-between"
                       class:disabled={drawPatternType !== DrawPatternType.fix}
                       for={WasmInputId.drawPatternOffset}>
                  <div> {i18n.t("offset", langVal)} </div>
                  <div> {drawPatternOffset} </div>
                </label>
                <input id={WasmInputId.drawPatternOffset}
                       type="range"
                       min={0}
                       max={15}
                       bind:value={drawPatternOffset}
                       disabled={drawPatternType !== DrawPatternType.fix}
                       step={1}/>
              </div>
            </div>
            <div slot="hiddenInput">
              <input id={WasmInputId.drawPatternType}
                     bind:value={drawPatternType}
                     class="hidden_input"/>
            </div>
          </DrawPatternContainer>
        {/if}
      </div>
        <!-- START GEOMETRY -->
      <div slot="geometry"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Geometry bind:shapes={shapes}>
            <div slot="shapes">
              <input id={WasmInputId.shapes} 
                     class="hidden_input"/>
            </div>
            <div  class="pl-5 pr-5 pb-5 grow flex flex-col justify-between items-stretch"
                  slot="radiusSliders">
              <div class="w-full flex flex-col justify-between items-stretch"
                       on:dblclick={() => handleRangeDoubleClick(WasmInputId.radiusBase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.radiusBase}>
                  <div> {i18n.t("base", langVal)} </div>
                  <div> {radiusBase} </div>
                </label>
                <input id={WasmInputId.radiusBase}
                       type="range"
                       min={-1.1}
                       max={1.1}
                       bind:value={radiusBase}
                       step={.01}/>
              </div>
              <div class="w-full flex flex-col justify-between items-stretch"
                       on:dblclick={() => handleRangeDoubleClick(WasmInputId.radiusStep)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.radiusStep}>
                  <div> {i18n.t("step", langVal)} </div>
                  <div> {radiusStep} </div>
                </label>
                <input id={WasmInputId.radiusStep}
                       type="range"
                       min={-0.2}
                       max={0.2}
                       bind:value={radiusStep}
                       step={.001}/>
              </div>
            </div>
          </Geometry>
        {/if}
      </div>
        <!-- END GEOMETRY -->
      <!-- START LFO  -->
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
            <div class="w-full h-full pl-5 pr-5 grow flex flex-col justify-around items-stretch"
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
                <button class="mt-2 mb-2"
                        class:active={lfo1Active}
                        on:click={() => handleLfoActiveToggle(Lfo.one)}>
                  <input id={WasmInputId.lfo1Active}
                         value={lfo1Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                       on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo1Freq)}>
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
              <div class="grow w-full flex flex-col justify-center items-stretch"
                       on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo1Amp)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo1Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo1Amp)} </div>
                </label>
                <input id={WasmInputId.lfo1Amp}
                       type="range"
                       min={-1}
                       max={1}
                       bind:value={lfo1Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo1Phase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo1Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {round2(lfo1Phase)} </div>
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
            <div class="w-full h-full pl-5 pr-5 grow flex flex-col justify-around items-stretch"
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
                <button class="mt-2 mb-2"
                        class:active={lfo2Active}
                        on:click={() => handleLfoActiveToggle(Lfo.two)}>
                  <input id={WasmInputId.lfo2Active}
                         value={lfo2Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo2Freq)}>
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
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo2Amp)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo2Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo2Amp)} </div>
                </label>
                <input id={WasmInputId.lfo2Amp}
                       type="range"
                       min={-1}
                       max={1}
                       bind:value={lfo2Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo2Phase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo2Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {round2(lfo2Phase)} </div>
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
            <div class="w-full h-full pl-5 pr-5 grow flex flex-col justify-around items-stretch"
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
                <button class="mt-2 mb-2"
                        class:active={lfo3Active}
                        on:click={() => handleLfoActiveToggle(Lfo.three)}>
                  <input id={WasmInputId.lfo3Active}
                         value={lfo3Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo3Freq)}>
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
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo3Amp)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo3Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo3Amp)} </div>
                </label>
                <input id={WasmInputId.lfo3Amp}
                       type="range"
                       min={-1}
                       max={1}
                       bind:value={lfo3Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                       on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo3Phase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo3Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {round2(lfo3Phase)} </div>
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
            <div class="w-full h-full pl-5 pr-5 grow flex flex-col justify-around items-stretch"
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
                <button class="mt-2 mb-2"
                        class:active={lfo4Active}
                        on:click={() => handleLfoActiveToggle(Lfo.four)}>
                  <input id={WasmInputId.lfo4Active}
                         value={lfo4Active}
                         class="hidden_input">
                    {i18n.t("active", langVal)}
                </button>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo4Freq)}>
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
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo4Amp)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo4Amp}>
                  <div> {i18n.t("amplitude", langVal)} </div>
                  <div> {round2(lfo4Amp)} </div>
                </label>
                <input id={WasmInputId.lfo4Amp}
                       type="range"
                       min={-1}
                       max={1}
                       bind:value={lfo4Amp}
                       step={.01}/>
              </div>
              <div class="grow w-full flex flex-col justify-center items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.lfo4Phase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.lfo4Phase}>
                  <div> {i18n.t("phase", langVal)} </div>
                  <div> {round2(lfo4Phase)} </div>
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
      <!-- LFO END -->
      <!-- PRESETS START -->
      <div slot="presets"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Presets  updateUiSettings={setAllSettingsFromPreset}
                    bind:preset={preset}>
            <div slot="preset">
              <input id={WasmInputId.preset}
                     value={preset}
                     class="hidden_input"/>
            </div>
          </Presets>
        {/if}
      </div>
      <!-- PRESETS ENDS -->

      <!-- TRANSLATION -->
      <div slot="translation"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Translation>
            <div  class="pl-5 pr-5 grow flex flex-col justify-around items-stretch"
                  slot="xSliders">
              <div class="w-full flex flex-col justify-between items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.translationXBase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationXBase}>
                  <div> {i18n.t("base", langVal)} </div>
                  <div> {translationXBase} </div>
                </label>
                <input id={WasmInputId.translationXBase}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={translationXBase}
                       step={.01}/>
              </div>
              <div class="w-full flex flex-col justify-between items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.translationXSpread)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationXSpread}>
                  <div> {i18n.t("spread", langVal)} </div>
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
            <div  class="pl-5 pr-5 grow flex flex-col justify-around items-stretch"
                  slot="ySliders">
              <div class="w-full flex flex-col justify-between items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.translationYBase)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationYBase}>
                  <div> {i18n.t("base", langVal)} </div>
                  <div> {translationYBase} </div>
                </label>
                <input id={WasmInputId.translationYBase}
                       type="range"
                       min={-2}
                       max={2}
                       bind:value={translationYBase}
                       step={.01}/>
              </div>
              <div class="w-full flex flex-col justify-between items-stretch"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.translationYSpread)}>
                <label class="slider_label flex justify-between" 
                       for={WasmInputId.translationYSpread}>
                  <div> {i18n.t("spread", langVal)} </div>
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
      <!-- TRANSFORMATION END -->

      <!-- ROTATION START -->
      <div slot="rotation"
           class="h-full">
        {#if !renderDataReady}
          <Loading />
        {:else}
          <Rotation>
            <div slot="pitch"
                 class="grow flex flex-col justify-around items-stretch p-2">
              <div class="flex flex-col"
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.pitchBase)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.pitchSpread)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.pitchX)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.pitchY)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.rollBase)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.rollSpread)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.rollX)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.rollY)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.yawBase)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.yawSpread)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.yawX)}>
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
                   on:dblclick={() => handleRangeDoubleClick(WasmInputId.yawY)}>
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

    &_canvas
      border-top: 5px double color.$blue-7
      border-bottom: 5px double color.$blue-7
      border-radius: 10px
      cursor: crosshair
      &_container
        background: color.$black-blue-horiz-grad
        border: 10px double color.$blue-7
        border-radius: 5px
        flex-grow: 1
        touch-action: auto

  .grid_col
    display: grid
    grid-template-columns: 1fr
    grid-template-rows: 2em calc(100% - 2em)
    gap: 5px

  .grid_row
    display: grid
    grid-template-columns: 1fr 1fr
    grid-template-rows: 100%
    gap: 5px

  .disabled
    color: #666

  .display
    align-items: center

  .slider_label
      width: 100%
      font-weight: text.$fw-l
      font-size: text.$fs-m
      padding-right: 5%

  .view_select_button
    border: 3px solid color.$blue-7
    box-shadow: none

  .hidden_input
    display: none

  .active
    background-color: color.$red-7

  .selected
    background-color: color.$green-7
</style>
