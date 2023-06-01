<script lang="ts">
  import {link} from "svelte-spa-router"
  export let href: string = "/"
  export let className: string = ""
  export let title: string = "Home"
  export let onClick: any = () => {}
  
</script>

<a href={href}
   use:link
   on:click={onClick}
   class={`link ${className}`}>
  <button class="link_button">
    {title}
  </button>
</a>

<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .link
    /* padding: 3px 7px 3px 7px */
    /* margin: 0 10px 0 10px */
    font-weight: text.$fw-l
    &_button
      margin: 0
      border-width: 0px
</style>
