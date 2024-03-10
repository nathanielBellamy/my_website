<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import Icon from '../lib/Icon.svelte'
  import { Icons } from '../lib/Icons.js'
  import Loading from "../lib/Loading.svelte"
  import ReposBanner from "./ReposBanner.svelte"
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

  function swapSortOrder(): void {
    switch (github.sortOrder) {
      case SortOrder.ASC:
        github.sortOrder = SortOrder.DESC
        break
      case SortOrder.DESC:
        github.sortOrder = SortOrder.ASC
        break
    }
  }

  function handleHeaderClick(col: SortColumn): void {
    if (col === github.sortColumn) swapSortOrder() // only swap if clicking already selected header
    github.sortReposBy(col)
  }

  $: [...github.reposVal]

  function handleHeaderDblClick(col: SortColumn): void {
    handleHeaderClick(col)
  }

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
      grid grid-rows-2 grid-cols-1 gap-4
      section_grid
    ">
    <div
      class="
        border-solid border-5 border-cyan-500
        overflow-y-scroll
        row-span-10 col-span-8
        mt-4
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
              repo-table-grid
              h-16
            ">
            <th
              class="
              ">
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumn.NAME)}
                on:dblclick={() => handleHeaderDblClick(SortColumn.NAME)}>
                Name
              </button>
            </th>
            <th
              class="
              ">
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumn.LANGUAGE)}
                on:dblclick={() => handleHeaderDblClick(SortColumn.LANGUAGE)}>
                Language
              </button>
            </th>
            <th
              class="
              ">
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumn.DESCRIPTION)}
                on:dblclick={() => handleHeaderDblClick(SortColumn.DESCRIPTION)}>
                Description
              </button>
            </th>
            <th
              class="
              ">
              Recent Commits
            </th>
            <th>
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumn.PUSHED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumn.PUSHED_AT)}>
                Latest Push
              </button>
            </th>
            <th>
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumn.CREATED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumn.CREATED_AT)}>
                Created
              </button>
            </th>
          </tr>
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
            <tr
              class="
                h-44
                pt-4
                hover:bg-slate-800 transition-colors
                border border-dashed border-b-2 border-r-0 border-cyan-500
                rounded-md
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
                      {name}
                      <ColorCircle color={colorData[i]}/>
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
                    flex flex-col justify-around
                    overflow-scroll
                  ">
                  {#if commitData.length}
                    {#each commitData.slice(0,3) as commit}
                      <ol
                        class="
                          h-full w-full
                          flex flex-col justify-around
                        ">
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
                      </ol>
                    {/each}
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
            </tr>
          {/each}
        </table>
      {/if}
    </div>
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
