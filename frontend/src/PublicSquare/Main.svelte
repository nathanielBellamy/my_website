<script lang="ts">
  import init, { PubSq, rust_init_message } from '../../pkg/src_rust.js'
  import { afterUpdate, beforeUpdate, onDestroy, onMount } from 'svelte'
  import { WebsocketBuilder } from 'websocket-ts'
  import { Toast } from 'flowbite-svelte'
  import { fly } from 'svelte/transition'

  // LIFECYCLE
  async function run(settings: any) {
    if (!hasBeenDestroyed) { 
      // resize + key block in Container.svelte may destroy component before wasm_bindgen can load
      // without this check, it is possible to load two wasm instances
      // since wasm retrieves the elements using .get_element_by_id
      // and since a new instance of the component will havee been mounted by the time wasm_bindgen loads
      // the result is two identical wasm instances listening to the same ui elements and drawing to the same context
      await init()
      rust_init_message("Public Square Wasm!")
      
      // init wasm process and set initial values
      const settings = await PubSq.run(settings, clientId, touchScreenVal)
      setAllSettings(settings)
      renderDataReady = true
    }
  }

  run()
</script>


<div>
  // TODO
</div>

<style lang="sass">

</style>
