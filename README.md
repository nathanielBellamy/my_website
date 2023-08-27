# Nate's Website

### build -- Go server, frontend SPA, auth SPA
`./build.sh`

### build only Go server (fast)
`./build.sh --server-only`

### serve
`./serve.sh`

### config.env
- `MODE=mode`
- `PW=my_password`

#### mode
- `localhost`
- `prod`
- `remotdev`

### cross-compiling
- `mode=localhost`
  - compiles Go for host architecture
- `mode=remotedev, prod`
  - compiles Go for linux
#### my_password
- password for site
- `localhost` and `remotedev` only

### local SPA development
- `cd` into root of SPA
- `npm run dev`
- serves hot-updated SPA on localhost:5173
- see SPA's `package.json` for more build options


### Made with: RustWasm, Go, Svelte, Typescript, WebGL, Vite, Tailwind, Flowbite
