import './app.css'
import App from './App.svelte'

const app = new App({
  target: document.getElementById('app'),
})

// force https
if (location.protocol !== 'https:') {
    console.dir('NOT HTTPS')
    // location.replace(`https:${location.href.substring(location.protocol.length)}`);
}

export default app
