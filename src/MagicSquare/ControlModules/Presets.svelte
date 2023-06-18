<script lang="ts">
  let bank: number = 0
  let preset: number = 0
</script>

<section class="h-full flex flex-col justify-around items-stretch">
  <div class="flex flex-col justify-between items-stretch">
    <div class="title text-left">
      Bank {bank + 1}
    </div>
    <div class="grid grid-cols-4">
      {#each {length: 4} as _, i}
        <button on:click={() => bank = i}
                class:selected={bank === i}>
          {i + 1}
        </button>
      {/each}
    </div>
  </div>
  <div class="grow flex flex-col justify-around items-stretch">
    <div class="title text-left">
      Preset {preset + 1}
    </div>
    <div class="grow grid grid-cols-4 grid-rows-4">
      {#if bank === 0}
        {#each {length: 16} as _, idx}
          <button on:click={() => preset = idx}
                  class:selected={preset === idx}>
            {idx + 1}
          </button>
        {/each}
      {:else if bank === 1}
        {#each {length: 16} as _, idx}
          <button on:click={() => preset = idx + 16}
                  class:selected={preset === idx + 16}>
            {idx + 17}
          </button>
        {/each}
      {:else if bank === 2}
        {#each {length: 16} as _, idx}
          <button on:click={() => preset = idx + 32}
                  class:selected={preset === idx + 32}>
            {idx + 33}
          </button>
        {/each}

      {:else if bank === 3}
        {#each {length: 16} as _, idx}
          <button on:click={() => preset = idx + 48}
                  class:selected={preset === idx + 48}>
            {idx + 49}
          </button>
        {/each}
      {/if}
    </div>
  </div>
</section>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .title
    color: color.$blue-7
    font-size: text.$fs-ml
    font-weight: text.$fw-l

  .selected
    background-color: color.$blue-7
</style>
