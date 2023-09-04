# Nate's Website

### build -- Go server, frontend SPA, auth SPA
`./build.sh`
- Reads `MODE=` from `config.env`
- NOTE: this script alters asset import paths in `index.html` files

### build only Go server (fast)
`./build.sh --server-only`

### serve
- `MODE=<mode> PW=<my_password> ./serve.sh`
- or more directly
  - `cd backend/go && MODE=<mode> PW=<my_password> ./main`

### config.env
- `MODE=mode`

#### mode
- `localhost`
- `prod`
- `remotdev`

#### my_password
- password for dev site
- `localhost` and `remotedev` only

### cross-compiling
- `mode=localhost`
  - compiles Go for localhost architecture
- `mode=remotedev, prod`
  - compiles Go for Linux

### local SPA development
- `cd` into root of SPA
- `npm run dev`
- serves hot-updated SPA on localhost:5173
- see SPA's `package.json` for more build options

### Made with: RustWasm, Go, Typescript, NixOS, Svelte, WebGL, Tailwind, Flowbite, Sass, Vite
