<script lang="ts">
</script>

<section class="h-full pb-5 flex flex-col justify-between items-stretch">
  <div class="grow flex flex-col justify-around items-stretch">
    <div class="axis_slider_group_title title text-left pl-5"> X </div>
    <div class="grow">
      <slot name="xSliders" />
    </div>
    <div class="axis_slider_group_title title text-left pl-5"> Y </div>
    <div class="grow">
      <slot name="ySliders" />
    </div>
  </div>
  <div class="grow flex flex-col justify-around items-stretch">
    <slot name="mouseTracking" />
  </div>
</section>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  @import "../styles/control_module_title.sass"

  .title
    @include control_module_title

  .axis_slider_group_title
    flex-grow: 0.1
</style>
