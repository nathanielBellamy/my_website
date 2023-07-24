<script lang="ts">
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import type { FeedMessage } from './FeedMessage'
  import { psFeed } from '../stores/psFeed'

  let psFeedVal: FeedMessage[]
  const unsubPsFeed = psFeed.subscribe((val: FeedMessage[]) => psFeedVal = [...val])

  export let sendFeedMessage: (body: string) => void
  export let toSendBody: string

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

  function formatClientId(id: number): string {
    var res: string = ""
    if (!!id) {
      res = `u-${id}:`
    }

    return res
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

<div class="w-full h-full p-2 grid grid-cols-1 grid-rows-2">
  <div  id="public_square_feed"
        class="grow p-2 overflow-y-scroll flex flex-col items-stretch">
    {#each psFeedVal as { clientId, body }, i} 
      {#if !!i}
        <div class="grow grid grid-cols-2 grid-rows-1 gap-2 overflow-y-scroll">
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
  <div class="grow p-2 grid grid-rows-2 grid-cols-1">
    <input bind:value={toSendBody}/>
    <button on:click={() => sendFeedMessage(toSendBody)}>
      SEND IT
    </button>
  </div>
</div>

<style lang="sass">
 /* TODO: template feed greed areas */
</style>
