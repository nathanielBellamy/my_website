<script lang="ts">
  import { onDestroy } from "svelte"
  import { push } from "svelte-spa-router"
  import Link from "./lib/Link.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from './stores/lang'
  import magicSquareExampleGif from './assets/magic_square_example.gif'
  import giveMeASineExampleGif from './assets/give_me_a_sine_example.gif'
  import AiMe from "./lib/AiMe.svelte";
  import { intoUrl, siteSection, SiteSection } from "./stores/siteSection";

  let siteSectionVal: SiteSection
  const unsubSiteSection = siteSection.subscribe((val: SiteSection) => siteSectionVal = val)

  // INIT LANG BOILER PLATE
  const i18n = new I18n("home")
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)

  let innerHeight: number
  let innerWidth: number
  $: imgSideLength = deriveImgSideLength(innerHeight, innerWidth)

  $: preview_title_font_size = computePreviewTitleFontSize(innerHeight, innerWidth)
  $: preview_text_font_size = computePreviewTextFontSize(innerHeight, innerWidth)

  $: showText = deriveShowText(innerWidth)

  function deriveImgSideLength(ih: number, iw: number): string {
    return Math.floor(Math.min(ih, iw) / 4.2).toString() + "px"
  }

  function computePreviewTitleFontSize(ih: number, iw: number): string {
    return Math.floor(Math.min(ih, iw) / 15.2).toString() + "px"
  }
  
  function computePreviewTextFontSize(ih: number, iw: number): string {
    if (iw > 768) {
      return Math.floor(Math.min(ih, iw) / 50.2).toString() + "px"
    } else {
      "auto"
    }
  }

  function deriveShowText(iw: number): Boolean {
    return iw > 768 // taliwind md: cutoff
  }

  function handlePreviewClick(s: SiteSection) {
    localStorage.setItem('ns_site_section', s)
    siteSection.update((_: SiteSection) => s)
    push(intoUrl(s))
  }

  onDestroy(() => {
    unsubLang()
    unsubSiteSection()
  })
</script>

<svelte:window bind:innerHeight
               bind:innerWidth />

