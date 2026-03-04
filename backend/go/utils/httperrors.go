package utils

import (
	"net"
	"net/http"
	"strings"
)

// HandleDBError analyzes an error, especially from a database operation,
// and writes an appropriate HTTP response. It defaults to 500 but returns 503
// if it detects a network connection issue.
func HandleDBError(w http.ResponseWriter, err error, defaultMsg string) {
	if err != nil {
		errStr := err.Error()
		// Common network connection and go-pg errors
		if strings.Contains(errStr, "dial tcp") ||
			strings.Contains(errStr, "connection refused") ||
			strings.Contains(errStr, "i/o timeout") ||
			strings.Contains(errStr, "EOF") ||
			strings.Contains(errStr, "broken pipe") ||
			strings.Contains(errStr, "connection reset by peer") ||
			strings.Contains(errStr, "no route to host") {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
	}

	http.Error(w, defaultMsg, http.StatusInternalServerError)
}
