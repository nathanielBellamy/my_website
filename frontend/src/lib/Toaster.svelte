<script lang="ts">
  import { onDestroy } from 'svelte'
  import { Toast } from 'flowbite-svelte'
  import { fly } from 'svelte/transition'
  import type { ToastColor } from './Toasty'

  export let color: ToastColor
  export let text: string
  export let open: boolean | null
  export let icon: Icons


  import { smallScreen } from '../stores/smallScreen'
  import Icon from './Icon.svelte';
  import type { Icons } from './Icons';
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
      <Icon {icon}/>
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
      <Icon {icon} />
    </svelte:fragment>
    {text}
  </Toast>
{/if}
