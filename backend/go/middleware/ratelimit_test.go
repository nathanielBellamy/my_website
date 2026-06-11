package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

// okHandler is a trivial handler that always responds 200.
var okHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
})

func newTestLogger() zerolog.Logger {
	return zerolog.Nop()
}

// newRequest builds a request with RemoteAddr set to the given address.
func newRequest(remoteAddr string) *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.RemoteAddr = remoteAddr
	return r
}

// TestRateLimitMiddleware_LoopbackExempt verifies that loopback IPs
// (both IPv4 and IPv6) are never throttled regardless of request volume.
func TestRateLimitMiddleware_LoopbackExempt(t *testing.T) {
	// Extremely tight limiter: 1 req/sec, burst of 1.
	// Any non-exempt IP would be blocked after the first request.
	limiter := NewIPRateLimiter(rate.Limit(1), 1)
	log := newTestLogger()
	handler := RateLimitMiddleware(limiter, &log, okHandler)

	loopbacks := []string{
		"127.0.0.1:12345",
		"[::1]:12345",
	}

	for _, addr := range loopbacks {
		t.Run(addr, func(t *testing.T) {
			for i := range 20 {
				w := httptest.NewRecorder()
				handler.ServeHTTP(w, newRequest(addr))
				if w.Code != http.StatusOK {
					t.Errorf("request %d: expected 200 for loopback %s, got %d", i+1, addr, w.Code)
				}
			}
		})
	}
}

// TestRateLimitMiddleware_NonLoopbackThrottled verifies that external IPs
// are throttled once the burst is exhausted.
func TestRateLimitMiddleware_NonLoopbackThrottled(t *testing.T) {
	// Burst of 3 — requests 1-3 pass, request 4 should be 429.
	limiter := NewIPRateLimiter(rate.Limit(0), 3)
	log := newTestLogger()
	handler := RateLimitMiddleware(limiter, &log, okHandler)

	addr := "203.0.113.42:9999" // TEST-NET-3, never a real loopback

	for i := range 4 {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, newRequest(addr))
		want := http.StatusOK
		if i == 3 {
			want = http.StatusTooManyRequests
		}
		if w.Code != want {
			t.Errorf("request %d: expected %d, got %d", i+1, want, w.Code)
		}
	}
}

// TestRateLimitMiddleware_ExemptPath verifies that paths matching an exempt
// prefix are never throttled, even for non-loopback IPs.
func TestRateLimitMiddleware_ExemptPath(t *testing.T) {
	limiter := NewIPRateLimiter(rate.Limit(0), 1)
	log := newTestLogger()
	handler := RateLimitMiddleware(limiter, &log, okHandler, "/grafana/")

	addr := "203.0.113.1:9999"
	paths := []string{"/grafana/", "/grafana/metrics", "/grafana/d/abc"}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			for i := range 5 {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, path, nil)
				r.RemoteAddr = addr
				handler.ServeHTTP(w, r)
				if w.Code != http.StatusOK {
					t.Errorf("request %d on %s: expected 200, got %d", i+1, path, w.Code)
				}
			}
		})
	}
}
