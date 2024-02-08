<script lang="ts">
  import { Months } from '../Months'
  import PaymentEventModal from '../PaymentEventModal.svelte'

  export let date: Date

  let showPaymentEventModal = false
</script>

<div
  class="
    m-3
    grid grid-rows-3 grid-cols-1 curr-day-grid
  ">
  <div
    class="
      p-3
      grid grid-rows-1 grid-cols-2
      current-day-heading-grid
    ">
    <h2
      class="
        grid grid-rows-1 grid-cols-4
        current-day-display-grid
        p-2
        font-bold
        text-xl
        text-cyan-500
      ">
      <p
        class="
          w-full
          text-left
        ">
        { date.getDay() }
      </p>
      <p
        class="
          w-full
          text-left
        ">
        { Months[date.getMonth()].abbreviation_3 }
      </p>
      <p
        class="
          w-full
          text-right
        ">
        { date.getDate() }
      </p>
      <p
        class="
          w-full
          text-right
        ">
        { date.getFullYear() }
      </p>
    </h2>
    <div
      class="
        grid grid-rows-2 grid-cols-1
        text-left
        font-bold
      ">
      <div
        class="

        ">
        Starting balance:
      </div>
      <div>
        Ending balance:
      </div>
    </div>
  </div>
  <div
    class="
      flex flex-col justify-between
    ">
    <h2
      class="
        pl-2
        font-bold text-lg
        text-left
      ">
      Payment Events
    </h2>
    <ol>
      <li> foo </li>
      <li> bar </li>
      <li> baz </li>
    </ol>
    <button on:click={() => showPaymentEventModal = true}>
      Add Payment Event
    </button>
    <PaymentEventModal bind:show={showPaymentEventModal}
                       date={date}
    />
  </div>
  <div
    class="
      h-full
      flex flex-col justify-between
    ">
    <h2
      class="
        pl-2
        font-bold text-lg
        text-left
      ">
      Recurring Payment Events
    </h2>
    <ol>
      <li> foo </li>
      <li> bar </li>
      <li> baz </li>
    </ol>
    <button>
      Add Recurring Payment Event
    </button>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"

  .curr-day-grid
    grid-template-rows: 16% 42% 42%
    border: 2px double color.$blue-7
    border-radius: 5px

  .current-day-heading-grid
    grid-template-columns: 35% 65%
    gap: 5px

  .current-day-display-grid
    grid-template-columns: 25% 20% 20% 35%
</style>
