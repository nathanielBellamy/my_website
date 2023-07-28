<script lang="ts">
  import { onMount } from "svelte"
  import Main from "./Main.svelte"

  import WarningModal from "../MagicSquare/WarningModal.svelte"
 
  let hasAcceptedWarning: boolean = false

  onMount(() => {
    hasAcceptedWarning = !!localStorage.getItem("magic_square_has_accepted_warning")
  })
</script>

<div id="magic_square_pub_container"
     class="w-full h-full overflow-hidden">
  {#if !hasAcceptedWarning}
    <WarningModal bind:hasAccepted={hasAcceptedWarning}/>
  {:else}
    <Main />
  {/if}
</div>

<style lang="sass">

</style>
