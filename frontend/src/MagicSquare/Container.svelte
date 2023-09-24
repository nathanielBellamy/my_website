<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import type { StorageSettings } from './StorageSettings'
  import Main from "./Main.svelte"

  // INIT Prev Settings
  import WarningModal from "./WarningModal.svelte"
  import Loading from "../lib/Loading.svelte"
  import { SquareType } from "../stores/currSquare"

  import { prevSettingsStore } from './PrevSettingsStore'
  let prevSettingsStoreVal: StorageSettings
  $: prevSettingsStoreVal
  const unsubPrevSettings = prevSettingsStore.subscribe(val => prevSettingsStoreVal = val)

  let dataReady: boolean = false
  let hasAcceptedWarning: boolean = false

  let counter: number = 2

  function waitOnLoad() {
    timeout()
  }

  function timeout() {
    if (--counter > 0)
      return setTimeout(timeout, 1000)
  }

  onMount(() => {
    hasAcceptedWarning = !!localStorage.getItem("magic_square_has_accepted_warning")

    let ses = localStorage.getItem("magic_square_settings")
    if (ses) {
      const res = JSON.parse(ses)
      prevSettingsStore.update((_: StorageSettings): StorageSettings => {
        return res
      })
    }
    dataReady = true
  })

  onDestroy(unsubPrevSettings)

  waitOnLoad()
</script>

<body id="magic_square_container"
      class="magic_square_container overscroll-none overflow-y-scroll">
  {#if counter > 0}
    <Loading />
  {:else}
    {#if !hasAcceptedWarning}
      <WarningModal bind:hasAccepted={hasAcceptedWarning}
                    squareType={SquareType.magic}/>
    {:else}
        {#if dataReady}
          <Main />
        {/if}
    {/if}
  {/if}
</body>

<style lang="sass">
  @use "./../styles/color"
  .magic_square_container
    height: 100%
    width: 100%
    overflow: hidden
</style>
