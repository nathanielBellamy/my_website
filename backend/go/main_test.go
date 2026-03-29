package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/admin"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	appLogs "github.com/nathanielBellamy/my_website/backend/go/logs"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/nathanielBellamy/my_website/backend/go/monitoring"
	"github.com/nathanielBellamy/my_website/backend/go/testutils"

	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

func TestMain(m *testing.M) {
	// Set MODE to localhost for testing purposes
	os.Setenv("MODE", "localhost")
	// Run tests
	code := m.Run()
	// Clean up
	os.Unsetenv("MODE")
	os.Exit(code)
}

func TestSetupBaseRoutes_MarketingBlogPosts(t *testing.T) {
	// Mock logger
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()

	// Mock dependencies
	cookieJar := cmap.New[auth.Cookie]()

	// Configure mock DB to return a sample blog post
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.BlogPost); ok {
				*v = []models.BlogPost{{ID: "1", Title: "Sample Blog Post"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.BlogPost); ok {
					*v = []models.BlogPost{{ID: "1", Title: "Sample Blog Post"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}

	// Create an http.ServeMux to register routes

	// Call SetupBaseRoutes to register handlers
	// Pass nil for oldSiteController as it's not relevant for this marketing test
	marketingService := marketing.NewService(mockDB)
	adminService := admin.NewService(mockDB, &log)
	marketingMux := http.NewServeMux()
	adminMux := http.NewServeMux()
	oldSiteMux := http.NewServeMux()
	startAt := time.Now()
	logDir := t.TempDir()
	logsController, err := appLogs.NewLogsController(&log, logDir, startAt)
	if err != nil {
		t.Fatalf("Failed to create logs controller: %v", err)
	}
	healthController := appLogs.NewHealthController(&log, startAt, mockDB)
	grafanaProxy := monitoring.NewGrafanaProxy(&log, "http://localhost:3000")
	SetupBaseRoutes(adminMux, oldSiteMux, marketingMux, &cookieJar, &log, nil, marketing.NewMarketingController(&log, marketingService), admin.NewAdminController(&log, adminService), logsController, healthController, grafanaProxy)

	// Create a request to the marketing blog posts endpoint
	req, err := http.NewRequest("GET", "/v1/api/marketing/blog?page=1&limit=5", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	marketingMux.ServeHTTP(rr, req)

	// Assert the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Assert the response body (basic check for non-empty array)
	var posts []models.BlogPost
	if err := json.Unmarshal(rr.Body.Bytes(), &posts); err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	if len(posts) == 0 {
		t.Errorf("Expected at least one blog post, got none")
	}

	t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
}
