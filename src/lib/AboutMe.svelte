<script lang="ts">
  import Embed from "./Embed.svelte"
  import GiveMeASine from "./GiveMeASine.svelte"
  import PolynomialConsoleGraph from "./PolynomialConsoleGraph.svelte"


  enum EmbeddedProgram {
    giveMeASign,
    polynomialConsoleGraph,
    none
  }

  let showEmbed: EmbeddedProgram = EmbeddedProgram.none

  interface PersonalProject {
    title: string,
    description: string,
    href: string,
    program: EmbeddedProgram
  }

  let personal_projects: PersonalProject[] = [
    {
      title: 'my_website (this)',
      description: 'Svelte, Typescript, Rust, WebAssembly, WebGL, Sass, Tailwind, Vite',
      href: 'https://github.com/nathanielBellamy/my_website',
      program: EmbeddedProgram.none
    },
    {
      title: 'monthly_budget',
      description: 'CSV Processing with Rust',
      href: 'https://github.com/nathanielBellamy/monthly_budget',
      program: EmbeddedProgram.none
    },
    {
      title: 'rustby',
      description: 'Inject Rust Optimizations into Ruby',
      href: 'https://github.com/nathanielBellamy/rustby',
      program: EmbeddedProgram.none
    },
    {
      title: 'trow',
      description: 'Multi-App React Redux Architecture in Typescript',
      href: 'https://github.com/nathanielBellamy/trow',
      program: EmbeddedProgram.none
    },
    {
      title: 'polynomial_console_graph',
      description: 'ASCII Graph Polynomials Using C++',
      href: 'https://github.com/nathanielBellamy/PolynomialConsoleGraph',
      program: EmbeddedProgram.polynomialConsoleGraph
    },
    {
      title: 'give_me_a_sine',
      description: 'ASCII Graph Sinusoidals Using Rust',
      href: 'https://github.com/nathanielBellamy/give_me_a_sine',
      program: EmbeddedProgram.giveMeASign
    }
  ]

  interface ProfessionalThing {
    title: string,
    description: string,
    href: string
  }

  let prefessional_things: ProfessionalThing[] = [
    {
      title: 'Ruby',
      description: 'on Rails, Rspec, Capybara',
      href: 'https://www.ruby-lang.org/en/'
    },
    {
      title: 'JS',
      description: 'React, Vue, Node, JQuery, Mocha',
      href: 'https://developer.mozilla.org/en-US/docs/Web/JavaScript'
    },
    {
      title: 'C#',
      description: '.NET',
      href: 'https://learn.microsoft.com/en-us/dotnet/csharp/'
    },
    {
      title: 'SQL',
      description: 'LINQ, ActiveRecord, Postgres',
      href: 'https://www.postgresql.org/'
    },
    {
      title: 'Heroku',
      description: 'Site Hosting',
      href: 'https://www.heroku.com/?'
    },
    {
      title: 'CircleCI',
      description: 'Testing and Deployment Pipelines',
      href: 'https://circleci.com/'
    },
    {
      title: 'Github',
      description: 'github.com/nathanielBellamy, Actions',
      href: 'https://github.com/'
    },
    {
      title: 'Azure',
      description: 'Cognitive Search',
      href: 'https://azure.microsoft.com/en-us'
    },
    {
      title: 'AWS S3',
      description: 'Prod Data Migrations',
      href: 'https://docs.aws.amazon.com/s3/?icmpid=docs_homepage_featuredsvcs'
    },
    {
      title: 'Postman',
      description: 'Api Stress Testing with Newman',
      href: 'https://www.postman.com/'
    },
    {
      title: 'Bootstrap',
      description: 'Responsive UI Design',
      href: 'https://getbootstrap.com/docs/3.4/css/'
    },

  ]

  const openLinkInNewTab = (href: string) => {
    window.open(href, '_blank');
  }

</script>

<div class="about_me flex flex-col justify-start items-stretch">
  <div class="section grid grid-rows-10 md:grid-cols-10 ">
    <div class="section_title text-xl font-extrabold row-span-2 md:col-span-2 md:row-span-1">
      Personal Projects
    </div>
    <div class="section_body row-span-8 md:col-span-8 md:row-span-1">
      {#each personal_projects as { title, description, href, program } }
        <div class="project grid grid-rows-1 md:grid-cols-4">
          <button class="project_title"
                  title="See It On Github"
                  on:click={() => openLinkInNewTab(href)}>
            {title}
          </button>
          <div class="project_description row-span-3 md:col-span-3 md:row-span-1">
            {description}
            {#if program != EmbeddedProgram.none}
              <button on:click={() => {
                if (showEmbed == program) {
                  showEmbed = EmbeddedProgram.none
                } else {
                  showEmbed = program
                }
              }}>
                Open WebAssembly Build
              </button>
            {/if}
          </div>
        </div>
      {/each}
    </div>
  </div>
  <Embed>
    {#if showEmbed == EmbeddedProgram.giveMeASign}
      <GiveMeASine />
    {:else if showEmbed == EmbeddedProgram.polynomialConsoleGraph}
      <PolynomialConsoleGraph />
    {/if}
  </Embed>
  <div class="section grid grid-rows-10 md:grid-cols-10">
    <div class="section_title text-xl font-extrabold row-span-2 md:col-span-2 md:row-span-1">
      Technical Knowledge
    </div>
    <div class = "section_body row-span-8 md:col-span-8 md:row-span-1">
      {#each prefessional_things as { title, description, href } }
        <div class="project grid grid-rows-1 md:grid-cols-4">
          <button class="project_title"
                  title={`Open In a New Tab: ${href}`}
                  on:click={() => openLinkInNewTab(href)}>
            {title}
          </button>
          <div class="project_description row-span-3 md:col-span-3 md:row-span-1">
            {description} 
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"

  .about_me
    width: 100%
    height: 100%

  .section
    align-items: stretch
    min-height: 200px
    border-top: 3px solid color.$yellow-4
    border-bottom: 3px solid color.$yellow-4
    margin-top: 5px
    margin-bottom: 5px

    &_title
      display: flex
      justify-content: space-around
      align-items: center
      color: color.$white
      flex-grow: .1
      font-size: 1.25em
      padding: 0 5px 0 5px
    
    &_body
      display: flex
      flex-direction: column
      justify-content: flex-start
      align-items: stretch
      color: color.$black-4
      flex-grow: .9
      padding: 5px 0 5px 0
   
  .project
    /* display: flex */
    /* justify-content: flex-start */
    align-items: stretch
    flex-grow: 1
    border-bottom: 2px solid color.$black-4
    
    &_title
      flex-grow: .1
      transition: background-color .25s
      text-align: left
      color: color.$white
      font-weight: 700
      padding: 0 5px 0 5px
      margin: 2px 5px 2px 5px
      cursor: pointer
      overflow-x: hidden
      border-bottom: 3px solid color.$blue-4
      border-right: 3px solid color.$blue-4
      border-top: 0px solid white
      border-left: 0px solid white
      border-radius: 5px
      &:hover
        background-color: color.$blue-3
        transition: background-color .25s

    &_description
      flex-grow: .9
      text-align: center
      display: flex
      align-items: center
      justify-content: space-between
      flex-wrap: wrap
      text-align: left
      padding-left: 1em
      font-weight: 700
      color: color.$white
      margin: 2px 5px 2px 5px
      border-bottom: 3px solid color.$green-4
      border-left: 3px solid color.$green-4

      border-radius: 5px

      
</style>
