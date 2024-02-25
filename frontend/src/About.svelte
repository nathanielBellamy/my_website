<script lang="ts">
  import { onDestroy, onMount } from 'svelte'

  import { lang } from "./stores/lang"
  import { I18n, Lang } from "./I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  interface ProfessionalThing {
    title: string,
    description: string,
    href: string
  }

  let prefessional_things: ProfessionalThing[] = [
    {
      title: 'Ruby',
      description: 'on Rails, RBS, Steep, Rspec, Capybara',
      href: 'https://www.ruby-lang.org/en/'
    },
    {
      title: 'JS',
      description: 'React, Vue, Vue Testing Library, Node, JQuery, Mocha',
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

  interface GithubRepo {
    created_at: Date,
    description: String,
    html_url: String,
    language: String,
    name: String,
    pushed_at: Date,
    updated_at: Date,
  }

  let githubRepos: GithubRepo[] = []
  function fetchGithubRepos() {
    const url: String = "https://api.github.com/users/nathanielBellamy/repos"
    fetch(url)
      .then((res) => res.json())
      .then((repos) => {
        // console.dir({ json })
        githubRepos = repos.map(repo => {
          return {
            created_at: new Date(repo.created_at),
            description: repo.description,
            html_url: repo.html_url,
            language: repo.language,
            name: repo.name,
            pushed_at: new Date(repo.pushed_at),
            updated_at: new Date(repo.updated_at),
          }
        })
      })
  }

  onMount(() => {
    fetchGithubRepos()
  })
  onDestroy(unsubLang)
</script>

<div class="about_me flex flex-col justify-start items-stretch gap-2">
  <div class="section grid grid-rows-10 md:grid-cols-10 gap-4">
    <div class="section_title text-xl font-extrabold row-span-2 md:col-span-2 md:row-span-1"
         data-testid="about_personal_projects">
      {i18n.t("personalProejects", langVal)}
    </div>
    <div class="section_body row-span-8 md:col-span-8 md:row-span-1">
      <!-- {#each personal_projects as { title, description, href } } -->
      <!--   <div class="project grid grid-rows-1 md:grid-cols-4"> -->
      <!--     <button class="project_title" -->
      <!--             title="See It On Github" -->
      <!--             on:click={() => openLinkInNewTab(href)}> -->
      <!--       {title} -->
      <!--     </button> -->
      <!--     <div class="project_description ml-10 md:ml-0 row-span-3 md:col-span-3 md:row-span-1"> -->
      <!--       {description} -->
      <!--     </div> -->
      <!--   </div> -->
      <!-- {/each} -->
      {#each githubRepos as { created_at, description, html_url, language, name, pushed_at, updated_at }}
        <div
          class="
            flex justify-around
          ">
          <button
            class="
              project_title
            "
            title="See It On Github"
            on:click={() => openLinkInNewTab(html_url)}>
            {name}
          </button>
          <div
            class="
              project_description
              flex flex-col justify-around
            ">
            <p>{language}</p>
            <p>{description}<p>
            <p> Created At: {created_at.toLocaleString()} </p>
            <p> Pushed At: {pushed_at.toLocaleString()} </p>
            <p> Updated At: {updated_at.toLocaleString()} </p>
          </div>
        </div>
      {/each}
    </div>
  </div>
  <div class="section grid grid-rows-10 md:grid-cols-10 gap-4">
    <div class="section_title text-xl font-extrabold row-span-2 md:col-span-2 md:row-span-1"
         data-testid="about_technical_experience">
      {i18n.t("technicalExperience", langVal)}
    </div>
    <div class = "section_body row-span-8 md:col-span-8 md:row-span-1">
      {#each prefessional_things as { title, description, href } }
        <div class="project grid grid-rows-1 md:grid-cols-4">
          <button class="project_title"
                  title={`Open In a New Tab: ${href}`}
                  on:click={() => openLinkInNewTab(href)}>
            {title}
          </button>
          <div class="project_description ml-10 md:ml-0 row-span-3 md:col-span-3 md:row-span-1">
            {description}
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>

<style lang="sass">
  @use "./styles/color"

  .about_me
    width: 100%
    overflow-y: scroll

  .section
    background: color.$black-blue-horiz-grad
    padding: 10px 0 10px 0

    &_title
      display: flex
      justify-content: space-around
      align-items: center
      color: color.$white
      flex-grow: .1
      font-size: 1.25em
      padding: 0 5px 0 5px
      color: color.$blue-4

    &_body
      display: flex
      flex-direction: column
      justify-content: flex-start
      align-items: stretch
      color: color.$black
      flex-grow: .9
      padding: 5px 0 5px 0

  .project
    /* display: flex */
    /* justify-content: flex-start */
    align-items: stretch
    flex-grow: 1
    border-bottom: 2px solid color.$black

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
      border-bottom: 3px solid color.$blue-7
      border-right: 3px solid color.$blue-7
      border-top: 0px solid white
      border-left: 0px solid white
      border-radius: 5px
      color: color.$blue-4
      &:hover
        background-color: color.$blue-6
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
      color: color.$green-2
      border-bottom: 3px solid color.$green-7
      border-left: 3px solid color.$green-7

      border-radius: 5px


</style>
