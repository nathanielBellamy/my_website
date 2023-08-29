<script lang="ts">
  import { onMount } from 'svelte'
  export let hasPassed: boolean = false
  var onloadCallback = function() {
    grecaptcha.render('html_recaptcha_element', {
      'sitekey' : import.meta.env.VITE_GRECAPTCHA_KEY
    });
  };

  let error: string
  let token: string

  async function handleCaptchaCallback() {
    await fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            recaptchaToken: token
        })
    })

    
    // reset recaptcha for future requests
    resetCaptcha()
  }

  function handleCaptchaError() {
    error = 'Recaptcha error. Please reload the page'
  }

  function resetCaptcha() {
    window.grecaptcha.reset()
  }

  const handleSubmit = () => {
    // reset any errors
    error = ''
    
    // tell recaptcha to process a request
    window.grecaptcha.execute()
  }

  onMount(() => {
    window.handleCaptchaCallback = handleCaptchaCallback;
    window.handleCaptchaError = handleCaptchaError;
    window.resetCaptcha = resetCaptcha;
  })
</script>

<body>
  <!-- TODO -->
  <!--   - create Go endpoint and have this hit it -->
  <form action="?" method="POST">
    {#if error}
      <div>
          <small class="text-yellow-300 font-bold">{error}</small>
      </div>
    {/if}
    <div class="g-recaptcha"
         data-sitekey={import.meta.env.VITE_RECAPTCHA_SITE_KEY}
         data-callback="handleCaptchaCallback"
         data-expired-callback="resetCaptcha"
         data-error-callback="handleCaptchaError"
         data-size="invisible"/>
    <div id="html_recaptcha_element"></div>
    <br>
    <input type="submit" value="Submit">
  </form>
  <script src="https://www.google.com/recaptcha/api.js?onload=onloadCallback&render=explicit"
      async defer>
  </script>
</body>

