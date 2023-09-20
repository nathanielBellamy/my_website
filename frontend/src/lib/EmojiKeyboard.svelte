<script lang="ts">
  import { onDestroy, onMount } from 'svelte'
  import { emojis, emojiKeymap } from '../../locales/emojis'
  export let value: string = "😎"
  export let valueSetIndicator: boolean = false

  interface EmojiVal {
    val: string,
    sortIdx: number
  }

  let sortedEmojiKeys: any[] = Object.entries(emojis)
                                     .sort(
    (a: any[], b: any[]) => {
      const aEmojiVal: EmojiVal = a[1]
      const bEmojiVal: EmojiVal = b[1]
      return aEmojiVal.sortIdx > bEmojiVal.sortIdx ? 1 : -1
    }).map(([key, _]) => key)
  
  function setVal(x: string) {
    valueSetIndicator = !valueSetIndicator
    value = x
  }

  function keyToEmoji(e: any) {
    const emoji = emojiKeymap[e.key]
    if (!!emoji) {
      e.preventDefault()
      setVal(emojis[emoji].val)
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
  {#each sortedEmojiKeys as emojiKey }
    <button class="p-2"
            on:click={() => setVal(emojis[emojiKey].val)}>
      {emojis[emojiKey].val}
    </button>
  {/each}
</div>
