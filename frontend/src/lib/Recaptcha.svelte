<script lang="ts">
  import { onMount } from 'svelte'
  export let title: string
  export let hasPassed: boolean = false

  function onClick(e) {
    e.preventDefault();
    grecaptcha.enterprise.ready(async () => {
      const token = await grecaptcha.enterprise.execute(
        import.meta.env.VITE_RECAPTCHA_SITE_KEY, 
        {action: 'LOGIN'}
      )
      const res = await sendTokenToServer(token)
    })
  }

  async function sendTokenToServer(token: string) {
    await fetch('recaptcha', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({token})
    })
    .then((res) => {
      console.log(res)
    })
  }
</script>

<button on:click={onClick}
        class="recaptcha_button font-mono">
  {title}
</button>

<style lang="sass">
  .recaptcha_button
    min-height: 50px
</style>

