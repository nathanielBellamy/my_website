<script lang="ts">
  import { link } from "svelte-spa-router"
  import { SiteSection } from "../stores/siteSection"
  import { OldSiteUrl } from "./OldSiteUrl"

  export let sameOrigin:boolean = true
  export let href: string = OldSiteUrl.Home
  export let className: string = ""
  export let title: string = "Home"
  export let onClick: any = (e:any) => {e.stopPropagation()}

  function setSiteSection(url: string){
    var newSection: SiteSection = SiteSection.home

    switch (url) {
      case OldSiteUrl.About:
        newSection = SiteSection.about
        break
      case OldSiteUrl.MagicSquare:
        newSection = SiteSection.magicSquare
        break
      case OldSiteUrl.GiveMeASine:
        newSection = SiteSection.giveMeASine
        break
      case OldSiteUrl.Home:
      default:
        break
    }

    localStorage.setItem('ns_site_section', newSection)
  }
</script>

{#if sameOrigin}
  <a href={href.toString()}
     use:link
     tabindex="0"
     on:click={(e) => {
      setSiteSection(href)
      onClick(e)
     }}
     class={`link ${className} rounded-md`}>
    <button class="link_button hover:bg-slate-700"
            tabindex="-1">
      {title}
    </button>
  </a>
{:else}
  <a href={href.toString()}
     tabindex="0"
     target="_blank"
     on:click={(e) => {
      onClick(e)
     }}
     class={`link ${className} rounded-md`}>
    <button class="link_button hover:bg-slate-700"
            tabindex="-1">
      {title}
    </button>
  </a>
{/if}

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .link
    font-weight: text.$fw-l
    &_button
      margin: 0
      border-width: 0px
      color: color.$blue-4
</style>
