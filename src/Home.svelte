<script lang="ts">
  import Link from "./lib/Link.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from './stores/lang'

  // INIT LANG BOILER PLATE
  const i18n = new I18n("home")
  let langVal: Lang
  lang.subscribe(val => langVal = val)
</script>

<body class="pl-5 pr-5 pb-5 flex flex-col justify-between items-stretch gap-2">
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
    What's Here?
  </div>
  <div class="grow pl-5 pr-5 pb-5 flex flex-col justify-between items-stretch">
    <div class="grow grid grid-cols-3 grid-rows-1">
      <div class="preview flex flex-col justify-between items-stretch">
        <h3>
          <Link href="/about" 
                title="About"
                sameOrigin={true}/>
        </h3>
        <div class="flex pl-5 pr-5 justify-around items-center">
          <ul class="preview_list">
            <li>
              An overview of what I do code-wise. 
            </li>
            <li>
              Please enjoy these AI-generated images of me.
            </li>
          </ul>
        </div>
      </div>
      <div class="preview flex flex-col justify-between items-stretch">
        <h3>
          <Link href="/magic_square" 
                title="Magic Square"
                sameOrigin={true}/>
        </h3>
        <div class="grow flex justify-around items-center">
          <img class="magic_square_img"
               src="/src/assets/magic_square_example.gif"
               alt="Magic Square Example"/>
        </div>
        <div class="flex pl-5 pr-5 justify-around items-center">
          <ul class="preview_list">
            <li>
              Extensive user control of WebAssembly memory.
            </li>
            <li>
              Implimenting
              <Link href="https://www.khronos.org/webgl/"
                    title="WebGL"
                    sameOrigin={false}/>
              +
              <Link href="https://rustwasm.github.io/wasm-bindgen/examples/webgl.html"
                    title="RustWasm"
                    sameOrigin={false}/>.
            </li>
            <li>
              Exploring concepts from
              <Link href="https://en.wikipedia.org/wiki/Modular_synthesizer"
                    title="modular synthesis"
                    sameOrigin={false}/>
              .
            </li>
          </ul>
        </div>
      </div>
      <div class="preview flex flex-col justify-between items-stretch">
        <h3>
          <Link href="/give_me_a_sine" 
                title="Give Me A Sine"
                sameOrigin={true}/>
        </h3>
        <div class="grow flex justify-around items-center">
          <img class="magic_square_img"
               src="/src/assets/give_me_a_sine_example.gif"
               alt="Give Me A Sine Example"/>
        </div>
        <div class="flex pl-5 pr-5 justify-around items-center">
          <ul class="preview_list">
            <li>
              Ascii-Graph Sine functions. 
            </li>
            <li>
              All graphs copy-and-paste-able. 
            </li>
            <li>
              A warm-up project for
              <Link href="/magic_square" 
                    title="Magic Square"
                    sameOrigin={true}/>.
            </li>
          </ul>
        </div>
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
      font-size: text.$fs-l
      &_dark
        color: color.$green-4
        font-weight: text.$fw-l
        font-size: text.$fs-l
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
    max-width: 250px
    min-height: 200px

  .preview
    border-radius: 5px
    margin: 0 5px 0 5px
    padding: 0 20px 0 20px
    color: color.$green-4
    font-weight: text.$fw-l
    font-size: text.$fs-s
    background: color.$black-blue-grad
    &_list
      padding: 0 10px 0 10px
      text-align: left
      list-style-type: square
    &_li
      
</style>
