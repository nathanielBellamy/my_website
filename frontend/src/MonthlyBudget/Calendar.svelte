<script lang="ts">
  import { type Day, Days } from './Days'
  import { type Month, Months } from './Months'
  import { type CalendarState } from './CalendarState'

  let currentDayIdx: number = 0
  let currentMonth: Month = contemporaneousMonthOnLoad()
  let currentYear: number = new Date().getFullYear()

  let calendarState: CalendarState = Array(35)
  setCalendarState()

  $: currentDay = calendarState[currentDayIdx]

  function contemporaneousMonthOnLoad(): Month {
    const monthId = new Date().getMonth()
    return Months[monthId]
  }

  function setCurrentMonth(e: any): void {
    const newMonthId = e.target.value
    currentMonth = Months[newMonthId]
    setCalendarState()
  }

  function setCalendarState() {
    const firstOfMonth = new Date(currentYear, currentMonth.id, 1)
    const dayFirstOfMonth = firstOfMonth.getDay()
    const dateLastOfMonth = new Date(currentYear, currentMonth.id, 0).getDate()

    // fill in main month
    for (let i = 0; i < 42; i++) {
      calendarState[i] = {
        day: Days[i % 7],
        date: new Date(currentYear, currentMonth.id, i + 1 - dayFirstOfMonth)
      }
    }
  }

  function setCurrentYear(e: any): void {
    currentYear = parseInt(e.target.value)
    setCalendarState()
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
        grid grid-cols-7 grid-rows-7
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
      {#each calendarState as datedDay, idx}
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
          <div
            class="
              flex
            ">
            <p>{ Months[datedDay.date.getMonth()].abbreviation_3 }.<p>
            <p>{ datedDay.date.getDate() }<p>
          </div>
        </button>
      {/each}
    </div>
  </div>

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
          flex gap-2
          font-bold
          text-xl
          text-cyan-500
        ">
        <p>{ currentDay.day.abbreviation_3 }</p>
        <p>{ Months[currentDay.date.getMonth()].abbreviation_3 }</p>
        <p>{ currentDay.date.getDate() }</p>
        <p>{ currentDay.date.getFullYear() }</p>
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

  .current-day-heading-grid
    grid-template-columns: 30% 70%

  .weekday
    background: color.$blue-6

  .weekend
    background: color.$blue-8

  .curr-day-grid
    grid-template-rows: 16% 42% 42%
    border: 2px double color.$blue-7
    border-radius: 5px
</style>
