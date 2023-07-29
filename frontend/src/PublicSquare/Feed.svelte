<script lang="ts">
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import type { FeedMessage } from './FeedMessage'
  import { psFeed } from '../stores/psFeed'
  import EmojiKeyboard from '../lib/EmojiKeyboard.svelte'

  let psFeedVal: FeedMessage[]
  const unsubPsFeed = psFeed.subscribe((val: FeedMessage[]) => psFeedVal = [...val])

  export let sendFeedMessage: (body: string) => void

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

  function formatClientId(id: number): string {
    var res: string = ""
    if (!!id) {
      res = `u-${id}:`
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

  // LIFECYCLE
  onMount(() => scrollFeedToBottom())
  onDestroy(() => unsubPsFeed())

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
      {#each psFeedVal as { clientId, body }, i} 
        {#if !!i}
          <div class="feed_message p-2 w-full h-fit rounded-md grid grid-cols-2 grid-rows-1 gap-2">
            <div class="feed_message_user text-sm h-full flex justify-around items-center">
              {formatClientId(clientId)}
            </div>
            <div class="feed_message_body font-bold text-lg mr-2 rounded-md">
              {body}
            </div>
          </div>
        {/if}
      {/each}
    </div>
  </div>
  <div class="public_square_feed_input p-2 h-full grid grid-rows-3 grid-cols-1 gap-4">
    <div class="public_square_feed_input_emoji_keyboard rounded-md overflow-y-scroll p-2">
      <EmojiKeyboard bind:value={nextEmoji}
                     bind:valueSetIndicator={emojiSetIndicator}/>
    </div>
    <div class="public_square_feed_input_body rounded-md flex justify-around items-center">
      {toSendBody}
    </div>
    <div class="public_square_feed_input_buttons w-full grid grid-cols-2 grid-rows-1 gap-2">
      <button class="public_square_feed_input_buttons_send"
              on:click={() => sendFeedMessage(toSendBody)}>
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
  @use "./../styles/color"
  @use "./../styles/text"

  .public_square_feed
    grid-template-areas: "messages" "input"
    grid-template-rows: 70% 30%
    &_messages
      grid-area: "messages"
      &_container
        background-color: color.$grey-transp
    &_input
      grid-area: "input"
      grid-template-areas: "keyboard" "body" "buttons"
      grid-template-rows: 50% 2em 1fr
      &_emoji_keyboard
        grid-area: "keyboard"
        border: color.$blue-7 2px solid
      &_body
        grid-area: "body"
        border: color.$blue-4 2px solid
      &_buttons
        grid-area: "buttons"
        grid-template-columns: 70% 30%
        /* &_send */
        /* &_clear */
      

  .feed_message
    background-color: color.$blue-7
    grid-template-areas: "user body"
    grid-template-columns: 40% 60%
    &_user
      grid-area: "user"
    &_body
      grid-area: "body"
      background-color: color.$black-7
</style>
