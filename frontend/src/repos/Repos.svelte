<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import Icon from '../lib/Icon.svelte'
  import { Icons } from '../lib/Icons.js'
  import ReposBanner from "./ReposBanner.svelte"
  import ReposTableHeader from "./ReposTableHeader.svelte"
  import ReposTableBody from "./ReposTableBody.svelte"
  import UserLangSummaryChart from "../integrations/github/UserLangSummaryChart.svelte"
  import {
    type GithubRepo,
    type GithubRepos,
    type GithubStore,
    SortColumn,
    SortOrder
  } from "../integrations/github/GithubTypes"

  import { lang } from "../stores/lang"
  import { I18n, Lang } from "../I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import GithubIntegration from "../integrations/github/GithubIntegration"
  let github: GithubIntegration = new GithubIntegration()

  let chartIdx: number = 0
  onMount(() => github.fetchRepos())

  onDestroy(unsubLang)
</script>

<div
  class="
    h-screen
    w-screen
    overflow-none
    p-5
  ">
  <ReposBanner
    bind:chartIdx={chartIdx}/>
  <div
    class="
      h-full w-full
      flex flex-col justify-between self-center
    ">
    <ReposTableHeader />
    <ReposTableBody
      bind:chartIdx={chartIdx}/>
  </div>
</div>
