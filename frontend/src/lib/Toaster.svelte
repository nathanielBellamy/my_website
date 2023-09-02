<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Toast } from 'flowbite-svelte'
  import { fly } from 'svelte/transition'
  import type { ToastColor } from './Toasty'

  export let color: ToastColor
  export let text: string
  export let open: boolean | null


  import { smallScreen } from '../stores/smallScreen'
  let smallScreenVal: boolean
  const unsubSmallScreen = smallScreen.subscribe((val: boolean | null) => smallScreenVal = val)

  enum Position {
    tr = "top-right",
    tl = "top-left",
    br = "bottom-right",
    bl = "bottom-left",
    none = "none"
  }

  let position: Position = Position.tr

  $: if (smallScreenVal) {
    position = Position.tr
  }

  $: if (!smallScreenVal) {
    position = Position.br
  }

  onDestroy(() => {
    unsubSmallScreen()
  })
</script>

{#if open === null}
  <Toast color={color}
         class="bg-slate-800 text-sky-600"
         transition={fly}
         params="{{x: 200}}"
         position={position}>
    <svelte:fragment slot="icon">
      <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path></svg>
      <span class="sr-only">Check icon</span>
    </svelte:fragment>
    {text}
  </Toast>
{:else}
  <Toast color={color}
         bind:open={open}
         class="bg-slate-800 text-sky-600"
         transition={fly}
         params="{{x: 200}}"
         position={position}>
    <svelte:fragment slot="icon">
      <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path></svg>
      <span class="sr-only">Check icon</span>
    </svelte:fragment>
    {text}
  </Toast>
{/if}
