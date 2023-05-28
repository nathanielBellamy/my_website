<script lang="ts">
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Link from "./lib/Link.svelte";

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
  <Link href="/" 
        className="nav_link"
        title="Home"/> 
  <Link href="/about_me" 
        className="nav_link" 
        title="About Me" />
  <Link href="/give_me_a_sine" 
        className="nav_link" 
        title="Give Me A Sine"/>
  <Link href="/magic_square" 
        className="nav_link"
        title="Magic Square"/> 
</nav>

<main class="rounded-md flex flex-col justify-start">
  <Router {routes}/>
</main>

<footer>
  <div>
    github.com/nathanielBellamy
  </div>
  <div>
    nbschieber@gmail.com
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
