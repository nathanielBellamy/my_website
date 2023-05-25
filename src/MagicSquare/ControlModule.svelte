<script lang="ts">
  export let title: string = ""
</script>

<div class="control_module flex flex-col">
  <div class="control_module_title">
    {title}
  </div>
  <slot class="control_module_slot"/> 
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .control_module
    border: 5px solid color.$blue-4
    align-items: stretch 
    border-radius: 5px
    margin: 5px
    padding: 5px
    overlfow-y: scroll
    height: 100%

    &_title
      text-align: center
      font-size: text.$fs_l
      font-weight: text.$fw_l
      text-align: center
</style>
