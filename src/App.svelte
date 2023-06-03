<script lang="ts">
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Link from "./lib/Link.svelte";
  import Language from "./lib/Language.svelte";

  const routes: { [key: string]: any } = {
    '/': wrap({
      asyncComponent: () => import('./Home.svelte')
    }),
    '/about': wrap({
      asyncComponent: () => import('./lib/About.svelte')
    }),
    '/give_me_a_sine': wrap({
      asyncComponent: () => import('./GiveMeASine.svelte')
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
          title="Home"
          onClick={() => handleClick('It\'s A Website')}/> 
    <Link href="/about" 
          title="About" 
          onClick={() => handleClick("About")}/>
    <Link href="/magic_square" 
          title="Magic Square"
          onClick={() => handleClick("Magic Square")}/> 
    <Link href="/give_me_a_sine" 
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

<footer class="flex flex-col space-between items-stretch pt-2 pb-2 md:flex-row md:pb-0">
  <div class="grow">
    <Link href="https://github.com/nathanielBellamy"
          title="github.com/nathanielBellamy"
          sameOrigin={false}/>
  </div>
  <div class="grow">
    <Link href="mailto:nbschieber@gmail.com"
          title="nbschieber@gmail.com"
          sameOrigin={false}/>
  </div>
  <div>
    PORTLAND, OR
  </div>
  <Language />
</footer>

<style lang="sass">
  @use "./styles/color"
  @use "./styles/text"

  .curr_section
    color: color.$blue-7
    font-size: text.$fs-l
    font-weight: text.$fw-l
    margin-top: -10px
</style>
