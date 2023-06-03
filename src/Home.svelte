<script lang="ts">
  import Link from "./lib/Link.svelte"
  import MagicBanner from "./lib/MagicBanner.svelte"
</script>

<body class="p-5 flex flex-col justify-between items-stretch gap-2">
  <div class="flex items-center">
    <h1 class="home_title text-left pl-5">
      It's A Website!
    </h1>
    <ul class="home_intro_list text-left p-5 flex flex-col justify-between items-stretch">
      <li>
        I like to learn.
      </li>
      <li>
        I wanted to learn how to combine
        <p>
          <Link href="https://www.rust-lang.org/"
                title="Rust"
                sameOrigin={false}/>
          with 
          <Link href="https://svelte.dev/"
                title="Svelte"
                sameOrigin={false}/>
          &
          <Link href="https://www.typescriptlang.org/"
                title="Typescript"
                sameOrigin={false}/>
          using
        </p>
        <Link href="https://webassembly.org/"
              title="WebAssembly"
              sameOrigin={false}/>
        &
        <Link href="https://crates.io/crates/wasm-bindgen"
              title="wasm-bindgen."
              sameOrigin={false}/>
      </li>
      <li>
        So I built this website.
      </li>
    </ul>
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
