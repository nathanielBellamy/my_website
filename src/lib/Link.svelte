<script lang="ts">
  import {link} from "svelte-spa-router"
  export let sameOrigin:boolean = true
  export let href: string = "/"
  export let className: string = ""
  export let title: string = "Home"
  export let onClick: any = () => {}
</script>

{#if sameOrigin}
  <a href={href}
     use:link
     on:click={onClick}
     class={`link ${className}`}>
    <button class="link_button">
      {title}
    </button>
  </a>
{:else}
  <a href={href}
     target="_blank"
     on:click={onClick}
     class={`link ${className}`}>
    <button class="link_button">
      {title}
    </button>
  </a>
{/if}


<style lang="sass">
  @use "./../styles/color"
  @use "./../styles/text"

  .link
    font-weight: text.$fw-l
    &_button
      margin: 0
      border-width: 0px
      color: color.$blue-4
</style>
