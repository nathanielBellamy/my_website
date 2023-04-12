<script lang="ts">
  import AboutMe from "./lib/AboutMe.svelte"
  import MagicSquare from "./lib/MagicSquare.svelte"
  import Title from "./lib/Title.svelte"

  const setCurrentSection = (newSection: string) => {
    currentSection = newSection
  }

  let currentSection: string = "aboutMe"
</script>

<main class="main rounded-md text-xl flex flex-col justify-start">
  <div class="title font-bold text-xl p-7 underline text-left">
    Nate B. Schieber's Website  
  </div>
  <div class="main_header flex flex-row gap-x-10 justify-start">
    <div class="title p-7 text-xl font-bold underline">
      Things You Can Find Here:
    </div>
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
    background: color.$green-grad
  
  .section_button
    padding: 0 1em 0 1em
    margin: .25em
    font-weight: 700
  
  .current_section
    background-color: color.$grey-3
    color: black
    transition: color .5s, background-color .5s
    
  .main
    width: 100vw
    height: 100vh
    &_header
      width: 100%
      border: 2px solid black
      background: color.$black-grad
      &_title
        max-width: 400px
        text-align: center
    &_body
      background: color.$blue-grad
      flex-grow: 1
</style>
