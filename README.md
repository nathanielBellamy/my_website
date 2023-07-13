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
      - starts a Vite dev server at `/5173` with hot updates for any changes in `/frontend` including new RustWasm builds
      - in our experience with Svelte, hot updates are nice but EventListeners often do not survive - a manual, hard refresh is still best
    - `npm run preview`
      - starts a Vite server at `/4173` meant to imitate a prod server pulling from the `/frontend/dist` directory
      - use this to ensure that Vite bundles all of your assets appropriately
  - from within `/backend/go`
    - `./main` starts the bare-bones Go server on `/8080`

### Made with: Rust, WebAssembly, Svelte, Typescript, WebGL, Vite
