<script lang="ts">
  import colors from './colors.json'

  export let lang: String
</script>

<div
  class="
    color-circle
    h-4 w-4
  "
  style={`background-color: ${colors[lang].color}`}/>

<style lang="sass">
  .color-circle
    border-radius: 50%
</style>
