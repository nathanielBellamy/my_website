<script lang="ts">
  export let color: String
  export let lang: String
</script>

<div
  class="
    flex items-center
  ">
  <div
    class="
      mr-2
      color-circle
      h-3 w-3
    "
    style={`background-color: ${color}`}/>
  {lang}
</div>
<style lang="sass">
  .color-circle
    border-radius: 50%
</style>
