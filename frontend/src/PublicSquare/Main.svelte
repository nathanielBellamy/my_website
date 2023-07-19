<script lang="ts">
  import init, { PubSq, rust_init_message } from '../../pkg/src_rust.js'
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'
  import { Toast } from 'flowbite-svelte'
  import { fly } from 'svelte/transition'

  // messaging
  interface Message {
    clientId: number,
    body: string
  }

  let toSendBody: string = ""
  $: toSend = {clientId: 1, body: toSendBody}

  let toReceive: Message | null = null

  // feed
  let feed: Message[] = []
  $: _feed = !toReceive ? [...feed] : [toReceive, ...feed]
  let feedWasScrolledToBottom: boolean = false
  
  function scrollFeedToBottom() {
    const feed = document.getElementById("public_square_feed")
    feed.scrollTop = feed.scrollHeight
  }

  function feedIsScrolledToBottom() {
    var res: boolean = false
    const feed = document.getElementById("public_square_feed")
    if (!feed) {
      res = false
    } else {
      // scroll is within 150px of bottom
      res = Math.abs(feed.scrollTop - (feed.scrollHeight - feed.offsetHeight)) < 150
    }

    return res
  }

  const FEED_LENGTH: number = 926
  function pushToFeed(m: Message) {
    if (feed.length > FEED_LENGTH) {
      feed.shift()
    }
    feed.push(m)
    feed = [...feed]
  }

  function formatClientId(id: number): string {
    var res: string = ""
    if (!!id) {
      res = `:: User ${id} :: `
    }

    return res
  }

  // websocket
  const ws = new WebsocketBuilder('ws://localhost:8080/public-square-feed-ws')
      .onOpen(() => {
        triggerShowConnected()
        pushToast(toastConnected)
      })
      .onClose(() => pushToast(toastDisconnected))
      .onError(() => pushToast(toastError))
      .onMessage((_i, ev) => { 
        const message: Message = JSON.parse(ev.data)
        pushToFeed(message)
      })
      .onRetry(() => {})
      .build()

  // alerts
  let showConnected: boolean = false;
  let counter: number = 6;

  function triggerShowConnected() {
    showConnected = true;
    counter = 6;
    timeout();
  }

  function timeout() {
    if (--counter > 0)
      return setTimeout(timeout, 1000);
    showConnected = false;
  }

  enum ToastColor {
    green = "green",
    blue = "blue",
    red = "red",
    gray = "gray",
    yellow = "yellow",
    indigo = "indigo",
    purple = "purple",
    orange = "orange",
    none = "none"
  }

  interface Toast {
    color: ToastColor,
    text: string,
  }

  const toastConnected: Toast = {
    color: ToastColor.green,
    text: "Connected"
  }

  const toastDisconnected = {
    color: ToastColor.blue,
    text: "Disconnected"
  }

  const toastError = {
    color: ToastColor.red,
    text: "Connection error"
  }
  
  let toasts: Toast[] = []
  function pushToast(t: Toast) {
    toasts = [t, ...toasts]
  }

  // LIFECYCLE
  onMount(() => scrollFeedToBottom())
  onDestroy(() => ws.close())

  beforeUpdate(() => {
    feedWasScrolledToBottom = feedIsScrolledToBottom()
  })
  afterUpdate(() => {
    if (feedWasScrolledToBottom) scrollFeedToBottom()
  })


  async function run() {
    if (!hasBeenDestroyed) { 
      // resize + key block in Container.svelte may destroy component before wasm_bindgen can load
      // without this check, it is possible to load two wasm instances
      // since wasm retrieves the elements using .get_element_by_id
      // and since a new instance of the component will havee been mounted by the time wasm_bindgen loads
      // the result is two identical wasm instances listening to the same ui elements and drawing to the same context
      await init()
      rust_init_message("Public Square Wasm!")
      
      // init wasm process and set initial values
      const settings = await PubSq.run(clientId, touchScreenVal)
      setAllSettings(settings)
      renderDataReady = true
    }
  }

  run()
</script>

{#each toasts as { color, text }}
  {#if text !== "Connected"}
    <Toast color={color}
           class="bg-slate-800 text-sky-600"
           transition={fly}
           params="{{x: 200}}"
           position="bottom-right">
      <svelte:fragment slot="icon">
        <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path></svg>
        <span class="sr-only">Check icon</span>
      </svelte:fragment>
      {text}
    </Toast>
  {:else}
    <Toast color={color}
           bind:open={showConnected}
           class="bg-slate-800 text-sky-600"
           transition={fly}
           params="{{x: 200}}"
           position="bottom-right">
      <svelte:fragment slot="icon">
        <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path></svg>
        <span class="sr-only">Check icon</span>
      </svelte:fragment>
      {text}
    </Toast>
  {/if}
{/each}

<div class="w-full h-full p-2 flex justify-between items-stretch">
  <div class="grow p-5 m-5 flex flex-col justify-between items-stretch">
    <input bind:value={toSendBody}/>
    <button on:click={() => ws.send(toSend.body)}>
      SEND IT
    </button>
  </div>
  <div  id="public_square_feed"
        class="grow p-5 m-5 overflow-y-scroll flex flex-col items-stretch">
    {#each _feed as { clientId, body }, i} 
      {#if !!i}
        <div class="grow p-5 m-5 flex items-center">
          <h4>
            {i}
          </h4>
          <div>
            {formatClientId(clientId)}
          </div>
          <div>
            {body}
          </div>
        </div>
      {/if}
    {/each}
  </div>


</div>

<style lang="sass">

</style>
