<script lang="ts">
  import { onMount } from 'svelte'
  export let hasPassed: boolean = false

  function onClick(e) {
    e.preventDefault();
    grecaptcha.enterprise.ready(async () => {
      const token = await grecaptcha.enterprise.execute(
        import.meta.env.VITE_RECAPTCHA_SITE_KEY, 
        {action: 'LOGIN'}
      )
    })
  }
</script>

<body>
  <h1> RECAP </h1>
  <!-- TODO -->
  <!--   - create Go endpoint and have this hit it -->
  <form action="?" method="POST">
    <button on:click={onClick}>
      Verify
    </button>
  </form>
</body>

