<script lang="ts">
  import MagicBanner from "./lib/MagicBanner.svelte"
</script>

<body>
  <div class="text_lines flex flex-col justify-between items-stretch">
    <p>
      FOO
    </p>
    <p>
      AND
    </p>
    <p>
      BAR
    </p>
  </div>
</body>

<style lang="sass">
  @use "./styles/color"
  @use "./styles/text"

  .home_header
    padding: 0 15px 0 15px

  .text_lines
    color: color.$blue-7
    text-align: left
    padding: 15px
    height: 100%
</style>
