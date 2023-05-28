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

<nav class="nav_bar flex items-stretch">
  <a href="/" 
     use:link
     class="nav_link"> 
    Home 
  </a>
  <a href="/about_me" 
     use:link
     class="nav_link"> 
    About Me 
  </a>
  <a href="/give_me_a_sine" 
     use:link
     class="nav_link"> 
    Give Me A Sine 
  </a>
  <a href="/magic_square" 
     use:link
     class="nav_link"> 
    Magic Square 
  </a>
</nav>

<main class="rounded-md flex flex-col justify-start">
  <Router {routes}/>
</main>

<footer>
  nbschieber@gmail.com
</footer>

<style lang="sass">
  @use "./styles/color"

  .nav_bar a:not(:last-child)
    border-right: 5px double color.$blue-7
  
  .main
    width: 100vw
    height: calc(100vh - 50px)

  .nav_link
    flex-grow: 1
    min-width: 150px
    max-width: 150px
</style>
