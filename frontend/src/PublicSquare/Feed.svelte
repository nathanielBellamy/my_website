<script lang="ts">
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import { Toast } from 'flowbite-svelte'
  import { WebsocketBuilder } from 'websocket-ts'
  import { ToastColor } from '../lib/Toaster'
  import type { ToasterProps } from '../lib/Toaster'
  import Toaster from '../lib/Toaster.svelte'
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

  const toastConnected: ToasterProps = {
    color: ToastColor.green,
    text: "Connected"
  }

  const toastDisconnected: ToasterProps = {
    color: ToastColor.blue,
    text: "Disconnected"
  }

  const toastError: ToasterProps = {
    color: ToastColor.red,
    text: "Connection error"
  }
  
  let toasts: ToasterProps[] = []
  function pushToast(t: ToasterProps) {
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
</script>

{#each toasts as { color, text }}
  {#if text !== "Connected"}
    <Toaster bind:open={showConnected}
             color={color}
             text={text}/>
  {:else}
    <Toaster open={null}
             color={color}
             text={text}/>
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
