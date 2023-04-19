import './app.css'
import App from './App.svelte'

import init, * as rust from "../src-rust/pkg/src_rust.js"

const app = new App({
  target: document.getElementById('app'),
})

// force https
if (location.protocol !== 'https:') {
    console.dir('NOT HTTPS')
    // location.replace(`https:${location.href.substring(location.protocol.length)}`);
}

init().then(() => {
  console.dir(rust.init_message("WASM WASM WASM"))
})


export default app

