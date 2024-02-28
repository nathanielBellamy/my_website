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

  export let id: String
  let name = githubReposVal[0].name
  let idSelf: String = `repo_chart_${id}`

  let data: any = githubReposVal[0].languageData

  function setupChart(): void {
    var chartDom = document.getElementById("wowzow")
    var myChart = echarts.init(chartDom, {height: "200px", width: "200px"})
    var option

    // This example requires ECharts v5.5.0 or later
    option = {
      tooltip: {
        trigger: 'item'
      },
      // legend: {
      //   top: '5%',
      //   left: 'center'
      // },
      series: [
        {
          name,
          type: 'pie',
          radius: ['40%', '70%'],
          center: ['50%', '70%'],
          // adjust the start and end angle
          startAngle: 180,
          endAngle: 360,
          data
          // [
          //   { value: 1048, name: 'Search Engine' },
          //   { value: 735, name: 'Direct' },
          //   { value: 580, name: 'Email' },
          //   { value: 484, name: 'Union Ads' },
          //   { value: 300, name: 'Video Ads' }
          // ]
        }
      ]
    }
    option && myChart.setOption(option)
  }

  onMount(() => {
    // console.dir({data})
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
    bg-blue-200
  ">
  <h1>
    {name}
  </h1>
  <div
    id="wowzow"
    class="
      repo-chart-dom
      ml-10
    "/>
</div>

<style lang="sass">
  .repo-chart-dom
    margin-top: -20px
    height: 150px
    width: 700px
</style>
