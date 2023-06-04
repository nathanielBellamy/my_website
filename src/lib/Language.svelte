<script lang="ts">
  import { Lang } from "../I18n"
  import { lang } from '../stores/lang'
  
  let langVal: Lang 
  lang.subscribe(val => langVal = val)
  function setLang(newLangKey:string) {
    lang.update((_: Lang) => Lang[newLangKey])
  }
</script>

<section>
  <div class="lang_select grow">
    {#each Object.keys(Lang) as langKey }
      <button class="lang_select_opt mt-0"
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
