<script lang="ts">
  import MagicSquare from "./MagicSquare.svelte"

  export let program: string = "magicSquare"
</script>

<div class="wasm_container">
  <script src="./src-rust/pkg/src_rust.js"></script>
  {#if program == "magicSquare"}
    <MagicSquare />
  {/if}
</div>


<style lang="sass">
  .wasm_container
    width: 100%
    height: 100%
</style>

