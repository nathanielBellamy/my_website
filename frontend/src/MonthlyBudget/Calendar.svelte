<script lang="ts">
  import { type Day, Days } from './Days'

  let currentDayIdx = 0
  $: currentDay = Days[currentDayIdx % 7]
</script>

<div
  class="
    h-full w-full
    grid grid-cols-2 grid-rows-1
  ">
  <div
    class="
      h-full
      flex flex-col justify-between
    ">
    <div class="
      flex justify-around
      w-full h-full
    ">
      <div
        class="
          flex flex-col justify-around
        ">
        <label
          class="
            font-bold
          "
          for='calendar-month-select'>
          Month
        </label>
        <select
          id='calendar-month-select'
          aria-label="Month"
          title="Month"
          class="
            h-1/2
          ">
          <option>
            Jan
          </option>
        </select>
      </div>
      <div
        class="
          flex flex-col justify-around
        ">
        <label
          class="
            font-bold
          "
          for='calendar-year-select'>
          Year
        </label>
        <select
          id='calendar-year-select'
          title="Year"
          class="
            h-1/2
          ">
          <option>
            2020
          </option>
        </select>
      </div>
    </div>
    <div 
      class="
        bg-green-700
        grow h-full w-min-1/5
        grid grid-cols-7 grid-rows-6
      ">
      {#each Object.values(Days) as day}
        <div 
          class="
            h-full font-bold
            flex flex-col justify-around
          ">
          { day.abbreviation_3 }
        </div>
      {/each}
      {#each {length: 35} as _, idx}
        <button
          class="
           p-2 m-2
           calendar-day
          " 
          class:current-day={currentDayIdx === idx}
          class:weekday={[1,2,3,4,5].includes(idx % 7)}
          class:weekend={[0,6].includes(idx % 7)}
          on:click={() => currentDayIdx = idx}>
          { idx }
        </button>
      {/each}
    </div>
  </div>

  <div
    class="
      bg-blue-400
      grid grid-rows-3 grid-cols-1 curr-day-grid
    ">
    <div
      class="
        flex justify-around
      ">
      <h1>
        { currentDay.abbreviation_3 }
      </h1>
      <div>
        Starting balance:
      </div>
      <div>
        Ending balance:
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
      <button>
        Add Payment Event
      </button>
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

</div>

<style lang="sass">
  @use "./../styles/color"

  .calendar-day:hover
    background: color.$blue-4
  
  .current-day
    background: color.$red-4 !important

  .weekday
    background: color.$blue-6

  .weekend
    background: color.$blue-8

  .curr-day-grid
    grid-template-rows: 16% 42% 42%
</style>
