<script lang="ts">
  import WebSocket from 'ws'

  const ws = new WebSocket('ws://localhost:8080/public-square')

  ws.on('error', console.error)

  ws.on('open', function open() {
    ws.send('something');
  })

  let curr_mess: string

  ws.on('message', function message(data: string) {
    curr_mess = data
    console.log('received: %s', data)
  });
</script>

<div>
  Curr Muss: {curr_mess}
</div>

<style lang="sass">

</style>
