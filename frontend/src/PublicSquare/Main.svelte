<script lang="ts">
  import { onDestroy } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'
  import Feed from './Feed.svelte'
  import type { ToasterProps } from '../lib/Toasty'
  import { ToastColor } from '../lib/Toasty'
  import type { FeedMessage } from './FeedMessage'
  import MagicSquarePub from './MagicSquarePub.svelte'
  import { FEED_LENGTH, psFeed } from '../stores/psFeed'
  import Toaster from '../lib/Toaster.svelte';

  let clientId: number

  // websocket
  const ws = new WebsocketBuilder('ws://localhost:8080/public-square-feed-ws')
      .onOpen(() => {
        triggerShowConnected()
        pushToast(toastConnected)
      })
      .onClose(() => pushToast(toastDisconnected))
      .onError(() => pushToast(toastError))
      .onMessage((_i, ev) => {
        const message: FeedMessage = JSON.parse(ev.data)
        if (message.body === "connected") {
          clientId = message.clientId
        } else {
          pushToFeed(message)
        }
      })
      .onRetry(() => {})
      .build()

  function sendFeedMessage(body: string) {
    ws.send(body)
  }

  // alerts
  const toastConnected: ToasterProps = {
    color: ToastColor.green,
    text: "Connected"
  }

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

  // messaging
  let toSendBody: string = ""
  
  function pushToFeed(m: FeedMessage) {
    psFeed.update((prevFeed: FeedMessage[]) => {
      if (prevFeed.length > FEED_LENGTH) {
        prevFeed.shift()
      }
      prevFeed.push(m)
      return [...prevFeed]
    })
  }

  onDestroy(() => {
    ws.close()
  })
</script>


<div class="h-full w-full overflow-hidden">
  <MagicSquarePub>
    <div slot="psFeed"
         class="h-full">
      <Feed sendFeedMessage={sendFeedMessage}
            bind:toSendBody={toSendBody}/>
    </div>
  </MagicSquarePub>
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
</div>

<style lang="sass">

</style>
