<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Modal } from 'flowbite-svelte'
  import { WebsocketBuilder } from 'websocket-ts'
  import Feed from '../MagicSquare/ControlModules/Feed.svelte'
  import type { ToasterProps } from '../lib/Toasty'
  import { ToastColor } from '../lib/Toasty'
  import type { FeedMessage } from './../MagicSquare/ControlModules/FeedMessage'
  import MagicSquarePub from './MagicSquarePub.svelte'
  import { FEED_LENGTH, psFeed } from '../stores/psFeed'
  import Toaster from '../lib/Toaster.svelte'
  import { ViteMode } from '../ViteMode'

  let clientId: number

  const baseUrl: string = import.meta.env.VITE_BASE_URL
  const protocol: string = import.meta.env.VITE_MODE == ViteMode.localhost
    ? "ws"
    : "wss"
  const fullUrl: string = `${protocol}://${baseUrl}/public-square-feed-ws`

  // websocket
  const ws = new WebsocketBuilder(fullUrl)
      .onOpen(() => {
        triggerShowConnected()
      })
      .onClose(() => pushToast(toastDisconnected))
      .onError(() => pushToast(toastError))
      .onMessage((_i, ev) => {
        const message: FeedMessage = JSON.parse(ev.data)
        if (message.body === "__init__connected__") {
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
  let showConnected: boolean = false;
  let counter: number = 0;

  function triggerShowConnected() {
    showConnected = true
    counter = 6
    timeout()
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
            bind:clientIdSelf={clientId}
            bind:toSendBody={toSendBody}/>
    </div>
  </MagicSquarePub>
  <Toaster bind:open={showConnected}
           color={ToastColor.green}
           text={"Connected"}/>
  {#each toasts as { color, text }}
      <Toaster open={null}
               color={color}
               text={text}/>
  {/each}
</div>

<style lang="sass">

</style>
