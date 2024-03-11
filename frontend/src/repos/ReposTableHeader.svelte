<script lang="ts">
  import {
    type GithubIntegration,
    SortColumn,
    SortOrder
  } from "../integrations/github/GithubTypes"

  import GithubIntegration from "../integrations/github/GithubIntegration"
  let github: GithubIntegration = new GithubIntegration()

  import { githubStore } from "../stores/githubStore"
  let sortColumnVal: SortOrder
  const unsubGithubStore = githubStore.subscribe((store: GithubStore) => sortColumnVal = store.sortColumn)

  function handleHeaderClick(col: SortColumn): void {
    if (col === sortColumnVal) github.swapSortOrder()
    github.sortReposBy(col)
  }

  function handleHeaderDblClick(col: SortColumn): void {
    handleHeaderClick(col)
  }
</script>
<div
  class="
    grid grid-rows-1 grid-cols-6
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
      h-full
      flex flex-col justify-around self-center
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
</div>
