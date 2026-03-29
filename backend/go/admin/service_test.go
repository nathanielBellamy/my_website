package admin

import (
	"errors" // Import errors package
	"fmt"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/rs/zerolog"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/testutils" // Import the new testutils package
)

func TestAdminGetAllBlogPosts(t *testing.T) {
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
	service := NewService(mockDB, &zerolog.Logger{})

	posts, _, err := service.GetAllBlogPosts(models.FilterOptions{Page: 1, Limit: 10})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
}

func TestAdminGetBlogPostByID(t *testing.T) {
	// Test case: Blog post found
	foundMockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.BlogPost); ok {
				*v = models.BlogPost{ID: "1", Title: "Test Post"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.BlogPost); ok {
					*v = models.BlogPost{ID: "1", Title: "Test Post"}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB, &zerolog.Logger{})

	post, err := foundService.GetBlogPostByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if post == nil {
		t.Error("expected a post, got nil")
	}
	if post.ID != "1" || post.Title != "Test Post" {
		t.Errorf("expected post with ID '1' and Title 'Test Post', got %v", post)
	}

	// Test case: Blog post not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB, &zerolog.Logger{})

	post, err = notFoundService.GetBlogPostByID("not-found")
	fmt.Printf("TEST: TestAdminGetBlogPostByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if post != nil {
		t.Error("expected nil post, got a post")
	}
}

func TestAdminGetBlogPostsByTag(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.BlogPost); ok {
				*v = []models.BlogPost{{ID: "1", Title: "Test Tagged Post"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.BlogPost); ok {
					*v = []models.BlogPost{{ID: "1", Title: "Test Tagged Post"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	posts, err := service.GetBlogPostsByTag("test-tag", 1, 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
	if posts[0].Title != "Test Tagged Post" {
		t.Errorf("expected title 'Test Tagged Post', got %s", posts[0].Title)
	}
}

func TestAdminCreateBlogPost(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		InsertFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.BlogPost); ok {
				// Simulate database setting an ID
				v.ID = "3"
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})
	post := &models.BlogPost{Title: "New Post"} // ID will be set by the mock InsertFunc
	newPost, err := service.CreateBlogPost(post)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if newPost.ID != "3" {
		t.Errorf("expected ID 3, got %s", newPost.ID)
	}
}

func TestAdminUpdateBlogPost(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		UpdateFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.BlogPost); ok {
				// Simulate database returning the updated object
				*v = models.BlogPost{ID: v.ID, Title: v.Title} // The passed 'post' already has the updated title
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		}}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})
	post := &models.BlogPost{ID: "1", Title: "Updated Post"}
	updatedPost, err := service.UpdateBlogPost(post)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if updatedPost.Title != "Updated Post" {
		t.Errorf("expected title 'Updated Post', got %s", updatedPost.Title)
	}
}

