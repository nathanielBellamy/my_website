# Nate's Website

### Running locally:

- at present the site is yet a statically served SPA

- `npm run build-rust-wasm` from `/frontend` after updating `/frontend/src-rust`
  - result put in `/frontend/pkg`
- `npm run build` from `/frontend` to bundle the SPA
  - result put in `/frontend/dist`

- multiple ways to serve locally
  - from within `/frontend`
    - `npm run dev`
      - starts a Vite dev server at `:5173` with hot updates for any changes in `/frontend` including new RustWasm builds
      - in our experience with Svelte, hot updates are nice but EventListeners often do not survive 
      - a manual, hard refresh is still best 
      - try one before any debugging when running dev
    - `npm run preview`
      - compile beforehand with `npm run build-frontend`
      - starts a Vite server at `:4173` meant to imitate a prod server pulling from the `/frontend/dist` directory
      - use this to ensure that Vite bundles all of your assets appropriately
  - from within `/backend/go`
    - `./main` starts the Go server on `/8080`
    - this should always be running during development
    - while `:5173` and `:4173` serve the assets, the front-end client will always listen to `:8080`
    - this includes production where an n-ginx reverse-proxy forwards external https traffic to `:8080`
    
  - thus far we have found it easiest to develop while
    - always running the Go server on :8080 
      - './main' from within '/backend/go'
      - rebuild as necessary with `go build main.go`
    - rebuild the whole front-end after any changes
      - 'npm run build-frontend' from '/frontend/'
      - the Go server will serve the new build w/o needing to restart
      - but you must refresh the browser
      - this isn't as fast as hot-updates but is much more reliable in reproducing a production environment
      - the front-end takes about 15-20seconds to build (RustWasm then esbuild)

  - running the dev server and wathing `:5173` in the browser is useful for CSS changes on a single page
    - it's a SPA, so refreshing the page gives you a true reload
    - if it doesn't, you're browser may be doing some extra caching

### Made with: RustWasm, Go, Typescript, Svelte, WebGL, Tailwind, Vite
