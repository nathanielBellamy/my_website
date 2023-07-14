<script lang="ts">
  import { onDestroy } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'

  let curr_mess: string = "!NOT YET!"
  const ws = new WebsocketBuilder('ws://localhost:8080/ws')
      .onOpen((i, ev) => { console.log("opened") })
      .onClose((i, ev) => { console.log("closed") })
      .onError((i, ev) => { console.log("error") })
      .onMessage((i, ev) => { 
        console.log("message")
        console.dir(ev)
        curr_mess = ev.data
      })
      .onRetry((i, ev) => { console.log("retry") })
      .build()

  onDestroy(() => ws.close())
</script>

<div class="w-full h-full flex justify-around items-center">
  Curr Mess: {curr_mess}
  <input bind:value={curr_mess}/>
  <button on:click={() => ws.send(curr_mess)}>
    SEND IT
  </button>
</div>

<style lang="sass">

</style>
