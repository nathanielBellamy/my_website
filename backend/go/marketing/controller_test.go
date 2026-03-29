package marketing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/nathanielBellamy/my_website/backend/go/testutils"

	"github.com/rs/zerolog"
)

type MockMarketingService struct {
	GetAllBlogPostsFunc        func(page, limit int, tags []string) ([]models.BlogPost, error)
	GetBlogPostByIDFunc        func(id string) (*models.BlogPost, error)
	GetBlogPostsByTagFunc      func(tag string, page, limit int) ([]models.BlogPost, error)
	GetTagsFunc                func(search string, limit int) ([]models.TagWithUsage, error)
	GetAllWorkContentFunc      func(page, limit int) ([]models.WorkContent, error)
	GetWorkContentByIDFunc     func(id string) (*models.WorkContent, error)
	GetAllGrooveJrContentFunc  func(page, limit int) ([]models.GrooveJrContent, error)
	GetGrooveJrContentByIDFunc func(id string) (*models.GrooveJrContent, error)
	GetAllAboutContentFunc     func(page, limit int) ([]models.AboutContent, error)
	GetAboutContentByIDFunc    func(id string) (*models.AboutContent, error)
	GetSitemapDataFunc         func() (*SitemapData, error)
}

func (m *MockMarketingService) GetAllBlogPosts(page, limit int, tags []string) ([]models.BlogPost, error) {
	return m.GetAllBlogPostsFunc(page, limit, tags)
}
func (m *MockMarketingService) GetBlogPostByID(id string) (*models.BlogPost, error) {
	return m.GetBlogPostByIDFunc(id)
}
func (m *MockMarketingService) GetBlogPostsByTag(tag string, page, limit int) ([]models.BlogPost, error) {
	return m.GetBlogPostsByTagFunc(tag, page, limit)
}
func (m *MockMarketingService) GetTags(search string, limit int) ([]models.TagWithUsage, error) {
	return m.GetTagsFunc(search, limit)
}
func (m *MockMarketingService) GetAllWorkContent(page, limit int) ([]models.WorkContent, error) {
	return m.GetAllWorkContentFunc(page, limit)
}
func (m *MockMarketingService) GetWorkContentByID(id string) (*models.WorkContent, error) {
	return m.GetWorkContentByIDFunc(id)
}
func (m *MockMarketingService) GetAllGrooveJrContent(page, limit int) ([]models.GrooveJrContent, error) {
	return m.GetAllGrooveJrContentFunc(page, limit)
}
func (m *MockMarketingService) GetGrooveJrContentByID(id string) (*models.GrooveJrContent, error) {
	return m.GetGrooveJrContentByIDFunc(id)
}
func (m *MockMarketingService) GetAllAboutContent(page, limit int) ([]models.AboutContent, error) {
	return m.GetAllAboutContentFunc(page, limit)
}
func (m *MockMarketingService) GetAboutContentByID(id string) (*models.AboutContent, error) {
	return m.GetAboutContentByIDFunc(id)
}
func (m *MockMarketingService) GetSitemapData() (*SitemapData, error) {
	return m.GetSitemapDataFunc()
}

func TestGetAllBlogPostsHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllBlogPostsFunc: func(page, limit int, tags []string) ([]models.BlogPost, error) {
			return []models.BlogPost{{ID: "1", Title: "Test Post"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, err := http.NewRequest("GET", "/v1/api/marketing/blog", nil)
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

	var posts []models.BlogPost
	if err := json.Unmarshal(rr.Body.Bytes(), &posts); err != nil {
		t.Fatal(err)
	}
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
}

func TestGetBlogPostByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetBlogPostByIDFunc: func(id string) (*models.BlogPost, error) {
			if id == "1" {
				return &models.BlogPost{ID: "1", Title: "Test Post"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/marketing/blog/{id}", controller.GetBlogPostByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/v1/api/marketing/blog/1", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var post models.BlogPost
	json.Unmarshal(rr.Body.Bytes(), &post)
	if post.ID != "1" {
		t.Errorf("expected post ID 1, got %s", post.ID)
	}

	// Test not found
	req, _ = http.NewRequest("GET", "/v1/api/marketing/blog/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetBlogPostsByTagHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetBlogPostsByTagFunc: func(tag string, page, limit int) ([]models.BlogPost, error) {
			if tag == "test" {
				return []models.BlogPost{{ID: "1", Title: "Tagged Post"}}, nil
			}
			return []models.BlogPost{}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/marketing/blog/tag/{tag}", controller.GetBlogPostsByTagHandler)

	req, _ := http.NewRequest("GET", "/v1/api/marketing/blog/tag/test", nil)
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var posts []models.BlogPost
	json.Unmarshal(rr.Body.Bytes(), &posts)
	if len(posts) != 1 {
		t.Errorf("expected 1 tagged post, got %d", len(posts))
	}
}

func TestGetAllWorkContentHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllWorkContentFunc: func(page, limit int) ([]models.WorkContent, error) {
			return []models.WorkContent{{ID: "1", Title: "Work Content"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/v1/api/marketing/work", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllWorkContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content []models.WorkContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if len(content) != 1 {
		t.Errorf("expected 1 work content, got %d", len(content))
	}
}

func TestGetWorkContentByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetWorkContentByIDFunc: func(id string) (*models.WorkContent, error) {
			if id == "1" {
				return &models.WorkContent{ID: "1", Title: "Work Content"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/marketing/work/{id}", controller.GetWorkContentByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/v1/api/marketing/work/1", nil)
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

	// Test not found
	req, _ = http.NewRequest("GET", "/v1/api/marketing/work/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetAllGrooveJrContentHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllGrooveJrContentFunc: func(page, limit int) ([]models.GrooveJrContent, error) {
			return []models.GrooveJrContent{{ID: "1", Title: "GrooveJr Content"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/v1/api/marketing/groovejr", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllGrooveJrContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content []models.GrooveJrContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if len(content) != 1 {
		t.Errorf("expected 1 GrooveJr content, got %d", len(content))
	}
}

func TestGetGrooveJrContentByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetGrooveJrContentByIDFunc: func(id string) (*models.GrooveJrContent, error) {
			if id == "1" {
				return &models.GrooveJrContent{ID: "1", Title: "GrooveJr Content"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/marketing/groovejr/{id}", controller.GetGrooveJrContentByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/v1/api/marketing/groovejr/1", nil)
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

	// Test not found
	req, _ = http.NewRequest("GET", "/v1/api/marketing/groovejr/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetAllAboutContentHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAllAboutContentFunc: func(page, limit int) ([]models.AboutContent, error) {
			return []models.AboutContent{{ID: "1", Title: "About Content"}}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/v1/api/marketing/about", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllAboutContentHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var content []models.AboutContent
	json.Unmarshal(rr.Body.Bytes(), &content)
	if len(content) != 1 {
		t.Errorf("expected 1 About content, got %d", len(content))
	}
}

func TestGetAboutContentByIDHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetAboutContentByIDFunc: func(id string) (*models.AboutContent, error) {
			if id == "1" {
				return &models.AboutContent{ID: "1", Title: "About Content"}, nil
			}
			return nil, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	// Create a test mux to handle path parameters
	testMux := http.NewServeMux()
	testMux.HandleFunc("/v1/api/marketing/about/{id}", controller.GetAboutContentByIDHandler)

	// Test found
	req, _ := http.NewRequest("GET", "/v1/api/marketing/about/1", nil)
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

	// Test not found
	req, _ = http.NewRequest("GET", "/v1/api/marketing/about/2", nil)
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for not found: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGetMarketingFileServerNoAuth(t *testing.T) {
	// Setup temporary build directory structure
	baseDir := "build/marketing/browser"
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("build")

	// Create index.html
	indexContent := "<html>index</html>"
	if err := os.WriteFile(baseDir+"/index.html", []byte(indexContent), 0644); err != nil {
		t.Fatal(err)
	}

	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	handler := GetMarketingFileServerNoAuth(&log)

	// Test serving index.html fallback for unknown file
	req, _ := http.NewRequest("GET", "/unknown", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for fallback: got %v want %v", status, http.StatusOK)
	}
	if rr.Body.String() != indexContent {
		t.Errorf("handler returned wrong content for fallback: got %v want %v", rr.Body.String(), indexContent)
	}
}

func TestSitemapHandler(t *testing.T) {
	mockService := &MockMarketingService{
		GetSitemapDataFunc: func() (*SitemapData, error) {
			now := time.Now()
			return &SitemapData{
				BlogPosts: []models.BlogPost{
					{ID: "post-1", UpdatedAt: now},
					{ID: "post-2", UpdatedAt: now},
				},
			}, nil
		},
	}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/sitemap.xml", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SitemapHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/xml" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/xml")
	}

	body := rr.Body.String()
	if !strings.Contains(body, "<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">") {
		t.Errorf("sitemap missing urlset tag")
	}
	if !strings.Contains(body, "<loc>https://nateschieber.dev/blog/post1</loc>") {
		t.Errorf("sitemap missing post-1 loc")
	}
	if !strings.Contains(body, "<loc>https://nateschieber.dev/blog/post2</loc>") {
		t.Errorf("sitemap missing post-2 loc")
	}
	if !strings.Contains(body, "<loc>https://nateschieber.dev/about</loc>") {
		t.Errorf("sitemap missing static page loc")
	}
}

func TestRobotsTxtHandler(t *testing.T) {
	mockService := &MockMarketingService{}
	mockLogOutput := &testutils.MockLogger{}
	log := zerolog.New(mockLogOutput)
	controller := NewMarketingController(&log, mockService)

	req, _ := http.NewRequest("GET", "/robots.txt", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.RobotsTxtHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "text/plain" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/plain")
	}

	body := rr.Body.String()
	if !strings.Contains(body, "User-agent: *") {
		t.Errorf("robots.txt missing User-agent: *")
	}
	if !strings.Contains(body, "Allow: /") {
		t.Errorf("robots.txt missing Allow: /")
	}
	if !strings.Contains(body, "Sitemap: https://nateschieber.dev/sitemap.xml") {
		t.Errorf("robots.txt missing Sitemap link")
	}
}
