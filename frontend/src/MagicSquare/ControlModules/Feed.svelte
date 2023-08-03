<script lang="ts">
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import type { FeedMessage } from './FeedMessage'
  import { psFeed } from '../../stores/psFeed'
  import EmojiKeyboard from '../../lib/EmojiKeyboard.svelte'

  let psFeedVal: FeedMessage[]
  const unsubPsFeed = psFeed.subscribe((val: FeedMessage[]) => psFeedVal = [...val])

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

  // LIFECYCLE
  onMount(() => {
    scrollFeedToBottom()
    window.addEventListener('keydown', keyboardListener)
  })
  onDestroy(() => {
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

<div class="public_square_feed w-full h-full pt-2 pl-2 pr-2 pb-8 grid grid-cols-1 grid-rows-2">
  <div  id="public_square_feed_messages_container"
        class="public_square_feed_messages_container h-full rounded-md overflow-y-scroll">
    <div  id="public_square_feed_messages"
          class="public_square_feed_messages h-fit p-2 flex flex-col items-center gap-2">
      {#if !psFeedVal.length}
        <h1> No New Messages </h1>
      {:else}
        {#each psFeedVal as { clientId, body }, i} 
          {#if !!i}
            <div class="feed_message p-2 w-full h-fit rounded-md"
                 class:feed_message_self={clientIdSelf === clientId}
                 class:feed_message_system={clientId === 0}
                 class:feed_message_other={clientIdSelf !== clientId}>
              {#if clientId === 0}
                <div class="feed_message_body font-bold text-lg p-2 mr-2 rounded-md w-full break-all">
                  {body}
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
                        me 
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
    <div class="public_square_feed_input_buttons w-full grid grid-cols-2 grid-rows-1 gap-2">
      <button class="public_square_feed_input_buttons_send"
              on:click={() => {
                if (toSendBody.length) {
                  sendFeedMessage(toSendBody)
                  clr()
                }
              }}>
        SEND IT
      </button>
      <button class="public_square_feed_input_buttons_clr flex justify-around items-center"
              on:click={() => clr()}>
        CLR
      </button>
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .public_square_feed
    grid-template-areas: "messages" "input"
    grid-template-rows: max(250px, 60%) 40%
    &_messages
      grid-area: "messages"
      &_container
        background-color: color.$grey-transp
    &_input
      grid-area: "input"
      grid-template-areas: "body" "keyboard" "buttons"
      grid-template-rows: 2em max(60%, 155px) 1fr
      &_emoji_keyboard
        grid-area: "keyboard"
        border: 3px color.$yellow-3 double
      &_body
        grid-area: "body"
        border: 3px color.$yellow-3 double
      &_buttons
        grid-area: "buttons"
        grid-template-columns: 70% 30%
        &_send
          border: 3px color.$green-4 double
        &_clr
          border: 3px color.$red-4 double
      
  .feed_message
    background-color: color.$blue-7
    grid-template-areas: "user body"
    grid-template-columns: 40% 60%
    &_user
      grid-area: "user"
      background-color: color.$purple-7
      height: fit-content
    &_body
      grid-area: "body"
      background-color: color.$black-7
    &_self
      background-color: color.$yellow-3
      display: grid
      grid-template-columns: 50% 50%
      grid-template-rows: 1fr
      gap: 4px
    &_other
      background-color: color.$green-4
      display: grid
      grid-template-columns: 50% 50%
      grid-template-rows: 1fr
      gap: 4px
    &_system
      background-color: color.$blue-7
      display: flex
      flex-direction: column
      justify-content: space-around
      align-items: stretch
</style>
