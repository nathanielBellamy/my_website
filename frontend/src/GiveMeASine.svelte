<script lang='ts' type="module">
  import { onDestroy } from 'svelte'
  import init, { GmasWasm, rust_init_message  } from '../pkg/src_rust.js'
  import Icon from './lib/Icon.svelte'
  import { Icons } from './lib/Icons.js'

  // INIT LANG BOILER PLATE
  import { I18n, Lang } from "./I18n"
  import { lang } from './stores/lang'
  const i18n = new I18n("gmas")
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)

  import { smallScreen } from './stores/smallScreen'
  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean | null) => smallScreenVal = val)

  enum GmasView {
    graph = "graph",
    control = "control"
  }

  let gmasView: GmasView = GmasView.graph

  function setGmasView(gv: GmasView) {
    gmasView = gv
  }

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

  async function run() {
    await init()

    rust_init_message("GMAS")
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
  }

  run()

  onDestroy(() => {
    unsubLang()
    unsubSmallScreen()
  })
</script>

<body class="give_me_a_sine overflow-y-scroll overscroll-none"
      class:gmas_flex={!smallScreenVal}
      class:gmas_grid={smallScreenVal}>
  {#if smallScreenVal}
   <div class="w-full text-sm grid grid-cols-2 grid-rows-1">
      <button on:click={() => setGmasView(GmasView.graph)}
              class="view_select_button text-xl pt-2 pb-2 flex justify-around items-center"
              class:selected={gmasView === GmasView.graph}>
        <span class="text-cyan-500">
          <Icon icon={Icons.EyeSolid} />
        </span>
      </button>
      <button on:click={() => setGmasView(GmasView.control)}
              class="view_select_button text-xl pt-2 pb-2 flex justify-around items-center"
              class:selected={gmasView === GmasView.control}>
        <span class="text-cyan-500">
          <Icon icon={Icons.UserSettingsSolid} />
        </span>
      </button>
   </div>
  {/if}
  <div id="give_me_a_sine_output"
       class="give_me_a_sine_output device_graph_font overscroll-none"
       class:flex={!smallScreenVal || gmasView !== GmasView.graph}
       class:hidden={smallScreenVal && gmasView !== GmasView.graph}>
  </div>
  <div id="give_me_a_sine_form"
       class="give_me_a_sine_form overflow-y-scroll flex flex-col overscroll-none"
       class:hidden={smallScreenVal && gmasView !== GmasView.control}>
    <div class="give_me_a_sine_form_header font-bold"
         data-testid="gmas_form_header">
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
          <p>{i18n.t("height", langVal)}</p>
          <p>{height}</p>
        </label>
        <input id={'gmas_form_input_height'}
               type="range"
               min={1}
               max={255}
               bind:value={height}
               step={1}/>
      </div>
      <div class="give_me_a_sine_form_cell">
        <label  class="give_me_a_sine_form_cell_label"
                for={'gmas_form_input_width'}>
          <p>{i18n.t("width", langVal)}</p>
          <p>{width}</p>
        </label>
        <input id={'gmas_form_input_width'}
               type="range"
               min={1}
               max={255}
               bind:value={width}
               step={1}/>
      </div>
      <div class="give_me_a_sine_form_cell">
        <label  class="give_me_a_sine_form_cell_label"
                for={'gmas_form_input_above_char'}>
          <p>{i18n.t("aboveColor", langVal)}</p>
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
          <p>{i18n.t("belowColor", langVal)}</p>
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
          <p>{i18n.t("graphColor", langVal)}</p>
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
</body>

<style lang="sass">
  @use "./styles/color"
  @use "./styles/text"

  @media (max-width : 700px)
      .device_graph_font
        font-size: 9px

  .gmas
    &_flex
      display: flex
      flex-direction: row-reverse
      justify-content: stretch
      align-items: center
    &_grid
      display: grid
      grid-template-rows: 10% 90%
      grid-template-columns: 100%

  .view_select_button
    border: 3px solid color.$blue-7
    box-shadow: none

  .selected
    background-color: color.$blue-7

  .give_me_a_sine
    flex-grow: 1
    overflow-y: scroll
    margin: 5px

    &_output
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
