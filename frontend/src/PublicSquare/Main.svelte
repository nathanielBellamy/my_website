<script lang="ts">
  import { onDestroy } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'

  let curr_mess: string
  
  let feed: string[] = []
  $: _feed = [curr_mess, ...feed]
  function pushToFeed(s: string) {
    feed.push(s)
  }

  const ws = new WebsocketBuilder('ws://localhost:8080/ws')
      .onOpen((i, ev) => { console.log("opened") })
      .onClose((i, ev) => { console.log("closed") })
      .onError((i, ev) => { console.log("error") })
      .onMessage((i, ev) => { 
        console.log("message")
        pushToFeed(ev.data)
        curr_mess = ev.data
      })
      .onRetry((i, ev) => { console.log("retry") })
      .build()

  onDestroy(() => ws.close())
</script>

<div class="w-full h-full flex justify-between items-stretch">
  <div class="grow flex flex-col justify-between items-stretch">
    Curr Mess: {curr_mess}
    <input bind:value={curr_mess}/>
    <button on:click={() => ws.send(curr_mess)}>
      SEND IT
    </button>
  </div>
  <div class="grow flex flex-col-reverse justify-between items-stretch">
    {#each _feed as hist, i} 
      <div class="grow">
        {i}:::{hist}
      </div>
    {/each}
  </div>
</div>

<style lang="sass">

</style>
