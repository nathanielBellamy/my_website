<script lang="ts">
  import { afterUpdate, onDestroy, onMount } from "svelte"
  import { watchResize } from "svelte-watch-resize"
  import MagicSquare from "./Main.svelte"
  
  let magicSquareInstance: number = 0
  let destroyChild = false
  let sideLength: number = 0

  async function handleResize() {
    magicSquareInstance += 1
    let element = document.getElementById("magic_square_container")
    sideLength = Math.floor(Math.min(element.offsetWidth, element.offsetHeight) / 1.3) - 25
  }

  afterUpdate(() => {
    return () => {
      window.removeEventListener('resize', handleResize)
    }
  })

  onMount(() => {
    window.addEventListener('resize', handleResize)
  })

  onDestroy(() => {
    magicSquareInstance += 1
    destroyChild = false
  })
</script>


<div id="magic_square_container"
     class="magic_square_container"
     use:watchResize={handleResize}>
    {#key magicSquareInstance}
      {#if !destroyChild}
        <MagicSquare sideLength={sideLength}/>
      {/if}
    {/key}
</div>

<style lang="sass">
  .magic_square_container
    height: 100%
    width: 100%
    overflow: hidden
</style>
