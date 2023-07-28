<script lang="ts">
  import { emojis } from '../../locales/emojis'
  export let value: string = "😎"
  export let valueSetIndicator: boolean = false
  
  function setVal(x: string) {
    valueSetIndicator = !valueSetIndicator
    value = x
  }
</script>

<div class="w-full h-full flex flex-wrap justify-evenly">
  {#each Object.keys(emojis) as emojiName, idx }
    <button on:click={() => setVal(emojis[emojiName])}
            on:keydown={(e) => {
              if (e.key === `${idx}`) {
                setVal(emojis[emojiName])
              }
            }}>
      {emojis[emojiName]}
    </button>
  {/each}
</div>

<style lang="sass">

</style>
