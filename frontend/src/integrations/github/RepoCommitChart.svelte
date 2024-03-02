<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import * as echarts from 'echarts'

  import { lang } from "../../stores/lang"
  import { I18n, Lang } from "../../I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import { type GithubRepo, type GithubRepos, githubRepos } from "../../stores/githubRepos"
  let githubReposVal: GithubRepos
  const unsubGithubRepos = githubRepos.subscribe((val: GithubRepos) => githubReposVal = [...val])

  export let idx: number
  let id: String = `repo_commit_chart_${idx}`

  function commitDataToDates(commitData: any): any {
    return commitData.reduce((dates, commit) => {
      const date: String = commit.commit.author.date.split('T')[0]
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
    const recent: Date = new Date(commitDates[0][0])
    const oldest: Date = new Date(commitDates[commitDates.length-1][0])
    return [oldest, recent].map(formatDate)
  }

  function setupChart(): void {
    var chartDom = document.getElementById(id)
    var myChart = echarts.init(chartDom, {height: "200px", width: "200px"})
    const data = commitDataToDates(githubReposVal[idx].commitData)
    var option

    // This example requires ECharts v5.5.0 or later
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
          color: "#1C263A",
          borderWidth: 0
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
    unsubGithubRepos()
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
