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

<section class="h-full w-full flex flex-col justify-around items-stretch gap-4">
  <div class="h-16 w-full flex justify-center items-center">
    <div class="font-mono text-cyan-500 text-4xl md:text-6xl w-fit flex justify-around items-center"> 
      {i18n.t("title", langVal)}
    </div>
    <Spinner color="purple" />
  </div>
</section>
