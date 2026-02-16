import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import { OldSiteUrl } from '../lib/OldSiteUrl'

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

export function intoUrl(s: SiteSection): string {
  switch (s) {
    case SiteSection.about:
      return OldSiteUrl.About.split('#')[1]
    case SiteSection.home:
      return OldSiteUrl.Home.split('#')[1]
    case SiteSection.giveMeASine:
      return OldSiteUrl.GiveMeASine.split('#')[1]
    case SiteSection.magicSquare:
      return OldSiteUrl.MagicSquare.split('#')[1]
    case SiteSection.publicSquare:
      return OldSiteUrl.PublicSquare.split('#')[1]
    case SiteSection.none:
    default:
      return OldSiteUrl.Home.split('#')[1]
  }
}

export const siteSection: Writable<SiteSection> = writable(SiteSection.home)
