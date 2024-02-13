<script lang="ts" type="module">
  // import { onDestroy, onMount } from "svelte"
  // import Loading from "../lib/Loading.svelte"
  // onMount(() => {})

  let newOrLoad: string | null = null
</script>

<body
  class="
    flex fle-col
  ">
  <h1
    class="
      flex flex-col justify-end
      font-mono font-bold text-cyan-500
    ">
    Monthly Budget
  </h1>

  {#if !newOrLoad}
    <div
      class="
        grow
        flex flex-col justify-around items-center
        gap-9
        m-9
      ">
      <button
        on:click={() => newOrLoad = 'new'}
        class="
          grow
          flex flex-col justify-around items-center
          rounded-md
          bg-emerald-700
          w-2/3
        ">
        New Budget
      </button>

      <button
        on:click={() => newOrLoad = 'load'}
        class="
          grow
          flex flex-col justify-around items-center
          rounded-md
          bg-blue-700
          w-2/3
        ">
        Load Budget
      </button>
    </div>
  {/if}
</body>

<style lang="sass">
  @use "./../styles/color"
</style>
