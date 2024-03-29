<script lang="ts">
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import type { FeedMessage } from './FeedMessage'
  import { SystemMessage } from './FeedMessage'

  import EmojiKeyboard from '../../lib/EmojiKeyboard.svelte'
  
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'
  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)
  let i18n = new I18n("magicSquare/feed")

  import { psFeed } from '../../stores/psFeed'
  let psFeedVal: FeedMessage[]
  const unsubPsFeed = psFeed.subscribe((val: FeedMessage[]) => psFeedVal = [...val])
  
  import { psConnected } from "../../stores/psConnected"
  let psConnectedVal: boolean
  const unsubPsConnected = psConnected.subscribe(val => psConnectedVal = val)

  export let sendFeedMessage: (body: string) => void
  export let clientIdSelf: number | null

  let feedWasScrolledToBottom: boolean = false

  function scrollFeedToBottom() {
    const feed = document.getElementById("public_square_feed_messages_container")
    feed.scrollTop = feed.scrollHeight
  }

  function feedIsScrolledToBottom() {
    var res: boolean = false
    const feed = document.getElementById("public_square_feed_messages_container")
    if (!feed) {
      res = false
    } else {
      // scroll is within 150px of bottom
      res = Math.abs(feed.scrollTop - (feed.scrollHeight - feed.offsetHeight)) < 150
    }
    return res
  }

  function formatClientId(id: number, clientIsSelf: boolean): string {
    var res: string = ""
    if (!!id) {
      if (!clientIsSelf){
        res = `sq-${id}`
      } else {
        res = `sq-${id}`
      }
    }

    return res
  }

  let nextEmoji: string = ""
  let emojiSetIndicator: boolean = false

  export let toSendBody: string = ""

  $: if (emojiSetIndicator || !emojiSetIndicator) {
    toSendBody = toSendBody + nextEmoji
  }

  $: if (toSendBody.length > 16){
       toSendBody = toSendBody.slice(2)
     } else {
       toSendBody = toSendBody
     }

  function clr(){
    nextEmoji = ""
    toSendBody = ""
  }

  // setup backspace
  function keyboardListener(e: any) {
    if (e.key === 'Backspace') clr()
    if (e.key === 'Enter' && !!toSendBody.length) {
      e.preventDefault()
      sendFeedMessage(toSendBody)
      clr()
    }
  }

  function formatSystemMessage(clientId: number, messageBody: string, lang: Lang): string {
    var res: string = ""
    switch (messageBody) {
      case SystemMessage.sqConnected:
        if (clientId === clientIdSelf) {
          res = `${i18n.t("sqConnectedSelf", lang)} sq-${clientId}`
        } else {
          res = `sq-${clientId}${i18n.t("sqConnected", lang)}`
        }
        break
      case SystemMessage.sqDisconnected:
        res = `sq-${clientId}${i18n.t("sqDisconnected", lang)}`
        break
    }
    return res
  }

  function systemConnected(body: string): boolean {
    switch (body) {
      case SystemMessage.init:
      case SystemMessage.sqConnected:
        return true
      case SystemMessage.sqDisconnected:
      default:
        return false
    }
  }

  // LIFECYCLE
  onMount(() => {
    scrollFeedToBottom()
    window.addEventListener('keydown', keyboardListener)
  })
  onDestroy(() => {
    unsubLang()
    unsubPsConnected()
    unsubPsFeed()
    window.removeEventListener('keydown', keyboardListener)
  })

  beforeUpdate(() => {
    feedWasScrolledToBottom = feedIsScrolledToBottom()
  })
  afterUpdate(() => {
    if (feedWasScrolledToBottom) scrollFeedToBottom()
  })
</script>

