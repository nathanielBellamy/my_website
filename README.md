# Nate's Website

### build -- Go server, frontend SPA, auth SPA
- `./build.sh`
  - outputs to `build` directory
  - reproduce a prod-like build locally
- `./build-dist.sh`
  - outputs to `dist` directory
  - compile locally, transfer build, run remotely
- Reads `MODE=` from `config.env`
- NOTE: these scripts alter asset import paths in `index.html` files

### build only Go server (fast)
`./build.sh --server-only`

### serve
- `MODE=<mode> PW=<my_password> ./serve.sh`
- or more directly
  - `cd backend/go && MODE=<mode> PW=<my_password> ./main`
- serves on `localhost:8080`

### config.env
- `MODE=mode`

#### mode
- `localhost`
- `prod`
- `remotdev`

#### my_password
- password for dev site
- `localhost` and `remotedev` modes only

### cross-compiling
- `mode=localhost`
  - build compiles Go for localhost architecture
- `mode=remotedev, prod`
  - build compiles Go for Linux

### local SPA development
- `cd` into root of SPA
- `npm run dev`
- serves hot-updated SPA on `localhost:5173`
- see SPA's `package.json` for more build options
- NOTE: in order to work with `PublicSquare` locally, you will need
    - to serve on `localhost:8080` or change `VITE_BASE_URL`
    - either
      - to disable Recaptcha manually in code
      - to establish a test Recaptcha Enterprise project, key, and Api Key (Credentials)
        - test Recaptcha Key protected by domain (localhost:8080)
        - test Api Key (Credentials) for the Project protected by IP


### specs
- build with `MODE=localhost`
- start Go server on `localhost:8080` (default) with password `foo`
```
MODE=localhost PW=foo ./serve.sh
```
- start cypress, use ui to run specs
```
cd spec && npx cypress open && cd ..
```

- NOTE: 
  - specs are in early development focused on running in Chrome 
  - [ticket for e2e testing](https://github.com/users/nathanielBellamy/projects/4?pane=issue&itemId=33246560)
  - [ticket for unit/component testing](https://github.com/users/nathanielBellamy/projects/4?pane=issue&itemId=39606773)

### Made with: RustWasm, Go, Typescript, NixOS, Svelte, WebGL, Tailwind, Flowbite, Sass, Vite, Cypress
