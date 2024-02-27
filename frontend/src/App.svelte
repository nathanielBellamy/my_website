<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Footer } from 'flowbite-svelte'
  import Device from 'svelte-device-info'
  import Router from "svelte-spa-router"
  import {wrap} from 'svelte-spa-router/wrap'
  import Language from "./lib/Language.svelte"
  import Navbar from './lib/Navbar.svelte'
  import SocialLinks from './lib/SocialLinks.svelte'

  import { smallScreen } from './stores/smallScreen'
  const unsubSmallScreen = smallScreen.subscribe((_: boolean | null) => {})

  import { touchScreen } from './stores/touchScreen'
  const unsubTouchScreen = touchScreen.subscribe((_: boolean) => {})
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

  const routes: { [key: string]: any } = {
    '/': wrap({
      asyncComponent: () => import('./Home.svelte')
    }),
    '/repos': wrap({
      asyncComponent: () => import('./Repos.svelte')
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

  onDestroy(() => {
    unsubSmallScreen()
    unsubTouchScreen()
  })
</script>

<svelte:window bind:innerWidth />

<Navbar />

<main class="main_container rounded-md flex flex-col justify-start overflow-hidden">
  <Router {routes}/>
</main>

<Footer class="w-full rounded-none flex justify-between items-center pt-2 pb-2 bg-black"
        data-testid="footer">
  <div class="grow h-full pl-2 pr-2 flex items-center">
    <SocialLinks />
  </div>
  <div class="grow pt-1 flex flex-row-reverse items-center">
    <Language />
  </div>
</Footer>

<style lang="sass">
  .main_container
    height: calc(100vh - 100px)
</style>
