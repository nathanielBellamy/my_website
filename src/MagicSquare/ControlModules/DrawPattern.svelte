<script lang="ts">
  //
</script>

<div class="draw_pattern_container justify-around">
  <slot></slot>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .draw_pattern
    &_container
      height: 100%
      width: 100%
      padding: 10px 0 10px 0
</style>
