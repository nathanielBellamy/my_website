<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { emojis, emojiKeymap } from '../../locales/emojis'
  export let value: string = "😎"
  export let valueSetIndicator: boolean = false
  
  function setVal(x: string) {
    valueSetIndicator = !valueSetIndicator
    value = x
  }

  function keyToEmoji(e: any) {
    const emoji = emojiKeymap[e.key]
    if (!!emoji) {
      e.preventDefault()
      setVal(emojis[emoji])
    }
  }

  function keyboardListener(e: any) {
    const ek_elem = document.getElementById("emoji_keyboard")
    if (!!ek_elem) keyToEmoji(e)
  }

  onMount(() => {
    window.addEventListener("keydown", keyboardListener)
  })

  onDestroy(() => {
    window.removeEventListener("keydown", keyboardListener)
  })
</script>

<div  id="emoji_keyboard"
      class="w-full h-full text-xs grid grid-cols-3 auto-rows-min">
  {#each Object.keys(emojis) as emojiName }
    <button class="p-2"
            on:click={() => setVal(emojis[emojiName])}>
      {emojis[emojiName]}
    </button>
  {/each}
</div>

<style lang="sass">

</style>
