<script lang="ts" type="module">
  import { onDestroy, onMount } from "svelte"
  // import Loading from "../lib/Loading.svelte"
  // onMount(() => {})
  import { newOrLoad, NewOrLoad } from './stores/newOrLoad'
  let newOrLoadVal: NewOrLoad
  const unsubNewOrLoad = newOrLoad.subscribe((val: NewOrLoad) => newOrLoadVal = val)

  function updateNewOrLoad(newVal: NewOrLoad) {
    newOrLoad.update(() => newVal)
  }

  onDestroy(() => {
    unsubNewOrLoad()
  })
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
    <div
      class:uninit={newOrLoadVal === NewOrLoad.uninit}
      class:init={newOrLoadVal !== NewOrLoad.uninit}>
      <button
        on:click={() => updateNewOrLoad(NewOrLoad.new)}
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
        on:click={() => updateNewOrLoad(NewOrLoad.load)}
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
    {#if newOrLoadVal === NewOrLoad.new}
      <form>
        new budget
      </form>
    {:else if newOrLoadVal === NewOrLoad.load}
      <form>
        load budget
      </form>
    {/if}
</body>

<style lang="sass">
  @use "./../styles/color"

  .uninit
    flex-grow: 1
    display: flex
    flex-direction: column
    justify-content: space-around
    align-items: center
    gap: 9px

  .init
    flex-grow: 0
    display: flex
    flex-direction: row
    justify-content: flex-end
    align-items: center
    gap: 9px
</style>
