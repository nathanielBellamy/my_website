<script lang="ts">
  import Link from "./lib/Link.svelte"
  import { I18n, Lang, toLang } from "./I18n"

  const i18n = new I18n
  let lang: Lang = toLang(localStorage.getItem('lang'))
  
  function handleLangSwitch() {
    const newLang = localStorage.getItem('lang')
    if (typeof newLang === 'string') {
      lang = toLang(newLang)
    }
  }
  window.addEventListener('lang', handleLangSwitch)
</script>

<body class="p-5 flex flex-col justify-between items-stretch gap-2">
  <div class="flex flex-col justify-between items-stretch md:flex-row md:justify-start md:items-center">
    <h1 class="home_title text-left pl-5">
      {i18n.t("home/title", lang)}
    </h1>
    <ul class="home_intro_list text-left p-5 flex flex-col justify-between items-stretch">
      <li>
        {i18n.t("home/intro/1", lang)}
      </li>
      <li>
        {i18n.t("home/intro/2", lang)}
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
          {i18n.t("home/intro/3", lang)}
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
        {i18n.t("home/intro/4", lang)}
      </li>
    </ul>
  </div>
  <div class="grow flex flex-col justify-between items-stretch">
    <div>
      {i18n.t("home/body/title", lang)}
    </div>
    <div class="grow flex justify-around items-center">
      <div>
        (TODO: cool-looking, clickable SVGs representing the different sections of the site)
      </div>
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

    &_intro_list
      color: color.$green-4
      font-weight: text.$fw-m

  .text_lines
    color: color.$blue-7
    text-align: left
    padding: 15px
    height: 100%
</style>