<body class="pl-5 pr-5 pb-5 flex flex-col justify-between items-stretch gap-2 overflow-y-scroll">
  <div class="home_title_container flex flex-col justify-between items-stretch md:flex-row md:justify-start md:items-center">
    <h2 class="home_title font-mono flex justify-around items-center pl-2 pr-2 mt-2 md:mt-0">
      {i18n.t("title", langVal)}
    </h2>
    <ul class="home_intro_list text-left p-5 flex flex-col justify-between items-stretch">
      <li>
        {i18n.t("intro/2", langVal)}
        <p>
          <Link href="https://rustwasm.github.io/"
                title="RustWasm"
                sameOrigin={false}/>
          +
          <Link href="https://www.typescriptlang.org/"
                title="Typescript"
                sameOrigin={false}/>
          +
          <Link href="https://go.dev/"
                title="Go"
                sameOrigin={false}/>
          {i18n.t("intro/3", langVal)}
        </p>
        <Link href="https://svelte.dev/"
              title="Svelte"
              sameOrigin={false}/>
        +
        <Link href="https://nixos.org/"
              title="NixOS"
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
  <div class="h-5/6 flex flex-col justify-between items-stretch md:grid md:grid-cols-3 md:grid-rows-1 gap-3">
    <button on:click={() => handlePreviewClick(SiteSection.about)}
            class="preview md:flex md:flex-col md:justify-between md:items-center md:h-5/6"
            class:pga_small_grid={!showText}>
      <div class="pga_title_and_pic grow flex flex-col justify-around items-stretch">
        <div class="h-full flex flex-col justify-between items-stretch">
          <div class="preview_title grow flex justify-around items-center"
               style:font-size={preview_title_font_size}>
            <div class="h-full flex justify-between items-center">
              {i18n.t("about", langVal)}
            </div>
          </div>
          <div class="grow flex justify-around self-center">
            <div class="h-full flex justify-between items-center">
              <div class="ai_me_container magic_square_img">
                <AiMe imgSideLength={imgSideLength}/>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="pga_text flex pl-5 pr-5 mb-2 justify-around items-center overflow-y-scroll">
        <ul class="preview_list h-5/6 flex flex-col justify-around items-stretch"
            style:font-size={preview_text_font_size}>
          <li>
            {i18n.t("about_1", langVal)}
          </li>
          <li>
            {i18n.t("about_2", langVal)}
          </li>
        </ul>
      </div>
    </button>
    <button on:click={() => handlePreviewClick(SiteSection.publicSquare)}
            class="preview md:flex md:flex-col md:justify-evenly md:items-center md:h-5/6"
            class:pga_small_grid={!showText}>
      <div class="pga_title_and_pic grow flex flex-col justify-around items-stretch">
        <div class="h-full flex flex-col justify-around items-stretch">
          <div class="preview_title flex justify-around items-center"
               style:font-size={preview_title_font_size}>
            {i18n.t("magicSquare", langVal)}
          </div>
          <div class="h-full grow flex justify-around items-center">
            <img class="magic_square_img ai_me"
                 src={magicSquareExampleGif}
                 style:height={imgSideLength}
                 style:width={imgSideLength}
                 alt="Magic Square Example"/>
          </div>
        </div>
      </div>
      <div class="pga_text flex pl-5 pr-5 mb-2 justify-around items-center overflow-y-scroll">
        <ul class="preview_list h-5/6 flex flex-col justify-around items-stretch"
            style:font-size={preview_text_font_size}>
          <li>
            {i18n.t("magicSquare_1", langVal)}
          </li>
          <li>
            {i18n.t("magicSquare_2", langVal)}
          </li>
          <li>
            {i18n.t("magicSquare_3", langVal)}
          </li>
        </ul>
      </div>
    </button>
    <button on:click={() => handlePreviewClick(SiteSection.giveMeASine)}
            class="preview md:flex md:flex-col md:justify-evenly md:items-center md:h-5/6"
            class:pga_small_grid={!showText}>
      <div class="pga_title_and_pic grow flex flex-col justify-around items-stretch">
        <div class="h-full flex flex-col justify-around items-stretch">
          <div class="preview_title flex justify-around items-center"
               style:font-size={preview_title_font_size}>
            {i18n.t("giveMeASine", langVal)}
          </div>
          <div class="h-full flex justify-around items-center">
            <img class="magic_square_img"
                 src={giveMeASineExampleGif}
                 style:height={imgSideLength}
                 style:width={imgSideLength}
                 alt="Give Me A Sine Example"/>
          </div>
        </div>
      </div>
      <div class="pga_text flex pl-5 pr-5 mb-2 justify-around items-center overflow-y-scroll">
        <ul class="preview_list h-5/6 flex flex-col justify-around items-stretch"
            style:font-size={preview_text_font_size}>
          <li>
            {i18n.t("giveMeASine_1", langVal)}
          </li>
          <li>
            {i18n.t("giveMeASine_2", langVal)}
          </li>
        </ul>
      </div>
    </button>
  </div>
</body>

<style lang="sass">
  @use "./styles/color"
  @use "./styles/text"

  .home
    &_title
      color: color.$blue-5
      font-size: text.$fs-l
      font-weight: text.$fw-l
      border-left: 5px double color.$blue-7
      border-right: 5px double color.$blue-7
      border-radius: 10px
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
    padding: 5px
    margin: 5px
    border-radius: 50% 

  .ai_me_container
    grid-template-areas: "img"

  .pga
    &_title_and_pic
      grid-area: title_and_pic
    &_text
      grid-area: text
    
    &_small_grid
      height: 100%
      display: grid
      grid-template-rows: 1fr 1fr
      grid-template-columns: 1fr 1fr
      grid-template-areas: "title_and_pic text" "title_and_pic text"
      
 
  .preview
    border-radius: 5px
    color: color.$green-4
    font-weight: text.$fw-l
    font-size: text.$fs-s
    background: color.$black-blue-grad
    min-height: 300px
    /* border-top: 5px double color.$blue-7 */
    /* border-bottom: 5px double color.$blue-7 */
    border-right: 0px solid black
    border-left: 0px solid black
    height: 100%
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
      width: 100%      
      
</style>
