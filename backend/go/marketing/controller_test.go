package marketing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
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
	err       error
	modelDest interface{}
	whereID   string // Track the ID from Where clause
}

// MockPgDB for pg.DB
type MockPgDB struct{}

// NewMockPgDB returns a new MockPgDB instance.
func NewMockPgDB() *MockPgDB {
	return &MockPgDB{}
}

// Model returns a new MockPgQuery
func (m *MockPgDB) Model(model ...interface{}) PgxQuerySeter {
	return &MockPgQuery{modelDest: model[0]}
}

func (mq *MockPgQuery) Relation(name string) PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Limit(count int) PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Offset(offset int) PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Where(query string, params ...interface{}) PgxQuerySeter {
	// Track the ID parameter for blog post queries
	if len(params) > 0 {
		if id, ok := params[0].(string); ok {
			mq.whereID = id
		}
	}
	return mq
}

func (mq *MockPgQuery) Join(join string, params ...interface{}) PgxQuerySeter {
	return mq
}

func (mq *MockPgQuery) Select(dest ...interface{}) error {
	if mq.err != nil {
		return mq.err
	}

	targetDest := mq.modelDest

	if len(dest) > 0 {
		targetDest = dest[0]
	}

	if targetDest != nil {
		if posts, ok := targetDest.(*[]BlogPost); ok {
			// Generate 10 mock blog posts for pagination tests
			*posts = make([]BlogPost, 10)
			for i := 0; i < 10; i++ {
				(*posts)[i] = BlogPost{
					ID:      "blog-post-id-" + string(rune('1'+i)),
					Title:   "Test Post",
					Content: "Content",
				}
			}
		} else if post, ok := targetDest.(*BlogPost); ok {
			// Single blog post - check if ID exists in our mock data
			if mq.whereID == "" || mq.whereID == "blog-id-1" {
				*post = BlogPost{
					ID:      "blog-id-1",
					Title:   "Test Post",
					Content: "Content",
				}
			} else {
				// Import pg package to use ErrNoRows
				return pg.ErrNoRows
			}
		}
	}
	return nil
}

func TestGetAllBlogPostsHandler(t *testing.T) {
	// Save original GetClientIpAddr and defer its restoration
	origGetClientIpAddr := auth.GetClientIpAddr
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	mockDB := NewMockPgDB()
	mc := NewMarketingController(&log, mockDB)

	t.Cleanup(func() {
		auth.GetClientIpAddr = origGetClientIpAddr
		t.Log(mockLogOutput.Buf.String()) // Log the buffer content
	})

	req, err := http.NewRequest("GET", "/api/blog?page=1&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mc.GetAllBlogPostsHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var posts []BlogPost
	if err := json.Unmarshal(rr.Body.Bytes(), &posts); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if len(posts) != 10 {
		t.Errorf("expected 10 blog posts, got %d", len(posts))
	}
	if posts[0].ID != "blog-post-id-1" || posts[1].ID != "blog-post-id-2" {
		t.Errorf("expected first two blog post IDs %q and %q, got %q and %q", "blog-post-id-1", "blog-post-id-2", posts[0].ID, posts[1].ID)
	}
}

func TestGetBlogPostByIDHandler(t *testing.T) {
	origGetClientIpAddr := auth.GetClientIpAddr
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	mockDB := NewMockPgDB()
	mc := NewMarketingController(&log, mockDB)
	t.Cleanup(func() {
		auth.GetClientIpAddr = origGetClientIpAddr
		t.Log(mockLogOutput.Buf.String()) // Log the buffer content
	})
	auth.GetClientIpAddr = func(r *http.Request) string { return "127.0.0.1" }

	// Create a test server
	router := http.NewServeMux()
	router.HandleFunc("/api/blog/{id}", mc.GetBlogPostByIDHandler)
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	// Test case for existing ID
	req, err := http.NewRequest("GET", testServer.URL+"/api/blog/blog-id-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := testServer.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var post BlogPost
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}
	if post.ID != "blog-id-1" {
		t.Errorf("expected blog post ID %q, got %q", "blog-id-1", post.ID)
	}

	// Test case for non-existing ID
	req, err = http.NewRequest("GET", testServer.URL+"/api/blog/99", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = testServer.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for non-existing ID: got %v want %v",
			status, http.StatusNotFound)
	}
}


