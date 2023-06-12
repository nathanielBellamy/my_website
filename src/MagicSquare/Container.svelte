<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import { watchResize } from "svelte-watch-resize";
  import Main from "./Main.svelte"
  
  let magicSquareInstance: number = 0
  let sideLength: number = 0

  async function handleResize() {
    incrementMagicSquareInstance()
    let element = document.getElementById("magic_square_container")
    sideLength = Math.floor(Math.min(element.offsetWidth, element.offsetHeight) / 1.3) - 25
  }

  onDestroy(() => {
    window.removeEventListener('resize', handleResize)
  })

  onMount(() => {
    window.addEventListener('resize', handleResize)
  })

  function incrementMagicSquareInstance() {
    magicSquareInstance += 1
  }
</script>

<div id="magic_square_container"
     class="magic_square_container"
     use:watchResize={handleResize}>
    {#key magicSquareInstance}
      <Main sideLength={sideLength}/>
    {/key}
</div>

<style lang="sass">
  .magic_square_container
    height: 100%
    width: 100%
    overflow: hidden
</style>
