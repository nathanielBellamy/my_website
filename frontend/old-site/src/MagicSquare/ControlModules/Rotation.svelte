<script lang="ts">
  import { onDestroy } from 'svelte'
  
  import { I18n, Lang } from '../../I18n'
  import { lang } from '../../stores/lang'
  let i18n = new I18n("magicSquare/rotation")
  let langVal: Lang 
  const unsubLang = lang.subscribe(val => langVal = val)

  enum Freedom {
    pitch = "pitch",
    roll = "roll",
    yaw = "yaw"
  }

  onDestroy(unsubLang)
</script>

<div class="h-full rotation_container w-full flex flex-col justify-between items-stretch">
  <div class="grow freedom_group flex gap-2 justify-between items-stretch">
    <div class="freedom_group_title flex justify-around items-center">
      {i18n.t(Freedom.pitch, langVal)}
    </div>
    <div class="h-full freedom_group_body flex flex-col justify-around items-stretch">
      <slot name="pitch"/>
    </div>
  </div>
  <div class="grow freedom_group flex gap-2 justify-stretch items-stretch">
    <div class="freedom_group_title flex justify-around items-center">
      {i18n.t(Freedom.roll, langVal)}
    </div>
    <div class="h-full freedom_group_body flex flex-col justify-around items-stretch">
      <slot name="roll"/>
    </div>
  </div>
  <div class="grow freedom_group flex gap-2 justify-stretch items-stretch">
    <div class="freedom_group_title flex justify-around items-center">
      {i18n.t(Freedom.yaw, langVal)}
    </div>
    <div class="h-full freedom_group_body flex flex-col justify-around items-stretch">
      <slot name="yaw" />
    </div>
  </div>
</div>

<style lang="sass">
  @use "../../styles/color"
  @use "../../styles/text"

  .rotation
    &_container
      padding: 5px 15px 5px 5px
      overflow-x: hidden
      overflow-y: scroll

  .freedom_group
    align-items: center
    padding: 5px
    margin: 5px

    &_title
      font-weight: text.$fw-m
      font-size: text.$fs-l
      transform: rotate(-90deg)
      max-width: 15px
      color: color.$blue-7

    &_body
      width: 90%
      margin: 0 0 0 5%
      border-top: 5px double color.$blue-7
      border-radius: 10px
      flex-grow: 1
</style>
