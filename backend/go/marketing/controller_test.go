package marketing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func TestGetAllBlogPostsHandler(t *testing.T) {
	// Save original GetClientIpAddr and defer its restoration
	origGetClientIpAddr := auth.GetClientIpAddr
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	mc := NewMarketingController(&log)

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

	if len(posts) != 2 {
		t.Errorf("expected 2 blog posts, got %d", len(posts))
	}
	if posts[0].ID != "blog-post-1" || posts[1].ID != "blog-post-2" {
		t.Errorf("expected blog post IDs 1 and 2, got %d and %d", posts[0].ID, posts[1].ID)
	}
}

func TestGetBlogPostByIDHandler(t *testing.T) {
	origGetClientIpAddr := auth.GetClientIpAddr
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	mc := NewMarketingController(&log)
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
	req, err := http.NewRequest("GET", testServer.URL+"/api/blog/1", nil)
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
	if post.ID != "blog-post-1]" {
		t.Errorf("expected blog post ID 1, got %d", post.ID)
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

func TestPostTrackerDataHandler(t *testing.T) {
	origGetClientIpAddr := auth.GetClientIpAddr
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	mc := NewMarketingController(&log)
	t.Cleanup(func() {
		auth.GetClientIpAddr = origGetClientIpAddr
		t.Log(mockLogOutput.Buf.String()) // Log the buffer content
	})

	req, err := http.NewRequest("POST", "/api/tracker", strings.NewReader("some tracking data"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mc.PostTrackerDataHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if body := rr.Body.String(); body != "Tracker data received" {
		t.Errorf("handler returned unexpected body: got %q want %q",
			body, "Tracker data received")
	}
}
