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

# Function for marketing SPA tests
test_marketing_spa() {
  cat << EOF

  📣  🧪   TESTING:
MARKETING SPA
EOF
  cd marketing && npm test
  cd ..
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
test_marketing_spa

cat << EOF

  📣  🏁  DONE:
TESTING COMPLETE
CHECK ABOVE OUTPUT FOR FAILURES

EOF
