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
    color: color.$cream
    font-weight: text.$fw-xl
    text-align: left
    padding: .1em 0 .1em 1em
    font-size: 1.1em
    border: solid
    border-width: 3px 0 3px 0
    border-color: color.$blue-4
    margin: .25em 0 .25em 0
    
</style>
