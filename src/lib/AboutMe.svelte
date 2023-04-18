<script lang="ts">
  
  interface PersonalProject {
    title: string,
    description: string,
    href: string
  }
  let personal_projects: PersonalProject[] = [
    {
      title: 'monthly_budget',
      description: 'CSV Processing in Rust',
      href: 'https://github.com/nathanielBellamy/monthly_budget'
    },
    {
      title: 'rustby',
      description: 'Inject Rust Optimizations into Ruby',
      href: 'https://github.com/nathanielBellamy/rustby'
    },
    {
      title: 'trow',
      description: 'Multi-App React Redux Architecture in Typescript',
      href: 'https://github.com/nathanielBellamy/trow'
    },
    {
      title: 'polynomial_console_graph',
      description: 'ASCII Graph Polynomials Using C++',
      href: 'https://github.com/nathanielBellamy/PolynomialConsoleGraph'
    },
    {
      title: 'give_me_a_sine',
      description: 'ASCII Graph Sinusoidals Using Rust',
      href: 'https://github.com/nathanielBellamy/give_me_a_sine'
    }
  ]

  interface GeneralFact {
    title: string,
    description: string
  }

  interface ProfessionalThing {
    title: string,
    description: string
  }

  let prefessional_things: ProfessionalThing[] = [
    {
      title: 'Ruby',
      description: 'Backend Rails'
    },
    {
      title: 'JS',
      description: 'React, Vue, Node, JQuery'
    },
    {
      title: 'Heroku',
      description: 'Dynos and stuff'
    }
  ]

</script>

<div class="about_me flex flex-col justify-start items-stretch">
  <div class="section grid grid-cols-10">
    <div class="section_title text-xl font-extrabold col-span-2">
      Personal Projects
    </div>
    <div class="section_body col-span-8">
      {#each personal_projects as { title, description, href } }
        <div class="project grid grid-cols-4">
          <a class="project_title"
             title="See It On Github"
             target="_blank"
             href={href}>
            {title}
          </a>
          <div class="project_description col-span-3">
            {description} 
          </div>
        </div>
      {/each}
    </div>
  </div>
  <div class="section grid grid-cols-10">
    <div class="section_title text-xl font-extrabold col-span-2">
      Professional Things
    </div>
    <div class = "section_body col-span-8">
      {#each prefessional_things as { title, description } }
        <div class="project grid grid-cols-4">
          <div class="project_title">
            {title}
          </div>
          <div class="project_description col-span-3">
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
    border-top: 5px solid color.$green-4
    border-bottom: 5px solid color.$green-4
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
      border-radius: 5px
      &:hover
        background-color: color.$blue-3
        transition: background-color .25s

    &_description
      flex-grow: .9
      text-align: center
      display: flex
      align-items: center
      text-align: left
      padding-left: 1em
      font-weight: 700
      color: color.$white
      margin: 2px 5px 2px 5px
      border-bottom: 3px solid color.$green-4
      border-left: 3px solid color.$green-4

      border-radius: 5px

      
</style>
