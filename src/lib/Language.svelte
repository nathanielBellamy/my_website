<script lang="ts">
  import { onMount } from 'svelte'
  import { Lang } from '../I18n'

  export let lang: Lang = Lang.en
  function setLang(newLangKey:string) {
    lang = Lang[newLangKey]
  }

  $: updateStorage(lang)

  function updateStorage(lang: Lang) {
    localStorage.setItem('lang', lang)
    dispatchEvent(new Event('lang', {bubbles: true}))
  }

  onMount(async () => {
    // set in storage so other components can access it w/o prop mining
    updateStorage(lang)
  })
</script>

<section>
  <div class="lang_select grow">
    {#each Object.keys(Lang) as langKey }
      <button class="lang_select_opt mt-0"
              class:selected="{Lang[langKey] === lang}"
              on:click={() => setLang(langKey)}>
        {langKey}
      </button>
    {/each}
  </div>
</section>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .lang_select_opt
    color: color.$blue-6
    border: none
  
  .selected
    color: color.$blue-4
</style>
