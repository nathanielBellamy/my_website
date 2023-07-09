import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export enum SiteSection {
  about = "about",
  giveMeASine = "giveMeASine",
  home = "home",
  magicSquare = "magicSquare",
  none = "none"
}

export function intoSiteSection(s: string | null | undefined): SiteSection {
  switch(s) {
    case "about":
      return SiteSection.about
    case "giveMeASine":
      return SiteSection.giveMeASine
    case "home":
      return SiteSection.home
    case "magicSquare":
      return SiteSection.magicSquare
    default:
      return SiteSection.none
  }
}

export function intoUrl(s: SiteSection) {
  switch (s) {
    case SiteSection.about:
      return '/about'
    case SiteSection.home:
      return '/'
    case SiteSection.giveMeASine:
      return '/give_me_a_sine'
    case SiteSection.magicSquare:
      return '/magic_square'
    case SiteSection.none:
    default:
      return '/'
  }
}

export const siteSection: Writable<SiteSection> = writable(SiteSection.home)
