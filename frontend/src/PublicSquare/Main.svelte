<script lang="ts">
  import { onDestroy } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'

  interface Message {
    clientId: number,
    body: string
  }

  let curr_message_body: string
  $: curr_mess = {type: 1, body: curr_message_body}
  
  let feed: Message[] = []
  $: _feed = [curr_mess, ...feed]

  const FEED_LENGTH: number = 926
  function pushToFeed(m: Message) {
    if (feed.length > FEED_LENGTH) {
      feed.shift()
    }
    feed.push(m)
  }

  const ws = new WebsocketBuilder('ws://localhost:8080/ws')
      .onOpen((i, ev) => { console.log("opened") })
      .onClose((i, ev) => { console.log("closed") })
      .onError((i, ev) => { console.log("error") })
      .onMessage((i, ev) => { 
        console.log("message")
        const message: Message = JSON.parse(ev.data)
        pushToFeed(message)
        curr_message_body = message.body
        curr_mess = {clientId: message.clientId, body: message.body}
      })
      .onRetry((i, ev) => { console.log("retry") })
      .build()

  onDestroy(() => ws.close())
</script>

<div class="w-full h-full flex justify-between items-stretch">
  <div class="grow flex flex-col justify-between items-stretch">
    Curr Mess: {curr_message_body}
    <input bind:value={curr_message_body}/>
    <button on:click={() => ws.send(curr_mess.body)}>
      SEND IT
    </button>
  </div>
  <div class="grow p-5 m-5 overflow-y-scroll flex flex-col items-stretch">
    {#each _feed as { clientId, body }, i} 
      {#if !!i}
        <div class="grow p-5 m-5 flex justify-between items-center">
          <h4>
            {i}
          </h4>
          <div>
            {`User ${clientId}`}:
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
