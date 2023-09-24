<script lang="ts">
  import { onDestroy } from 'svelte'
  import { push } from "svelte-spa-router"
  import { SquareType } from '../stores/currSquare'

  import { I18n, Lang } from "../I18n"
  import { lang } from '../stores/lang'
  const i18n = new I18n("magicSquare/warning")
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)

  export let hasAccepted: boolean = false
  export let squareType: SquareType

  function handleAccept() {
    localStorage.setItem('magic_square_has_accepted_warning', "true")
    hasAccepted = true
  }

  function handleGoBack() {
    push("/")
  }

  $: body_3 = body3(langVal)

  function body3(lv: Lang): String {
    switch (squareType) {
      case SquareType.magic:
        return i18n.t('body_3_ms', lv)
      case SquareType.public:
        return i18n.t('body_3_ps', lv)
      case SquareType.none:
        return ""
    }
  }

  onDestroy(unsubLang)
</script>

<div class="warning_main h-full flex flex-col justify-between items-stretch"
     data-testid="epilepsy_warning">
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
      {body_3}
    </p>
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <div class="grow grid grid-cols-2 grid-rows-1">
      <button on:click={handleGoBack}
              class="green"
              data-testid="epilepsy_warning_go_home">
        {i18n.t('go_home', langVal)}
      </button>
      <button on:click={handleAccept}
              class="red"
              data-testid="epilepsy_warning_accept">
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
