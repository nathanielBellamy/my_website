<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import * as echarts from 'echarts'

  import { lang } from "./stores/lang"
  import { I18n, Lang } from "./I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import { type GithubRepo, type GithubRepos, githubRepos } from "./stores/githubRepos"
  let githubReposVal: GithubRepos
  const unsubGithubRepos = githubRepos.subscribe((val: GithubRepos) => githubReposVal = [...val])

  export let idx: number
  let id: String = `repo_chart_${idx}`

  function setupChart(): void {
    var chartDom = document.getElementById(id)
    var myChart = echarts.init(chartDom, {height: "200px", width: "200px"})
    var option

    // This example requires ECharts v5.5.0 or later
    option = {
      tooltip: {
        trigger: 'item'
      },
      legend: {
        show: true,
        right: "10%",
        bottom: 0,
        textStyle: {
          color: "#73DACA",
          fontWeight: "bolder",
          fontSize: 20
        }
      },
      title: {
        text: githubReposVal[idx].name,
        textAlign: 'left',
        textStyle: {
          color: "#73DACA",
          fontWeight: "bolder",
          fontSize: 22
        }
      },
      series: [
        {
          type: 'pie',
          radius: ['40%', '70%'],
          center: ['30%', '50%'],
          // adjust the start and end angle
          startAngle: 180,
          endAngle: 360,
          width: 300,
          height: 250,
          data: githubReposVal[idx].languageData.map(obj => {

            return {
              label: {show: false},
              labelLine: {show: false},
              ...obj}
          })
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
    height={150}
    width={500}
    class="
      ml-10
    "/>
</div>

<style lang="sass">
</style>
