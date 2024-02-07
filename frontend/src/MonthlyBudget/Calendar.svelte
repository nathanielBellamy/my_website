<script lang="ts">
  import { type Day, Days } from './Days'
  import { type Month, Months } from './Months'

  let currentDayIdx = 0
  $: currentDay = Days[currentDayIdx % 7]

  let currentMonth: Month = contemporaneousMonthOnLoad()
  let currentYear: number = new Date().getFullYear()


  function contemporaneousMonthOnLoad(): Month {
    const monthId = new Date().getMonth()
    return Months[monthId]
  }

  function setCurrentMonth(e: any): void {
    const newMonthId = e.target.value
    currentMonth = Months[newMonthId]
  }

  function setCurrentYear(e: any): void {
    currentYear = parseInt(e.target.value)
  }
</script>

<div
  class="
    h-full w-full
    grid grid-rows-1 grid-cols-2
  ">
  <div
    class="
      h-full
      grid grid-rows-2 grid-cols-1 calendar-grid
    ">
    <div class="
      flex justify-around
      w-full h-full
      pt-2 pb-2
    ">
      <h1
        class="
          flex flex-col justify-end
          font-mono font-bold text-cyan-500
        ">
        Monthly Budget
      </h1>
      <div
        class="
          flex flex-col justify-end
        ">
        <label
          class="
            font-bold
            text-left
            text-cyan-300
          "
          for='calendar-month-select'>
          Month
        </label>
        <select
          id='calendar-month-select'
          aria-label="Month"
          title="Month"
          value={currentMonth.id}
          on:change={(e) => setCurrentMonth(e)}
          class="
            h-1/2
          ">
          {#each Object.values(Months) as month}
            <option
              value={month.id}
              class="
                w-full
                flex justify-between
              ">
              <p> {month.id + 1}. </p>
              <p> {month.name} </p>
            </option>
          {/each}
        </select>
      </div>
      <div
        class="
          flex flex-col justify-end
          text-left
        ">
        <label
          class="
            text-cyan-300
            font-bold
          "
          for='calendar-year-select'>
          Year
        </label>
        <select
          id='calendar-year-select'
          title="Year"
          value={currentYear}
          on:change={setCurrentYear}
          class="
            h-1/2
          ">
          {#each {length: 100} as _, year}
            <option
              value={year + 1990}>
              {year + 1990}
            </option>
          {/each}
        </select>
      </div>
    </div>
    <div
      class="
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
            flex flex-col
            text-left
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
        grid grid-rows-1 grid-cols-2
      ">
      <h1
        class="
          grow
          flex justify-around
        ">
        <p>{ currentDay.abbreviation_3 }</p>
        <p>{ currentMonth.abbreviation_3}</p>
        <p>{ currentYear }</p>
      </h1>
      <div
        class="
          grid grid-rows-2 grid-cols-1
          text-left
        ">
        <div>
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

  .calendar-grid
    grid-template-rows: 12% 88%

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
