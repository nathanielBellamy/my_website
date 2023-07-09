<script lang="ts">
  import { onDestroy } from 'svelte'

  export let imgSideLength: string

  import aiMe0 from '../assets/ai_me/0.png'
  import aiMe1 from '../assets/ai_me/1.png'
  import aiMe2 from '../assets/ai_me/2.png'
  import aiMe3 from '../assets/ai_me/3.png'
  import aiMe4 from '../assets/ai_me/4.png'
  import aiMe5 from '../assets/ai_me/5.png'
  import aiMe6 from '../assets/ai_me/6.png'
  import aiMe7 from '../assets/ai_me/7.png'
  import aiMe8 from '../assets/ai_me/8.png'
  import aiMe9 from '../assets/ai_me/9.png'
  import aiMe10 from '../assets/ai_me/10.png'
  import aiMe11 from '../assets/ai_me/11.png'
  import aiMe12 from '../assets/ai_me/12.png'
  import aiMe13 from '../assets/ai_me/13.png'
  import aiMe14 from '../assets/ai_me/14.png'
  import aiMe15 from '../assets/ai_me/15.png'
  import aiMe16 from '../assets/ai_me/16.png'
  import aiMe17 from '../assets/ai_me/17.png'
  import aiMe18 from '../assets/ai_me/18.png'
  import aiMe19 from '../assets/ai_me/19.png'
  import aiMe20 from '../assets/ai_me/20.png'
  import aiMe21 from '../assets/ai_me/21.png'
  import aiMe22 from '../assets/ai_me/22.png'
  import aiMe23 from '../assets/ai_me/23.png'
  import aiMe24 from '../assets/ai_me/24.png'
  import aiMe25 from '../assets/ai_me/25.png'
  import aiMe26 from '../assets/ai_me/26.png'
  import aiMe27 from '../assets/ai_me/27.png'
  import aiMe28 from '../assets/ai_me/28.png'
  import aiMe29 from '../assets/ai_me/29.png'
  import aiMe30 from '../assets/ai_me/30.png'
  import aiMe31 from '../assets/ai_me/31.png'
  import aiMe32 from '../assets/ai_me/32.png'
  import aiMe33 from '../assets/ai_me/33.png'
  import aiMe34 from '../assets/ai_me/34.png'
  import aiMe35 from '../assets/ai_me/35.png'
  import aiMe36 from '../assets/ai_me/36.png'
  import aiMe37 from '../assets/ai_me/37.png'
  import aiMe38 from '../assets/ai_me/38.png'
  import aiMe39 from '../assets/ai_me/39.png'
  import aiMe40 from '../assets/ai_me/40.png'
  import aiMe41 from '../assets/ai_me/41.png'
  import aiMe42 from '../assets/ai_me/42.png'
  import aiMe43 from '../assets/ai_me/43.png'
  import aiMe44 from '../assets/ai_me/44.png'
  import aiMe45 from '../assets/ai_me/45.png'
  import aiMe46 from '../assets/ai_me/46.png'

  const aiMes: any[] = [
    aiMe0,
    aiMe1,
    aiMe2,
    aiMe3,
    aiMe4,
    aiMe5,
    aiMe6,
    aiMe7,
    aiMe8,
    aiMe9,
    aiMe10,
    aiMe11,
    aiMe12,
    aiMe13,
    aiMe14,
    aiMe15,
    aiMe16,
    aiMe17,
    aiMe18,
    aiMe19,
    aiMe20,
    aiMe21,
    aiMe22,
    aiMe23,
    aiMe24,
    aiMe25,
    aiMe26,
    aiMe27,
    aiMe28,
    aiMe29,
    aiMe30,
    aiMe31,
    aiMe32,
    aiMe33,
    aiMe34,
    aiMe35,
    aiMe36,
    aiMe37,
    aiMe38,
    aiMe39,
    aiMe40,
    aiMe41,
    aiMe42,
    aiMe43,
    aiMe44,
    aiMe45,
    aiMe46,
  ]

  let ai_me_counter: number = randomIntFromInterval(0, 46)
  $: ai_me_curr = ai_me_counter % 47 // we have 47 ai-generated images

  const incr_curr_ai_me = () => ai_me_counter = randomIntFromInterval(0, 46)
  const ai_me_interval: any = setInterval(incr_curr_ai_me, 5000)

  function randomIntFromInterval(min: number, max: number): number { // min and max included 
    return Math.floor(Math.random() * (max - min + 1) + min)
  }

  onDestroy(() => {
    clearInterval(ai_me_interval)
  })
</script>

<div>
  {#each aiMes as aiMe, idx}
    <img class="h-full w-full ai_me ai_me_img"
         class:ai_me_img_hide={ai_me_curr !== idx}
         class:ai_me_img_show={ai_me_curr === idx}
         style:height={imgSideLength}
         style:width={imgSideLength}
         src={aiMe}
         alt={`AI ME #${idx}`}/>
  {/each}
</div>

<style lang="sass">
  .ai_me
    grid-area: img
    &_img
      border-radius: 50%
      filter: brightness(80%)
      &_hide
        display: none
        border-radius: 50%
        opacity: 0%
        transition: opacity 1s
      &_show
        visibility: visible
        border-radius: 50%
        opacity: 100%
        transition: opacity 1.25s ease-in
</style>
