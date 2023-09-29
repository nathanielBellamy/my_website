<script lang="ts">
  import { onDestroy } from 'svelte'
  import realMe from '../assets/real_me.png'
  import Icon from './Icon.svelte';
  import { Icons } from './Icons';

  // flowbite drawer
  import { Drawer, CloseButton, SidebarItem, Sidebar, SidebarWrapper, SidebarGroup} from 'flowbite-svelte';
  import { sineIn } from 'svelte/easing';
  let hiddenDrawer = true;
  let transitionParams = {
    x: -320,
    duration: 200,
    easing: sineIn
  };

  import { I18n, Lang } from "../I18n"
  import { lang } from "../stores/lang"
  let i18n = new I18n("app")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  onDestroy(() => {
    unsubLang()
  })
</script>

<nav class="nav_bar bg-black mx-0 flex justify-between items-center gap-2 pt-2 pb-2 pr-4"
     data-testid="nav_bar">
  <button id="navList"
          class="!shadow-none !p-0 border-1 border-slate-800 hover:border-slate-700 flex justify-around self-center"
          on:click={() => hiddenDrawer = false}
          data-testid="nav_button">
    <div class="text-cyan-500 pb-1 pr-1 pl-1 font-mono font-bold text-2xl flex justify-around items-center">
      ☰
    </div>
  </button>
  <Drawer bind:hidden={hiddenDrawer} 
          class="bg-black pl-6"
          transitionType="fly" 
          {transitionParams} 
          id="navList">
    <div class="flex items-center">
      <h3 id="drawer-navigation-label-3" 
          class="text-base font-semibold text-gray-500 uppercase">
        Menu
      </h3>
      <CloseButton on:click={() => (hiddenDrawer = true)} 
                   class="mb-4 bg-black text-cyan-700 hover:bg-slate-800" />
    </div>
    <Sidebar>
      <SidebarWrapper divClass="overflow-y-auto py-4 px-3 rounded dark:bg-gray-800">
        <SidebarGroup>
          <SidebarItem class="hover:bg-transparent !list-none flex items-center font-bold text-blue-200"
                       tabIndex="0"
                       on:click={() => hiddenDrawer = true}
                       href="/#/"
                       data-testid="nav_dropdown_home"
                       label={i18n.t("nav/home", langVal)}>
            <svelte:fragment slot="icon">
              <Icon icon={Icons.HomeOutline} />
            </svelte:fragment>
          </SidebarItem>
          <SidebarItem class="hover:bg-transparent flex items-center font-bold text-blue-200"
                       tabIndex="0"
                       on:click={() => hiddenDrawer = true}
                       href="/#/public-square"
                       data-testid="nav_dropdown_public_square"
                       label={i18n.t("nav/publicSquare", langVal)}>
            <svelte:fragment slot="icon">
              <Icon icon={Icons.UsersGroupOutline} />
            </svelte:fragment>
          </SidebarItem>
          <SidebarItem class="hover:bg-transparent flex items-center font-bold text-blue-200"
                       tabIndex="0"
                       on:click={() => hiddenDrawer = true}
                       href="/#/magic-square"
                       data-testid="nav_dropdown_magic_square"
                       label={i18n.t("nav/magicSquare", langVal)}>
            <svelte:fragment slot="icon">
              <Icon icon={Icons.UserOutline} />
            </svelte:fragment>
          </SidebarItem>
          <SidebarItem class="hover:bg-transparent flex items-center font-bold text-blue-200"
                       tabIndex="0"
                       on:click={() => hiddenDrawer = true}
                       href="/#/give-me-a-sine"
                       data-testid="nav_dropdown_give_me_a_sine"
                       label={i18n.t("nav/giveMeASine", langVal)}>
            <svelte:fragment slot="icon">
              <Icon icon={Icons.GridOutline} />
            </svelte:fragment>
          </SidebarItem>
          <SidebarItem class="hover:bg-transparent flex items-center font-bold text-blue-200"
                       tabIndex="0"
                       on:click={() => hiddenDrawer = true}
                       href="/#/about"
                       data-testid="nav_dropdown_about"
                       label={i18n.t("nav/about", langVal)}>
            <svelte:fragment slot="icon">
              <Icon icon={Icons.InfoCircleOutline} />
            </svelte:fragment>
          </SidebarItem>
          <SidebarItem class="hover:bg-transparent flex items-center font-bold text-blue-200"
                       tabIndex="0"
                       on:click={() => hiddenDrawer = true}
                       href="/#/system-diagram"
                       data-testid="nav_dropdown_system_diagram"
                       label={i18n.t("nav/systemDiagram", langVal)}>
            <svelte:fragment slot="icon">
              <Icon icon={Icons.DnaOutline} />
            </svelte:fragment>
          </SidebarItem>
        </SidebarGroup>
      </SidebarWrapper>
    </Sidebar>
  </Drawer>

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
