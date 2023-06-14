<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import { watchResize } from "svelte-watch-resize"
  import type { StorageSettings } from './StorageSettings'
  import Main from "./Main.svelte"

  // INIT Prev Settings
  import { prevSettingsStore } from './PrevSettingsStore'
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


  onMount(() => {
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
    {#key magicSquareInstance}
      {#if dataReady}
        <Main bind:instance={magicSquareInstance}
              sideLength={sideLength}/>
      {/if}
    {/key}
</div>

<style lang="sass">
  .magic_square_container
    height: 100%
    width: 100%
    overflow: hidden
</style>
