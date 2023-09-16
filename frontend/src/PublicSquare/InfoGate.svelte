<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Modal, Popover } from 'flowbite-svelte'
  import { slide } from 'svelte/transition'
  import Recaptcha from "../lib/Recaptcha.svelte"
  import { ViteMode } from "../ViteMode"
  import megaphone from '../assets/megaphone.png'
  import Link from '../lib/Link.svelte'
  import { smallScreen } from '../stores/smallScreen'
  import { I18n, type Lang } from '../I18n'
  import { lang } from '../stores/lang'
  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean | null) => smallScreenVal = val)
  
  const i18n = new I18n('publicSquare/infoGate')
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)

  export let hasPassedGate: boolean

  onDestroy(() => {
    unsubLang()
    unsubSmallScreen()
  })
</script>

<div class="info_gate font-mono w-full h-full overflow-y-scroll pb-4">
  <div class="h-full w-full flex flex-col justify-between items-stretch gap-4">
    <div class="w-full flex justify-around items-center">
      <div class="info_gate_title flex justify-around items-center">
        <img class="info_gate_megaphone flip mr-5 h-24 w-24"
             class:hidden={smallScreenVal}
             src={megaphone}
             alt={"Megaphone"}/>
        <Popover class="bg-slate-800 text-2xl"
                 offset={-10}
                 transition={slide}>
          <p class="font-extrabold text-cyan-500">
            {i18n.t("wow", langVal)}
          </p>
        </Popover>
        <div class="text-3xl md:text-6xl font-extrabold text-cyan-500">
          <p>
            ===============
          </p>
          <p>
            {i18n.t("welcome", langVal)}
          </p>
          <p>
            {i18n.t("publicSquare", langVal)}
          </p>
          <p>
            ===============
          </p>
        </div>
        <img class="info_gate_megaphone ml-5 h-24 w-24"
             class:hidden={smallScreenVal}
             src={megaphone}
             alt={"Megaphone"}/>
        <Popover class="bg-slate-800 text-2xl"
                 offset={-10}
                 transition={slide}>
          <p class="font-extrabold text-cyan-500">
            {i18n.t("zoinks", langVal)}
          </p>
        </Popover>
      </div>
    </div>
    <div class="h-full w-full pl-4 pr-4 flex justify-around items-center text-left text-2xl">
      <ul class="info_gate_intro w-10/12 font-bold">
        <li>
          {i18n.t("freeArt", langVal)}
        </li>
        <li>
          {i18n.t("shapesEmojis", langVal)}
        </li>
        <li>
          {i18n.t("offlineVersion", langVal)}
        </li>
      </ul>
    </div>
    <Recaptcha title={i18n.t("enter", langVal)}
               action="PSLOGIN"
               bind:hasPassed={hasPassedGate}/>
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .info_gate
    background-color: color.$blue-9
    &_intro
      color: color.$blue-4
      list-style-type: circle
    &_title 
      width: 60%
    &_megaphone
      filter: hue-rotate(90deg) drop-shadow(2px 2px color.$blue-7)

  .flip
    transform: scaleX(-1)
</style>

