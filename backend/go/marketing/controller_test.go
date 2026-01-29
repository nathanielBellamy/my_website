package marketing_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"
	"github.com/rs/zerolog"
)

// MockLogger is a mock implementation of zerolog.Logger for testing.
type MockLogger struct {
	Buf bytes.Buffer
}

func (m *MockLogger) Write(p []byte) (n int, err error) {
	return m.Buf.Write(p)
}

// MockService is a mock implementation of the Service for testing.
type MockService struct{}

func (m *MockService) GetAllBlogPosts(page, limit int) ([]marketing.BlogPost, error) {
	var posts []marketing.BlogPost
	for i := 1; i <= limit; i++ {
		posts = append(posts, marketing.BlogPost{ID: "blog-post-id-" + strconv.Itoa(i)})
	}
	return posts, nil
}

func (m *MockService) GetBlogPostByID(id string) (*marketing.BlogPost, error) {
	if id == "blog-id-1" {
		return &marketing.BlogPost{ID: "blog-id-1"}, nil
	}
	return nil, nil
}

func (m *MockService) GetBlogPostsByTag(tag string, page, limit int) ([]marketing.BlogPost, error) {
	return nil, nil
}
func (m *MockService) GetAllHomeContent(page, limit int) ([]marketing.HomeContent, error) {
	return nil, nil
}
func (m *MockService) GetHomeContentByID(id string) (*marketing.HomeContent, error) {
	return nil, nil
}
func (m *MockService) GetAllGrooveJrContent(page, limit int) ([]marketing.GrooveJrContent, error) {
	return nil, nil
}
func (m *MockService) GetGrooveJrContentByID(id string) (*marketing.GrooveJrContent, error) {
	return nil, nil
}
func (m *MockService) GetAllAboutContent(page, limit int) ([]marketing.AboutContent, error) {
	return nil, nil
}
func (m *MockService) GetAboutContentByID(id string) (*marketing.AboutContent, error) {
	return nil, nil
}

func TestGetAllBlogPostsHandler(t *testing.T) {
	// Save original GetClientIpAddr and defer its restoration
	origGetClientIpAddr := auth.GetClientIpAddr
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	service := &MockService{}
	mc := marketing.NewMarketingController(&log, service)

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

	var posts []marketing.BlogPost
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
	service := &MockService{}
	mc := marketing.NewMarketingController(&log, service)
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

	var post marketing.BlogPost
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
