<script lang="ts">
  import { onMount } from "svelte"
  import Main from "./Main.svelte"
  import WarningModal from "../MagicSquare/WarningModal.svelte"
  import InfoGate from "./InfoGate.svelte"
  import Loading from "../lib/Loading.svelte";
  import { SquareType } from "../stores/currSquare";
 
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
      class="w-full h-full overflow-hidden overscroll-none">
  {#if counter > 0}
    <Loading />
  {:else}
    {#if !hasPassedGate}
      <InfoGate bind:hasPassedGate={hasPassedGate}/>
    {:else if !hasAcceptedWarning}
      <WarningModal bind:hasAccepted={hasAcceptedWarning}
                    squareType={SquareType.public}/>
    {:else}
      <Main />
    {/if}
  {/if}
</body>
