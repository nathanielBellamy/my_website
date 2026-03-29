# Nate's Website

### Project Structure

```text
my_website/
├── backend/            # Go backend application (API, models, DB access)
├── frontend/           # Frontend SPA applications
│   ├── admin/          # New Angular admin dashboard
│   ├── auth/           # Authentication app guarding the admin area
│   ├── marketing/      # New Angular public-facing application
│   └── old-site/       # Legacy Svelte application
├── database/           # PostgreSQL initialization and migration scripts
├── e2e/                # Cypress end-to-end tests
├── lifecycle/          # Build, deploy, and orchestration shell scripts
├── nixos/              # NixOS VM configuration files
├── docker/             # Docker configuration (Dockerfile)
└── .github/            # CI/CD GitHub Action workflows
```

### dev environment
- suggest [Nix](https://nixos.org/guides/how-nix-works)
- [rustup](https://rustup.rs/)
- [wasm-bindgen](https://github.com/wasm-bindgen/wasm-bindgen)
- [Go](https://go.dev/)
- [npm](https://www.npmjs.com/)

### build -- Go server, marketing SPA, auth SPA, and old-site SPA
- `./lifecycle/backup.sh`
  - backs up dg
- `./lifecycle/build.sh`
  - outputs to `build` directory
  - reproduce a prod-like build locally
- `./lifecycle/build-dist.sh`
  - outputs to `dist` directory
  - compile locally, transfer build, run remotely
- `./lifecycle/e2e.sh`
  - runs Cypress tests against locally running server
- `./lifecycle/kill.sh`
  - copies logs out of `my_website_backend` onto host machine
  - tears down docker containers according to arg
- `./lifecycle/serve.sh`
  - build + serve site using two docker containers: `my_website_backend` and `my_website_db`
- `./lifecycle/test.sh`
  - run unit tests for all projects
- `./lifescycle/update.sh`
  - updates dependencies across projects
    

### build only Go server (fast)
`./lifecycle/build.sh --server-only`

### serve
- `./lifecycle/serve.sh`
- serves on `localhost:8080`

### .env/.env.${MODE}
```bash
MODE=localhost # | remotedev | production

# url
BASE_URL=http://localhost:8080
BASE_URL_API=http://localhost:8080/api
BASE_URL_OLD_SITE=http://old-site.localhost:8080
BASE_URL_GRAFANA=https://admin.localhost:8080/grafana/
GRAFANA_DOMAIN=admin.localhost:8080

# recaptcha
GOOGLE_API_KEY=xxxx
RECAPTCHA_PROJECT_ID=xxxx
RECAPTCHA_SITE_KEY=xxxx

# totp
ENABLE_AUTH_LOCAL=false 
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
- `production`
- `remotdev`

### cross-compiling
- set `MODE` in `config.env`
- `MODE=localhost`
  - build compiles Go for localhost architecture
- `MODE=remotedev, prod`
  - build compiles Go for Linux

### local development
- `./lifecycle/serve.sh`
- cd into project, either `frontend/marketing` or `frontend/admin`
- `ng serve` on :4200
- hmr frontend talks to backend service on :8080

### old-site local SPA development
- `cd old-site`
- `npm run dev`
- serves hot-updated SPA on `localhost:5173`
- see SPA's `package.json` for more build options

### Made with: Rust (WASM via wasm-bindgen), Go, Typescript, NixOS, Angular, Svelte, WebGL, Tailwind, Flowbite, Sass, Vite, Cypress
