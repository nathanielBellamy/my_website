<script lang="ts">
  import MagicBanner from "./lib/MagicBanner.svelte"
</script>

<header class="home_header flex items-center">
  <div class="title">
    My Website  
  </div>
</header>

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
    
  .title
    color: color.$blue-7
    font-size: text.$fs-l
    font-weight: text.$fw-m
</style>
