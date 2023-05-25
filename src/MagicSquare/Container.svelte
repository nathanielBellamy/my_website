<script lang="ts">
  import { watchResize } from "svelte-watch-resize"
  import MagicSquare from "./Main.svelte"
  
  let magicSquareInstance = 0

  // To handle resize
  // we simply alternate between two identical instances
  // on each resize, one is destroyed and one is mounted
  // the internals of MagicSquare handle getting the height/width onMount
  // TODO: persist settings during resize
  // - this likely means storing a settings object here
  // - this component becomes the "env" for MagicSquare
  // - while MagicSquare holds the engine itself
  const handleResize = async () => {
    magicSquareInstance = (magicSquareInstance + 1) % 2
  }
</script>

<div class="magic_square_container"
     use:watchResize={handleResize}>
    <MagicSquare />
</div>

<style lang="sass">
  .magic_square_container
    height: 100%
    width: 100%
</style>
