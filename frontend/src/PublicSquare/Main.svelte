<script lang="ts">
  import { onDestroy } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'
  import Feed from '../MagicSquare/ControlModules/Feed.svelte'
  import type { ToasterProps } from '../lib/Toasty'
  import { ToastColor } from '../lib/Toasty'
  import type { FeedMessage } from './../MagicSquare/ControlModules/FeedMessage'
  import MagicSquarePub from './MagicSquarePub.svelte'
  import { FEED_LENGTH, psFeed } from '../stores/psFeed'
  import Toaster from '../lib/Toaster.svelte'
  import { ViteMode } from '../ViteMode'
  import { Icons } from '../lib/Icons'

  import { I18n, Lang } from "../I18n"
  import { lang } from "../stores/lang"
  let i18n = new I18n("publicSquare/main")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  let clientId: number

  const baseUrl: string = import.meta.env.VITE_BASE_URL
  const protocol: string = import.meta.env.VITE_MODE == ViteMode.localhost
    ? "ws"
    : "wss"
  const fullUrl: string = `${protocol}://${baseUrl}/public-square-feed-ws`

  let showConnectionError: boolean = false
  let showDisconnected: boolean = false

  // websocket
  const ws = new WebsocketBuilder(fullUrl)
      .onOpen(() => {
        triggerShowConnected()
      })
      .onClose(() => showDisconnected = true)
      .onError(() => showConnectionError = true)
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
  
  $: connectedText = i18n.t("connected", langVal)
  $: connectionErrorText = i18n.t("connectionError", langVal)
  $: disconnectedText = i18n.t("disconnected", langVal)

  const toastDisconnected: ToasterProps = {
    color: ToastColor.red,
    text: disconnectedText,
    icon: Icons.ExclamationCircleSolid
  }

  const toastError: ToasterProps = {
    color: ToastColor.red,
    text: connectionErrorText,
    icon: Icons.ExclamationCircleSolid
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
    unsubLang()
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
           icon={Icons.CheckCircleSolid}
           bind:text={connectedText}/>
  <Toaster bind:open={showConnectionError}
           icon={Icons.ExclamationCircleSolid}
           color={ToastColor.red}
           bind:text={connectionErrorText}/>
  <Toaster bind:open={showDisconnected}
           icon={Icons.ExclamationCircleSolid}
           color={ToastColor.red}
           bind:text={connectionErrorText}/>
</div>
