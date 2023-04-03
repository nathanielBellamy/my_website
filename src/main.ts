import './app.css'
import App from './App.svelte'

import init, * as rust from "../src-rust/pkg/src_rust.js"

const app = new App({
  target: document.getElementById('app'),
})

init().then(() => {
  console.dir(rust.AppState.foo("main.ts bar"))
})


export default app

