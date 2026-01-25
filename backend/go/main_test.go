package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	// "github.com/go-pg/pg/v10" // Removed unused import
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"

	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

// MockLogger is a mock implementation of zerolog.Logger for testing.
type MockLogger struct {
	Buf bytes.Buffer
}

func (m *MockLogger) Write(p []byte) (n int, err error) {
	return m.Buf.Write(p)
}

// MockPgQuery for pg.Query
type MockPgQuery struct {
	err error
	modelDest interface{} // New field to store the destination passed to Model
}

// Model returns a new MockPgQuery
func (m *MockPgDB) Model(model ...interface{}) marketing.PgxQuerySeter {
	// In a real mock, you might inspect 'model' to return different data
	// For this test, we'll return a generic mock query.
	return &MockPgQuery{modelDest: model[0]} // Store the model destination
}

func (mq *MockPgQuery) Relation(name string) marketing.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Limit(count int) marketing.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Offset(offset int) marketing.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Where(query string, params ...interface{}) marketing.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Join(join string, params ...interface{}) marketing.PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Select(dest ...interface{}) error {
	if mq.err != nil {
		return mq.err
	}

	targetDest := mq.modelDest // Default to modelDest

	if len(dest) > 0 {
		targetDest = dest[0] // If dest is provided, use it
	}

	if targetDest != nil {
		if posts, ok := targetDest.(*[]marketing.BlogPost); ok {
			*posts = []marketing.BlogPost{
				{ID: "1", Title: "Test Post 1", Content: "Content 1", Author: &marketing.Author{Name: "Author 1"}},
				{ID: "2", Title: "Test Post 2", Content: "Content 2", Author: &marketing.Author{Name: "Author 2"}},
			}
		}
	}
	return nil
}

// MockPgDB for pg.DB
type MockPgDB struct {
	// You might embed pg.DB if you only want to mock certain methods
	// For now, we'll mock just what's needed.
}

// NewMockPgDB returns a new MockPgDB instance.
func NewMockPgDB() *MockPgDB {
	return &MockPgDB{}
}

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
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()

	// Mock dependencies
	cookieJar := cmap.New[auth.Cookie]()
	mockDB := NewMockPgDB()
	
	// Create an http.ServeMux to register routes
	mux := http.NewServeMux()

	// Call SetupBaseRoutes to register handlers
	// Pass nil for oldSiteController as it's not relevant for this marketing test
	SetupBaseRoutes(mux, &cookieJar, &log, nil, marketing.NewMarketingController(&log, mockDB))

	// Create a request to the marketing blog posts endpoint
	req, err := http.NewRequest("GET", "/api/marketing/blog?page=1&limit=5", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	// Assert the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Assert the response body (basic check for non-empty array)
	var posts []marketing.BlogPost
	if err := json.Unmarshal(rr.Body.Bytes(), &posts); err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	if len(posts) == 0 {
		t.Errorf("Expected at least one blog post, got none")
	}

	t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
}