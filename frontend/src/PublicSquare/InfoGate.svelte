<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Popover } from 'flowbite-svelte'
  import { slide } from 'svelte/transition'
  import Recaptcha from "../lib/Recaptcha.svelte"
  import { ViteMode } from "../ViteMode"
  import megaphone from '../assets/megaphone.png'
  import Link from '../lib/Link.svelte'
  import { smallScreen } from '../stores/smallScreen'
  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean | null) => smallScreenVal = val)

  export let hasPassedGate: boolean
  let showRecaptcha = false

  function onEnterSquareClick() {
    if (import.meta.env.MODE !== ViteMode.localhost) {
      showRecaptcha = true
    } else {
      hasPassedGate = true
    }
  }

  onDestroy(unsubSmallScreen)
</script>

<div class="info_gate w-full h-full overflow-y-scroll pb-4">
  <div class="h-full w-full flex flex-col justify-between items-stretch gap-4">
    <div class="w-full flex justify-around items-center">
      <div class="info_gate_title flex justify-around items-center">
        <img class="info_gate_megaphone flip mr-5 h-24 w-24"
             src={megaphone}
             alt={"Megaphone"}/>
        <Popover class="bg-slate-800 text-2xl"
                 transition={slide}>
          <p>
            Come one! Come all,
          </p>
          <p>
            to The Public Square!
          </p>
        </Popover>
        <div class:text-6xl={!smallScreenVal}
             class:text-2xl={smallScreenVal}>
          <p>
            ===============
          </p>
          <p>
            == Welcome to ==
          </p>
          <p>
            The Public Square
          </p>
          <p>
            ===============
          </p>
        </div>
        <img class="info_gate_megaphone ml-5 h-24 w-24"
             src={megaphone}
             alt={"Megaphone"}/>
        <Popover class="bg-slate-800 text-2xl"
                 transition={slide}>
          <p>
            Mysteries abound in
          </p>
          <p>
            The Public Square!
          </p>
        </Popover>
      </div>
    </div>
    <div class="h-full w-full pl-4 pr-4 flex justify-around items-center text-left text-2xl">
      <ul class="info_gate_intro w-10/12 font-bold">
        <li>
          <span class="abelone">
            The Public Square 
          </span>
          is a free, anonymously collaborative art project
        </li>
        <li>
          A digital wall for geometric graffiti animations
        </li>
        <li>
          Make colorful shapes and send emojis
        </li>
        <li>
          Everyone has equal voice to change the art in
          <span class="abelone">
            THE PUBLIC SQUARE
          </span>
        </li>
        <li>
          For a solo version where you can save presets, nagivate to
          <span class="abelone">
            <Link href="/magic_square"
                  title="The Magic Square"
                  sameOrigin={true}/>
          </span>
        </li>
      </ul>
    </div>
    <div class="grow w-full flex justify-around items-center">
      <button class="abelone text-3xl"
              on:click={onEnterSquareClick}>
        Enter The Public Square
      </button>
    </div>
    {#if showRecaptcha && import.meta.env.MODE !== ViteMode.localhost}
      <Recaptcha bind:hasPassed={hasPassedGate}/>
    {/if}
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"
  @use "./../styles/font"

  .info_gate
    background-color: color.$blue-9
    &_intro
      color: color.$blue-4
      list-style-type: circle
    &_title 
      font-family: 'Abelone'
      width: 60%
    &_megaphone
      filter: hue-rotate(90deg) drop-shadow(2px 2px color.$blue-7)

  .abelone
    font-family: 'Abelone'

  .flip
    transform: scaleX(-1)
</style>

