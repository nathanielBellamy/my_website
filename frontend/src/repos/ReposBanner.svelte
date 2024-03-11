<script lang="ts">
  import { onDestroy } from 'svelte'
  import {
    type GithubRepo,
    type GithubRepos,
    type GithubStore,
    SortColumn,
    SortOrder
  } from "../integrations/github/GithubTypes"
  import RepoLangChart from "../integrations/github/RepoLangChart.svelte"
  import RepoCommitChart from "../integrations/github/RepoCommitChart.svelte"
  import Loading from "../lib/Loading.svelte"

  import GithubIntegration from "../integrations/github/GithubIntegration"
  import { githubStore } from "../stores/githubStore"
  let reposReadyVal: boolean
  let reposVal: GithubRepos
  let sortColumnVal: SortColumn
  let sortOrderVal: SortOrder
  const unsubGithubStore = githubStore.subscribe((store: GithubStore) => {
    reposReadyVal = store.reposReady
    reposVal = [...store.repos]
    sortColumnVal = store.sortColumn
    sortOrderVal = store.sortOrder
  })
  let github: GithubIntegration = new GithubIntegration()

  export let chartIdx: number

  onDestroy(unsubGithubStore)
</script>

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
    <div
      class="
        font-xxl
        text-left
        text-cyan-500
      ">
      Repos
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
        value={sortColumnVal}
        on:change={(e) => github.sortReposBy(e.target.value)}>
        {#each Object.values(SortColumn) as col}
          <option
            value={col}>
            {col}
          </option>
        {/each}
      </select>
      <select
        id="repos-sort-by"
        data-testid="repos-sort-by"
        value={sortOrderVal}
        on:change={(e) => github.setSortOrder(e.target.value)}>
        {#each Object.values(SortOrder) as order}
          <option
            value={order}>
            {order}
          </option>
        {/each}
      </select>
    </div>
  </div>
  {#if !reposReadyVal}
    <Loading />
  {:else}
    <div
      class="
        w-full
        rounded-md
        grid grid-rows-3 grid-cols-2
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
          {reposVal[chartIdx].name}
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
