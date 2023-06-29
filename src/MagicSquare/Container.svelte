<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import { watchResize } from "svelte-watch-resize"
  import type { StorageSettings } from './StorageSettings'
  import Main from "./Main.svelte"

  // INIT Prev Settings
  import { prevSettingsStore } from './PrevSettingsStore'
  import WarningModal from "./WarningModal.svelte"

  let prevSettingsStoreVal: StorageSettings
  $: prevSettingsStoreVal
  const unsubscribe = prevSettingsStore.subscribe(val => prevSettingsStoreVal = val)
  
  let magicSquareInstance: number = 0
  $: magicSquareInstance
  let sideLength: number = 0

  async function handleResize() {
    incrementMagicSquareInstance()
    prevSettingsStoreVal = prevSettingsStoreVal
    let element = document.getElementById("magic_square_container")
    sideLength = Math.floor(Math.min(element.offsetWidth, element.offsetHeight) / 1.3) - 25
  }

  onDestroy(() => {
    window.removeEventListener('resize', handleResize)
  })

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
      incrementMagicSquareInstance()
    }
    window.addEventListener('resize', handleResize)
    dataReady = true
  })

  onDestroy(unsubscribe)

  function incrementMagicSquareInstance() {
    magicSquareInstance += 1
  }
</script>

<div id="magic_square_container"
     class="magic_square_container"
     use:watchResize={handleResize}>
  {#if !hasAcceptedWarning}
    <WarningModal bind:hasAccepted={hasAcceptedWarning}/>
  {:else}
    {#key magicSquareInstance}
      {#if dataReady}
        <Main sideLength={sideLength}/>
      {/if}
    {/key}
  {/if}
</div>

<style lang="sass">
  .magic_square_container
    height: 100%
    width: 100%
    overflow: hidden
</style>
