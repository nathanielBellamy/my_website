<script lang="ts">
  function handleRadiusChange(e: any, id: string, idx: number) {
    var input = document.getElementById(id)
    input.value = e.detail.value
    input.dispatchEvent(new Event('input', {bubbles: true}))
  }
</script>

<div class="h-full flex flex-col justify-around items-stretch">
  <slot name="minStepSliders"/>
</div>

