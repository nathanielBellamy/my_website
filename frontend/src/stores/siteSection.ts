import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export enum SiteSection {
  about = "about",
  giveMeASine = "giveMeASine",
  home = "home",
  magicSquare = "magicSquare",
  none = "none",
  publicSquare = "publicSquare",
  systemDiagram = "systemDiagram",
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
    case "publicSquare":
      return SiteSection.publicSquare
    default:
      return SiteSection.none
  }
}

export function intoUrl(s: SiteSection) {
  switch (s) {
    case SiteSection.about:
      return '/v1/about'
    case SiteSection.home:
      return '/v1/'
    case SiteSection.giveMeASine:
      return '/v1/give-me-a-sine'
    case SiteSection.magicSquare:
      return '/v1/magic-square'
    case SiteSection.publicSquare:
      return '/v1/public-square'
    case SiteSection.none:
    default:
      return '/v1/'
  }
}

export const siteSection: Writable<SiteSection> = writable(SiteSection.home)
