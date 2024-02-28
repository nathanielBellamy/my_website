<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import Loading from "./lib/Loading.svelte";
  import RepoChart from "./RepoChart.svelte";

  import Icon from './lib/Icon.svelte'
  import { Icons } from './lib/Icons.js'

  import { lang } from "./stores/lang"
  import { I18n, Lang } from "./I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import { githubRepos } from "./stores/githubRepos"
  import { type GithubRepo, type GithubRepos, GithubIntegration, SortColumns, SortOrder } from "./GithubIntegration"
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
            flex justify-around
            bg-cyan-900
            rounded-md
            pt-2
          ">
          <RepoChart idx={0}/>
          <RepoChart idx={1}/>
          <RepoChart idx={2}/>
        </div>
      {/if}
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
              <button
                on:click={() => handleHeaderClick(SortColumns.NAME)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.NAME)}>
                Name
              </button>
            </th>
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
              <button
                on:click={() => handleHeaderClick(SortColumns.LANGUAGE)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.LANGUAGE)}>
                Language
              </button>
            </th>
            <th
              class="
                w-96
              ">
              <button
                on:click={() => handleHeaderClick(SortColumns.DESCRIPTION)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.DESCRIPTION)}>
                Description
              </button>
            </th>
            <th>
              <button
                on:click={() => handleHeaderClick(SortColumns.PUSHED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.PUSHED_AT)}>
                Lastest Push
              </button>
            </th>
            <th>
              <button
                on:click={() => handleHeaderClick(SortColumns.UPDATED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.UPDATED_AT)}>
                Lastest Update
              </button>
            </th>
            <th>
              <button
                on:click={() => handleHeaderClick(SortColumns.CREATED_AT)}
                on:dblclick={() => handleHeaderDblClick(SortColumns.CREATED_AT)}>
                Created
              </button>
            </th>
          </tr>
          {#each githubReposVal as { created_at, description, html_url, languageBreakdown, name, pushed_at, updated_at }}
            <tr
              class="
                h-24
                border border-dashed border-b-2 border-cyan-500
                rounded-lg
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
