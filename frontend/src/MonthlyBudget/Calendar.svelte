<script lang="ts">

</script>

<div 
  class="
    w-full h-full
    grid grid-cols-7 grid-rows-5
  ">
  {#each {length: 35} as _, i}
    <div
      class="
       p-2 m-2 cursor-pointer
       bg-cyan-800 hover:bg-cyan-500
      ">
      { i }
    </div>
  {/each}
</div>
