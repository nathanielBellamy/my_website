<script lang="ts">
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Link from "./lib/Link.svelte";

  const routes: { [key: string]: any } = {
    '/': wrap({
      asyncComponent: () => import('./Home.svelte')
    }),
    '/about': wrap({
      asyncComponent: () => import('./lib/About.svelte')
    }),
    '/give_me_a_sine': wrap({
      asyncComponent: () => import('./lib/GiveMeASine.svelte')
    }),
    '/magic_square': wrap({
      asyncComponent: () => import('./MagicSquare/Container.svelte')
    })
  }
</script>

<nav class="nav_bar flex items-stretch">
  <Link href="/" 
        className="nav_link"
        title="Home"/> 
  <Link href="/about" 
        className="nav_link" 
        title="About" />
  <div>
    Rust Wasm: 
  </div>
  <Link href="/magic_square" 
        className="nav_link"
        title="Magic Square"/> 
  <Link href="/give_me_a_sine" 
        className="nav_link" 
        title="Give Me A Sine"/>

</nav>

<main class="rounded-md flex flex-col justify-start">
  <Router {routes}/>
</main>

<footer>
  <div>
    <a href="https://github.com/nathanielBellamy">
      github.com/nathanielBellamy
    </a>
  </div>
  <div>
    <a href="mailto:nbschieber@gmail.com">
      nbschieber@gmail.com
    </a>
  </div>
  <div>
    PORTLAND, OR
  </div>
</footer>

<style lang="sass">
  @use "./styles/color"

  .nav_link
    flex-grow: 1
    min-width: 150px
    max-width: 150px
</style>
