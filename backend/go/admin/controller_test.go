package admin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"

	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/rs/zerolog"
)

// MockLogger is a mock implementation of zerolog.Logger for testing.
type MockLogger struct {
	Buf bytes.Buffer
}

func (m *MockLogger) Write(p []byte) (n int, err error) {
	return m.Buf.Write(p)
}

type MockAdminService struct {
	GetAllBlogPostsFunc      func(filter models.FilterOptions) ([]models.BlogPost, int, error)
	GetBlogPostByIDFunc      func(id string) (*models.BlogPost, error)
	GetBlogPostsByTagFunc    func(tag string, page, limit int) ([]models.BlogPost, error)
	CreateBlogPostFunc       func(post *models.BlogPost) (*models.BlogPost, error)
	UpdateBlogPostFunc       func(post *models.BlogPost) (*models.BlogPost, error)
	DeleteBlogPostFunc       func(id string) error
	GetTagsFunc              func(search string, limit int) ([]models.TagWithUsage, error)
	GetAllHomeContentFunc    func(filter models.FilterOptions) ([]models.HomeContent, int, error)
	GetHomeContentByIDFunc   func(id string) (*models.HomeContent, error)
	CreateHomeContentFunc    func(content *models.HomeContent) (*models.HomeContent, error)
	UpdateHomeContentFunc    func(content *models.HomeContent) (*models.HomeContent, error)
	DeleteHomeContentFunc    func(id string) error
	GetAllGrooveJrContentFunc func(filter models.FilterOptions) ([]models.GrooveJrContent, int, error)
	GetGrooveJrContentByIDFunc func(id string) (*models.GrooveJrContent, error)
	CreateGrooveJrContentFunc func(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	UpdateGrooveJrContentFunc func(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	DeleteGrooveJrContentFunc func(id string) error
	GetAllAboutContentFunc   func(filter models.FilterOptions) ([]models.AboutContent, int, error)
	GetAboutContentByIDFunc  func(id string) (*models.AboutContent, error)
	CreateAboutContentFunc   func(content *models.AboutContent) (*models.AboutContent, error)
	UpdateAboutContentFunc   func(content *models.AboutContent) (*models.AboutContent, error)
	DeleteAboutContentFunc   func(id string) error
}

func (m *MockAdminService) GetAllBlogPosts(filter models.FilterOptions) ([]models.BlogPost, int, error) {
	return m.GetAllBlogPostsFunc(filter)
}
func (m *MockAdminService) GetBlogPostByID(id string) (*models.BlogPost, error) {
	return m.GetBlogPostByIDFunc(id)
}
func (m *MockAdminService) GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error) {
	return m.GetBlogPostsByTagFunc(tag, page, limit)
}
func (m *MockAdminService) CreateBlogPost(post *models.BlogPost) (*models.BlogPost, error) {
	return m.CreateBlogPostFunc(post)
}
func (m *MockAdminService) UpdateBlogPost(post *models.BlogPost) (*models.BlogPost, error) {
	return m.UpdateBlogPostFunc(post)
}
func (m *MockAdminService) DeleteBlogPost(id string) error {
	return m.DeleteBlogPostFunc(id)
}
func (m *MockAdminService) GetTags(search string, limit int) ([]models.TagWithUsage, error) {
	return m.GetTagsFunc(search, limit)
}
func (m *MockAdminService) GetAllHomeContent(filter models.FilterOptions) ([]models.HomeContent, int, error) {
	return m.GetAllHomeContentFunc(filter)
}
func (m *MockAdminService) GetHomeContentByID(id string) (*models.HomeContent, error) {
	return m.GetHomeContentByIDFunc(id)
}
func (m *MockAdminService) CreateHomeContent(content *models.HomeContent) (*models.HomeContent, error) {
	return m.CreateHomeContentFunc(content)
}
func (m *MockAdminService) UpdateHomeContent(content *models.HomeContent) (*models.HomeContent, error) {
	return m.UpdateHomeContentFunc(content)
}
func (m *MockAdminService) DeleteHomeContent(id string) error {
	return m.DeleteHomeContentFunc(id)
}
func (m *MockAdminService) GetAllGrooveJrContent(filter models.FilterOptions) ([]models.GrooveJrContent, int, error) {
	return m.GetAllGrooveJrContentFunc(filter)
}
func (m *MockAdminService) GetGrooveJrContentByID(id string) (*models.GrooveJrContent, error) {
	return m.GetGrooveJrContentByIDFunc(id)
}
func (m *MockAdminService) CreateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error) {
	return m.CreateGrooveJrContentFunc(content)
}
func (m *MockAdminService) UpdateGrooveJrContent(content *models.GrooveJrContent) (*models.GrooveJrContent, error) {
	return m.UpdateGrooveJrContentFunc(content)
}
func (m *MockAdminService) DeleteGrooveJrContent(id string) error {
	return m.DeleteGrooveJrContentFunc(id)
}
func (m *MockAdminService) GetAllAboutContent(filter models.FilterOptions) ([]models.AboutContent, int, error) {
	return m.GetAllAboutContentFunc(filter)
}
func (m *MockAdminService) GetAboutContentByID(id string) (*models.AboutContent, error) {
	return m.GetAboutContentByIDFunc(id)
}
func (m *MockAdminService) CreateAboutContent(content *models.AboutContent) (*models.AboutContent, error) {
	return m.CreateAboutContentFunc(content)
}
func (m *MockAdminService) UpdateAboutContent(content *models.AboutContent) (*models.AboutContent, error) {
	return m.UpdateAboutContentFunc(content)
}
func (m *MockAdminService) DeleteAboutContent(id string) error {
	return m.DeleteAboutContentFunc(id)
}

func TestAdminGetAllBlogPostsHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetAllBlogPostsFunc: func(filter models.FilterOptions) ([]models.BlogPost, int, error) {
			return []models.BlogPost{{ID: "1", Title: "Test Post Admin"}}, 1, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	req, err := http.NewRequest("GET", "/api/admin/blog", nil)
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

	var resp models.PaginatedResponse[models.BlogPost]
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 post, got %d", len(resp.Data))
	}
}

func TestAdminGetBlogPostByIDHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetBlogPostByIDFunc: func(id string) (*models.BlogPost, error) {
			if id == "1" {
				return &models.BlogPost{ID: "1", Title: "Test Post"}, nil
			}
			return nil, nil // Controller expects nil, nil for not found
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/blog/{id}", controller.GetBlogPostByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/api/admin/blog/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req) // Use the testMux

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var post models.BlogPost
	json.Unmarshal(rr.Body.Bytes(), &post)
	if post.ID != "1" {
		t.Errorf("expected post ID 1, got %s", post.ID)
	}

	// Test not found
	req, _ = http.NewRequest("GET", "/api/admin/blog/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req) // Use the testMux

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestAdminGetBlogPostsByTagHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetBlogPostsByTagFunc: func(tag string, page, limit int) ([]models.BlogPost, error) {
			if tag == "test" {
				return []models.BlogPost{{ID: "1", Title: "Tagged Post"}}, nil
			}
			return []models.BlogPost{}, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/blog/tag/{tag}", controller.GetBlogPostsByTagHandler)

	req, _ := http.NewRequest("GET", "/api/admin/blog/tag/test", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req) // Use the testMux

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var posts []models.BlogPost
	json.Unmarshal(rr.Body.Bytes(), &posts)
	if len(posts) != 1 {
		t.Errorf("expected 1 tagged post, got %d", len(posts))
	}
}

func TestAdminCreateBlogPostHandler(t *testing.T) {
	mockService := &MockAdminService{
		CreateBlogPostFunc: func(post *models.BlogPost) (*models.BlogPost, error) {
			post.ID = "new-id"
			return post, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	postData := models.BlogPost{Title: "New Post"}
	jsonBody, _ := json.Marshal(postData)
	req, _ := http.NewRequest("POST", "/api/admin/blog", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateBlogPostHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var newPost models.BlogPost
	json.Unmarshal(rr.Body.Bytes(), &newPost)
	if newPost.ID != "new-id" {
		t.Errorf("expected new post ID 'new-id', got %s", newPost.ID)
	}
}

func TestAdminUpdateBlogPostHandler(t *testing.T) {
	mockService := &MockAdminService{
		UpdateBlogPostFunc: func(post *models.BlogPost) (*models.BlogPost, error) {
			if post.ID == "1" {
				post.Title = "Updated Title"
				return post, nil
			}
			return nil, errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/blog/{id}", controller.UpdateBlogPostHandler)

	postData := models.BlogPost{ID: "1", Title: "Updated Post"}
	jsonBody, _ := json.Marshal(postData)
	req, _ := http.NewRequest("PUT", "/api/admin/blog/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req) // Use the testMux

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var updatedPost models.BlogPost
	json.Unmarshal(rr.Body.Bytes(), &updatedPost)
	if updatedPost.Title != "Updated Title" {
		t.Errorf("expected updated title 'Updated Title', got %s", updatedPost.Title)
	}
}

func TestAdminDeleteBlogPostHandler(t *testing.T) {
	mockService := &MockAdminService{
		DeleteBlogPostFunc: func(id string) error {
			if id == "1" {
				return nil
			}
			return errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/blog/{id}", controller.DeleteBlogPostHandler)

	req, _ := http.NewRequest("DELETE", "/api/admin/blog/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req) // Use the testMux

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestAdminGetAllHomeContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetAllHomeContentFunc: func(filter models.FilterOptions) ([]models.HomeContent, int, error) {
			return []models.HomeContent{{ID: "1", Title: "Test Home Admin"}}, 1, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	req, _ := http.NewRequest("GET", "/api/admin/home", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllHomeContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var resp models.PaginatedResponse[models.HomeContent]
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 home content, got %d", len(resp.Data))
	}
}

func TestAdminGetHomeContentByIDHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetHomeContentByIDFunc: func(id string) (*models.HomeContent, error) {
			if id == "1" {
				return &models.HomeContent{ID: "1", Title: "Test Home"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/home/{id}", controller.GetHomeContentByIDHandler)

	req, _ := http.NewRequest("GET", "/api/admin/home/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content models.HomeContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if content.ID != "1" {
		t.Errorf("expected home content ID 1, got %s", content.ID)
	}

	req, _ = http.NewRequest("GET", "/api/admin/home/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestAdminCreateHomeContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		CreateHomeContentFunc: func(content *models.HomeContent) (*models.HomeContent, error) {
			content.ID = "new-id-home"
			return content, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	contentData := models.HomeContent{Title: "New Home Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("POST", "/api/admin/home", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateHomeContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var newContent models.HomeContent
	json.Unmarshal(rr.Body.Bytes(), &newContent)
	if newContent.ID != "new-id-home" {
		t.Errorf("expected new home content ID 'new-id-home', got %s", newContent.ID)
	}
}

func TestAdminUpdateHomeContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		UpdateHomeContentFunc: func(content *models.HomeContent) (*models.HomeContent, error) {
			if content.ID == "1" {
				content.Title = "Updated Home Title"
				return content, nil
			}
			return nil, errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/home/{id}", controller.UpdateHomeContentHandler)

	contentData := models.HomeContent{ID: "1", Title: "Updated Home Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("PUT", "/api/admin/home/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var updatedContent models.HomeContent
	json.Unmarshal(rr.Body.Bytes(), &updatedContent)
	if updatedContent.Title != "Updated Home Title" {
		t.Errorf("expected updated title 'Updated Home Title', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteHomeContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		DeleteHomeContentFunc: func(id string) error {
			if id == "1" {
				return nil
			}
			return errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/home/{id}", controller.DeleteHomeContentHandler)

	req, _ := http.NewRequest("DELETE", "/api/admin/home/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestAdminGetAllGrooveJrContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetAllGrooveJrContentFunc: func(filter models.FilterOptions) ([]models.GrooveJrContent, int, error) {
			return []models.GrooveJrContent{{ID: "1", Title: "Test GrooveJr Admin"}}, 1, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	req, _ := http.NewRequest("GET", "/api/admin/groovejr", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllGrooveJrContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var resp models.PaginatedResponse[models.GrooveJrContent]
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 GrooveJr content, got %d", len(resp.Data))
	}
}

func TestAdminGetGrooveJrContentByIDHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetGrooveJrContentByIDFunc: func(id string) (*models.GrooveJrContent, error) {
			if id == "1" {
				return &models.GrooveJrContent{ID: "1", Title: "Test GrooveJr"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/groovejr/{id}", controller.GetGrooveJrContentByIDHandler)

	req, _ := http.NewRequest("GET", "/api/admin/groovejr/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content models.GrooveJrContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if content.ID != "1" {
		t.Errorf("expected GrooveJr content ID 1, got %s", content.ID)
	}

	req, _ = http.NewRequest("GET", "/api/admin/groovejr/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestAdminCreateGrooveJrContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		CreateGrooveJrContentFunc: func(content *models.GrooveJrContent) (*models.GrooveJrContent, error) {
			content.ID = "new-id-groovejr"
			return content, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	contentData := models.GrooveJrContent{Title: "New GrooveJr Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("POST", "/api/admin/groovejr", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateGrooveJrContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var newContent models.GrooveJrContent
	json.Unmarshal(rr.Body.Bytes(), &newContent)
	if newContent.ID != "new-id-groovejr" {
		t.Errorf("expected new GrooveJr content ID 'new-id-groovejr', got %s", newContent.ID)
	}
}

func TestAdminUpdateGrooveJrContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		UpdateGrooveJrContentFunc: func(content *models.GrooveJrContent) (*models.GrooveJrContent, error) {
			if content.ID == "1" {
				content.Title = "Updated GrooveJr Title"
				return content, nil
			}
			return nil, errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/groovejr/{id}", controller.UpdateGrooveJrContentHandler)

	contentData := models.GrooveJrContent{ID: "1", Title: "Updated GrooveJr Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("PUT", "/api/admin/groovejr/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var updatedContent models.GrooveJrContent
	json.Unmarshal(rr.Body.Bytes(), &updatedContent)
	if updatedContent.Title != "Updated GrooveJr Title" {
		t.Errorf("expected updated title 'Updated GrooveJr Title', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteGrooveJrContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		DeleteGrooveJrContentFunc: func(id string) error {
			if id == "1" {
				return nil
			}
			return errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/groovejr/{id}", controller.DeleteGrooveJrContentHandler)

	req, _ := http.NewRequest("DELETE", "/api/admin/groovejr/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
func TestAdminGetAllAboutContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetAllAboutContentFunc: func(filter models.FilterOptions) ([]models.AboutContent, int, error) {
			return []models.AboutContent{{ID: "1", Title: "Test About Admin"}}, 1, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	req, _ := http.NewRequest("GET", "/api/admin/about", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllAboutContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var resp models.PaginatedResponse[models.AboutContent]
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 About content, got %d", len(resp.Data))
	}
}

func TestAdminGetAboutContentByIDHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetAboutContentByIDFunc: func(id string) (*models.AboutContent, error) {
			if id == "1" {
				return &models.AboutContent{ID: "1", Title: "Test About"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/about/{id}", controller.GetAboutContentByIDHandler)

	req, _ := http.NewRequest("GET", "/api/admin/about/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content models.AboutContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if content.ID != "1" {
		t.Errorf("expected About content ID 1, got %s", content.ID)
	}

	req, _ = http.NewRequest("GET", "/api/admin/about/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestAdminCreateAboutContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		CreateAboutContentFunc: func(content *models.AboutContent) (*models.AboutContent, error) {
			content.ID = "new-id-about"
			return content, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	contentData := models.AboutContent{Title: "New About Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("POST", "/api/admin/about", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateAboutContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var newContent models.AboutContent
	json.Unmarshal(rr.Body.Bytes(), &newContent)
	if newContent.ID != "new-id-about" {
		t.Errorf("expected new About content ID 'new-id-about', got %s", newContent.ID)
	}
}

func TestAdminUpdateAboutContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		UpdateAboutContentFunc: func(content *models.AboutContent) (*models.AboutContent, error) {
			if content.ID == "1" {
				content.Title = "Updated About Title"
				return content, nil
			}
			return nil, errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/about/{id}", controller.UpdateAboutContentHandler)

	contentData := models.AboutContent{ID: "1", Title: "Updated About Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("PUT", "/api/admin/about/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var updatedContent models.AboutContent
	json.Unmarshal(rr.Body.Bytes(), &updatedContent)
	if updatedContent.Title != "Updated About Title" {
		t.Errorf("expected updated title 'Updated About Title', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteAboutContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		DeleteAboutContentFunc: func(id string) error {
			if id == "1" {
				return nil
			}
			return errors.New("not found")
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/api/admin/about/{id}", controller.DeleteAboutContentHandler)

	req, _ := http.NewRequest("DELETE", "/api/admin/about/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
