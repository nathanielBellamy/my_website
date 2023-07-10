<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { push } from "svelte-spa-router"
  import { Button as DropdownButton, Dropdown, DropdownItem, DropdownHeader, Chevron } from 'flowbite-svelte'
  import Device from 'svelte-device-info'
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Link from "./lib/Link.svelte"
  import Language from "./lib/Language.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from "./stores/lang"
  import { intoSiteSection, intoUrl, SiteSection, siteSection } from "./stores/siteSection"
  import { touchScreen } from './stores/touchScreen'
  	import { Alert } from 'flowbite-svelte';
    import Home from './Home.svelte';

  let touchScreenVal: boolean
  const unsubTouchScreen = touchScreen.subscribe((val: boolean) => touchScreenVal = val)
  touchScreen.update((_: boolean) => isTouchScreen())

  function isTouchScreen(): boolean {
    return Device.isPhone || Device.isTablet || Device.isLegacyTouchDevice
  }

  let siteSectionVal: SiteSection
  const unsubSiteSection = siteSection.subscribe((val: SiteSection) => siteSectionVal = val)
  
  let i18n = new I18n("app")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  const routes: { [key: string]: any } = {
    '/': wrap({
      asyncComponent: () => import('./Home.svelte')
    }),
    '/about': wrap({
      asyncComponent: () => import('./About.svelte')
    }),
    '/give_me_a_sine': wrap({
      asyncComponent: () => import('./GiveMeASine.svelte')
    }),
    '/magic_square': wrap({
      asyncComponent: () => import('./MagicSquare/Container.svelte')
    })
  }

  onMount(() => {
    let storageSiteSection: SiteSection = intoSiteSection(localStorage.getItem('ns_site_section'))
    siteSection.update((_:SiteSection) => storageSiteSection)
  })

  onDestroy(() => {
    unsubLang()
    unsubSiteSection()
    unsubTouchScreen()
  })

  function handleDropdownClick(s: SiteSection) {
    push(intoUrl(s))
  }
</script>

<nav class="nav_bar flex items-center gap-2 pt-2 pb-2">
  <!-- <Dropdown> -->
    <!-- <DropdownHeader> -->

  <DropdownButton color="none"
                  size='xs'>
    <div class="dropdown_icon flex justify-around items-center pb-1">
      ☰
    </div>
  </DropdownButton>
  <Dropdown placement="right">
    <DropdownItem class="w-11/12 flex items-center"
                  on:click={() => handleDropdownClick(SiteSection.home)}>
      {i18n.t("nav/home", langVal)} 
    </DropdownItem>
    <DropdownItem class="w-11/12 flex items-center"
                  on:click={() => handleDropdownClick(SiteSection.about)}>
      {i18n.t("nav/about", langVal)}
    </DropdownItem>
    <DropdownItem class="w-11/12 flex items-center"
                  on:click={() => handleDropdownClick(SiteSection.magicSquare)}>
      {i18n.t("nav/magicSquare", langVal)}
    </DropdownItem>
    <DropdownItem class="w-11/12 flex items-center"
                  on:click={() => handleDropdownClick(SiteSection.giveMeASine)}>
      {i18n.t("nav/giveMeASine", langVal)}
    </DropdownItem>
  </Dropdown>
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
  <div class="city">
    PORTLAND, OR
  </div>
  <Language />
</footer>

<style lang="sass">
  @use "./styles/color"
  @use "./styles/text"

  .dropdown
    &_icon
      font-size: text.$fs-ml
      font-weight: text.$fw-l
      color: color.$blue-4
  
  

  .nav_bar
    width: 100%
  /* .curr_section */
  /*   color: color.$blue-7 */
  /*   font-size: text.$fs-l */
  /*   font-weight: text.$fw-l */

  .city
    font-weight: text.$fw-l
</style>
