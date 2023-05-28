<script lang="ts">
  import {link} from "svelte-spa-router"
  export let href: string = "/"
  export let className: string = ""
  export let title: string = "Home"
  
</script>

<a href={href}
   use:link
   class={`link ${className}`}> 
  {title}
</a>

<style lang="sass">
  @use "./../styles/color"
  .link
    padding: 3px 7px 3px 7px
</style>
