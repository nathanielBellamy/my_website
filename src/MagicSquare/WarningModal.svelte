<script lang="ts">
  import { push } from "svelte-spa-router"
  import { siteSection, SiteSection } from "../stores/siteSection"

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
    Epilepsy Warning
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <p class="content">
      Please be aware that this software contains sequences of flashing lights which may trigger seizures for people with photosensitive epilepsy. Viewer discretion is advised.
    </p>
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <p class="content">
      If you or anyone in your household has an epileptic condition, please consult a medical professional before using this software. If you experience dizziness, altered vision, eye or muscle twitches, loss of awareness, disorientation, or any involuntary movement or convulsion while using the software, immediately discontinue use and consult a medical professional.
    </p>
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <p class="content">
      Please ensure you enjoy Magic Square in a well-lit environment and take frequent breaks. Your safety is important to us.
    </p>
  </div>
  <div class="grow w-full p-5 flex justify-around items-stretch">
    <div class="grow grid grid-cols-2 grid-rows-1">
      <button on:click={handleGoBack}
              class="green">
        Go To Home
      </button>
      <button on:click={handleAccept}
              class="red">
        Accept & Continue
      </button>
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .green
    color: color.$green-4
    border-color: color.$green-4

  .red
    color: color.$red-4
    border-color: color.$red-4

  .warning_main
    overflow-y: scroll

  .title
    font-size: text.$fs-l
    color: color.$red-4
    font-weight: text.$fw-xl

  .content
    text-align: left
    font-size: text.$fs-m
    font-weight: text.$fw-l
    color: color.$blue-4
    width: 90%
</style>
