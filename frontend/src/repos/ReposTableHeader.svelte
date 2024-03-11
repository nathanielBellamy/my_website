<script lang="ts">
  import {
    type GithubIntegration,
    SortColumn,
    SortOrder
  } from "../integrations/github/GithubTypes"

  export let github: GithubIntegration

  function handleHeaderClick(col: SortColumn): void {
    if (col === github.sortColumn) swapSortOrder() // only swap if clicking already selected header
    github.sortReposBy(col)
  }


  function handleHeaderDblClick(col: SortColumn): void {
    handleHeaderClick(col)
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
