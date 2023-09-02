<script lang="ts">
  export let action: string
  export let title: string
  export let hasPassed: boolean = false

  function onClick(e: any) {
    e.preventDefault();
    grecaptcha.enterprise.ready(async () => {
      const token = await grecaptcha.enterprise.execute(
        import.meta.env.VITE_RECAPTCHA_SITE_KEY, 
        {action: 'LOGIN'}
      )
      await sendTokenToServer(token)
    })
  }

  interface RecaptchaPayload {
    action: String
    token: String
  }

  async function sendTokenToServer(token: string) {
    const payload: RecaptchaPayload = { action, token }
    await fetch('recaptcha', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload)
    })
    .then((res) => { hasPassed = res.status === 200 })
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

