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

<nav class="nav_bar bg-black mx-0 flex justify-between items-center gap-2 pt-2 pb-2 pr-4">
  <button id="siteSectionDropdown"
          class="dropdown_button shadow-none h-5/6 flex justify-around items-center text-xs border-transparent">
    <div class="dropdown_icon text-xl pb-2 flex justify-around items-center">
      ☰
    </div>
  </button>
  <Dropdown triggeredBy="#siteSectionDropdown"
            bind:open={dropdownOpen}
            placement={touchScreenVal ? "bottom" : "right"}
            classUl="w-40"
            ulClass="pt-2 pb-2 rounded-lg bg-black text-blue-200">
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/">
      {i18n.t("nav/home", langVal)} 
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/public-square">
      {i18n.t("nav/publicSquare", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/magic-square">
      {i18n.t("nav/magicSquare", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/give-me-a-sine">
      {i18n.t("nav/giveMeASine", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/about">
      {i18n.t("nav/about", langVal)}
    </DropdownItem>
    <DropdownItem class="hover:bg-transparent w-11/12 flex items-center font-bold text-blue-200"
                  tabIndex="0"
                  on:click={() => dropdownOpen = false}
                  href="/#/system-diagram">
      {i18n.t("nav/systemDiagram", langVal)}
    </DropdownItem>
  </Dropdown>

  <div id="user-drop"
       class="dropdown_button border-transparent font-bold flex justify-between items-center gap-4">
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
  @use "./../styles/text"

  .nav_bar
    border-bottom: 5px double color.$blue-7

  .dropdown_button
    color: color.$blue-4

  .dropdown_icon
    font-size: text.$fs-l
    font-weight: text.$fw-l
    color: color.$blue-4
</style>



