#!/bin/bash
#
# init.sh — Full dev environment setup for my_website.
#
# After running this script, you should be ready to develop locally.
# Start the project with: ./lifecycle/serve.sh
#
set -e

cd "$(dirname "$0")/.."
ROOT_DIR=$(pwd)

# ─── Formatting helpers ───────────────────────────────────────────────

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m' # No Color

ok()   { echo -e "  ${GREEN}✔${NC} $1"; }
warn() { echo -e "  ${YELLOW}⚠${NC} $1"; }
fail() { echo -e "  ${RED}✘${NC} $1"; }
info() { echo -e "  ${CYAN}ℹ${NC} $1"; }

# ─── Header ───────────────────────────────────────────────────────────

cat << 'EOF'

  ┌──────────────────────────────────────┐
  │                                      │
  │   🏗️  my_website — Dev Env Setup     │
  │                                      │
  └──────────────────────────────────────┘

EOF

# ─── 1. Check required system dependencies ────────────────────────────

echo -e "${BOLD}📋 Checking system dependencies...${NC}"
echo ""

MISSING=()

# Node.js
if command -v node &> /dev/null; then
  NODE_VERSION=$(node --version | sed 's/v//')
  NODE_MAJOR=$(echo "$NODE_VERSION" | cut -d. -f1)
  if [ "$NODE_MAJOR" -ge 20 ]; then
    ok "Node.js v${NODE_VERSION}"
  else
    fail "Node.js v${NODE_VERSION} found — v20+ required"
    MISSING+=("node")
  fi
else
  fail "Node.js not found — v20+ required"
  MISSING+=("node")
fi

# npm
if command -v npm &> /dev/null; then
  ok "npm $(npm --version)"
else
  fail "npm not found"
  MISSING+=("npm")
fi

# Go
if command -v go &> /dev/null; then
  GO_VERSION=$(go version | grep -oE '[0-9]+\.[0-9]+(\.[0-9]+)?' | head -1)
  ok "Go v${GO_VERSION}"
else
  fail "Go not found — v1.24+ required"
  MISSING+=("go")
fi

# Rust & wasm toolchain
if command -v rustc &> /dev/null; then
  RUST_VERSION=$(rustc --version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)
  ok "Rust v${RUST_VERSION}"
else
  fail "Rust not found — required for WASM builds"
  MISSING+=("rust")
fi

if command -v wasm-pack &> /dev/null; then
  ok "wasm-pack $(wasm-pack --version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' || echo '(installed)')"
else
  warn "wasm-pack not found — needed for WASM builds (install: cargo install wasm-pack)"
fi

if command -v wasm-bindgen &> /dev/null; then
  ok "wasm-bindgen $(wasm-bindgen --version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' || echo '(installed)')"
else
  warn "wasm-bindgen not found — needed for WASM builds (install: cargo install wasm-bindgen-cli)"
fi

# Docker
if command -v docker &> /dev/null; then
  DOCKER_VERSION=$(docker --version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)
  ok "Docker v${DOCKER_VERSION}"
else
  fail "Docker not found — required to run services"
  MISSING+=("docker")
fi

# Docker Compose
if docker compose version &> /dev/null; then
  COMPOSE_VERSION=$(docker compose version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)
  ok "Docker Compose v${COMPOSE_VERSION}"
elif command -v docker-compose &> /dev/null; then
  COMPOSE_VERSION=$(docker-compose --version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1)
  ok "Docker Compose v${COMPOSE_VERSION} (legacy)"
else
  fail "Docker Compose not found — required to run services"
  MISSING+=("docker-compose")
fi

echo ""

