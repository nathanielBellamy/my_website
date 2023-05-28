<script lang="ts">
  import MagicBanner from "./lib/MagicBanner.svelte"
</script>

<header class="main_header grid grid-cols-10">
  <div class="title p-3 text-lg font-bold col-span-2">
    It's A Website  
  </div>
</header>

<body>
  There is a body
</body>

<style lang="sass">
  @use "./styles/color"

  .title
    color: color.$blue-4
</style>
