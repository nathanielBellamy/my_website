<script lang="ts">
  import AboutMe from "./lib/AboutMe.svelte"
  import MagicBanner from "./lib/MagicBanner.svelte"
  import MagicSquare from "./lib/MagicSquare.svelte"
  import Title from "./lib/Title.svelte"

  const setCurrentSection = (newSection: string) => {
    currentSection = newSection
  }

  let currentSection: string = "aboutMe"
</script>

<main class="main rounded-md text-xl flex flex-col justify-start">
  <div class="main_header grid grid-cols-10">
    <div class="title p-3 text-lg font-bold col-span-2">
      It's A Website  
    </div>
    <div class="col-span-8">
      <MagicBanner />
    </div>
  </div>
  <div class="section_select flex flex-row justify-start items-stretch">
    <button class="section_button"
            class:current_section={currentSection == 'aboutMe'}
            on:click={()=>setCurrentSection("aboutMe")}>
      About Me
    </button>
    <button class="section_button"
            class:current_section={currentSection == 'magicSquare'}
            on:click={()=>setCurrentSection("magicSquare")}>
      Magic Square
    </button>
  </div>
  <div class="main_body">
    {#if currentSection == "magicSquare"}
      <Title title="Magic Square"/>
      <MagicSquare />
    {:else}
      <div>
        <Title title="About Me"/>
        <AboutMe />
      </div>
    {/if}
  </div>
</main>

<style lang="sass">
  @use "./styles/color"
  
  .title
    color: color.$white
    font-weight: 900
    font-size: 1.25em
    background: color.$blue-black-grad
    flex-grow: 1
  
  .section_select
    background: color.$black-grad

  .current_section
    background: color.$grey-grad
    color: color.$black-4
    transition: color .5s, background-color .5s
    
  .main
    width: 100vw
    height: 100vh
    &_header
      border: 2px solid black
      background: color.$black-grad
      overflow: hidden
      max-height: 100px

    &_body
      background: color.$blue-grad
      flex-grow: 1
</style>