# Bail out if critical deps are missing
if [ ${#MISSING[@]} -ne 0 ]; then
  echo -e "${RED}${BOLD}❌ Missing required dependencies: ${MISSING[*]}${NC}"
  echo "   Please install the above before re-running this script."
  exit 1
fi

# ─── 2. Check environment files ──────────────────────────────────────

echo -e "${BOLD}📂 Checking environment files...${NC}"
echo ""

if [ -f ".env/.env.localhost" ]; then
  ok ".env/.env.localhost exists"
else
  fail ".env/.env.localhost not found"
  echo ""
  echo "   The local env file is required for Docker Compose."
  echo "   Ask a project maintainer for a copy of .env/.env.localhost"
  exit 1
fi

if [ -f "config.env" ]; then
  ok "config.env exists"
else
  warn "config.env not found — creating default (MODE=localhost)"
  echo "MODE=localhost" > config.env
  ok "config.env created with MODE=localhost"
fi

echo ""

# ─── 3. Install dependencies (with user confirmation) ────────────────

FRONTEND_PROJECTS=(
  "frontend/marketing"
  "frontend/admin"
  "frontend/auth/admin"
  "frontend/old-site"
)

echo -e "${BOLD}📦 The following dependency installs will be performed:${NC}"
echo ""
info "Go modules           — go mod download"
for project in "${FRONTEND_PROJECTS[@]}"; do
  info "npm install           — ${project}/"
done
info "npm install           — e2e/"
echo ""

read -p "   Proceed with dependency installation? [Y/n] " CONFIRM
CONFIRM=${CONFIRM:-Y}

if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
  echo ""
  warn "Dependency installation skipped. You can install manually later."
  echo ""
else
  echo ""

  # ── Go modules ──
  echo -e "${BOLD}🔧 Installing Go modules...${NC}"
  go mod download
  ok "Go modules downloaded"
  echo ""

  # ── Frontend npm installs ──
  for project in "${FRONTEND_PROJECTS[@]}"; do
    if [ -d "$project" ] && [ -f "$project/package.json" ]; then
      echo -e "${BOLD}🔧 Installing npm dependencies — ${project}/${NC}"
      if [[ "$project" == "frontend/old-site" ]]; then
        (cd "$project" && npm install --legacy-peer-deps)
      else
        (cd "$project" && npm install)
      fi
      ok "${project} dependencies installed"
      echo ""
    else
      warn "${project} not found — skipping"
      echo ""
    fi
  done

  # ── E2E npm install ──
  if [ -d "e2e" ] && [ -f "e2e/package.json" ]; then
    echo -e "${BOLD}🔧 Installing npm dependencies — e2e/${NC}"
    (cd e2e && npm install)
    ok "e2e dependencies installed"
    echo ""
  fi
fi

# ─── 4. Verify Docker is running ─────────────────────────────────────

echo -e "${BOLD}🐳 Checking Docker daemon...${NC}"
echo ""

if docker info &> /dev/null; then
  ok "Docker daemon is running"
else
  warn "Docker daemon is not running — start Docker before using serve.sh"
fi

echo ""

# ─── 5. Summary ──────────────────────────────────────────────────────

cat << EOF

  ┌──────────────────────────────────────┐
  │                                      │
  │   ✅  Dev environment is ready!      │
  │                                      │
  └──────────────────────────────────────┘

EOF

echo -e "${BOLD}  Lifecycle scripts:${NC}"
echo ""
echo -e "    ${CYAN}./lifecycle/serve.sh${NC}    Start all services (Docker Compose)"
echo -e "    ${CYAN}./lifecycle/test.sh${NC}     Run unit tests (Go + Angular)"
echo -e "    ${CYAN}./lifecycle/e2e.sh${NC}      Run Cypress E2E tests"
echo -e "    ${CYAN}./lifecycle/kill.sh${NC}     Teardown containers (-f full, -d db, -b backend)"
echo -e "    ${CYAN}./lifecycle/update.sh${NC}   Safe dependency updates (npm audit fix + go patch)"
echo ""
echo -e "  ${BOLD}Quick start:${NC}"
echo -e "    ${GREEN}\$ ./lifecycle/serve.sh${NC}"
echo -e "    Then visit ${CYAN}http://localhost:8080${NC}"
echo ""
