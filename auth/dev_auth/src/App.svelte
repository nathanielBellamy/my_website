<script lang="ts">
  let password: string

  async function authorize() {
    // TODO:
    //  POST password
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
  <form>
    <input bind:value={password}>
  </form>
</main>

<style>

</style>
