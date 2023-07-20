<script lang="ts">
  import init, { PubSq, rust_init_message } from '../../pkg/src_rust.js'
  import { watchResize } from "svelte-watch-resize"
  import { onDestroy, onMount } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'
  import Feed from './Feed.svelte'
  import type { ToasterProps } from '../lib/Toaster'
  import { ToastColor } from '../lib/Toaster'
  import type { FeedMessage } from './FeedMessage'
  import { touchScreen } from '../stores/touchScreen'
  import MagicSquarePub from './MagicSquarePub.svelte'
  import { FEED_LENGTH, psFeed } from '../stores/psFeed'
  import WarningModal from '../MagicSquare/WarningModal.svelte';

  let psFeedVal: FeedMessage[]
  const unsubPsFeed = psFeed.subscribe((val: FeedMessage[]) => psFeedVal = [...val])

  let magicSquareInstance: number = 0
  $: magicSquareInstance
  let sideLength: number = 0

  function incrementMagicSquareInstance() {
    magicSquareInstance += 1
  }

  async function handleResize() {
    incrementMagicSquareInstance()
    let element = document.getElementById("magic_square_container")
    if (!!element){
      sideLength = Math.floor(Math.min(element.offsetWidth, element.offsetHeight) / 1.3) - 25
    }
  }

  // TODO:
  // this combination of touchSreen store and value updates works 
  // to ensure that the touchScreen value is updated by the time it is passed to RustWasm
  // this includes the hidden element using touchScreenVal in the html
  // not sure why this magic combo gets it done
  // but we won't worry about it right now
  const id = (x: any): any => x
  let touchScreenVal: boolean
  const unsubTouchScreen = touchScreen.subscribe((val: boolean) => touchScreenVal = val)
  $: isTouchScreen = id(touchScreenVal)

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

  function triggerShowConnected() {
    showConnected = true;
    counter = 6;
    timeout();
  }
  let showConnected: boolean = false;
  let counter: number = 6;

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
  $: toSend = {clientId: 1, body: toSendBody}
  
  function pushToFeed(m: FeedMessage) {
    psFeed.update((prevFeed: FeedMessage[]) => {
      if (prevFeed.length > FEED_LENGTH) {
        prevFeed.shift()
      }
      prevFeed.push(m)
      return [...prevFeed]
    })
  }

  // LIFECYCLE
  let renderDataReady: boolean = false
  let hasBeenDestroyed: boolean = false
  let hasAcceptedWarning: boolean = false
  onMount(() => {
    hasAcceptedWarning = !!localStorage.getItem("magic_square_has_accepted_warning")
  })
  onDestroy(() => {
    hasBeenDestroyed = true
    unsubPsFeed()
    unsubTouchScreen()
    ws.close()
  })

  let settings: any

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
      settings = await PubSq.run(touchScreenVal)
      // setAllSettings(settings)
      renderDataReady = true
    }
  }

  onMount(async () => {
    run()
  })
</script>


<div class="flex justify-around items-center"
     use:watchResize={handleResize}>
  <div style="display: none"> {touchScreenVal}  </div>
  {#if !hasAcceptedWarning}
    <WarningModal bind:hasAccepted={hasAcceptedWarning}/>
  {:else if renderDataReady}
    {#key magicSquareInstance}
      <MagicSquarePub  bind:renderDataReady={renderDataReady}
                       settings={settings}
                       sideLength={sideLength}>
        <div slot="psFeed"
             class="h-full">
          <Feed sendFeedMessage={sendFeedMessage}
                bind:showConnected={showConnected}
                bind:toasts={toasts}
                bind:toSendBody={toSendBody}/>
        </div>
      </MagicSquarePub>
    {/key}
  {/if}
</div>

<style lang="sass">

</style>
