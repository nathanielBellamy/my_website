package monitoring

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestGrafanaProxy_ForwardsRequests(t *testing.T) {
	// Create a mock Grafana server
	grafanaServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify auth proxy header
		user := r.Header.Get("X-WEBAUTH-USER")
		if user != "admin" {
			t.Errorf("expected X-WEBAUTH-USER=admin, got %q", user)
		}

		// Verify path stripping
		w.Header().Set("X-Received-Path", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("grafana response"))
	}))
	defer grafanaServer.Close()

	log := zerolog.New(os.Stdout).Level(zerolog.Disabled)
	proxy := NewGrafanaProxy(&log, grafanaServer.URL)

	req := httptest.NewRequest("GET", "/grafana/d/system-overview", nil)
	rr := httptest.NewRecorder()

	proxy.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	if string(body) != "grafana response" {
		t.Errorf("expected 'grafana response', got %q", string(body))
	}

	// Verify path was stripped
	receivedPath := rr.Header().Get("X-Received-Path")
	if receivedPath != "/d/system-overview" {
		t.Errorf("expected stripped path '/d/system-overview', got %q", receivedPath)
	}
}

func TestGrafanaProxy_StripsGrafanaPrefix(t *testing.T) {
	grafanaServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Received-Path", r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer grafanaServer.Close()

	log := zerolog.New(os.Stdout).Level(zerolog.Disabled)
	proxy := NewGrafanaProxy(&log, grafanaServer.URL)

	tests := []struct {
		name         string
		inputPath    string
		expectedPath string
	}{
		{
			name:         "root grafana path",
			inputPath:    "/grafana",
			expectedPath: "/",
		},
		{
			name:         "grafana with trailing slash",
			inputPath:    "/grafana/",
			expectedPath: "/",
		},
		{
			name:         "grafana dashboard",
			inputPath:    "/grafana/d/system-overview?orgId=1",
			expectedPath: "/d/system-overview",
		},
		{
			name:         "grafana api",
			inputPath:    "/grafana/api/health",
			expectedPath: "/api/health",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tc.inputPath, nil)
			rr := httptest.NewRecorder()

			proxy.ServeHTTP(rr, req)

			receivedPath := rr.Header().Get("X-Received-Path")
			if receivedPath != tc.expectedPath {
				t.Errorf("for %q: expected path %q, got %q", tc.inputPath, tc.expectedPath, receivedPath)
			}
		})
	}
}

func TestGrafanaProxy_SetsAuthHeader(t *testing.T) {
	var receivedHeaders http.Header

	grafanaServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeaders = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	}))
	defer grafanaServer.Close()

	log := zerolog.New(os.Stdout).Level(zerolog.Disabled)
	proxy := NewGrafanaProxy(&log, grafanaServer.URL)

	req := httptest.NewRequest("GET", "/grafana/", nil)
	rr := httptest.NewRecorder()

	proxy.ServeHTTP(rr, req)

	if receivedHeaders.Get("X-WEBAUTH-USER") != "admin" {
		t.Errorf("X-WEBAUTH-USER header not set to 'admin'")
	}
}
