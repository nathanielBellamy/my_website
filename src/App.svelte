<script lang="ts">
  import Router from "svelte-spa-router"
  import {link} from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'

  const routes: { [key: string]: any } = {
    '/': wrap({
      asyncComponent: () => import('./Home.svelte')
    }),
    '/about_me': wrap({
      asyncComponent: () => import('./lib/AboutMe.svelte')
    }),
    '/give_me_a_sine': wrap({
      asyncComponent: () => import('./lib/GiveMeASine.svelte')
    }),
    '/magic_square': wrap({
      asyncComponent: () => import('./MagicSquare/Container.svelte')
    })
  }
</script>

<nav>
  <a href="/" use:link> Home </a>
  <a href="/about_me" use:link> About Me </a>
  <a href="/give_me_a_sine" use:link> Give Me A Sine </a>
  <a href="/magic_square" use:link> Magic Square </a>
</nav>

<main class="main rounded-md flex flex-col justify-start">
  <Router {routes}/>
</main>

<style lang="sass">
  @use "./styles/color"
  
  .main
    width: 100vw
    height: calc(100vh - 50px)
</style>
