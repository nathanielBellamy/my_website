<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import { watchResize } from "svelte-watch-resize"
  import Main from "./Main.svelte"

  // INIT Prev Settings
  import WarningModal from "../MagicSquare/WarningModal.svelte"

  let magicSquareInstance: number = 0
  $: magicSquareInstance
  let sideLength: number = 0
  let sideLengthReady: boolean = false

  async function handleResize() {
    incrementMagicSquareInstance()
    let element = document.getElementById("magic_square_pub_container")
    sideLength = Math.floor(Math.min(element.offsetWidth, element.offsetHeight) / 1.3) - 25
    sideLengthReady = true
  }

  onDestroy(() => {
    window.removeEventListener('resize', handleResize)
  })

  let hasAcceptedWarning: boolean = false

  onMount(() => {
    hasAcceptedWarning = !!localStorage.getItem("magic_square_has_accepted_warning")
    incrementMagicSquareInstance()
    window.addEventListener('resize', handleResize)
  })

  function incrementMagicSquareInstance() {
    magicSquareInstance += 1
  }
</script>

<div id="magic_square_pub_container"
     class="grow w-full h-full flex justify-around items-center"
     use:watchResize={handleResize}>
  {#if !hasAcceptedWarning}
    <WarningModal bind:hasAccepted={hasAcceptedWarning}/>
  {:else}
    {#key magicSquareInstance}
      {#if sideLengthReady}
        <Main sideLength={sideLength}/>
      {/if}
    {/key}
  {/if}
</div>

<!-- <div id="magic_square_container" -->
<!--      class="magic_square_container" -->
<!--      use:watchResize={handleResize}> -->
<!--   {#if !hasAcceptedWarning} -->
<!--     <WarningModal bind:hasAccepted={hasAcceptedWarning}/> -->
<!--   {:else} -->
<!--     {#key magicSquareInstance} -->
<!--       {#if dataReady} -->
<!--         <MagicSquarePub sideLength={sideLength}/> -->
<!--       {/if} -->
<!--     {/key} -->
<!--   {/if} -->
<!-- </div> -->

<style lang="sass">
  .magic_square_container
    height: 100%
    width: 100%
    overflow: hidden
</style>
