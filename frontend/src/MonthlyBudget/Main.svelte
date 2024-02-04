<script lang="ts" type="module">
  // import { onDestroy, onMount } from "svelte"

  // INIT Prev Settings
  // import Loading from "../lib/Loading.svelte"
  import { BottomNav, BottomNavItem } from 'flowbite-svelte'
  import { Views } from './Views'
  import Calendar from './Calendar.svelte'

  let currentView: Views = Views.Calendar
  // onMount(() => {})
</script>

<div 
  class="
    flex flex-col
    h-[calc(100%-90px)]
  ">
  <body id="monthly_budget_container"
        class="h-full grow overscroll-none overflow-y-scroll">
    {#if currentView === Views.Calendar}
      <Calendar />
    {:else if currentView === Views.PaymentEents}
      <h2>
        PaymentEvents
      </h2>
    {:else if currentView === Views.RecurringPaymentEvents}
      <h2>
        RecurringPaymentEvents
      </h2>
    {:else if currentView === Views.Settings}
      <h2>
        Settings
      </h2>
    {/if}
  </body>

  <footer
    class="w-full h-fit pt-4 pb-4
           flex justify-stretch
           absolute bottom-0 
           text-emerald-700 text-xs font-mono
    ">
   <button
      class="grow hover:text-emerald-500 p-2"
      on:click={() => currentView = Views.Calendar}>
     Calendar
   </button>
   <button 
      class="grow hover:text-emerald-500 p-2"
      on:click={() => currentView = Views.PaymentEents}>
     Payment Events
   </button>
   <button 
      class="grow hover:text-emerald-500 p-2"
      on:click={() => currentView = Views.RecurringPaymentEvents}>
     Recurring Payment Events
   </button>
   <button 
      class="grow hover:text-emerald-500 p-2"
      on:click={() => currentView = Views.Settings}>
     Settings
   </button>
  </footer>
</div>

<style lang="sass">
  @use "./../styles/color"

  .current
    background-color: color.$blue-7 
</style>
