<script lang="ts">
  export let title: string = "EMPTY"
</script>

<div class="control_module flex flex-col">
  <div class="control_module_title">
    {title}
  </div>
  <div class="control_module_slot_container">
    <slot class="control_module_slot flex flex-col"/> 
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .control_module
    border: 5px solid color.$blue-7
    margin: 2px 0 2px 0
    align-items: stretch 
    border-radius: 5px
    height: 100%
    min-width: 226px
    &_slot
      height: 100%
      
      &_container
        flex-grow: 1
        overflow-y: scroll
        padding: 10px 0 10px 0
        width: 100%

    &_title
      color: color.$cream
      text-align: center
      font-size: text.$fs_m
      font-weight: text.$fw_xl
      text-align: center
      border-bottom: 5px solid color.$blue-7
</style>
