<script lang="ts">
  import { Popover } from 'flowbite-svelte'
  import { slide } from 'svelte/transition';
  import Recaptcha from "../lib/Recaptcha.svelte"
  import { ViteMode } from "../ViteMode"
  import megaphone from '../assets/megaphone.png'

  export let hasPassedGate: boolean
  let showRecaptcha = false

  function onEnterSquareClick() {
    if (import.meta.env.MODE !== ViteMode.localhost) {
      showRecaptcha = true
    } else {
      hasPassedGate = true
    }
  }
</script>

<div class="info_gate w-full h-full">
  <div class="w-full flex flex-col justify-between items-stretch">
    <div class="w-full flex justify-around items-center">
      <div class="info_gate_title flex justify-around items-center">
        <img class="info_gate_megaphone mr-5 h-24 w-24"
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
        <div class="text-7xl">
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
    <div class="w-full pl-4 pr-4 flex flex-col justify-between items-streth">
    </div>
    <div class="grow w-full bg-slate-800 flex justify-around items-center">
      <button on:click={onEnterSquareClick}>
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
    &_title 
      font-family: 'Abelone', 'Impact'
      width: 60%
    &_megaphone
      filter: hue-rotate(90deg) drop-shadow(2px 2px color.$blue-7)
</style>

