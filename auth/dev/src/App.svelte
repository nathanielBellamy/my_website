<script lang="ts">
  import { Drawer, Button, CloseButton } from 'flowbite-svelte'
  import { sineIn } from 'svelte/easing'

  let nopeHidden: boolean = true
  let transitionParams: any = {
    x: -320,
    duration: 200,
    easing: sineIn
  }

  let password: string

  async function authorize(password: string) {
    await fetch('dev-auth', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: `pw=${password}`
    })
    .then((res) => {
      if (res.ok) {
        window.location.href = "/"
      } else {
        nopeHidden = false
      }
    })
  }

  function onInput() {
    nopeHidden = true
  }
</script>

<main class="h-full w-full grid grid-cols-1 grid-rows-2">
  <div class="h-full w-full">
    <div class="w-full flex flex-col justify-between items-stretch text-left">
      <div class="text-left">
        Answer me these questions three to get across --
      </div>
      <div>
        Who's the boss? What size is them shoes of yours?
      </div>
      <div>
        We the wrong crew to cross, true or false?
      </div>
    </div>
  <div>
    - MF DOOM
  </div>
    <form on:submit={(e) => {
      e.preventDefault()
      authorize(password)
    }}>
    <input type="password"
           bind:value={password}
           on:input={onInput}>
  </form>
  <Drawer transitionType="fly" {transitionParams} 
          bind:hidden={nopeHidden} 
          placement="bottom"
          id="nope">
    <div class="h-32 w-full">
        <h3 class="nope text-amber-800">
          Nice Try
        </h3>
    </div>
  </Drawer>
</main>

<style>

</style>
