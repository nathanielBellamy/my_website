<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import type { StorageSettings } from './StorageSettings'
  import Main from "./Main.svelte"

  // INIT Prev Settings
  import { prevSettingsStore } from './PrevSettingsStore'
  import WarningModal from "./WarningModal.svelte"

  let prevSettingsStoreVal: StorageSettings
  $: prevSettingsStoreVal
  const unsubscribe = prevSettingsStore.subscribe(val => prevSettingsStoreVal = val)

  let dataReady: boolean = false
  let hasAcceptedWarning: boolean = false

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

  onDestroy(unsubscribe)
</script>

<div id="magic_square_container"
     class="magic_square_container">
  {#if !hasAcceptedWarning}
    <WarningModal bind:hasAccepted={hasAcceptedWarning}/>
  {:else}
      {#if dataReady}
        <Main />
      {/if}
  {/if}
</div>

<style lang="sass">
  .magic_square_container
    height: 100%
    width: 100%
    overflow: hidden
</style>
