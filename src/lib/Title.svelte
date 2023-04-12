<script lang="ts">
  export let title: string = '--'
</script>

<div class="title">
  {title}
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"
  
  .title
    background: color.$blue-grad
    color: color.$black-4
    font-weight: text.$fw-xl
    text-align: left
    padding: .1em 0 .1em 1em
    font-size: 1.1em
    
</style>
