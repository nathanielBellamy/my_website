<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Spinner } from 'flowbite-svelte'
  import { I18n, Lang } from "../I18n"
  import { lang } from "../stores/lang"
  let i18n = new I18n("loading")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  onDestroy(unsubLang)
</script>

<section class="h-full w-full flex flex-col justify-around items-stretch">
  <div class="h-16 w-full flex justify-center items-center gap-6">
    <Spinner color="purple" 
             size="10"/>
    <div class="font-mono text-cyan-500 text-4xl md:text-6xl w-fit flex justify-around items-center"
         data-testid="loading_title"> 
      {i18n.t("title", langVal)}
    </div>
  </div>
</section>