func TestAdminDeleteBlogPost(t *testing.T) {
	mockDB := &testutils.MockPgDB{}
	service := NewService(mockDB, &zerolog.Logger{})

	err := service.DeleteBlogPost("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAdminGetAllWorkContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.WorkContent); ok {
				*v = []models.WorkContent{{ID: "1", Title: "Test Work Content"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.WorkContent); ok {
					*v = []models.WorkContent{{ID: "1", Title: "Test Work Content"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	content, _, err := service.GetAllWorkContent(models.FilterOptions{Page: 1, Limit: 10})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(content) != 1 {
		t.Errorf("expected 1 content, got %d", len(content))
	}
}

func TestAdminGetWorkContentByID(t *testing.T) {
	// Test case: Work content found
	foundMockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.WorkContent); ok {
				*v = models.WorkContent{ID: "1", Title: "Test Work Content"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.WorkContent); ok {
					*v = models.WorkContent{ID: "1", Title: "Test Work Content"}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB, &zerolog.Logger{})

	content, err := foundService.GetWorkContentByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content == nil {
		t.Error("expected content, got nil")
	}
	if content.ID != "1" || content.Title != "Test Work Content" {
		t.Errorf("expected content with ID '1' and Title 'Test Work Content', got %v", content)
	}

	// Test case: Work content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB, &zerolog.Logger{})

	content, err = notFoundService.GetWorkContentByID("not-found")
	fmt.Printf("TEST: TestAdminGetWorkContentByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if content != nil {
		t.Error("expected nil content, got content")
	}
}

func TestAdminCreateWorkContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		InsertFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.WorkContent); ok {
				// Simulate database setting an ID
				v.ID = "3"
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	content := &models.WorkContent{Title: "New Work Content"} // ID will be set by the mock InsertFunc
	newContent, err := service.CreateWorkContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if newContent.ID != "3" {
		t.Errorf("expected ID 3, got %s", newContent.ID)
	}
}

func TestAdminUpdateWorkContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		UpdateFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.WorkContent); ok {
				// Simulate database returning the updated object
				*v = models.WorkContent{ID: v.ID, Title: v.Title} // The passed 'content' already has the updated title
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	content := &models.WorkContent{ID: "1", Title: "Updated Work Content"}
	updatedContent, err := service.UpdateWorkContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if updatedContent.Title != "Updated Work Content" {
		t.Errorf("expected title 'Updated Work Content', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteWorkContent(t *testing.T) {
	mockDB := &testutils.MockPgDB{}
	service := NewService(mockDB, &zerolog.Logger{})

	err := service.DeleteWorkContent("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAdminGetAllGrooveJrContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.GrooveJrContent); ok {
				*v = []models.GrooveJrContent{{ID: "1", Title: "Test GrooveJr Content"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.GrooveJrContent); ok {
					*v = []models.GrooveJrContent{{ID: "1", Title: "Test GrooveJr Content"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	content, _, err := service.GetAllGrooveJrContent(models.FilterOptions{Page: 1, Limit: 10})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(content) != 1 {
		t.Errorf("expected 1 content, got %d", len(content))
	}
}

func TestAdminGetGrooveJrContentByID(t *testing.T) {
	// Test case: GrooveJr content found
	foundMockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.GrooveJrContent); ok {
				*v = models.GrooveJrContent{ID: "1", Title: "Test GrooveJr Content"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.GrooveJrContent); ok {
					*v = models.GrooveJrContent{ID: "1", Title: "Test GrooveJr Content"}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB, &zerolog.Logger{})

	content, err := foundService.GetGrooveJrContentByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content == nil {
		t.Error("expected content, got nil")
	}
	if content.ID != "1" || content.Title != "Test GrooveJr Content" {
		t.Errorf("expected content with ID '1' and Title 'Test GrooveJr Content', got %v", content)
	}

	// Test case: GrooveJr content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB, &zerolog.Logger{})

	content, err = notFoundService.GetGrooveJrContentByID("not-found")
	fmt.Printf("TEST: TestAdminGetGrooveJrContentByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if content != nil {
		t.Error("expected nil content, got content")
	}
}

func TestAdminCreateGrooveJrContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		InsertFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.GrooveJrContent); ok {
				// Simulate database setting an ID
				v.ID = "3"
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})
	content := &models.GrooveJrContent{Title: "New GrooveJr Content"} // ID will be set by the mock InsertFunc
	newContent, err := service.CreateGrooveJrContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if newContent.ID != "3" {
		t.Errorf("expected ID 3, got %s", newContent.ID)
	}
}

func TestAdminUpdateGrooveJrContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		UpdateFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.GrooveJrContent); ok {
				// Simulate database returning the updated object
				*v = models.GrooveJrContent{ID: v.ID, Title: v.Title} // The passed 'content' already has the updated title
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})
	content := &models.GrooveJrContent{ID: "1", Title: "Updated GrooveJr Content"}
	updatedContent, err := service.UpdateGrooveJrContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if updatedContent.Title != "Updated GrooveJr Content" {
		t.Errorf("expected title 'Updated GrooveJr Content', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteGrooveJrContent(t *testing.T) {
	mockDB := &testutils.MockPgDB{}
	service := NewService(mockDB, &zerolog.Logger{})

	err := service.DeleteGrooveJrContent("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAdminGetAllAboutContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.AboutContent); ok {
				*v = []models.AboutContent{{ID: "1", Title: "Test About Content"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]models.AboutContent); ok {
					*v = []models.AboutContent{{ID: "1", Title: "Test About Content"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	content, _, err := service.GetAllAboutContent(models.FilterOptions{Page: 1, Limit: 10})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(content) != 1 {
		t.Errorf("expected 1 content, got %d", len(content))
	}
}

func TestAdminGetAboutContentByID(t *testing.T) {
	// Test case: About content found
	foundMockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*models.AboutContent); ok {
				*v = models.AboutContent{ID: "1", Title: "Test About Content"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*models.AboutContent); ok {
					*v = models.AboutContent{ID: "1", Title: "Test About Content"}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB, &zerolog.Logger{})

	content, err := foundService.GetAboutContentByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content == nil {
		t.Error("expected content, got nil")
	}
	if content.ID != "1" || content.Title != "Test About Content" {
		t.Errorf("expected content with ID '1' and Title 'Test About Content', got %v", content)
	}

	// Test case: About content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB, &zerolog.Logger{})

	content, err = notFoundService.GetAboutContentByID("not-found")
	fmt.Printf("TEST: TestAdminGetAboutContentByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if content != nil {
		t.Error("expected nil content, got content")
	}
}

func TestAdminCreateAboutContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		InsertFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.AboutContent); ok {
				// Simulate database setting an ID
				v.ID = "3"
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	content := &models.AboutContent{Title: "New About Content"} // ID will be set by the mock InsertFunc
	newContent, err := service.CreateAboutContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if newContent.ID != "3" {
		t.Errorf("expected ID 3, got %s", newContent.ID)
	}
}

func TestAdminUpdateAboutContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		UpdateFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*models.AboutContent); ok {
				// Simulate database returning the updated object
				*v = models.AboutContent{ID: v.ID, Title: v.Title} // The passed 'content' already has the updated title
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	content := &models.AboutContent{ID: "1", Title: "Updated About Content"}
	updatedContent, err := service.UpdateAboutContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if updatedContent.Title != "Updated About Content" {
		t.Errorf("expected title 'Updated About Content', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteAboutContent(t *testing.T) {
	mockDB := &testutils.MockPgDB{}
	service := NewService(mockDB, &zerolog.Logger{})

	err := service.DeleteAboutContent("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAdminExportBlogPosts(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]models.BlogPost); ok {
				*v = []models.BlogPost{
					{ID: "1", Title: "Post 1", Ordering: 1},
					{ID: "2", Title: "Post 2", Ordering: 2},
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	posts, err := service.ExportBlogPosts()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(posts) != 2 {
		t.Errorf("expected 2 posts, got %d", len(posts))
	}
}

func TestAdminImportBlogPosts(t *testing.T) {
	insertCount := 0
	mockQuery := &testutils.MockPgQuery{
		InsertFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			insertCount++
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			// Simulate finding existing author/tags if searched by name
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB, &zerolog.Logger{})

	posts := []models.BlogPost{
		{
			ID:     "1",
			Title:  "Imported Post",
			Author: &models.Author{Name: "New Author"},
			Tags:   []*models.Tag{{Name: "Tag1"}},
		},
	}

	err := service.ImportBlogPosts(posts)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	// Expected inserts:
	// 1. Author (if ID not present and not found - logic is complex, assumes found if select returns nil err for now)
	// Actually logic says: if Author.ID == "", search by name. If found (err==nil), use it. Else Insert.
	// My mock SelectFunc returns nil, so it finds author.
	// 2. Tags: same logic. Finds Tag1.
	// 3. Post: Upsert (Insert).
	// 4. PostTags: Insert.

	// With SelectFunc returning nil (found):
	// Author: No insert.
	// Tag: No insert.
	// Post: Insert (Upsert).
	// PostTags: Insert.
	// Total 2 inserts?
	// Wait, code deletes PostTags then Inserts new ones. Delete is mocked too.

	if insertCount == 0 {
		t.Error("expected inserts to happen")
	}
}
