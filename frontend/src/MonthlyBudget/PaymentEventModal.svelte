<script lang="ts">
  import { onMount } from 'svelte'
  import { Button, Modal, Label, Input, Radio } from 'flowbite-svelte'

  import { selectedDate } from './stores/selectedDate'
  let selectedDateVal: Date
  const unsubSelectedDate = selectedDate.subscribe((val: Date) => selectedDateVal = val)

  $: initDate = selectedDateVal.toJSON().split('T')[0]

  export let show: boolean = false
  export let disableDatePicker: boolean = false

  const paymentEventTypes: any = [
    {value:'payment', name: 'Payment'},
    {value:'payment_received', name: 'Payment Received'},
  ]

  let paymentTypeGroup = "payment"
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
      New Payment Event
    </h3>
    <Label
      class="
        space-y-2
        text-left
      ">
      <span>Date</span>
      <Input
        type="date"
        name="date"
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
      <span>Amount</span>
      <Input
        type="number"
        name="memo"
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
      Create Payment Event
    </Button>
  </form>
</Modal>

