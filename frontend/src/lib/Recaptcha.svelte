<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Modal, Spinner } from "flowbite-svelte"
  import rImg from "../assets/recaptcha_logo.svg"

  import { I18n, Lang } from "../I18n"
  import { lang } from "../stores/lang"
  let i18n = new I18n("recaptcha")
  let langVal: Lang
  const unsubLang = lang.subscribe( val => langVal = val)

  export let action: string
  export let title: string
  export let hasPassed: boolean = false

  let showModal: boolean = false
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
      showModal = true
      await sendTokenToServer(token)
      showModal = false
    })
  }

  interface RecaptchaPayload {
    action: String
    token: String
  }

  async function sendTokenToServer(token: string) {
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
    .then((res) => { hasPassed = res.status === 200 })
  }

  onDestroy(unsubLang)
</script>

<Modal bind:open={showModal}
       class="w-2/3 bg-slate-800 text-slate-300">
  <div class="h-5/6 w-5/6 bg-slate-800 flex items-center">
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

