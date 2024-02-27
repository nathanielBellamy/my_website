<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import Loading from "./lib/Loading.svelte";

  import { lang } from "./stores/lang"
  import { I18n, Lang } from "./I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  const openLinkInNewTab = (href: string) => {
    window.open(href, '_blank');
  }

  interface GithubRepoLangBreakdown { [key: String]: number }

  interface GithubRepo {
    created_at: Date,
    description: String,
    html_url: String,
    language: String,
    languageBreakdown: GithubRepoLangBreakdown,
    name: String,
    pushed_at: Date,
    updated_at: Date,
  }


  let reposReady: boolean = false
  let githubRepos: GithubRepo[] = []
  function fetchGithubRepos() {
    const url: String = "https://api.github.com/users/nathanielBellamy/repos"
    fetch(url)
      .then((res) => res.json())
      .then(async (repos) => {
        const repoLangDict: { [key: String]: GithubRepoLangBreakdown[] } = {}
        const languagesPromises: Promise[] = repos.map(repo => {
          const repoLanguagesUrl = `https://api.github.com/repos/nathanielBellamy/${repo.name}/languages`
          return fetch(repoLanguagesUrl)
                   .then(res => res.json())
                   .then(res => repoLangDict[repo.name] = res)
        })

        await Promise.all(languagesPromises)

        // console.dir({ repoLangDict })
        githubRepos = repos.map(repo => {
          return {
            created_at: new Date(repo.created_at),
            description: repo.description,
            html_url: repo.html_url,
            language: repo.language,
            languageBreakdown: repoLangDict[repo.name],
            name: repo.name,
            pushed_at: new Date(repo.pushed_at),
            updated_at: new Date(repo.updated_at),
          }
        })
      })
      .then(() => { setTimeout(() => {reposReady = true}, 500) })
  }

  onMount(fetchGithubRepos)
  onDestroy(unsubLang)
</script>

<div
  class="
    h-screen
    overflow-none
  ">
  <div
    class="
      my-10
      h-full w-full
      grid grid-rows-10 grid-cols-10 gap-4
    ">
    <div
      class="
        section_title
        overflow-scroll
        text-xl font-extrabold
        row-span-10 col-span-2
      "
      data-testid="about_personal_projects">
      {i18n.t("personalProejects", langVal)}
    </div>
    <div
      class="
        border-solid border-5 border-cyan-500
        overflow-y-scroll
        row-span-10 col-span-8
      ">
      {#if !reposReady}
        <Loading />
      {:else}
        <table
          class="
            w-full
            border-collapse
            table-fixed
          ">
          <tr
            class="
              h-16
            ">
            <th> Project </th>
            <th> Languages </th>
            <th> Description </th>
            <th> Last Push </th>
            <th> Created At </th>
          </tr>
          {#each githubRepos as { created_at, description, html_url, languageBreakdown, name, pushed_at, updated_at }}
            <tr
              class="
                h-24
                border-4 border-cyan-500
              ">
              <td>
                <button
                  class="
                    project_title
                    w-full h-full
                  "
                  title="See It On Github"
                  on:click={() => openLinkInNewTab(html_url)}>
                  {name}
                </button>
              </td>
              <td
                class="
                  w-full
                  overflow-hidden
                  project_description
                  text-wrap
                ">
                <p>{Object.keys(languageBreakdown).join(', ')}</p>
              </td>
              <td
                class="
                  ml-2 mr-2
                  text-left
                ">
                {description}
              </td>
              <td>
                {pushed_at.toLocaleString().split(',')[0]}
              </td>
              <td>
                {created_at.toLocaleString().split(',')[0]}
              </td>
            </tr>
          {/each}
        </table>
      {/if}
    </div>
  </div>
</div>

<style lang="sass">
  @use "./styles/color"
</style>
