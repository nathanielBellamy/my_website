<script lang="ts">
  import { afterUpdate, beforeUpdate, onMount } from 'svelte'
  import type { ToasterProps } from '../lib/Toaster'
  import Toaster from '../lib/Toaster.svelte'
  import type { FeedMessage } from './FeedMessage'

  export let feed: FeedMessage[]
  export let sendFeedMessage: (body: string) => void
  export let showConnected: boolean
  export let toasts: ToasterProps[]
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
      res = `:: User ${id} :: `
    }

    return res
  }

  // LIFECYCLE
  onMount(() => scrollFeedToBottom())

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
    <button on:click={() => sendFeedMessage(toSendBody)}>
      SEND IT
    </button>
  </div>
  <div  id="public_square_feed"
        class="grow p-5 m-5 overflow-y-scroll flex flex-col items-stretch">
    {#each feed as { clientId, body }, i} 
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
