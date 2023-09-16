<script lang="ts">
  import sd from './assets/SYSTEM_DIAGRAM.png'

  let innerHeight: number
  $: imgHeight = deriveImgHeight(innerHeight)

  function deriveImgHeight(ih: number): string {
    // TODO: unmagic this number
    //  but for now
    //  the navbar and footer heights are static
    //  and that is for what this number is accounting
    //  plus adding a small margin
    return (ih - 150).toString() + "px"
  }
</script>

<svelte:window bind:innerHeight />

<body class="w-full h-full bg-black flex justify-around items-center">
  <img src={sd}
       style:height={imgHeight}
       alt="System Diagram"/>
</body>
