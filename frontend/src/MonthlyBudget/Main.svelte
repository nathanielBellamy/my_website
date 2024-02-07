<script lang="ts" type="module">
  // import { onDestroy, onMount } from "svelte"

  // INIT Prev Settings
  // import Loading from "../lib/Loading.svelte"
  import { BottomNav, BottomNavItem } from 'flowbite-svelte'
  import { Views } from './Views'
  import Calendar from './Calendar.svelte'

  let currentView: Views = Views.Calendar
  let views = Object.values(Views)
  // onMount(() => {})
</script>

<div 
  class="
    flex flex-col
    h-[calc(100%-84px)]
  ">
  <body id="monthly_budget_container"
        class="h-full grow overscroll-none overflow-y-scroll">
    {#if currentView === Views.Calendar}
      <Calendar />
    {:else if currentView === Views.PaymentEents}
      <h2>
        {Views.PaymentEents}
      </h2>
    {:else if currentView === Views.RecurringPaymentEvents}
      <h2>
        {Views.RecurringPaymentEvents}
      </h2>
    {:else if currentView === Views.Reports}
      <h2>
        {Views.Reports}
      </h2>
    {:else if currentView === Views.Settings}
      <h2>
        {Views.Settings}
      </h2>
    {/if}
  </body>

  <footer
    class="w-full h-fit pt-4 pb-4
           flex justify-stretch
           absolute bottom-0 
           text-emerald-700 text-xs font-mono
    ">
  {#each Object.values(Views) as view}
    <button
        class="grow hover:text-emerald-500 p-2"
        class:current={currentView == view}
        on:click={() => currentView = view}>
      { view }
    </button>
  {/each}
  </footer>
</div>

<style lang="sass">
  @use "./../styles/color"

  .current
    background-color: color.$green-7
    color: color.$green-2
</style>
