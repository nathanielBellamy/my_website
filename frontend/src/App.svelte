<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { push } from "svelte-spa-router"
  import { Button as DropdownButton, Dropdown, DropdownItem, Chevron } from 'flowbite-svelte'
  import { Footer, FooterCopyright, FooterLinkGroup, FooterLink, FooterBrand, FooterIcon } from 'flowbite-svelte'
  import Device from 'svelte-device-info'
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Link from "./lib/Link.svelte"
  import Language from "./lib/Language.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from "./stores/lang"
  import { intoSiteSection, intoUrl, SiteSection, siteSection } from "./stores/siteSection"
  import githubLogo from './assets/github_logo.png'
  import linkedInLogo from './assets/linked_in_logo.png'
  import gmailLogo from './assets/gmail_logo.png'


  import { smallScreen } from './stores/smallScreen'
  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean | null) => smallScreenVal = val)

  import { touchScreen } from './stores/touchScreen'
  let touchScreenVal: boolean
  const unsubTouchScreen = touchScreen.subscribe((val: boolean) => touchScreenVal = val)
  touchScreen.update((_: boolean) => isTouchScreen())

  function isTouchScreen(): boolean {
    return Device.isPhone || Device.isTablet || Device.isLegacyTouchDevice
  }

  let innerWidth: number
  $: if (innerWidth < 1000) {
    smallScreen.update((_: boolean | null) => true)
  }
  $: if (innerWidth > 1000) {
    smallScreen.update((_: boolean | null) => false)
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
    }),
    '/public-square': wrap({
      asyncComponent: () => import('./PublicSquare/Container.svelte')
    })
  }

  onMount(() => {
    let storageSiteSection: SiteSection = intoSiteSection(localStorage.getItem('ns_site_section'))
    siteSection.update((_:SiteSection) => storageSiteSection)
  })

  onDestroy(() => {
    unsubLang()
    unsubSiteSection()
    unsubSmallScreen()
    unsubTouchScreen()
  })

  function handleDropdownClick(s: SiteSection) {
    push(intoUrl(s))
  }
</script>

<svelte:window bind:innerWidth />

<nav class="nav_bar flex justify-between items-center gap-2 pt-2 pb-2">
  <DropdownButton id="siteSectionDropdown"
                  class="border-transparent"
                  color="none"
                  size='xs'>
    <div class="dropdown_icon pb-1 flex justify-around items-center">
      ☰
    </div>
  </DropdownButton>
  <Dropdown triggeredBy="#siteSectionDropdown"
            placement={touchScreenVal ? "bottom" : "right"}
            ulClass="pt-2 pb-2 rounded-lg bg-zinc-800 text-blue-200">
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold"
                  on:click={() => handleDropdownClick(SiteSection.home)}>
      {i18n.t("nav/home", langVal)} 
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold"
                  on:click={() => handleDropdownClick(SiteSection.publicSquare)}>
      Public square
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold"
                  on:click={() => handleDropdownClick(SiteSection.magicSquare)}>
      {i18n.t("nav/magicSquare", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold"
                  on:click={() => handleDropdownClick(SiteSection.giveMeASine)}>
      {i18n.t("nav/giveMeASine", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold"
                  on:click={() => handleDropdownClick(SiteSection.about)}>
      {i18n.t("nav/about", langVal)}
    </DropdownItem>
  </Dropdown>
  <DropdownButton id="contactInfo"
                  class="border-transparent"
                  color="none"
                  size='xs'>
    <div class="dropdown_personal flex justify-around items-center">
      Nate Schieber
      <div class="pb-1">
        <Chevron />
      </div>
    </div>
  </DropdownButton>
  <Dropdown triggeredBy="#contactInfo"
            placement="bottom"
            ulClass="pt-3 pb-3 rounded-lg bg-zinc-800 text-blue-200">
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold">
      <Link href="mailto:nbschieber@gmail.com"
            title="nbschieber@gmail.com"
            sameOrigin={false}/>
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold">
      <Link href="https://linkedin.com/in/nateschieber"
            title="in/nateschieber"
            sameOrigin={false}/>
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold">
      <Link href="https://github.com/nathanielBellamy"
            title="github.com/nathanielBellamy"
            sameOrigin={false}/>
    </DropdownItem>
    <DropdownItem class="text-blue-200 hover:bg-transparent w-11/12 flex items-center font-bold">
      <Link href="https://www.travelportland.com"
            title="PORTLAND, OR"
            sameOrigin={false}/>
    </DropdownItem>
  </Dropdown>
</nav>

<main class="rounded-md flex flex-col justify-start overflow-hidden">
  <Router {routes}/>
</main>

<Footer class="w-full rounded-none flex justify-between items-center pt-2 pb-2 bg-black">
  <div class="grow h-full pl-2 pr-2 flex items-center">
    <div class="h-full w-fit flex justify-between items-center gap-4">
      <FooterIcon href="https://github.com/nathanielBellamy" 
                  class="text-gray-400 hover:text-gray-900 h-10 flex justify-around items-center">
        <img src={githubLogo}
             title="GitHub"
             class="invert"
             style:width="30px"
             style:height="30px"
             alt="GitHub"/>
      </FooterIcon>
      <FooterIcon href="https://www.linkedin.com/in/nateschieber/" 
                  class="text-gray-400 hover:text-gray-900 h-10 flex justify-around items-center">
        <img src={linkedInLogo}
             title="LinkedIn"
             style:width="30px"
             style:height="30px"
             alt="LinkedIn"/>
      </FooterIcon>
      <FooterIcon href="mailto:nbschieber@gmail.com" 
                  class="w-fit text-gray-400 hover:text-gray-900 h-10 flex justify-around items-center">
        <img src={gmailLogo}
             title="gmail"
             class="invert"
             style:width="30px"
             style:height="30px"
             alt="gmail"/>
      </FooterIcon>
    </div>
  </div>
  <div class="grow pt-1 flex flex-row-reverse items-center">
    <Language />
  </div>
</Footer>

<style lang="sass">
  @use "./styles/color"
  @use "./styles/text"
  @use "./styles/font"

  .dropdown
    &_icon
      font-size: text.$fs-l
      font-weight: text.$fw-l
      color: color.$blue-4

    &_personal
      font-size: text.$fs-m
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
