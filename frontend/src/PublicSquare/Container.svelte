<script lang="ts">
  import { onMount } from "svelte"
  import Main from "./Main.svelte"
  import WarningModal from "../MagicSquare/WarningModal.svelte"
  import InfoGate from "./InfoGate.svelte"
  import { Spinner } from "flowbite-svelte"

  import { smallScreen } from '../stores/smallScreen'
  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean | null) => smallScreenVal = val)
 
  let hasPassedGate: boolean = false
  let hasAcceptedWarning: boolean = false

  let counter: number = 2

  function waitOnLoad() {
    timeout()
  }

  function timeout() {
    if (--counter > 0)
      return setTimeout(timeout, 750)
  }

  waitOnLoad()

  onMount(() => {
    hasAcceptedWarning = !!localStorage.getItem("magic_square_has_accepted_warning")
  })

</script>

<body id="magic_square_pub_container"
      class="w-full h-full overflow-hidden overscroll-none"
      class:mb-10={smallScreenVal}>
  {#if counter > 0}
    <div class="h-full w-full flex justify-center items-center gap-4">
      <div class="info_gate_loading font-mono text-4xl md:text-6xl w-fit flex justify-around items-center"> 
        Loading...
      </div>
      <Spinner color="purple" />
    </div>
  {:else}
    {#if !hasPassedGate}
      <InfoGate bind:hasPassedGate={hasPassedGate}/>
    {:else if !hasAcceptedWarning}
      <WarningModal bind:hasAccepted={hasAcceptedWarning}/>
    {:else}
      <Main />
    {/if}
  {/if}
</body>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .info_gate_loading
    color: color.$blue-7
</style>

