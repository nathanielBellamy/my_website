#!/bin/bash
#
# Runs safe dependency updates across all projects.
#   - npm audit fix   for each frontend project and e2e
#   - go get -u=patch for the Go backend (patch-level updates only)
#
cd "$(dirname "$0")/.."

FRONTEND_PROJECTS=("frontend/marketing" "frontend/admin" "frontend/auth" "frontend/old-site")

for project in "${FRONTEND_PROJECTS[@]}"; do
  if [ -d "$project" ] && [ -f "$project/package.json" ]; then
    cat << EOF

  📣  🔧  UPDATING:
${project}

EOF
    if [[ "$project" == "frontend/old-site" || "$project" == "frontend/auth" ]]; then
      (cd "$project" && npm audit fix --legacy-peer-deps)
    else
      (cd "$project" && npm audit fix)
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
