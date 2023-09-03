<script lang="ts">
  import { onDestroy } from 'svelte'
  import { push } from "svelte-spa-router"
  import { siteSection, SiteSection } from "../stores/siteSection"

  import { I18n, Lang } from "../I18n"
  import { lang } from '../stores/lang'

  // INIT LANG BOILER PLATE
  const i18n = new I18n("magicSquare/warning")
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)
  onDestroy(unsubLang)

  export let hasAccepted: boolean = false

  function handleAccept() {
    localStorage.setItem('magic_square_has_accepted_warning', "true")
    hasAccepted = true
  }

  function handleGoBack() {
    siteSection.update((_:SiteSection) => SiteSection.home)
    push("/")
  }
</script>

<div class="warning_main h-full flex flex-col justify-between items-stretch">
  <div class="title p-5">
    {i18n.t('title', langVal)}
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <p class="content">
      {i18n.t('body_1', langVal)}
    </p>
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <p class="content">
      {i18n.t('body_2', langVal)}
    </p>
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <p class="content">
      {i18n.t('body_3', langVal)}
    </p>
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <div class="grow grid grid-cols-2 grid-rows-1">
      <button on:click={handleGoBack}
              class="green">
        {i18n.t('go_home', langVal)}
      </button>
      <button on:click={handleAccept}
              class="red">
        {i18n.t('accept_and_continue', langVal)}
      </button>
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .green
    color: color.$green-7
    border-color: color.$green-7

  .red
    color: color.$red-7
    border-color: color.$red-7

  .warning_main
    overflow-y: scroll

  .title
    font-size: text.$fs-l
    color: color.$red-7
    font-weight: text.$fw-xl

  .content
    text-align: left
    font-size: text.$fs-m
    font-weight: text.$fw-l
    color: color.$blue-4
    width: 90%
</style>
