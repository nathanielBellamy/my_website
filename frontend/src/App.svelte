<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { push } from "svelte-spa-router"
  import { Footer, FooterCopyright, FooterLinkGroup, FooterLink, FooterBrand, FooterIcon } from 'flowbite-svelte'
  import Device from 'svelte-device-info'
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Link from "./lib/Link.svelte"
  import Language from "./lib/Language.svelte"
  import { I18n, Lang } from "./I18n"
  import { lang } from "./stores/lang"
  import { intoSiteSection, intoUrl, SiteSection, siteSection } from "./stores/siteSection"
  import Navbar from './lib/Navbar.svelte'
  import SocialLinks from './lib/SocialLinks.svelte'
  import { ViteMode } from './ViteMode'

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

  const routes: { [key: string]: any } = {
    '/': wrap({
      asyncComponent: () => import('./Home.svelte')
    }),
    '/about': wrap({
      asyncComponent: () => import('./About.svelte')
    }),
    '/give-me-a-sine': wrap({
      asyncComponent: () => import('./GiveMeASine.svelte')
    }),
    '/magic-square': wrap({
      asyncComponent: () => import('./MagicSquare/Container.svelte')
    }),
    '/public-square': wrap({
      asyncComponent: () => import('./PublicSquare/Container.svelte')
    }),
    '/system-diagram': wrap({
      asyncComponent: () => import('./SystemDiagram.svelte')
    })
  }

  onMount(() => {
    let storageSiteSection: SiteSection = intoSiteSection(localStorage.getItem('ns_site_section'))
    siteSection.update((_:SiteSection) => storageSiteSection)
  })

  onDestroy(() => {
    unsubSiteSection()
    unsubSmallScreen()
    unsubTouchScreen()
  })
</script>

<svelte:window bind:innerWidth />

<Navbar />

<main class="rounded-md flex flex-col justify-start overflow-hidden">
  <Router {routes}/>
</main>

<Footer class="w-full rounded-none flex justify-between items-center pt-2 pb-2 bg-black">
  <div class="grow h-full pl-2 pr-2 flex items-center">
    <SocialLinks />
  </div>
  <div class="grow pt-1 flex flex-row-reverse items-center">
    <Language />
  </div>
</Footer>
