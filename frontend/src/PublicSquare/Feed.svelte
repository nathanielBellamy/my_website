<script lang="ts">
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import type { FeedMessage } from './FeedMessage'
  import { psFeed } from '../stores/psFeed'
  import EmojiKeyboard from '../lib/EmojiKeyboard.svelte';

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

  let nextEmoji: string = "😎"
  let emojiSetIndicator: boolean = false

  export let toSendBody: string = ""

  $: if (emojiSetIndicator || !emojiSetIndicator) {
    toSendBody = toSendBody + nextEmoji
  }

  $: if (toSendBody.length > 12){
       toSendBody = toSendBody.slice(2)
     } else {
       toSendBody = toSendBody
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

<div class="public_square_feed w-full h-full p-2 grid grid-cols-1 grid-rows-2">
  <div  id="public_square_feed_messages_container"
        class="h-full overflow-y-scroll">
    <div  id="public_square_feed_messages"
          class="public_square_feed_messages h-fit p-2 flex flex-col items-center gap-2">
      {#each psFeedVal as { clientId, body }, i} 
        {#if !!i}
          <div class="feed_message w-full h-fit rounded-md grid grid-cols-2 grid-rows-1 gap-2">
            <div class="feed_message_user text-sm">
              {formatClientId(clientId)}
            </div>
            <div class="feed_message_body font-bold text-lg">
              {body}
            </div>
          </div>
        {/if}
      {/each}
    </div>
  </div>
  <div class="public_square_feed_input p-2 h-full grid grid-rows-3 grid-cols-1 gap-2">
    <div class="public_square_feed_input_emoji_keyboard overflow-y-scroll">
      <EmojiKeyboard bind:value={nextEmoji}
                     bind:valueSetIndicator={emojiSetIndicator}/>
    </div>
    <div class="public_square_feed_input_body">
      {toSendBody}
    </div>
    <button class="public_square_feed_input_button"
            on:click={() => sendFeedMessage(toSendBody)}>
      SEND IT
    </button>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .public_square_feed
    grid-template-areas: "messages" "input"
    &_messages
      grid-area: "messages"
    &_input
      grid-area: "input"
      grid-template-areas: "keyboard" "body" "button"
      grid-template-rows: 70% 15% 15%
      &_emoji_keyboard
        grid-area: "keyboard"
      &_body
        grid-area: "body"
      &_button
        grid-area: "button"

  .feed_message
    background-color: color.$blue-transp
    grid-template-areas: "user body"
    grid-template-columns: "20% 80%"
    &_user
      grid-area: "user"
    &_body
      grid-area: "body"
</style>
