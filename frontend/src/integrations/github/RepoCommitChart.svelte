<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import * as echarts from 'echarts'

  import { lang } from "../../stores/lang"
  import { I18n, Lang } from "../../I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import {
    type GithubRepo,
    type GithubRepos,
    type GithubStore
  } from "./GithubTypes"
  import { githubStore } from "../../stores/githubStore"
  let githubReposVal: GithubRepos
  const unsubGithubStore = githubStore.subscribe((store: GithubStore) => githubReposVal = [...store.repos])

  export let idx: number
  let id: String = `repo_commit_chart_${idx}`

  function commitDataToDates(commitData: any): any {
    return commitData.reduce((dates, commit) => {
      const date: String = commit.date.split('T')[0]
      const dateIndex = dates.findIndex(d => d[0] === date)
      if (dateIndex > -1)
      {
        dates[dateIndex][1] += 1
      }
      else
      {
        dates.push([date, 1])
      }
      return dates
    }, [])
  }

  function formatDate(date: Date): String {
    return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
  }

  function dateRange(commitDates: any): String[] {
    let recent: Date
    let oldest: Date
    if (commitDates.length)
    {
      recent = new Date(commitDates[0][0])
      oldest = new Date(commitDates[commitDates.length-1][0])
    }
    else
    {
      recent = new Date()
      oldest = new Date().setYear(today.getFullYear()-1)
    }
    return [oldest, recent].map(formatDate)
  }

  function setupChart(): void {
    var chartDom = document.getElementById(id)
    var myChart = echarts.init(chartDom, {height: "200px", width: "200px"})
    const data = commitDataToDates(githubReposVal[idx].commitData)
    var option

    option = {
      visualMap: {
        show: false,
        min: -5,
        max: 7,
        inRange: {
          color: ['#192F2B', '#73DACA']
        }
      },
      legend: {
        show: true,
      },
      calendar: {
        range: dateRange(data),
        itemStyle: {
          color: "#000000",
          borderWidth: 0.3,
        },
        monthLabel: {
          show: true,
          textStyle: {
            color: "#73DACA",
            fontWeight: "bolder",
            fontSize: 15
          }
        },
        yearLabel: {
          show: true,
          textStyle: {
            color: "#73DACA",
            fontWeight: "bolder",
            fontSize: 15
          }
        },
        width: 400,
        height: 100
      },
      title: {
        text: "Commits",
        textAlign: 'left',
        textStyle: {
          color: "#73DACA",
          fontWeight: "bolder",
          fontSize: 22
        }
      },
      series: [
        {
          type: 'heatmap',
          coordinateSystem: 'calendar',
          data
        }
      ]
    }
    option && myChart.setOption(option)
  }

  let mounted: boolean = false
  $: if (mounted) [...githubReposVal] && idx + 1 && setupChart()

  onMount(() => {
    mounted = true
    setupChart()
  })

  onDestroy(() => {
    unsubGithubStore()
  })

</script>

<div
  class="
    w-full h-full
    pl-20
    overflow-hidden
  ">
  <h2
    class="
      text-cyan-500
    ">
    {name}
  </h2>
  <canvas
    id={id}
    height={250}
    width={300}
    class="
      ml-10
    "/>
</div>

<style lang="sass">
</style>
