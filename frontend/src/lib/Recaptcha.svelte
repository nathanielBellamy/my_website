<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Modal, Spinner } from "flowbite-svelte"
  import rImg from "../assets/recaptcha_logo.svg"

  import { I18n, Lang } from "../I18n"
  import { lang } from "../stores/lang"
    import Link from './Link.svelte';
  let i18n = new I18n("recaptcha")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  export let action: string
  export let title: string
  export let hasPassed: boolean = false

  let showModalVerify: boolean = false
  let showModalFailed: boolean = false
  function timeout(ms: number) {
      return new Promise(resolve => setTimeout(resolve, ms));
  }
  async function sleep(ms: number) {
      await timeout(ms)
  }

  function onClick(e: any) {
    e.preventDefault();
    grecaptcha.enterprise.ready(async () => {
      const token = await grecaptcha.enterprise.execute(
        import.meta.env.VITE_RECAPTCHA_SITE_KEY, 
        {action}
      )
      await sendTokenToServer(token)
    })
  }

  interface RecaptchaPayload {
    action: String
    token: String
  }

  async function sendTokenToServer(token: string) {
    showModalVerify = true

    const payload: RecaptchaPayload = { action, token }
    await fetch('recaptcha', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload)
    })
    .then(async (res) => {
        await sleep(1626)
        return res
      })
    .then((res) => {
      showModalVerify = false
      if (res.status === 200) {
        hasPassed = true
      } else {
        showModalFailed = true
      }
    })
  }

  onDestroy(unsubLang)
</script>

<Modal bind:open={showModalVerify}
       class="w-2/3 bg-slate-800 text-slate-300">
  <div class="h-5/6 w-5/6 bg-slate-800 flex items-center gap-4">
    <img src={rImg}
         style:height="70px"
         style:width="70px"
         alt="Google Recaptcha"/>
    <h3 class="text-cyan-700 mt-4 pl-4 pr-4 font-mono font-extrabold flex items-center">
      {i18n.t("verifying", langVal)}
    </h3>
    <Spinner color="blue" 
             size="5"/>
  </div>
</Modal>

<Modal bind:open={showModalFailed}
       class="w-2/3 bg-slate-800 text-slate-300">
  <div class="h-5/6 w-full bg-slate-800 flex items-center gap-4">
    <img src={rImg}
         style:height="70px"
         style:width="70px"
         alt="Google Recaptcha"/>
    <div class="font-extrabold flex flex-col justify-between items-stretch gap-4 text-left">
      <p class="text-xl text-red-700">
        {i18n.t("failed_1", langVal)}
      </p>
      <p class="text-cyan-700">
        {i18n.t("failed_2", langVal)}
      </p>
      <p class="text-cyan-700">
        {i18n.t("failed_3", langVal)}
        <Link href="/magic-square"
              title={i18n.t("magicSquare", langVal)}
              sameOrigin={true}/>
      </p>
    </div>
  </div>
</Modal>

<div class="w-full flex justify-around items-center">
  <button on:click={onClick}
          class="text-cyan-500 recaptcha_button font-mono w-5/6">
    {title}
  </button>
</div>

<style lang="sass">
  .recaptcha_button
    min-height: 50px
</style>

