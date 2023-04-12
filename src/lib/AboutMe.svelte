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

</script>

<div class="about_me flex flex-col justify-start items-stretch">
  <div class="section">
    <div class="section_title text-xl font-extrabold">
      Personal Projects
    </div>
    <div class="section_body">
      {#each personal_projects as { title, description, href } }
        <div class="project grid grid-cols-4">
          <a class="project_title"
             title="See It On Github"
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
</div>

<style lang="sass">
  @use "./../styles/color"

  .about_me
    width: 100%
    height: 100%

  .section
    display: flex
    justify-content: flex-start
    align-items: stretch
    min-height: 200px

    &_title
      display: flex
      justify-content: space-around
      align-items: center
      background: color.$black-grad
      color: color.$white
      flex-grow: .1
      font-size: 1.25em
      padding: 0 5px 0 5px
    
    &_body
      display: flex
      flex-direction: column
      justify-content: flex-start
      align-items: stretch
      background-color: color.$grey-1
      color: color.$black-4
      flex-grow: .9

  .project
    /* display: flex */
    /* justify-content: flex-start */
    align-items: stretch
    flex-grow: 1
    border-bottom: 2px solid color.$black-4
    
    &_title
      display: flex
      justify-content: space-around
      align-items: center
      flex-grow: .1
      background-color: color.$blue-4
      transition: background-color .25s
      color: color.$white
      font-weight: 900
      padding: 0 5px 0 5px
      cursor: pointer
      &:hover
        background-color: color.$blue-3
        transition: background-color .25s

    &_description
      flex-grow: .9
      text-align: center
      background: color.$yellow-grad
      display: flex
      align-items: center
      text-align: left
      padding-left: 1em
      font-weight: 700
</style>
