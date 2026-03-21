package admin

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

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
	GetAllBlogPostsFunc        func(filter models.FilterOptions) ([]models.BlogPost, int, error)
	GetBlogPostByIDFunc        func(id string) (*models.BlogPost, error)
	GetBlogPostsByTagFunc      func(tag string, page, limit int) ([]models.BlogPost, error)
	CreateBlogPostFunc         func(post *models.BlogPost) (*models.BlogPost, error)
	UpdateBlogPostFunc         func(post *models.BlogPost) (*models.BlogPost, error)
	DeleteBlogPostFunc         func(id string) error
	GetTagsFunc                func(search string, limit int) ([]models.TagWithUsage, error)
	GetAllWorkContentFunc      func(filter models.FilterOptions) ([]models.WorkContent, int, error)
	GetWorkContentByIDFunc     func(id string) (*models.WorkContent, error)
	CreateWorkContentFunc      func(content *models.WorkContent) (*models.WorkContent, error)
	UpdateWorkContentFunc      func(content *models.WorkContent) (*models.WorkContent, error)
	DeleteWorkContentFunc      func(id string) error
	GetAllGrooveJrContentFunc  func(filter models.FilterOptions) ([]models.GrooveJrContent, int, error)
	GetGrooveJrContentByIDFunc func(id string) (*models.GrooveJrContent, error)
	CreateGrooveJrContentFunc  func(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	UpdateGrooveJrContentFunc  func(content *models.GrooveJrContent) (*models.GrooveJrContent, error)
	DeleteGrooveJrContentFunc  func(id string) error
	GetAllAboutContentFunc     func(filter models.FilterOptions) ([]models.AboutContent, int, error)
	GetAboutContentByIDFunc    func(id string) (*models.AboutContent, error)
	CreateAboutContentFunc     func(content *models.AboutContent) (*models.AboutContent, error)
	UpdateAboutContentFunc     func(content *models.AboutContent) (*models.AboutContent, error)
	DeleteAboutContentFunc     func(id string) error

	ExportBlogPostsFunc       func() ([]models.BlogPost, error)
	ImportBlogPostsFunc       func(posts []models.BlogPost) error
	ExportWorkContentFunc     func() ([]models.WorkContent, error)
	ImportWorkContentFunc     func(content []models.WorkContent) error
	ExportGrooveJrContentFunc func() ([]models.GrooveJrContent, error)
	ImportGrooveJrContentFunc func(content []models.GrooveJrContent) error
	ExportAboutContentFunc    func() ([]models.AboutContent, error)
	ImportAboutContentFunc    func(content []models.AboutContent) error
	ExportTagsFunc            func() ([]models.Tag, error)
	ImportTagsFunc            func(tags []models.Tag) error
	ExportAuthorsFunc         func() ([]models.Author, error)
	ImportAuthorsFunc         func(authors []models.Author) error
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
func (m *MockAdminService) GetAllWorkContent(filter models.FilterOptions) ([]models.WorkContent, int, error) {
	return m.GetAllWorkContentFunc(filter)
}
func (m *MockAdminService) GetWorkContentByID(id string) (*models.WorkContent, error) {
	return m.GetWorkContentByIDFunc(id)
}
func (m *MockAdminService) CreateWorkContent(content *models.WorkContent) (*models.WorkContent, error) {
	return m.CreateWorkContentFunc(content)
}
func (m *MockAdminService) UpdateWorkContent(content *models.WorkContent) (*models.WorkContent, error) {
	return m.UpdateWorkContentFunc(content)
}
func (m *MockAdminService) DeleteWorkContent(id string) error {
	return m.DeleteWorkContentFunc(id)
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

func (m *MockAdminService) ExportBlogPosts() ([]models.BlogPost, error) {
	if m.ExportBlogPostsFunc != nil {
		return m.ExportBlogPostsFunc()
	}
	return nil, nil
}
func (m *MockAdminService) ImportBlogPosts(posts []models.BlogPost) error {
	if m.ImportBlogPostsFunc != nil {
		return m.ImportBlogPostsFunc(posts)
	}
	return nil
}
func (m *MockAdminService) ExportWorkContent() ([]models.WorkContent, error) {
	if m.ExportWorkContentFunc != nil {
		return m.ExportWorkContentFunc()
	}
	return nil, nil
}
func (m *MockAdminService) ImportWorkContent(content []models.WorkContent) error {
	if m.ImportWorkContentFunc != nil {
		return m.ImportWorkContentFunc(content)
	}
	return nil
}
func (m *MockAdminService) ExportGrooveJrContent() ([]models.GrooveJrContent, error) {
	if m.ExportGrooveJrContentFunc != nil {
		return m.ExportGrooveJrContentFunc()
	}
	return nil, nil
}
func (m *MockAdminService) ImportGrooveJrContent(content []models.GrooveJrContent) error {
	if m.ImportGrooveJrContentFunc != nil {
		return m.ImportGrooveJrContentFunc(content)
	}
	return nil
}
func (m *MockAdminService) ExportAboutContent() ([]models.AboutContent, error) {
	if m.ExportAboutContentFunc != nil {
		return m.ExportAboutContentFunc()
	}
	return nil, nil
}
func (m *MockAdminService) ImportAboutContent(content []models.AboutContent) error {
	if m.ImportAboutContentFunc != nil {
		return m.ImportAboutContentFunc(content)
	}
	return nil
}
func (m *MockAdminService) ExportTags() ([]models.Tag, error) {
	if m.ExportTagsFunc != nil {
		return m.ExportTagsFunc()
	}
	return nil, nil
}
func (m *MockAdminService) ImportTags(tags []models.Tag) error {
	if m.ImportTagsFunc != nil {
		return m.ImportTagsFunc(tags)
	}
	return nil
}
func (m *MockAdminService) ExportAuthors() ([]models.Author, error) {
	if m.ExportAuthorsFunc != nil {
		return m.ExportAuthorsFunc()
	}
	return nil, nil
}
func (m *MockAdminService) ImportAuthors(authors []models.Author) error {
	if m.ImportAuthorsFunc != nil {
		return m.ImportAuthorsFunc(authors)
	}
	return nil
}

func (m *MockAdminService) UploadImage(filename, originalName, altText string) (*models.Image, error) {
	return nil, nil
}
func (m *MockAdminService) ListImages() ([]models.Image, error) {
	return nil, nil
}
func (m *MockAdminService) DeleteImage(id string) error {
	return nil
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

	req, err := http.NewRequest("GET", "/v1/api/admin/blog", nil)
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
	testMux.HandleFunc("/v1/api/admin/blog/{id}", controller.GetBlogPostByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/v1/api/admin/blog/1", nil)
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
	req, _ = http.NewRequest("GET", "/v1/api/admin/blog/2", nil)
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
	testMux.HandleFunc("/v1/api/admin/blog/tag/{tag}", controller.GetBlogPostsByTagHandler)

	req, _ := http.NewRequest("GET", "/v1/api/admin/blog/tag/test", nil)
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
	req, _ := http.NewRequest("POST", "/v1/api/admin/blog", bytes.NewBuffer(jsonBody))
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
	testMux.HandleFunc("/v1/api/admin/blog/{id}", controller.UpdateBlogPostHandler)

	postData := models.BlogPost{ID: "1", Title: "Updated Post"}
	jsonBody, _ := json.Marshal(postData)
	req, _ := http.NewRequest("PUT", "/v1/api/admin/blog/1", bytes.NewBuffer(jsonBody))
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
	testMux.HandleFunc("/v1/api/admin/blog/{id}", controller.DeleteBlogPostHandler)

	req, _ := http.NewRequest("DELETE", "/v1/api/admin/blog/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req) // Use the testMux

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestAdminGetAllWorkContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetAllWorkContentFunc: func(filter models.FilterOptions) ([]models.WorkContent, int, error) {
			return []models.WorkContent{{ID: "1", Title: "Test Work Admin"}}, 1, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	req, _ := http.NewRequest("GET", "/v1/api/admin/work", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllWorkContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var resp models.PaginatedResponse[models.WorkContent]
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 work content, got %d", len(resp.Data))
	}
}

func TestAdminGetWorkContentByIDHandler(t *testing.T) {
	mockService := &MockAdminService{
		GetWorkContentByIDFunc: func(id string) (*models.WorkContent, error) {
			if id == "1" {
				return &models.WorkContent{ID: "1", Title: "Test Work"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/admin/work/{id}", controller.GetWorkContentByIDHandler)

	req, _ := http.NewRequest("GET", "/v1/api/admin/work/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content models.WorkContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if content.ID != "1" {
		t.Errorf("expected work content ID 1, got %s", content.ID)
	}

	req, _ = http.NewRequest("GET", "/v1/api/admin/work/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestAdminCreateWorkContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		CreateWorkContentFunc: func(content *models.WorkContent) (*models.WorkContent, error) {
			content.ID = "new-id-work"
			return content, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	contentData := models.WorkContent{Title: "New Work Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("POST", "/v1/api/admin/work", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateWorkContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var newContent models.WorkContent
	json.Unmarshal(rr.Body.Bytes(), &newContent)
	if newContent.ID != "new-id-work" {
		t.Errorf("expected new work content ID 'new-id-work', got %s", newContent.ID)
	}
}

func TestAdminUpdateWorkContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		UpdateWorkContentFunc: func(content *models.WorkContent) (*models.WorkContent, error) {
			if content.ID == "1" {
				content.Title = "Updated Work Title"
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
	testMux.HandleFunc("/v1/api/admin/work/{id}", controller.UpdateWorkContentHandler)

	contentData := models.WorkContent{ID: "1", Title: "Updated Work Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("PUT", "/v1/api/admin/work/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var updatedContent models.WorkContent
	json.Unmarshal(rr.Body.Bytes(), &updatedContent)
	if updatedContent.Title != "Updated Work Title" {
		t.Errorf("expected updated title 'Updated Work Title', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteWorkContentHandler(t *testing.T) {
	mockService := &MockAdminService{
		DeleteWorkContentFunc: func(id string) error {
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
	testMux.HandleFunc("/v1/api/admin/work/{id}", controller.DeleteWorkContentHandler)

	req, _ := http.NewRequest("DELETE", "/v1/api/admin/work/1", nil)
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

	req, _ := http.NewRequest("GET", "/v1/api/admin/groovejr", nil)
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
	testMux.HandleFunc("/v1/api/admin/groovejr/{id}", controller.GetGrooveJrContentByIDHandler)

	req, _ := http.NewRequest("GET", "/v1/api/admin/groovejr/1", nil)
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

	req, _ = http.NewRequest("GET", "/v1/api/admin/groovejr/2", nil)
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
	req, _ := http.NewRequest("POST", "/v1/api/admin/groovejr", bytes.NewBuffer(jsonBody))
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
	testMux.HandleFunc("/v1/api/admin/groovejr/{id}", controller.UpdateGrooveJrContentHandler)

	contentData := models.GrooveJrContent{ID: "1", Title: "Updated GrooveJr Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("PUT", "/v1/api/admin/groovejr/1", bytes.NewBuffer(jsonBody))
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
	testMux.HandleFunc("/v1/api/admin/groovejr/{id}", controller.DeleteGrooveJrContentHandler)

	req, _ := http.NewRequest("DELETE", "/v1/api/admin/groovejr/1", nil)
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

	req, _ := http.NewRequest("GET", "/v1/api/admin/about", nil)
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
	testMux.HandleFunc("/v1/api/admin/about/{id}", controller.GetAboutContentByIDHandler)

	req, _ := http.NewRequest("GET", "/v1/api/admin/about/1", nil)
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

	req, _ = http.NewRequest("GET", "/v1/api/admin/about/2", nil)
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
	req, _ := http.NewRequest("POST", "/v1/api/admin/about", bytes.NewBuffer(jsonBody))
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
	testMux.HandleFunc("/v1/api/admin/about/{id}", controller.UpdateAboutContentHandler)

	contentData := models.AboutContent{ID: "1", Title: "Updated About Content"}
	jsonBody, _ := json.Marshal(contentData)
	req, _ := http.NewRequest("PUT", "/v1/api/admin/about/1", bytes.NewBuffer(jsonBody))
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
	testMux.HandleFunc("/v1/api/admin/about/{id}", controller.DeleteAboutContentHandler)

	req, _ := http.NewRequest("DELETE", "/v1/api/admin/about/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestAdminExportCSVHandler(t *testing.T) {
	mockService := &MockAdminService{
		ExportBlogPostsFunc: func() ([]models.BlogPost, error) {
			return []models.BlogPost{{ID: "1", Title: "CSV Post"}}, nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/admin/csv/{entity}", controller.ExportCSVHandler)

	req, _ := http.NewRequest("GET", "/v1/api/admin/csv/blog", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr.Header().Get("Content-Type") != "text/csv" {
		t.Errorf("handler returned wrong content type: got %v want %v", rr.Header().Get("Content-Type"), "text/csv")
	}
	if rr.Header().Get("Content-Disposition") != "attachment;filename=blog.csv" {
		t.Errorf("handler returned wrong content disposition: got %v want %v", rr.Header().Get("Content-Disposition"), "attachment;filename=blog.csv")
	}
}

func TestAdminImportCSVHandler(t *testing.T) {
	mockService := &MockAdminService{
		ImportBlogPostsFunc: func(posts []models.BlogPost) error {
			if len(posts) != 1 || posts[0].Title != "Imported Post" {
				return errors.New("unexpected post data")
			}
			return nil
		},
	}
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewAdminController(&log, mockService)

	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/admin/csv/{entity}", controller.ImportCSVHandler)

	csvContent := "title,content,ordering,created_at,updated_at,activated_at,deactivated_at,tags\nImported Post,Content,1,2023-01-01T00:00:00Z,2023-01-01T00:00:00Z,,,"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.csv")
	part.Write([]byte(csvContent))
	writer.Close()

	req, _ := http.NewRequest("POST", "/v1/api/admin/csv/blog", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
