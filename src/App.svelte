<script lang="ts">
  import MagicSquare from "./lib/MagicSquare.svelte"

  const setCurrentSection = (newSection: string) => {
    currentSection = newSection
  }

  let currentSection: string = "main"
</script>

<main class="main rounded-md text-xl flex flex-col justify-start">
  <div class="main_header rounded-md flex flex-row justify-start">
    <button class="p-5 m-5"
            on:click={()=>setCurrentSection("main")}>
      It's My Website
    </button>
    <button class="p-5 m-5"
            on:click={()=>setCurrentSection("magicSqaure")}>
      Magic Square
    </button>
  </div>
  <div class="main_body rounded-md">
    {#if currentSection == "magicSqaure"}
      <MagicSquare />
    {:else}
      <div>
        My Website
      </div>
    {/if}
  </div>
</main>

<style lang="sass">
  .main
    width: 100vw
    height: 100vh
    background-color: lightblue
    &_header
      width: 100%
      border: 2px solid black
      &_title
        max-width: 400px
        text-align: center
        background-color: black
    &_body
      background-color: blue
      flex-grow: 1
</style>
