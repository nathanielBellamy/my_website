# Nate's Website

### Running locally:
- default password is `guest`
- which bcrypt hashes to:
```
$2a$10$eoka2klp4SoOA4mXyiHkQuctdKkXXJfalLotfvX7hbuiryu5fQA.G

```
- compiling + building
  - `npm run build-rust-wasm` from `/frontend` after updating `/frontend/src-rust` to build the wasm modules + RustWasm bindings using [wasm-pack](https://rustwasm.github.io/wasm-pack/)
    - result put in `/frontend/pkg`
  - `npm run build` from `/frontend` to bundle the SPA
    - result put in `/frontend/dist`
  - `npm run build-frontend` from `/frontend` to do both
    - runs `npm run build-rust-wasm`
    - then runs `npm run build`

- multiple ways to serve
  - from within `/frontend`
    - `npm run dev`
      - starts a Vite dev server at `:5173` with hot updates for any changes in `/frontend` including new RustWasm builds
      - in our experience with Svelte, hot updates are nice but EventListeners and store subscriptions may not survive 
      - a manual, hard refresh is still best 
      - try one before any debugging when running dev
    - `npm run preview`
      - compile beforehand with `npm run build-frontend`
      - starts a Vite server at `:4173` meant to imitate a prod server pulling from the `/frontend/dist` directory
      - use this to ensure that Vite bundles all of your assets appropriately
  - from within `/backend/go`
    - `MODE=<runtime_env> ./main` to start the Go server on `/8080`
      - `MODE=localhost PW_HASH=$2a$10$eoka2klp4SoOA4mXyiHkQuctdKkXXJfalLotfvX7hbuiryu5fQA.G ./main` 
        - will serve `auth/dev_auth/dist`
        - require login to access `frontend/dist`
      - `MODE=production ./main`
        - servers `frontend/dist`
      - `MODE=remotedev PW_HASH=$2a$10$eoka2klp4SoOA4mXyiHkQuctdKkXXJfalLotfvX7hbuiryu5fQA.G ./main`
        - will serve `auth/dev_auth/dist`
        - require login to access `frontend/dist`
      - TODO: script these
    - `./main` starts the Go server on `/8080`
    - will serve latest result of `npm run build-frontend` w/o needing to restart
    - refresh the browser, though
    
  - thus far we have found it easiest to develop while:
    - always running the Go server on :8080 
      - `MODE=<runtime_env> ./main` from within `/backend/go`
        - `MODE=localhost ./main`
        - `MODE=production ./main`
        - `MODE=remotedev ./main`

      - rebuild as necessary with `go build main.go`
    - rebuild the whole front-end after any changes
      - `npm run build-frontend` from `/frontend/`
      - this isn't as fast as hot-updates but is more reliable in reproducing a production environment
      - the front-end takes about 15-20seconds to build (RustWasm then Vite/esbuild)

  - running the dev server while using `:5173` in the browser is useful for CSS changes

### Made with: RustWasm, Go, Typescript, Svelte, WebGL, Tailwind, Vite