<div class="public_square_feed h-full w-full pt-2 pl-2 pr-2 pb-8 grid grid-cols-1 grid-rows-2">
  <div  id="public_square_feed_messages_container"
        class="public_square_feed_messages_container h-full rounded-md overflow-y-scroll">
    <div  id="public_square_feed_messages"
          class="public_square_feed_messages h-fit p-2 flex flex-col items-center gap-2">
      {#if !psConnectedVal}
        <div class="w-full text-center text-sm font-bold font-mono text-red-700">
          {i18n.t("disconnected", langVal)}
        </div>
      {:else if !psFeedVal.length}
        <div class="w-full text-center text-sm font-bold font-mono text-red-700">
          {i18n.t("noMessages", langVal)}
        </div>
      {:else}
        {#each psFeedVal as { clientId, body, system }, i} 
          {#if !!i}
            <div class="feed_message p-2 w-full h-fit rounded-md"
                 class:feed_message_self={clientIdSelf === clientId}
                 class:feed_message_system={clientId === 0}
                 class:feed_message_other={clientIdSelf !== clientId}>
              {#if system}
                <div class="font-bold text-sm p-2 mr-2 rounded-md w-full col-span-2 break-all flex justify-around"
                     class:text-red-700={!systemConnected(body)}
                     class:text-emerald-700={systemConnected(body)}>
                  {formatSystemMessage(clientId, body, langVal)}
                </div>
              {:else if clientIdSelf !== clientId}
                <div class="feed_message_body font-bold text-lg p-2 mr-2 rounded-md w-full break-all">
                  {body}
                </div>
                <div class="h-full w-full flex flex-col justify-around items-center">
                  <div class="feed_message_user pl-2 pr-2 pt-4 pb-4 rounded-md text-sm font-semibold h-full flex justify-around items-center">
                    {formatClientId(clientId, false)}
                  </div>
                </div>
              {:else}
                <div class="h-full w-full flex flex-col justify-around items-center">
                  <div class="feed_message_user pl-2 pr-2 pt-4 pb-4 rounded-md text-sm font-semibold h-full flex justify-around items-center">
                    <div class="flex flex-col justify-between items-stretch">
                      <div class="w-full flex justify-around items-center"> 
                        {i18n.t('me', langVal)}
                      </div>
                      <div class="w-full flex justify-around items-center text-xs"> 
                        {formatClientId(clientId, true)}
                      </div>
                    </div>
                  </div>
                </div>
                <div class="feed_message_body font-bold text-lg p-2 mr-2 rounded-md w-full break-all">
                  {body}
                </div>
              {/if}
            </div>
          {/if}
        {/each}
      {/if}
    </div>
  </div>
  <div class="public_square_feed_input p-2 h-full grid grid-rows-3 grid-cols-1 gap-4">
    <div class="public_square_feed_input_body rounded-md flex justify-around items-center">
      {toSendBody}
    </div>
    <div class="public_square_feed_input_emoji_keyboard rounded-md overflow-y-scroll p-2">
      <EmojiKeyboard bind:value={nextEmoji}
                     bind:valueSetIndicator={emojiSetIndicator}/>
    </div>
    <div class="public_square_feed_input_buttons w-full grid grid-cols-2 grid-rows-1 gap-2 text-cyan-700">
      <button class="public_square_feed_input_buttons_send"
              on:click={() => {
                if (toSendBody.length) {
                  sendFeedMessage(toSendBody)
                  clr()
                }
              }}>
        {i18n.t('send', langVal)}
      </button>
      <button class="public_square_feed_input_buttons_clr flex justify-around items-center"
              on:click={() => clr()}>
        {i18n.t('clear', langVal)}
      </button>
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .public_square_feed
    grid-template-areas: "messages" "input"
    grid-template-rows: 20em 
    &_messages
      grid-area: "messages"
      &_container
        background-color: color.$grey-transp
    &_input
      grid-area: "input"
      grid-template-areas: "body" "buttons" "keyboard" 
      grid-template-rows: 2em 4em 226px
      &_emoji_keyboard
        grid-area: keyboard
        border-left: 3px color.$blue-3 double
        border-right: 3px color.$blue-3 double
      &_body
        grid-area: body
        border-right: 3px color.$blue-3 double
      &_buttons
        grid-area: buttons
        grid-template-columns: 70% 30%
        &_send
          border-left: 3px color.$green-7 double
          border-right: 3px color.$green-7 double
        &_clr
          border-left: 3px color.$red-7 double
          border-right: 3px color.$red-7 double
      
  .feed_message
    grid-template-areas: "user body"
    grid-template-columns: 20% 80%
    &_user
      grid-area: "user"
      border-top: 5px double color.$blue-7
      border-bottom: 5px double color.$blue-7
      height: fit-content
    &_body
      grid-area: "body"
      background-color: color.$grey-transp
    &_self
      border-right: 5px double color.$blue-3
      display: grid
      grid-template-columns: 50% 50%
      grid-template-rows: 1fr
      gap: 4px
    &_other
      border-left: 5px double color.$green-7
      display: grid
      grid-template-columns: 50% 50%
      grid-template-rows: 1fr
      gap: 4px
    &_system
      display: flex
      flex-direction: column
      justify-content: space-around
      align-items: stretch
</style>
