<script lang="ts">
  enum Lang {
    en = "english",
    es = "spanish",
    fr = "french"
  }
  
  let lang: Lang = Lang.en

  function setLang(newLangKey:string) {
    lang = Lang[newLangKey]
  }
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
