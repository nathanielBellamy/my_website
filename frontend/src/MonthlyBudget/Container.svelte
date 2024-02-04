<script lang="ts">
  import { onDestroy, onMount } from "svelte"
  import Main from "./Main.svelte"

  // INIT Prev Settings
  // import Loading from "../lib/Loading.svelte"


  onMount(() => {
  })

  // onDestroy()
</script>

<body id="monthly_budget_container"
      class="monthly_budget_container overscroll-none overflow-y-scroll">
  <Main />
</body>

<style lang="sass">
  @use "./../styles/color"
</style>
