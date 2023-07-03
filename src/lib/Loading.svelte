<script lang="ts">
  export let message: string = ''
</script>

<section>
  <h1 class="loading">
    {!!message ? message : "Loading..."}
  </h1>
  <slot />
</section>

<style lang="sass">
  @use '../styles/color'
  @use '../styles/text'

  .loading
    color: color.$blue-7
    font-size: text.$fs-l
    font-weight: text.$fw-l
</style>
