<script lang="ts">
  import { onDestroy } from 'svelte'

  export let imgSideLength: string

  import aiMe0 from '../assets/ai_me/0.png'
  import aiMe1 from '../assets/ai_me/1.png'
  import aiMe2 from '../assets/ai_me/2.png'
  import aiMe3 from '../assets/ai_me/3.png'
  import aiMe4 from '../assets/ai_me/4.png'
  
  const aiMes: any[] = [
    aiMe0,
    aiMe1,
    aiMe2,
    aiMe3,
    aiMe4,
  ]

  let ai_me_counter: number = randomIntFromInterval(0, 4)
  $: ai_me_curr = ai_me_counter % 4 // we have 4 ai-generated images

  const incr_curr_ai_me = () => ai_me_counter = randomIntFromInterval(0, 4)
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
