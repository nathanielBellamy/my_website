package metrics

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/rs/zerolog"
)

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "plain path",
			input:    "/v1/api/admin/blog",
			expected: "/v1/api/admin/blog",
		},
		{
			name:     "uuid segment",
			input:    "/v1/api/admin/blog/550e8400-e29b-41d4-a716-446655440000",
			expected: "/v1/api/admin/blog/:id",
		},
		{
			name:     "numeric segment",
			input:    "/v1/api/admin/work/42",
			expected: "/v1/api/admin/work/:id",
		},
		{
			name:     "root path",
			input:    "/",
			expected: "/",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "multiple ids",
			input:    "/v1/api/admin/blog/123/comments/456",
			expected: "/v1/api/admin/blog/:id/comments/:id",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizePath(tc.input)
			if result != tc.expected {
				t.Errorf("normalizePath(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsID(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"uuid", "550e8400-e29b-41d4-a716-446655440000", true},
		{"numeric", "123", true},
		{"word", "blog", false},
		{"empty", "", false},
		{"mixed", "abc123", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isID(tc.input)
			if result != tc.expected {
				t.Errorf("isID(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestInstrumentHandler_RecordsMetrics(t *testing.T) {
	log := zerolog.New(os.Stdout).Level(zerolog.Disabled)

	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	handler := InstrumentHandler(&log, inner)

	req := httptest.NewRequest("GET", "/v1/api/admin/health", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	// Verify that the counter was incremented
	count := testutil.ToFloat64(HttpRequestsTotal.WithLabelValues("GET", "/v1/api/admin/health", "200"))
	if count < 1 {
		t.Errorf("expected http_requests_total to be at least 1, got %f", count)
	}
}

func TestStatusWriter_CapturesStatusCode(t *testing.T) {
	rr := httptest.NewRecorder()
	sw := &statusWriter{ResponseWriter: rr, statusCode: http.StatusOK}

	sw.WriteHeader(http.StatusNotFound)

	if sw.statusCode != http.StatusNotFound {
		t.Errorf("expected statusCode 404, got %d", sw.statusCode)
	}
}
