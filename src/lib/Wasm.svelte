<script lang="ts">
  import GiveMeASine from "./GiveMeASine.svelte"
  import MagicSquare from "./MagicSquare.svelte"

  export let program: string = "magicSquare"


</script>

<div class="wasm_container">
  {#if program == "giveMeASine"}
    <GiveMeASine />
  {:else if program == "magicSquare"}
    <MagicSquare />
  {/if}
</div>

<style lang="sass">
  .wasm_container
    width: 100%
    height: 100%
</style>

