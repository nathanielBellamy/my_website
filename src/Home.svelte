<script lang="ts">
  import Link from "./lib/Link.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from './stores/lang'

  // INIT LANG BOILER PLATE
  const i18n = new I18n("home")
  let langVal: Lang
  lang.subscribe(val => langVal = val)
</script>

<body class="p-5 flex flex-col justify-between items-stretch gap-2">
  <div class="flex flex-col justify-between items-stretch md:flex-row md:justify-start md:items-center">
    <h1 class="home_title text-left pl-5">
      {i18n.t("title", langVal)}
    </h1>
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
  <div class="grow flex flex-col justify-between items-stretch">
    <div>
      {i18n.t("title", langVal)}
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
</style>
