<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import Loading from "./lib/Loading.svelte";

  import Icon from './lib/Icon.svelte'
  import { Icons } from './lib/Icons.js'

  import { lang } from "./stores/lang"
  import { I18n, Lang } from "./I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  const openLinkInNewTab = (href: string) => {
    window.open(href, '_blank');
  }

  interface GithubRepoLangBreakdown { [key: String]: number }

  enum SortColumns {
    PROJECT = "PROJECT",
    LANGUAGES = "LANGUAGES",
    DESCRIPTION = "DESCRIPTION",
    LAST_PUSH = "LAST_PUSH",
    UPDATED_AT = "UPDATED_AT",
    CREATED_AT = "CREATED_AT"
  }

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
    w-screen
    overflow-none
    p-5
  ">
  <div
    class="
      h-full w-full
      grid grid-rows-2 grid-cols-1 gap-4
      section_grid
    ">
    <div
      class="
        w-full
        text-xl
        font-extrabold
        grid grid-rows-1 grid-cols-2 gap-4
        repos_banner_grid
      "
      data-testid="
        repos_banner
      ">
      <div
        class="
          flex flex-col justify-around
        ">
        <h2>
          {i18n.t("personalProejects", langVal)}
        </h2>
        <select
          placeholder="Sort By">
          {#each Object.values(SortColumns) as col}
            <option
              value={col}>
              {col}
            </option>
          {/each}
        </select>
      </div>
      <span
        class="
          bg-emerald-500
        ">
        Graph
      </span>
    </div>
    <div
      class="
        border-solid border-5 border-cyan-500
        overflow-y-scroll
        row-span-10 col-span-8
        pb-72
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
            <th
              class="
                w-56
              ">
              Project </th>
            <th
              class="
                w-16
              ">
              Code
            </th>
            <th
              class="
                w-72
              ">
              Languages
            </th>
            <th
              class="
                w-96
              ">
              Description
            </th>
            <th>
              Last Push
            </th>
            <th>
              Updated At
            </th>
            <th>
              Created At
            </th>
          </tr>
          {#each githubRepos as { created_at, description, html_url, languageBreakdown, name, pushed_at, updated_at }}
            <tr
              class="
                h-24
                border-4 border-cyan-500
              ">
              <td
                class="
                  font-bold
                  text-left
                  pl-5
                ">
                {name}
              </td>
              <td>
                <button
                  class="
                    w-fit h-full
                  "
                  title="See It On Github"
                  on:click={() => openLinkInNewTab(html_url)}>
                  <Icon icon={Icons.GithubSolid} />
                </button>
              </td>
              <td
                class="
                  w-full
                  px-5
                  overflow-hidden
                  project_description
                  text-left text-wrap
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
                {updated_at.toLocaleString().split(',')[0]}
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
  .repos_banner_grid
    grid-template-columns: 15% 85%
  .section_grid
    grid-template-rows: 20% 80%
</style>
