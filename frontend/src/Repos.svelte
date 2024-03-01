<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import Icon from './lib/Icon.svelte'
  import { Icons } from './lib/Icons.js'
  import Loading from "./lib/Loading.svelte";
  import RepoLangChart from "./integrations/github/RepoLangChart.svelte";
  import RepoCommitChart from "./integrations/github/RepoCommitChart.svelte";
  import {
    type GithubRepo,
    type GithubRepos,
    SortColumns,
    SortOrder
  } from "./integrations/github/GithubTypes"
  import GithubIntegration from "./integrations/github/GithubIntegration"
  import ColorCircle from "./integrations/github/ColorCircle.svelte"

  import { lang } from "./stores/lang"
  import { I18n, Lang } from "./I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import { githubRepos } from "./stores/githubRepos"
  let githubReposVal: GithubRepos
  const unsubGithubRepos = githubRepos.subscribe((val: GithubRepos) => githubReposVal = [...val])

  const openLinkInNewTab = (href: string) => {
    window.open(href, '_blank');
  }

  function setSortOrder(order: String): void {
    github.sortOrder = order
    sortGithubReposBy()
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

  function handleHeaderClick(col: SortColumns): void {
    swapSortOrder()
    github.sortReposBy(col)
  }

  $: [...github.reposVal]

  function handleHeaderDblClick(col: SortColumns): void {
    handleHeaderClick(col)
  }

  let reposReady: boolean = false
  function updateReposReady(val: boolean): void {
    reposReady = val
  }

  let github: GithubIntegration = new GithubIntegration(githubRepos, updateReposReady)

  let chartIdx: number = 0
  function setChartIdx(repoName: String) {
    chartIdx = githubReposVal.findIndex(r => r.name === repoName)
  }

  onMount(() => {
    github.fetchRepos()
  })

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
          flex flex-col justify-between
        ">
        <div
          class="
            font-xxl
            text-left
            text-cyan-500
            grow
          ">
          {i18n.t("personalProejects", langVal)}
        </div>
        <div
          class="
            flex flex-col justify-around gap-2
          ">
          <label
            for="repos-sort-by"
            class="
              text-sm
              text-left
              text-blue-500
            ">
            Sort By:
          </label>
          <select
            id="repos-sort-by"
            data-testid="repos-sort-by"
            value={github.sortColumn}
            on:change={(e) => github.sortReposBy(e.target.value)}>
            {#each Object.values(SortColumns) as col}
              <option
                value={col}>
                {col}
              </option>
            {/each}
          </select>
          <select
            id="repos-sort-by"
            data-testid="repos-sort-by"
            value={github.sortOrder}
            on:change={(e) => {github.sortOrder = e.target.value; github.sortReposBy()}}>
            {#each Object.values(SortOrder) as order}
              <option
                value={order}>
                {order}
              </option>
            {/each}
          </select>
        </div>
      </div>
      {#if reposReady}
        <div
          class="
            w-full
            rounded-md
            grid grid-rows-2 grid-cols-1
            repos-charts-grid
          ">
          <h2
            class="
              col-span-2
              text-xl
              font-extrabold
              text-right
              text-cyan-500
            ">
            <p>
              {githubReposVal[chartIdx].name}
            </p>
          </h2>
          <div
            class="
              w-full flex
            ">
            <RepoLangChart bind:idx={chartIdx}/>
          </div>
          <div
            class="
              w-full flex
            ">
            <RepoCommitChart bind:idx={chartIdx}/>
          </div>
        </div>
      {/if}
    </div>
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
                on:click={() => handleHeaderClick(SortColumns.NAME)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.NAME)}>
                Name
              </button>
            </th>
            <th
              class="
              ">
              Code
            </th>
            <th
              class="
              ">
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumns.LANGUAGE)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.LANGUAGE)}>
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
                on:click={() => handleHeaderClick(SortColumns.DESCRIPTION)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.DESCRIPTION)}>
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
                on:click={() => handleHeaderClick(SortColumns.PUSHED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.PUSHED_AT)}>
                Lastest Push
              </button>
            </th>
            <th>
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumns.UPDATED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.UPDATED_AT)}>
                Lastest Update
              </button>
            </th>
            <th>
              <button
                class="
                  w-full
                  border-x-0
                "
                on:click={() => handleHeaderClick(SortColumns.CREATED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.CREATED_AT)}>
                Created
              </button>
            </th>
          </tr>
          {#each githubReposVal as {
            commitData,
            created_at,
            description,
            html_url,
            languageBreakdown,
            name,
            pushed_at,
            updated_at
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
                  text-wrap
                  pl-5
                ">
                {name}
              </td>
              <td
                class="
                  h-full w-full
                ">
                <div
                  class="
                    h-full w-full
                    flex justify-around items-center
                  ">
                  <button
                    class="
                      w-1/2 h-1/2
                      flex justify-around items-center
                    "
                    title="See The Code On Github"
                    on:click={() => openLinkInNewTab(html_url)}>
                    <Icon icon={Icons.GithubSolid} />
                  </button>
                </div>
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
                    mt-6
                    flex flex-wrap gap-2
                  ">
                  {#each Object.keys(languageBreakdown) as lang}
                    <span
                      class="
                        font-bold
                        flex justify-between gap-2
                      ">
                      {lang}
                      <ColorCircle lang={lang}/>
                    </span>
                  {/each}
                </div>
              </td>
              <td
                class="
                  ml-2 mr-2
                  text-left
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
  .repos-charts-grid
    grid-template-rows: 10% 90%
    grid-template-columns: 50% 50%
  .repos_banner_grid
    grid-template-columns: 15% 85%
  .section_grid
    grid-template-rows: max(226px, 30%) 70%
</style>
