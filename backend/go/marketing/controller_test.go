package marketing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nathanielBellamy/my_website/backend/go/testutils"

	"github.com/rs/zerolog"
)

type MockMarketingService struct {
	GetAllBlogPostsFunc      func(page, limit int) ([]BlogPost, error)
	GetBlogPostByIDFunc      func(id string) (*BlogPost, error)
	GetBlogPostsByTagFunc    func(tag string, page, limit int) ([]BlogPost, error)
	GetAllHomeContentFunc    func(page, limit int) ([]HomeContent, error)
	GetHomeContentByIDFunc   func(id string) (*HomeContent, error)
	GetAllGrooveJrContentFunc func(page, limit int) ([]GrooveJrContent, error)
	GetGrooveJrContentByIDFunc func(id string) (*GrooveJrContent, error)
	GetAllAboutContentFunc   func(page, limit int) ([]AboutContent, error)
	GetAboutContentByIDFunc  func(id string) (*AboutContent, error)
}

func (m *MockMarketingService) GetAllBlogPosts(page, limit int) ([]BlogPost, error) {
	return m.GetAllBlogPostsFunc(page, limit)
}
func (m *MockMarketingService) GetBlogPostByID(id string) (*BlogPost, error) {
	return m.GetBlogPostByIDFunc(id)
}
func (m *MockMarketingService) GetBlogPostsByTag(tag string, page, limit int) ([]BlogPost, error) {
	return m.GetBlogPostsByTagFunc(tag, page, limit)
}
func (m *MockMarketingService) GetAllHomeContent(page, limit int) ([]HomeContent, error) {
	return m.GetAllHomeContentFunc(page, limit)
}
func (m *MockMarketingService) GetHomeContentByID(id string) (*HomeContent, error) {
	return m.GetHomeContentByIDFunc(id)
}
func (m *MockMarketingService) GetAllGrooveJrContent(page, limit int) ([]GrooveJrContent, error) {
	return m.GetAllGrooveJrContentFunc(page, limit)
}
func (m *MockMarketingService) GetGrooveJrContentByID(id string) (*GrooveJrContent, error) {
	return m.GetGrooveJrContentByIDFunc(id)
}
func (m *MockMarketingService) GetAllAboutContent(page, limit int) ([]AboutContent, error) {
	return m.GetAllAboutContentFunc(page, limit)
}
func (m *MockMarketingService) GetAboutContentByID(id string) (*AboutContent, error) {
	return m.GetAboutContentByIDFunc(id)
}

func TestGetAllBlogPostsHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllBlogPostsFunc: func(page, limit int) ([]BlogPost, error) {
			return []BlogPost{{ID: "1", Title: "Test Post"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, err := http.NewRequest("GET", "/api/marketing/blog", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllBlogPostsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var posts []BlogPost
	if err := json.Unmarshal(rr.Body.Bytes(), &posts); err != nil {
		t.Fatal(err)
	}
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
}

func TestGetBlogPostByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetBlogPostByIDFunc: func(id string) (*BlogPost, error) {
			if id == "1" {
				return &BlogPost{ID: "1", Title: "Test Post"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/marketing/blog/{id}", controller.GetBlogPostByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/api/marketing/blog/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var post BlogPost
	json.Unmarshal(rr.Body.Bytes(), &post)
	if post.ID != "1" {
		t.Errorf("expected post ID 1, got %s", post.ID)
	}

	// Test not found
	req, _ = http.NewRequest("GET", "/api/marketing/blog/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetBlogPostsByTagHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetBlogPostsByTagFunc: func(tag string, page, limit int) ([]BlogPost, error) {
			if tag == "test" {
				return []BlogPost{{ID: "1", Title: "Tagged Post"}}, nil
			}
			return []BlogPost{}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/marketing/blog/tag/{tag}", controller.GetBlogPostsByTagHandler)

	req, _ := http.NewRequest("GET", "/api/marketing/blog/tag/test", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var posts []BlogPost
	json.Unmarshal(rr.Body.Bytes(), &posts)
	if len(posts) != 1 {
		t.Errorf("expected 1 tagged post, got %d", len(posts))
	}
}

func TestGetAllHomeContentHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllHomeContentFunc: func(page, limit int) ([]HomeContent, error) {
			return []HomeContent{{ID: "1", Title: "Home Content"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/api/marketing/home", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllHomeContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content []HomeContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if len(content) != 1 {
		t.Errorf("expected 1 home content, got %d", len(content))
	}
}

func TestGetHomeContentByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetHomeContentByIDFunc: func(id string) (*HomeContent, error) {
			if id == "1" {
				return &HomeContent{ID: "1", Title: "Home Content"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/marketing/home/{id}", controller.GetHomeContentByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/api/marketing/home/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content HomeContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if content.ID != "1" {
		t.Errorf("expected home content ID 1, got %s", content.ID)
	}

	// Test not found
	req, _ = http.NewRequest("GET", "/api/marketing/home/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetAllGrooveJrContentHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllGrooveJrContentFunc: func(page, limit int) ([]GrooveJrContent, error) {
			return []GrooveJrContent{{ID: "1", Title: "GrooveJr Content"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/api/marketing/groovejr", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllGrooveJrContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content []GrooveJrContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if len(content) != 1 {
		t.Errorf("expected 1 GrooveJr content, got %d", len(content))
	}
}

func TestGetGrooveJrContentByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetGrooveJrContentByIDFunc: func(id string) (*GrooveJrContent, error) {
			if id == "1" {
				return &GrooveJrContent{ID: "1", Title: "GrooveJr Content"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/marketing/groovejr/{id}", controller.GetGrooveJrContentByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/api/marketing/groovejr/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content GrooveJrContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if content.ID != "1" {
		t.Errorf("expected GrooveJr content ID 1, got %s", content.ID)
	}

	// Test not found
	req, _ = http.NewRequest("GET", "/api/marketing/groovejr/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetAllAboutContentHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllAboutContentFunc: func(page, limit int) ([]AboutContent, error) {
			return []AboutContent{{ID: "1", Title: "About Content"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/api/marketing/about", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllAboutContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content []AboutContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if len(content) != 1 {
		t.Errorf("expected 1 About content, got %d", len(content))
	}
}

func TestGetAboutContentByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAboutContentByIDFunc: func(id string) (*AboutContent, error) {
			if id == "1" {
				return &AboutContent{ID: "1", Title: "About Content"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/marketing/about/{id}", controller.GetAboutContentByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/api/marketing/about/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content AboutContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if content.ID != "1" {
		t.Errorf("expected About content ID 1, got %s", content.ID)
	}

	// Test not found
	req, _ = http.NewRequest("GET", "/api/marketing/about/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}