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
  let userLanguageSummaryVal: UserLanguageSummary
  const unsubGithubStore = githubStore.subscribe((store: GithubStore) => {
    userLanguageSummaryVal = store.userLanguageSummary
  })

  export let sideLength: number
  let id: String = `user_lang_summary_chart`

  function setupChart(): void {
    var chartDom = document.getElementById( id )
    var myChart = echarts.init(chartDom, {height: `${sideLength}px`, width: `${sideLength}px`})
    var option

    option = {
      legend: {
        show: true,
        left: "37%",
        top: "25%",
        width: sideLength / 1.5,
        textStyle: {
          color: "#73DACA",
          fontWeight: "bold",
          fontSize: 12
        }
      },
      // title: {
      //   text: "Language Summary",
      //   textAlign: 'left',
      //   textStyle: {
      //     color: "#73DACA",
      //     fontWeight: "bolder",
      //     fontSize: 22
      //   }
      // },
      color: userLanguageSummaryVal.color_data,
      series: [
        {
          type: 'pie',
          radius: ['10%', '70%'],
          center: ['37%', '50%'],
          colorBy: 'data',
          // adjust the start and end angle
          startAngle: 270,
          endAngle: 90,
          width: sideLength /2,
          height: sideLength,
          data: userLanguageSummaryVal.language_data.map(obj => {
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
  $: if (mounted) sideLength && setupChart()

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
    overflow-hidden
    flex justify-around
  ">
  <canvas
    id={id}
    height={sideLength}
    width={sideLength *2}
    class="
      ml-10
    "/>
</div>

<style lang="sass">
</style>
