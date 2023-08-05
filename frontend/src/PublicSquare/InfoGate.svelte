<script lang="ts">
  import Recaptcha from "../lib/Recaptcha.svelte";

  export let hasPassedGate: boolean
</script>

<div class="w-full h-full bg-slate-800">
  <div class="w-full flex flex-col justify-between items-stretch">
    <h1 class="text-gray-400">
      Welcome to The Public Square
    </h1>
    <div class="w-full pl-4 pr-4 flex flex-col justify-between items-streth">
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Step right up, step right up, ladies and gentlemen, boys and girls of all ages! 
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Gather 'round and feast your eyes upon the spectacle that has dazzled kings and confounded philosophers, a mystery wrapped in an enigma, known only as the MAGIC SQUARE!
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Presenting, for the first time in this town and perhaps in the whole world, the ineffable, the inexplicable, the utterly miraculous MAGIC SQUARE! 
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Yes, you heard it right, a geometric wonder that defies explanation and transcends comprehension!
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Is it a mere illusion or a reality-shattering phenomenon? 
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        A trick of the eye or a gateway to a world unknown? Why, only the bravest of souls dare to look upon it and discern its secrets!
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        This, dear folks, is no ordinary square! Within its four equal sides lies an enigma so profound, it has puzzled the greatest minds and confounded the most learned scholars. 
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Legends tell of ancient civilizations brought to their knees in trying to understand its arcane properties!
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Come, come, don't be shy! Step forward and be a part of history! 
      </p>
      <p class="text-slate-400 font-bold text-lg text-left m-2 p-2">
        Experience a marvel that will be talked about for generations to come!
      </p>
    </div>
    <div class="grow w-full bg-slate-800 flex justify-around items-center">
      <button on:click={() => hasPassedGate = true}>
        Enter The Public Square
      </button>
    </div>
    <Recaptcha />
  </div>
</div>

<style lang="sass">

</style>

