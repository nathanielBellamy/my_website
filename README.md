# Nate's Website

### dev environment
- suggest [Nix](https://nixos.org/guides/how-nix-works)
- [rustup](https://rustup.rs/)
- [wasm-bindgen](https://github.com/wasm-bindgen/wasm-bindgen)
- [Go](https://go.dev/)
- [npm](https://www.npmjs.com/)

### build -- Go server, marketing SPA, auth SPA, and old-site SPA
- `./lifecycle/build.sh`
  - outputs to `build` directory
  - reproduce a prod-like build locally
- `./lifecycle/build-dist.sh`
  - outputs to `dist` directory
  - compile locally, transfer build, run remotely
- `./lifecycle/serve.sh`
  - build + serve site using two docker containers: `my_website_backend` and `my_website_db`
- `./lifecycel/teardown.sh`
  - copies logs out of `my_website_backend` onto host machine
  - tears down docker containers
- Reads `MODE=` from `config.env`
- NOTE: these scripts alter asset import paths in `index.html` files

### build only Go server (fast)
`./lifecycle/build.sh --server-only`

### serve
- `MODE=<mode> PW=<my_password> ./lifecycle/serve.sh`
- or more directly
  - `cd backend/go && MODE=<mode> PW=<my_password> ./main`
- serves on `localhost:8080`

### .env/.env.${MODE}
```bash
MODE=localhost # | remotedev | production

GOOGLE_API_KEY=xxxx
RECAPTCHA_PROJECT_ID=xxxx # test google project
RECAPTCHA_SITE_KEY=xxxx # test site key

# totp
ENABLE_AUTH_LOCAL=false # Optional: Set to true to test the login flow on localhost
TOTP_SECRET=xxxx
ADMIN_EMAIL=xxxx
SMTP_HOST=xxxx
SMTP_PORT=587
SMTP_USER=xxxx
SMTP_PASS=xxxx

# postgres
DATABASE_URL=postgres://admin:admin@localhost:5432/mw_db?sslmode=disable
POSTGRES_USER=admin
POSTGRES_PASSWORD=password
POSTGRES_DB=my_db
```

#### mode
- `localhost`
- `prod`
- `remotdev`

#### my_password
- password for dev site
- `localhost` and `remotedev` modes only

### cross-compiling
- set `MODE` in `config.env`
- `MODE=localhost`
  - build compiles Go for localhost architecture
- `MODE=remotedev, prod`
  - build compiles Go for Linux

### old-site local SPA development
- `cd old-site`
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

### Made with: RustWasm, Go, Typescript, NixOS, Angular, Svelte, WebGL, Tailwind, Flowbite, Sass, Vite, Cypress
