<script lang="ts">
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'

  interface Message {
    clientId: number,
    body: string
  }

  let toSendBody: string = ""
  $: toSend = {clientId: 1, body: toSendBody}

  let toReceive: Message | null = null
  
  let feed: Message[] = []
  $: _feed = !toReceive ? [...feed] : [toReceive, ...feed]
  let feedWasScrolledToBottom: boolean = false
  
  function scrollFeedToBottom() {
    const feed = document.getElementById("public_square_feed")
    feed.scrollTop = feed.scrollHeight
  }

  function updateFeedIsScrolledToBottom() {
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

  const FEED_LENGTH: number = 926
  function pushToFeed(m: Message) {
    if (feed.length > FEED_LENGTH) {
      feed.shift()
    }
    feed.push(m)
  }

  function formatClientId(id: number): string {
    var res: string = ""
    if (!!id) {
      res = `:: User ${id} :: `
    }

    return res
  }


  const ws = new WebsocketBuilder('ws://localhost:8080/ws')
      .onOpen((i, ev) => { console.log("opened") })
      .onClose((i, ev) => { console.log("closed") })
      .onError((i, ev) => { console.log("error") })
      .onMessage((i, ev) => { 
        console.log("message")
        const message: Message = JSON.parse(ev.data)
        pushToFeed(message)
        toReceive = message
      })
      .onRetry((i, ev) => { console.log("retry") })
      .build()


  // LIFECYCLE

  onMount(() => scrollFeedToBottom())
  onDestroy(() => ws.close())

  beforeUpdate(() => {
    feedWasScrolledToBottom = updateFeedIsScrolledToBottom()
  })
  afterUpdate(() => {
    if (feedWasScrolledToBottom) scrollFeedToBottom()
  })
</script>

<div class="w-full h-full p-2 flex justify-between items-stretch">
  <div class="grow p-5 m-5 flex flex-col justify-between items-stretch">
    <input bind:value={toSendBody}/>
    <button on:click={() => ws.send(toSend.body)}>
      SEND IT
    </button>
  </div>
  <div  id="public_square_feed"
        class="grow p-5 m-5 overflow-y-scroll flex flex-col items-stretch">
    {#each _feed as { clientId, body }, i} 
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
