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

  let currentSection: string = 'Home'
  function handleClick (newSection: string) {
    currentSection = newSection
  }
</script>

<nav class="nav_bar flex justify-between items-stretch">
  <div class="links flex justify-between items-stretch">
    <Link href="/" 
          className="nav_link"
          title="Home"
          onClick={() => handleClick('It\'s A Website')}/> 
    <Link href="/about" 
          className="nav_link" 
          title="About" 
          onClick={() => handleClick("About")}/>
    <Link href="/magic_square" 
          className="nav_link"
          title="Magic Square"
          onClick={() => handleClick("Magic Square")}/> 
    <Link href="/give_me_a_sine" 
          className="nav_link" 
          title="Give Me A Sine"
          onClick={() => handleClick("Give Me A Sine")}/>
  </div>
  <div class="curr_section">
    {currentSection}
  </div>
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
  @use "./styles/text"

  .curr_section
    color: color.$blue-4
    font-size: text.$fs-l
    font-weight: text.$fw-l
    margin-top: -10px
</style>
