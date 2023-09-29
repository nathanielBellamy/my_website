<script lang="ts">
  import { onDestroy } from 'svelte'
  import { push } from "svelte-spa-router"
  import { Dropdown, DropdownItem } from 'flowbite-svelte'
  import { intoUrl, SiteSection  } from "../stores/siteSection"
  import realMe from '../assets/real_me.png'

  import { I18n, Lang } from "../I18n"
  import { lang } from "../stores/lang"
  let i18n = new I18n("app")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  import { touchScreen } from '../stores/touchScreen'
  let touchScreenVal: boolean
  const unsubTouchScreen = touchScreen.subscribe((val: boolean) => touchScreenVal = val)

  let dropdownOpen: boolean = false;

  function handleDropdownClick(s: SiteSection) {
    push(intoUrl(s))
  }

  onDestroy(() => {
    unsubLang()
    unsubTouchScreen()
  })
</script>

<nav class="nav_bar bg-black mx-0 flex justify-between items-center gap-2 pt-2 pb-2 pr-4"
     data-testid="nav_bar">
  <button id="siteSectionDropdown"
          class="!shadow-none !p-0 border-1 border-slate-800 hover:border-slate-700 flex justify-around self-center"
          data-testid="nav_button">
    <div class="text-cyan-500 pb-1 pr-1 pl-1 font-mono font-bold text-2xl flex justify-around items-center">
      ☰
    </div>
  </button>
  <Dropdown triggeredBy="#siteSectionDropdown"
            bind:open={dropdownOpen}
            placement={touchScreenVal ? "bottom" : "right"}
            classUl="w-40"
            ulClass="pt-2 pb-2 rounded-lg bg-black text-blue-200"
            data-testid="nav_dropdown">
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/"
                  data-testid="nav_dropdown_home">
      {i18n.t("nav/home", langVal)} 
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/public-square"
                  data-testid="nav_dropdown_public_square">
      {i18n.t("nav/publicSquare", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/magic-square"
                  data-testid="nav_dropdown_magic_square">
      {i18n.t("nav/magicSquare", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/give-me-a-sine"
                  data-testid="nav_dropdown_give_me_a_sine">
      {i18n.t("nav/giveMeASine", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/about"
                  data-testid="nav_dropdown_about">
      {i18n.t("nav/about", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/system-diagram"
                  data-testid="nav_dropdown_system_diagram">
      {i18n.t("nav/systemDiagram", langVal)}
    </DropdownItem>
  </Dropdown>

  <div class="text-cyan-500 font-bold flex justify-between items-center gap-4">
    <div>
      Nate Schieber
    </div>
    <img src={realMe}
         alt="RealMe"
         height="30px"
         width="30px"
         class="rounded-full" />
  </div>
</nav>

<style lang="sass">
  @use "./../styles/color"

  .nav_bar
    border-bottom: 5px double color.$blue-7
</style>



