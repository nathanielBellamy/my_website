<script lang="ts">
  export let color: String
</script>

<div
  class="
    color-circle
    h-4 w-4
  "
  style={`background-color: ${color}`}/>

<style lang="sass">
  .color-circle
    border-radius: 50%
</style>
