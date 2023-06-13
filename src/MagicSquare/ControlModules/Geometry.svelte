<script lang="ts">
  function handleRadiusChange(e: any, id: string) {
    var input = document.getElementById(id)
    input.value = e.detail.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<div class="h-full flex flex-col justify-around items-stretch">
  <div class="title flex items-stretch pl-5">
    shape
  </div>
  <div class="grow">
    <!-- TODO   -->
  </div>
  <div class="title flex items-stretch pl-5">
    transform order
  </div>
  <div class="grow">
    <!-- TODO   -->
  </div>
  <div class="title flex items-stretch pl-5">
    radius
  </div>
  <slot name="radiusSliders"/>
</div>

<style lang="sass">
  @use "./../../styles/color"
  @use "./../../styles/text"

  .title
    color: color.$blue-7
    font-size: text.$fs-l
    font-weight: text.$fw-l
</style>
