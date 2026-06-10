#!/bin/bash
#
# install.sh
#
# Installs dependencies for all active projects.
#
# Usage:
#   ./lifecycle/install.sh              # skips frontend/old-site (default)
#   ./lifecycle/install.sh --include-old-site
#
cd "$(dirname "$0")/.."

INCLUDE_OLD_SITE=false

for arg in "$@"; do
  case "$arg" in
    --include-old-site)
      INCLUDE_OLD_SITE=true
      ;;
    *)
      echo "Unknown argument: $arg"
      echo "Usage: $0 [--include-old-site]"
      exit 1
      ;;
  esac
done

# ─────────────────────────────────────────
# Helpers
# ─────────────────────────────────────────

print_header() {
  cat << EOF

  📦  INSTALLING: $1
EOF
}

npm_install() {
  local dir="$1"
  print_header "$dir"
  (cd "$dir" && npm install) || {
    echo "❌  npm install failed in $dir"
    exit 1
  }
  echo "✅  $dir"
}

# ─────────────────────────────────────────
# Go (root module)
# ─────────────────────────────────────────

print_header "backend/go  (go mod download)"
go mod download || {
  echo "❌  go mod download failed"
  exit 1
}
echo "✅  backend/go"

# ─────────────────────────────────────────
# Rust / WASM  (frontend/marketing/src-rust)
#
# src-rust is a cdylib library crate (WASM).
# `cargo fetch` downloads all dependencies
# without building — the Cargo equivalent of
# `npm install`.
# ─────────────────────────────────────────

print_header "frontend/marketing/src-rust  (cargo fetch)"
(cd frontend/marketing/src-rust && cargo fetch) || {
  echo "❌  cargo fetch failed in frontend/marketing/src-rust"
  exit 1
}
echo "✅  frontend/marketing/src-rust"

# ─────────────────────────────────────────
# npm projects
# ─────────────────────────────────────────

npm_install "frontend/marketing"
npm_install "frontend/admin"
npm_install "frontend/auth"
npm_install "frontend/auth/admin"

# ─────────────────────────────────────────
# old-site (opt-in)
# ─────────────────────────────────────────

if [ "$INCLUDE_OLD_SITE" = true ]; then
  npm_install "frontend/old-site"
else
  echo ""
  echo "  ⏭️   Skipping frontend/old-site  (pass --include-old-site to include)"
fi

echo ""
echo "  🎉  All installs complete."
echo ""
