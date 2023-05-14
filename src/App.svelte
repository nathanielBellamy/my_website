<script lang="ts">
  import AboutMe from "./lib/AboutMe.svelte"
  import MagicBanner from "./lib/MagicBanner.svelte"
  import Title from "./lib/Title.svelte"
  import Wasm from "./lib/Wasm.svelte"
  
  const setCurrentSection = (newSection: string) => {
    currentSection = newSection
  }

  let currentSection: string = "aboutMe"// "magicSquare"
</script>

<main class="main rounded-md flex flex-col justify-start">
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
      <Wasm program="magicSquare"/>
    {:else}
      <div>
        <Title title="About Me"
               subTitle="nbschieber@gmail.com -- Software Engineer -- Portland, OR"/>
        <AboutMe />
      </div>
    {/if}
  </div>
</main>

<style lang="sass">
  @use "./styles/color"
  
  .title
    color: color.$white
    border: 5px solid color.$blue-4
    border-radius: 5px
    margin-bottom: 5px
    font-weight: 900
    font-size: 1.25em
    flex-grow: 1
    text-align: center
    display: flex
    flex-direction: column
    justify-content: space-around
  
  .current_section
    background: color.$blue-black-grad
    color: color.$cream
    transition: color .5s, background .5s
    
  .main
    width: 100vw
    height: 100vh
    &_header
      border: 2px solid black
      background: color.$black-grad
      overflow: hidden
      max-height: 100px
      min-height: 100px

    &_body
      padding-bottom: 3em
      margin-bottom: 5px
      height: calc(100vh - 200px)
</style>
