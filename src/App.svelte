<script lang="ts">
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Link from "./lib/Link.svelte"
  import Language from "./lib/Language.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from "./stores/lang"
  
  let i18n = new I18n
  let langVal: Lang

  lang.subscribe( val => langVal = val)

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

  let currentSection: string = 'home'
  function handleClick (newSection: string) {
    currentSection = newSection
  }
</script>

<nav class="nav_bar flex flex-row justify-between items-stretch">
  <div class="links flex justify-between items-stretch">
    <Link href="/" 
          title={i18n.t("app/nav/home", langVal)}
          onClick={() => handleClick("home")}/> 
    <Link href="/about" 
          title={i18n.t("app/nav/about", langVal)}
          onClick={() => handleClick("about")}/>
    <Link href="/magic_square" 
          title={i18n.t("app/nav/magicSquare", langVal)}
          onClick={() => handleClick("magicSquare")}/> 
    <Link href="/give_me_a_sine" 
          title={i18n.t("app/nav/giveMeASine", langVal)}
          onClick={() => handleClick("giveMeASine")}/>
  </div>
  <div class="curr_section hidden md:block">
    {i18n.t(`app/nav/${currentSection}`, langVal)}
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
  
  .nav_bar
    min-width: 500px
    overflow-x: scroll
  .curr_section
    color: color.$blue-7
    font-size: text.$fs-l
    font-weight: text.$fw-l
    margin-top: -10px
</style>
