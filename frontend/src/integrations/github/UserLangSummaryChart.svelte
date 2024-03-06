<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import * as echarts from 'echarts'

  import { lang } from "../../stores/lang"
  import { I18n, Lang } from "../../I18n"
  let i18n = new I18n("about")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import { type GithubRepo, type GithubRepos, type GithubStore } from "./GithubTypes"
  import { githubStore } from "../../stores/githubStore"
  let githubReposVal: GithubRepos
  const unsubGithubStore = githubStore.subscribe((store: GithubStore) => githubReposVal = [...store.repos])

  export let idx: number
  let id: String = `repo_lang_chart_${idx}`

  function setupChart(): void {
    var chartDom = document.getElementById( id )
    var myChart = echarts.init(chartDom, {height: "200px", width: "200px"})
    const repo = (githubReposVal[idx] || {})
    var option

    option = {
      legend: {
        show: true,
        right: "10%",
        bottom: 0,
        textStyle: {
          color: "#73DACA",
          fontWeight: "bolder",
          fontSize: 12
        }
      },
      title: {
        text: "Languages",
        textAlign: 'left',
        textStyle: {
          color: "#73DACA",
          fontWeight: "bolder",
          fontSize: 22
        }
      },
      color: repo.colorData,
      series: [
        {
          type: 'pie',
          radius: ['40%', '70%'],
          center: ['30%', '60%'],
          colorBy: 'data',
          color: repo.colorData,
          // adjust the start and end angle
          startAngle: 180,
          endAngle: 360,
          width: 300,
          height: 250,
          data: repo.languageData.map(obj => {
            return {
              label: {show: false},
              labelLine: {show: false},
              ...obj
            }
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
    height={150}
    width={450}
    class="
      ml-10
    "/>
</div>

<style lang="sass">
</style>
