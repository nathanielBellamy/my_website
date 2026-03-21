package metrics

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.statusCode = code
	sw.ResponseWriter.WriteHeader(code)
}

// Flush delegates to the underlying ResponseWriter if it supports flushing (needed for SSE)
func (sw *statusWriter) Flush() {
	if f, ok := sw.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// Unwrap returns the underlying ResponseWriter so http.NewResponseController can reach it
func (sw *statusWriter) Unwrap() http.ResponseWriter {
	return sw.ResponseWriter
}

// InstrumentHandler wraps an http.Handler to record request metrics
func InstrumentHandler(log *zerolog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(sw, r)

		duration := time.Since(start).Seconds()
		path := normalizePath(r.URL.Path)

		HttpRequestsTotal.WithLabelValues(r.Method, path, strconv.Itoa(sw.statusCode)).Inc()
		HttpRequestDuration.WithLabelValues(r.Method, path).Observe(duration)
	})
}

// normalizePath reduces cardinality by collapsing dynamic path segments
func normalizePath(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		// Replace UUIDs and numeric IDs with a placeholder
		if isID(part) {
			parts[i] = ":id"
		}
	}
	return strings.Join(parts, "/")
}

func isID(s string) bool {
	if s == "" {
		return false
	}
	// Check for UUID format (8-4-4-4-12 hex)
	if len(s) == 36 && s[8] == '-' && s[13] == '-' && s[18] == '-' && s[23] == '-' {
		return true
	}
	// Check for numeric IDs
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}
