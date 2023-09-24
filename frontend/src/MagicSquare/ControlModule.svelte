<script lang="ts">
  import { onDestroy } from 'svelte'
  
  import { I18n, Lang } from '../I18n'
  import { lang } from '../stores/lang'
  const i18n = new I18n('magicSquare/controlModule')
  let langVal: Lang
  const unsubLang = lang.subscribe(val => langVal = val)

  export let title: string = ''
  export let side: string = ''

  onDestroy(unsubLang)
</script>

<div class="control_module flex flex-col">
  <div class="control_module_title"
       class:left="{side == 'left'}"
       class:right="{side == 'right'}">
    {#if title.length}
      {title}
    {:else}
      {i18n.t("empty", langVal)}
    {/if}
  </div>
  <div class="control_module_slot_container">
    <slot class="control_module_slot flex flex-col justify-around items-stretch"/> 
  </div>
</div>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .control_module
    border: 5px double color.$blue-7
    margin: 2px 0 2px 0
    align-items: stretch 
    border-radius: 5px
    height: 98%
    /* min-width: 226px */
    &_slot
      height: 98%
      
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

  .left
    background-color: color.$green-7
  .right
    background-color: color.$purple-7
</style>
