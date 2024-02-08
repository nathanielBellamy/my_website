
<script lang="ts">
  import { onMount } from 'svelte'
  import { Button, Modal, Label, Input, Radio, Select } from 'flowbite-svelte'

  import { selectedDate } from './stores/selectedDate'
  let selectedDateVal: Date
  const unsubSelectedDate = selectedDate.subscribe((val: Date) => selectedDateVal = val)

  $: initDate = selectedDateVal.toJSON().split('T')[0]

  export let show: boolean = false
  export let disableDatePicker: boolean = false

  let selectedRecurrenceEvery: any;
  const recurrenceEvery: any = [
    { value: 'day', name: 'Day(s)' },
    { value: 'week', name: 'Week(s)' },
    { value: 'month', name: 'Month(s)'},
    { value: 'year', name: 'Year(s)' }
  ]

  let paymentTypeGroup: string = "payment"
</script>

<Modal
  bind:open={show}
  size="md"
  autoclose={false}
  class="
    w-full
    bg-gray-200
    text-cyan-500
  ">
  <form class="flex flex-col space-y-6" action="#">
    <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
      New Recurring Payment Event
    </h3>
    <Label
      class="
        space-y-2
        text-left
      ">
      <span>Start Date</span>
      <Input
        type="date"
        name="start-date"
        value={initDate}
        class="
          bg-blue-200
        "
        disabled={disableDatePicker}
        required />
    </Label>
    <Label
      class="
        space-y-2
        text-left
      ">
      <span>End Date</span>
      <Input
        type="date"
        name="end-date"
        class="
          bg-blue-200
        "
        optional />
    </Label>
    <Label
      class="
        space-y-2
        text-left
      ">
      <span>Recurs Every</span>
      <div
        class="
          flex justify-between gap-2
        ">
        <Input
          type="number"
          name="every"
          step="any"
          min="0"
          placeholder="1"
          class="
            bg-blue-200
          "
          required />
        <Select
          name="recurrence"
          items={recurrenceEvery}
          bind:value={selectedRecurrenceEvery}
          class="
            bg-blue-200
          "
          required />
      </div>
    </Label>
    <Label
      class="
        space-y-2
        text-left
      ">
      <span>Amount</span>
      <Input
        type="number"
        name="amount"
        step="any"
        min="0"
        placeholder="$0.00"
        class="
          bg-blue-200
        "
        required />
    </Label>
    <Label
      class="
        space-y-2
        text-left
      "
    >
      <span>Type</span>
      <div class="grid grid-cols-2 gap-6">
        <div
          class="rounded border border-gray-200 dark:border-gray-700">
          <Radio
            name="payment-event-type"
            value="payment"
            bind:group={paymentTypeGroup}
            class="
              w-full p-4
              bg-red-300
            ">
            Payment
          </Radio>
        </div>
        <div
          class="rounded border border-gray-200 dark:border-gray-700">
          <Radio
            name="payment-event-type"
            value="payment_received"
            bind:group={paymentTypeGroup}
            class="
              w-full p-4
              bg-green-300
            ">
            Payment Received
          </Radio>
        </div>
      </div>
    </Label>
    <Label
      class="
        space-y-2
        text-left
      ">
      <span>Memo</span>
      <Input
        type="text"
        name="memo"
        placeholder="Describe this payment event"
        class="
          bg-blue-200
        "
        required />
    </Label>
    <Button
      type="submit"
      class="
        w-full
        bg-red-500 hover:bg-red-400
      ">
      Create Recurring Payment Event
    </Button>
  </form>
</Modal>

