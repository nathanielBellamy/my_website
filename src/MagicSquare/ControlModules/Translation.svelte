<script lang="ts">
</script>

<section class="h-full flex flex-col justify-between items-stretch">
  <div class="grow flex flex-col justify-around">
    <div class="axis_slider_group_title text-left pl-5"> X </div>
    <slot name="xSliders" />
    <div class="axis_slider_group_title text-left pl-5"> Y </div>
    <slot name="ySliders" />
  </div>
  <div class="grow flex flex-col justify-around items-stretch">
    <slot name="mouseTracking" />
  </div>
</section>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .axis_slider_group_title
    color: color.$blue-7
    font-size: text.$fs-l
    font-weight: text.$fw-l
</style>
