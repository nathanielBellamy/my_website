#!/bin/bash
#
cd "$(dirname "$0")/.."

# Function for Go server tests
test_go_server() {
    cat << EOF
  📣  🧪   TESTING:
GO SERVER
EOF
  go test ./backend/go/...
  cat << EOF
  📣  🏁  DONE:
GO SERVER TESTED
EOF
}

# TODO_NS: test future angular auth spa

# Function for auth admin SPA tests
test_auth_admin_spa() {
  cat << EOF

  📣  🧪   TESTING:
AUTH ADMIN SPA
EOF
  cd frontend/auth/admin && npm test
  cd ../../..
  cat << EOF

  📣  🏁  DONE:
AUTH ADMIN SPA TESTED
EOF
}

# Function for admin SPA tests
test_admin_spa() {
  cat << EOF

  📣  🧪   TESTING:
ADMIN SPA
EOF
  cd frontend/admin && npm test
  cd ../..
  cat << EOF

  📣  🏁  DONE:
ADMIN SPA TESTED
EOF
}

# Function for marketing SPA tests
test_marketing_spa() {
  cat << EOF

  📣  🧪   TESTING:
MARKETING SPA
EOF
  cd frontend/marketing && npm test
  cd ../..
  cat << EOF

  📣  🏁  DONE:
MARKETING SPA TESTED
EOF
}

######

cat << EOF

  📣  🧪   TESTING WEBSITE

EOF

test_go_server
test_auth_admin_spa
test_admin_spa
test_marketing_spa

cat << EOF

  📣  🏁  DONE:
TESTING COMPLETE
CHECK ABOVE OUTPUT FOR FAILURES

EOF
