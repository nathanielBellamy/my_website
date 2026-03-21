package marketing

import (
	"errors" // Import errors package
	"fmt"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/testutils" // Import the new testutils package
)

func TestGetAllBlogPosts(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.BlogPost); ok {
				*v = []models.BlogPost{{ID: "1", Title: "Test Post"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.BlogPost); ok {
					*v = []models.BlogPost{{ID: "1", Title: "Test Post"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	posts, err := service.GetAllBlogPosts(1, 10, nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
}

func TestGetBlogPostByID(t *testing.T) {
	// Test case: Blog post found
	var foundMockQuery *testutils.MockPgQuery // Declare first
	foundMockQuery = &testutils.MockPgQuery{  // Initialize
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.BlogPost); ok {
				*v = models.BlogPost{ID: foundMockQuery.WhereID, Title: "Test Post " + foundMockQuery.WhereID}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.BlogPost); ok {
					*v = models.BlogPost{ID: foundMockQuery.WhereID, Title: "Test Post " + foundMockQuery.WhereID}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB)

	post, err := foundService.GetBlogPostByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if post == nil {
		t.Error("expected a post, got nil")
	}
	if post.ID != "1" || post.Title != "Test Post 1" {
		t.Errorf("expected post with ID '1' and Title 'Test Post 1', got %v", post)
	}

	// Test case: Blog post not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

	post, err = notFoundService.GetBlogPostByID("not-found")
	fmt.Printf("TEST: TestGetBlogPostByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if post != nil {
		t.Error("expected nil post, got a post")
	}
}

func TestGetBlogPostsByTag(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.BlogPost); ok {
				*v = []models.BlogPost{{ID: "1", Title: "Tagged Post"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.BlogPost); ok {
					*v = []models.BlogPost{{ID: "1", Title: "Tagged Post"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	posts, err := service.GetBlogPostsByTag("test-tag", 1, 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
}

func TestGetAllWorkContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.WorkContent); ok {
				*v = []models.WorkContent{{ID: "1", Title: "Test Work"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.WorkContent); ok {
					*v = []models.WorkContent{{ID: "1", Title: "Test Work"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content, err := service.GetAllWorkContent(1, 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(content) != 1 {
		t.Errorf("expected 1 content, got %d", len(content))
	}
}

func TestGetWorkContentByID(t *testing.T) {
	// Test case: Home content found
	var foundMockQuery *testutils.MockPgQuery // Declare first
	foundMockQuery = &testutils.MockPgQuery{  // Initialize
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.WorkContent); ok {
				*v = models.WorkContent{ID: foundMockQuery.WhereID, Title: "Test Work " + foundMockQuery.WhereID}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.WorkContent); ok {
					*v = models.WorkContent{ID: foundMockQuery.WhereID, Title: "Test Work " + foundMockQuery.WhereID}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB)

	content, err := foundService.GetWorkContentByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content == nil {
		t.Error("expected content, got nil")
	}
	if content.ID != "1" || content.Title != "Test Work 1" {
		t.Errorf("expected content with ID '1' and Title 'Test Work 1', got %v", content)
	}

	// Test case: Home content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

	content, err = notFoundService.GetWorkContentByID("not-found")
	fmt.Printf("TEST: TestGetWorkContentByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if content != nil {
		t.Error("expected nil content, got content")
	}
}

func TestGetAllGrooveJrContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.GrooveJrContent); ok {
				*v = []models.GrooveJrContent{{ID: "1", Title: "Test GrooveJr"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.GrooveJrContent); ok {
					*v = []models.GrooveJrContent{{ID: "1", Title: "Test GrooveJr"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content, err := service.GetAllGrooveJrContent(1, 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(content) != 1 {
		t.Errorf("expected 1 content, got %d", len(content))
	}
}

func TestGetGrooveJrContentByID(t *testing.T) {
	// Test case: GrooveJr content found
	var foundMockQuery *testutils.MockPgQuery // Declare first
	foundMockQuery = &testutils.MockPgQuery{  // Initialize
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.GrooveJrContent); ok {
				*v = models.GrooveJrContent{ID: foundMockQuery.WhereID, Title: "Test GrooveJr " + foundMockQuery.WhereID}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.GrooveJrContent); ok {
					*v = models.GrooveJrContent{ID: foundMockQuery.WhereID, Title: "Test GrooveJr " + foundMockQuery.WhereID}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB)

	content, err := foundService.GetGrooveJrContentByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content == nil {
		t.Error("expected content, got nil")
	}
	if content.ID != "1" || content.Title != "Test GrooveJr 1" {
		t.Errorf("expected content with ID '1' and Title 'Test GrooveJr 1', got %v", content)
	}

	// Test case: GrooveJr content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

	content, err = notFoundService.GetGrooveJrContentByID("not-found")
	fmt.Printf("TEST: TestGetGrooveJrContentByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if content != nil {
		t.Error("expected nil content, got content")
	}
}

func TestGetAllAboutContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.AboutContent); ok {
				*v = []models.AboutContent{{ID: "1", Title: "Test About"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.AboutContent); ok {
					*v = []models.AboutContent{{ID: "1", Title: "Test About"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content, err := service.GetAllAboutContent(1, 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(content) != 1 {
		t.Errorf("expected 1 content, got %d", len(content))
	}
}

func TestGetAboutContentByID(t *testing.T) {
	// Test case: About content found
	var foundMockQuery *testutils.MockPgQuery // Declare first
	foundMockQuery = &testutils.MockPgQuery{  // Initialize
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.AboutContent); ok {
				*v = models.AboutContent{ID: foundMockQuery.WhereID, Title: "Test About " + foundMockQuery.WhereID}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.AboutContent); ok {
					*v = models.AboutContent{ID: foundMockQuery.WhereID, Title: "Test About " + foundMockQuery.WhereID}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB)

	content, err := foundService.GetAboutContentByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content == nil {
		t.Error("expected content, got nil")
	}
	if content.ID != "1" || content.Title != "Test About 1" {
		t.Errorf("expected content with ID '1' and Title 'Test About 1', got %v", content)
	}

	// Test case: About content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

	content, err = notFoundService.GetAboutContentByID("not-found")
	fmt.Printf("TEST: TestGetAboutContentByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if content != nil {
		t.Error("expected nil content, got content")
	}
}
