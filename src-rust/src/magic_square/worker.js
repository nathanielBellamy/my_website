// js loaded using workers to achieve multithreading
// https://www.tweag.io/blog/2022-11-24-wasm-threads-and-messages/

importScripts("../../pkg/src_rust.js")

const { child_entry_point } = wasm_bindgen

self.onmessage = async event => {
  // Expect message to have form [module, memory, ptr]
  // - module is a WebAssembly.Module
  // - memory is the WebAssembly.Memory object that the main thread is using
  // - ptr is the pointer to the function we want to execute on the spawned worker thread
  
  let init = await wasm_bindgen('../../pkg/src_rust_bg.wasm', event.data[1]).catch(err => {
    // propogate to main `onerror`
    setTimeout(() => {
      throw err
    });
    // Rethrow to keep promise rejected and prevent execute of further commands
    throw err
  })
 
  child_entry_point(Number(event.data[2]))

  // Free memory (stack, thread-locals) held (in the wasm linear memory) by the thread.
  init.__wbindgen_thread_destroy();
  // Tell the browser to stop the thread.
  close();
}
