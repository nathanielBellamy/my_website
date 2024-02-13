<script lang="ts">
  import { Views } from './Views'
  import Calendar from './Calendar/Calendar.svelte'
  import Main from "./Main.svelte"

  let currentView: Views = Views.Main
</script>

<body id="monthly_budget_container"
      class="monthly_budget_container overscroll-none overflow-y-scroll">
  <div
    class="
      flex flex-col
      h-[calc(100%-84px)]
    ">
    <body id="monthly_budget_container"
          class="h-full grow overscroll-none overflow-y-scroll">
      {#if currentView === Views.Main}
        <Main />
      {:else if currentView === Views.Calendar}
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
</body>

<style lang="sass">
  @use "./../styles/color"

  .current
    background-color: color.$green-7
    color: color.$green-2
</style>
