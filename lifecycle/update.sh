#!/bin/bash
#
# Runs safe dependency updates across all projects.
#   - npm update, then npm audit --omit=dev for Angular frontends
#   - npm audit fix --legacy-peer-deps for legacy frontend
#   - go get -u=patch for the Go backend (patch-level updates only)
#
cd "$(dirname "$0")/.."

FRONTEND_PROJECTS=("frontend/marketing" "frontend/admin" "frontend/auth/admin" "frontend/old-site")

for project in "${FRONTEND_PROJECTS[@]}"; do
  if [ -d "$project" ] && [ -f "$project/package.json" ]; then
    cat << EOF

  📣  🔧  UPDATING:
${project}

EOF
    if [[ "$project" == "frontend/old-site" ]]; then
      (cd "$project" && npm audit fix --legacy-peer-deps)
    else
      (cd "$project" && npm update --audit=false && npm audit --omit=dev)
    fi
  fi
done

# Go backend — patch-level updates only (equivalent of npm audit fix)
if [ -d "backend/go" ]; then
  cat << EOF

  📣  🔧  UPDATING:
backend/go

EOF
  (cd backend/go && go get -u=patch ./... && go mod tidy)
fi

cat << EOF

  📣  🏁  DONE:
ALL DEPENDENCY UPDATES COMPLETE

EOF
