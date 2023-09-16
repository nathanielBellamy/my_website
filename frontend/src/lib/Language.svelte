<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { Lang } from "../I18n"
  import { lang } from '../stores/lang'
  
  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)
  function setLang(newLangKey:string) {
    localStorage.setItem('lang', Lang[newLangKey])
    lang.update((_: Lang) => {
      return Lang[newLangKey]
    })
  }
  
  onMount(() => {
    const oldLang: string = localStorage.getItem('lang')
    if (typeof oldLang === 'string') {
      setLang(oldLang)
    } else {
      setLang(Lang.en)
    }
  })

  onDestroy(() => {
    unsubLang()
  })
</script>

<section class="w-fit">
  <div class="lang_select flex justify-between items-center">
    {#each Object.keys(Lang) as langKey }
      <button class="lang_select_opt mt-0 flex justify-around items-center"
              class:selected="{Lang[langKey] === langVal}"
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
    color: color.$cream
    background-color: color.$blue-7
</style>
