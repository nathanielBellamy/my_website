<script lang="ts">
</script>

<section class="h-full flex flex-col justify-between items-stretch">
  <div class="grow flex flex-col justify-around">
    <slot name="xyzSliders" />
  </div>
  <div class="grow flex flex-col justify-around items-stretch">
    <slot name="mouseTracking" />
  </div>
</section>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"
</style>
