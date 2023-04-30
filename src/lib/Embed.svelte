<script lang="ts">
  export let title: string = '--'
  export let subTitle: string = ''
  export let css_class: string = ''
</script>

<div class={`embed ${css_class}`}>
  <slot/>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .embed
    display: flex
    flex-direction: column
    justify-content: flex-start
    align-items: stretch
</style>
