<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import Icon from '../lib/Icon.svelte'
  import { Icons } from '../lib/Icons.js'
  import Loading from "../lib/Loading.svelte"
  import ReposBanner from "./ReposBanner.svelte"
  import ReposTableHeader from "./ReposTableHeader.svelte"
  import UserLangSummaryChart from "../integrations/github/UserLangSummaryChart.svelte"
  import {
    type GithubRepo,
    type GithubRepos,
    type GithubStore,
    SortColumn,
    SortOrder
  } from "../integrations/github/GithubTypes"
  import ColorCircle from "../integrations/github/ColorCircle.svelte"

  import { lang } from "../stores/lang"
  import { I18n, Lang } from "../I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import GithubIntegration from "../integrations/github/GithubIntegration"
  import { githubStore } from "../stores/githubStore"
  let githubReposVal: GithubRepos = []
  const unsubGithubStore = githubStore.subscribe((store: GithubStore) => githubReposVal = [...store.repos])
  let github: GithubIntegration = new GithubIntegration(githubStore, updateReposReady)

  const openLinkInNewTab = (href: string) => {
    window.open(href, '_blank');
  }

  function setSortOrder(order: String): void {
    github.sortOrder = order;
    github.sortReposBy()
  }

  $: [...github.reposVal]

  let reposReady: boolean = false
  function updateReposReady(val: boolean): void {
    reposReady = val
  }

  let chartIdx: number = 0
  function setChartIdx(repoName: String) {
    chartIdx = githubReposVal.findIndex(r => r.name === repoName)
  }

  onMount(() => github.fetchRepos())

  onDestroy(() => {
    unsubGithubStore()
    unsubLang()
  })
</script>

<div
  class="
    h-screen
    w-screen
    overflow-none
    p-5
  ">
  <ReposBanner
    bind:chartIdx={chartIdx}
    bind:github={github}
    bind:reposReady={reposReady}
  />
  <div
    class="
      h-full w-full
      flex flex-col justify-between self-center
    ">
    <ReposTableHeader
      bind:github={github} />
    {#if !reposReady}
      <Loading />
    {:else}
      <div
        class="
          grow w-full
          overflow-y-scroll
          mb-32 pb-72
          flex flex-col justify-between self-center
        ">
        {#each githubReposVal as {
          colorData,
          commitData,
          created_at,
          description,
          html_url,
          languageData,
          name,
          pushed_at,
        }}
          <div
            class="
              h-fit
              pt-2 pb-2
              grid grid-rows-1 grid-cols-6
              hover:bg-slate-800 transition-colors
              border border-dashed
              border-b-2
              border-r-0
              border-t-0
              border-l-0
              border-cyan-500
            "
            on:mouseenter={() => setChartIdx(name)}>
            <td
              class="
                font-bold
                text-left
                text-xl
                break-words
                pl-5
              ">
              <a
                title="See The Code on Github"
                href={html_url}
                target="_blank">
                {name}
              </a>
            </td>
            <td
              class="
                w-full h-full
                px-5
                overflow-hidden
                flex flex-col jusity-around
              ">
              <div
                class="
                  w-full h-full
                  m-2
                  flex flex-wrap gap-2
                ">
                {#each languageData as {name}, i}
                  <span
                    class="
                      text-sm
                      font-bold
                      flex justify-between gap-2
                    ">
                    <ColorCircle
                      lang={name}
                      color={colorData[i]}/>
                  </span>
                {/each}
              </div>
            </td>
            <td
              class="
                ml-2 mr-2
                text-left
                text-lg
                font-bold
              ">
              {description}
            </td>
            <td>
              <div
                class="
                  h-full w-full
                  flex flex-col justify-around self-center
                  overflow-scroll
                ">
                {#if commitData.length}
                  <ol
                    class="
                      h-full w-full
                      flex flex-col justify-around
                    ">
                    {#each commitData.slice(0,3) as commit}
                        <li
                          class="
                            break-words
                          ">
                          <a
                            title="See Commit On Github"
                            target="_blank"
                            href={commit.html_url}>
                            {`${commit.sha.substring(0,7)}`}
                          </a>
                        </li>
                    {/each}
                  </ol>
                {/if}
              </div>
            </td>
            <td
              class="
                text-cyan-600
                font-bold
                text-lg
              "
              >
              {pushed_at.toLocaleString().split(',')[0]}
            </td>
            <td
              class="
                text-cyan-600
                font-bold
                text-lg
              "
              >
              {created_at.toLocaleString().split(',')[0]}
            </td>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>

<style lang="sass">
  @use "../styles/color"
  .repos-charts-grid
    grid-template-rows: 10% 90%
    grid-template-columns: 50% 50%
  .repos_banner_grid
    grid-template-columns: 15% 85%
  .section_grid
    grid-template-rows: max(226px, 30%) 70%
</style>
