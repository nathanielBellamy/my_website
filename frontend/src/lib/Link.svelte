<script lang="ts">
  import { onDestroy } from 'svelte'
  import { link } from "svelte-spa-router"
  import { siteSection, SiteSection } from "../stores/siteSection"

  let siteSectionVal: SiteSection
  const unsubSiteSection = siteSection.subscribe(val => siteSectionVal = val)

  function updateSiteSection(s: SiteSection) {
    siteSection.update((_: SiteSection) => s)
  }

  function setSiteSection(href: string){
    var newSection: SiteSection = SiteSection.home

    switch (href) {
      case "/about":
        newSection = SiteSection.about
        break
      case "/magic_square":
        newSection = SiteSection.magicSquare
        break
      case "/give_me_a_sine":
        newSection = SiteSection.giveMeASine
        break
      case "/":
      default:
        break
    }

    localStorage.setItem('ns_site_section', newSection)
    updateSiteSection(newSection)
  }
  export let sameOrigin:boolean = true
  export let href: string = "/"
  export let className: string = ""
  export let title: string = "Home"
  export let onClick: any = (e:any) => {e.stopPropagation()}

  onDestroy(() => {
    unsubSiteSection()
  })
</script>

{#if sameOrigin}
  <a href={href}
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
  <a href={href}
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
