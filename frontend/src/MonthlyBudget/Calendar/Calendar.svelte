<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import { type Day, Days } from '../Days'
  import { type Month, Months } from '../Months'
  import { type CalendarState, initialCalendarState, CALENDAR_STATE_LENGTH } from './CalendarState'
  import { type FixedLengthArray } from '../FixedLengthArray'
  import PaymentEventModal from '../PaymentEventModal.svelte'
  import CurrentDay from './CurrentDay.svelte'

  import { selectedDate } from '../stores/selectedDate'
  let selectedDateVal: Date
  const unsubSelectedDate = selectedDate.subscribe((val: Date) => selectedDateVal = val)

  function contemporaneousMonthOnLoad(): Month {
    const monthId = new Date().getMonth()
    return Months[monthId]
  }

  let selectedDateIdx: number = 0
  let currentMonth: Month = contemporaneousMonthOnLoad()
  let currentYear: number = new Date().getFullYear()

  let calendarState: CalendarState = initialCalendarState

  function setSelectedDate(idx: number) {
    selectedDateIdx = idx
    selectedDate.update(() => calendarState[idx])
  }

  function setCurrentMonth(e: any): void {
    const newMonthId = e.target.value
    currentMonth = Months[newMonthId]
    setCalendarState()
  }

  function setCalendarState(init = false) {
    const firstOfMonth = new Date(currentYear, currentMonth.id, 1)
    const dayFirstOfMonth = firstOfMonth.getDay()

    // fill in main month
    for (let i = 0; i < 42; i++) {
      calendarState[i] = new Date(currentYear, currentMonth.id, i + 1 - dayFirstOfMonth)
    }

    //TODO: eventually this should see if the user has pre-selected a date
    const today = new Date().setHours(0,0,0,0)
    if (init) {
      selectedDateIdx = calendarState.findIndex(d => d.setHours(0,0,0,0) === today)
    }
  }

  function setCurrentYear(e: any): void {
    currentYear = parseInt(e.target.value)
    setCalendarState()
  }

  onMount(() => {
    setCalendarState(true)
  })

  onDestroy(() => {
    unsubSelectedDate()
  })
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
      {#each calendarState as date, idx}
        <button
          class="
            m-2
            calendar-day
            flex flex-col
            text-left
          "
          class:current-day={selectedDateIdx === idx}
          class:weekday={[1,2,3,4,5].includes(idx % 7)}
          class:weekend={[0,6].includes(idx % 7)}
          on:click={() => setSelectedDate(idx)}>
          <div
            class="
              w-full
              text-left
            ">
            <p>{ date.getDate() }<p>
          </div>
        </button>
      {/each}
    </div>
  </div>
  <CurrentDay />
</div>

<style lang="sass">
  @use "./../../styles/color"

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
</style>
