<script lang="ts">
  import { push } from "svelte-spa-router"
  import Link from "./lib/Link.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from './stores/lang'
  import magicSquareExampleGif from './assets/magic_square_example.gif'
  import giveMeASineExampleGif from './assets/give_me_a_sine_example.gif'

  import { intoUrl, siteSection, SiteSection } from "./stores/siteSection";
  import AiMe from "./lib/AiMe.svelte";

  let siteSectionVal: SiteSection
  siteSection.subscribe((val: SiteSection) => siteSectionVal = val)

  // INIT LANG BOILER PLATE
  const i18n = new I18n("home")
  let langVal: Lang
  lang.subscribe(val => langVal = val)

  let innerHeight: number
  $: imgSideLength = deriveImgSideLength(innerHeight)

  function deriveImgSideLength(ih: number): string {
    return Math.floor(ih / 4.2).toString() + "px"
  }

  function handlePreviewClick(s: SiteSection) {
    localStorage.setItem('ns_site_section', s)
    siteSection.update((_: SiteSection) => s)
    push(intoUrl(s))
  }
</script>

<svelte:window bind:innerHeight />

<body class="pl-5 pr-5 pb-5 flex flex-col justify-between items-stretch gap-2 overflow-y-scroll">
  <div class="home_title_container flex flex-col justify-between items-stretch md:flex-row md:justify-start md:items-center">
    <h2 class="home_title text-left pl-5">
      {i18n.t("title", langVal)}
    </h2>
    <ul class="home_intro_list text-left p-5 flex flex-col justify-between items-stretch">
      <li>
        {i18n.t("intro/2", langVal)}
        <p>
          <Link href="https://www.rust-lang.org/"
                title="Rust"
                sameOrigin={false}/>
          +
          <Link href="https://svelte.dev/"
                title="Svelte"
                sameOrigin={false}/>
          +
          <Link href="https://www.typescriptlang.org/"
                title="Typescript"
                sameOrigin={false}/>
          {i18n.t("intro/3", langVal)}
        </p>
        <Link href="https://webassembly.org/"
              title="WebAssembly"
              sameOrigin={false}/>
        +
        <Link href="https://crates.io/crates/wasm-bindgen"
              title="wasm-bindgen."
              sameOrigin={false}/>
      </li>
      <li>
        {i18n.t("intro/4", langVal)}
      </li>
    </ul>
  </div>
  <div class="home_title_dark text-left pl-5">
    {i18n.t("whatsHere", langVal)}
  </div>
  <div class="grow pl-5 pr-5 pb-5 flex flex-col justify-between items-stretch">
    <div class="grow grid grid-cols-1 grid-rows-3 md:grid-cols-3 md:grid-rows-1 gap-3">
      <button on:click={() => handlePreviewClick(SiteSection.about)}
              class="preview grid grid-cols-1 grid-rows-4 gap-2">
        <div class="preview_title">
          {i18n.t("about", langVal)}
        </div>
        <div class="row-span-2 w-full flex justify-around items-center">
          <div class="ai_me_container magic_square_img grid grid-rows-1 grid-cols-1">
            <AiMe imgSideLength={imgSideLength}/>
          </div>
        </div>
        <div class="flex pl-5 pr-5 mb-2 justify-around items-center overflow-y-scroll">
          <ul class="preview_list">
            <li>
              {i18n.t("about_1", langVal)}
            </li>
            <li>
              {i18n.t("about_2", langVal)}
            </li>
          </ul>
        </div>
      </button>
      <button on:click={() => handlePreviewClick(SiteSection.magicSquare)}
              class="preview grid grid-cols-1 grid-rows-4 gap-2">
        <div class="preview_title">
          {i18n.t("magicSquare", langVal)}
        </div>
        <div class="row-span-2 flex justify-around items-center">
          <img class="magic_square_img ai_me"
               src={magicSquareExampleGif}
               style:height={imgSideLength}
               style:width={imgSideLength}
               alt="Magic Square Example"/>
        </div>
        <div class="flex pl-5 pr-5 mb-2 justify-around items-center overflow-y-scroll">
          <ul class="preview_list">
            <li>
              {i18n.t("magicSquare_1", langVal)}
              <Link href="https://www.khronos.org/webgl/"
                    title="WebGL"
                    sameOrigin={false}/>
              +
              <Link href="https://rustwasm.github.io/wasm-bindgen/examples/webgl.html"
                    title="RustWasm"
                    sameOrigin={false}/>
            </li>
            <li>
              {i18n.t("magicSquare_2", langVal)}
              <Link href="https://en.wikipedia.org/wiki/Modular_synthesizer"
                    title={i18n.t("magicSquare_3", langVal)}
                    sameOrigin={false}/>
            </li>
          </ul>
        </div>
      </button>
      <button on:click={() => handlePreviewClick(SiteSection.giveMeASine)}
              class="preview grid grid-cols-1 grid-rows-4 gap-2">
        <div class="preview_title">
          {i18n.t("giveMeASine", langVal)}
        </div>
        <div class="row-span-2 flex justify-around items-center">
          <img class="magic_square_img"
               src={giveMeASineExampleGif}
               style:height={imgSideLength}
               style:width={imgSideLength}
               alt="Give Me A Sine Example"/>
        </div>
        <div class="flex flex-col pl-5 pr-5 mb-2 justify-around items-stretch overflow-y-scroll">
          <ul class="preview_list">
            <li>
              {i18n.t("giveMeASine_1", langVal)}
            </li>
            <li>
              <Link href="https://rustwasm.github.io/wasm-bindgen"
                    title="RustWasm"
                    sameOrigin={false}/>
              {i18n.t("giveMeASine_2", langVal)}
              <Link href="/magic_square" 
                    title="Magic Square"
                    sameOrigin={true}/>
            </li>
          </ul>
        </div>
      </button>
    </div>
  </div>
</body>

<style lang="sass">
  @use "./styles/color"
  @use "./styles/text"

  .home
    &_title
      color: color.$blue-5
      font-weight: text.$fw-l
      font-size: text.$fs-l
      &_dark
        color: color.$green-4
        font-weight: text.$fw-l
        font-size: text.$fs-ml
      &_container
        border-bottom: 5px double color.$blue-5

    &_intro_list
      color: color.$green-4
      font-weight: text.$fw-m

  .magic_square_img
    border: 5px double color.$blue-5
    border-radius: 5px
    padding: 5px
    margin: 5px
    border-radius: 50% 

  .ai_me_container
    grid-template-areas: "img"
 
  .preview
    border-radius: 5px
    margin: 0 5px 0 5px
    padding: 0 20px 5px 20px
    color: color.$green-4
    font-weight: text.$fw-l
    font-size: text.$fs-s
    background: color.$black-blue-grad
    border-width: 0
    grid-template-rows: 0.5fr 1fr 1fr 80px
    &_title
      color: color.$blue-6
      font-weight: text.$fw-l
      font-size: text.$fs-l
      display: flex
      justify-content: space-around
      align-items: center

    &_list
      padding: 0 10px 0 10px
      text-align: left
      list-style-type: square
      height: 100%
      
</style>
