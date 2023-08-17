<script lang="ts">
  let url: string
  let password: string

  async function authorize(password: string): any {
    await fetch('dev-auth', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: `pw=${password}`
    })
  }
</script>

<main class="h-full w-full">
  <div class="h-full w-full">
    <div class="w-full flex flex-col justify-between items-stretch text-left">
      <div class="text-left">
        Answer me these questions three to get across
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
    <input bind:value={password}>
  </form>
</main>

<style>

</style>
