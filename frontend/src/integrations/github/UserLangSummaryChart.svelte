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

  let id: String = `user_lang_summary_chart`

  function setupChart(): void {
    var chartDom = document.getElementById( id )
    var myChart = echarts.init(chartDom, {height: "200px", width: "200px"})
    var option

    option = {
      legend: {
        show: true,
        left: 0,
        bottom: 0,
        width: 300,
        textStyle: {
          color: "#73DACA",
          fontWeight: "bolder",
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
          radius: ['20%', '60%'],
          center: ['50%', '50%'],
          colorBy: 'data',
          color: userLanguageSummaryVal.color_data,
          // adjust the start and end angle
          startAngle: 0,
          endAngle: 360,
          width: 300,
          height: 300,
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
    height={300}
    width={300}
    class="
      ml-10
    "/>
</div>

<style lang="sass">
</style>
