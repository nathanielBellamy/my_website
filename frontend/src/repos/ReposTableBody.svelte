<script lang="ts">
  import { onDestroy } from 'svelte'
  import Loading from "../lib/Loading.svelte"
  import ColorCircle from "../integrations/github/ColorCircle.svelte"
  import GithubIntegration from '../integrations/github/GithubIntegration'
  import {
    type GithubRepo,
    type GithubRepos,
    type GithubStore,
    SortColumn,
    SortOrder
  } from "../integrations/github/GithubTypes"

  import { githubStore } from "../stores/githubStore"
  let reposReadyVal: boolean = false
  let reposVal: GithubRepos = []
  const unsubGithubStore = githubStore.subscribe((store: GithubStore) => {
    reposVal = [...store.repos]
    reposReadyVal = store.reposReady
  })

  export let chartIdx: number
  function setChartIdx(repoName: String) {
    chartIdx = reposVal.findIndex(r => r.name === repoName)
  }

  function formatDate(date: Date): String {
    return date.toLocaleString().split(',')[0]
  }

  onDestroy(unsubGithubStore)
</script>

{#if !reposReadyVal}
  <Loading />
{:else}
  <div
    class="
      grow w-full
      overflow-y-scroll
      pb-72 mb-32
      flex flex-col justify-between self-center
    ">
    {#each reposVal as {
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
          {formatDate(pushed_at)}
        </td>
        <td
          class="
            text-cyan-600
            font-bold
            text-lg
          "
          >
          {formatDate(created_at)}
        </td>
      </div>
    {/each}
  </div>
{/if}
