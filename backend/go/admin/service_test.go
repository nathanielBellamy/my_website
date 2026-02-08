package admin

import (
	"errors" // Import errors package
	"fmt"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/nathanielBellamy/my_website/backend/go/testutils" // Import the new testutils package
)

func TestAdminGetAllBlogPosts(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]BlogPost); ok {
				*v = []BlogPost{{ID: "1", Title: "Test Post"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]BlogPost); ok {
					*v = []BlogPost{{ID: "1", Title: "Test Post"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	posts, err := service.GetAllBlogPosts(1, 10)
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
			if v, ok := modelDest.(*BlogPost); ok {
				*v = BlogPost{ID: "1", Title: "Test Post"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*BlogPost); ok {
					*v = BlogPost{ID: "1", Title: "Test Post"}
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
	if post.ID != "1" || post.Title != "Test Post" {
		t.Errorf("expected post with ID '1' and Title 'Test Post', got %v", post)
	}

	// Test case: Blog post not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

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
			if v, ok := modelDest.(*[]BlogPost); ok {
				*v = []BlogPost{{ID: "1", Title: "Test Tagged Post"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]BlogPost); ok {
					*v = []BlogPost{{ID: "1", Title: "Test Tagged Post"}}
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
	if posts[0].Title != "Test Tagged Post" {
		t.Errorf("expected title 'Test Tagged Post', got %s", posts[0].Title)
	}
}

func TestAdminCreateBlogPost(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		InsertFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			            if v, ok := modelDest.(*BlogPost); ok {
			                // Simulate database setting an ID
			                v.ID = "3"
			            }
			            return &testutils.MockPgResult{NumRowsAffected: 1}, nil
			        },
			    }
			    mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
			    service := NewService(mockDB)
	post := &BlogPost{Title: "New Post"} // ID will be set by the mock InsertFunc
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
			if v, ok := modelDest.(*BlogPost); ok {
				// Simulate database returning the updated object
							                *v = BlogPost{ID: v.ID, Title: v.Title} // The passed 'post' already has the updated title
							            }
							            return &testutils.MockPgResult{NumRowsAffected: 1}, nil
							        },	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)
	post := &BlogPost{ID: "1", Title: "Updated Post"}
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
	service := NewService(mockDB)

	err := service.DeleteBlogPost("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAdminGetAllHomeContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]HomeContent); ok {
				*v = []HomeContent{{ID: "1", Title: "Test Home Content"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]HomeContent); ok {
					*v = []HomeContent{{ID: "1", Title: "Test Home Content"}}
				}
			}
			return nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content, err := service.GetAllHomeContent(1, 10)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(content) != 1 {
		t.Errorf("expected 1 content, got %d", len(content))
	}
}

func TestAdminGetHomeContentByID(t *testing.T) {
	// Test case: Home content found
	foundMockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*HomeContent); ok {
				*v = HomeContent{ID: "1", Title: "Test Home Content"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*HomeContent); ok {
					*v = HomeContent{ID: "1", Title: "Test Home Content"}
				}
			}
			return nil
		},
	}
	foundMockDB := &testutils.MockPgDB{MockQuery: foundMockQuery}
	foundService := NewService(foundMockDB)

	content, err := foundService.GetHomeContentByID("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content == nil {
		t.Error("expected content, got nil")
	}
	if content.ID != "1" || content.Title != "Test Home Content" {
		t.Errorf("expected content with ID '1' and Title 'Test Home Content', got %v", content)
	}

	// Test case: Home content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

	content, err = notFoundService.GetHomeContentByID("not-found")
	fmt.Printf("TEST: TestAdminGetHomeContentByID - err from service: %v\n", err)
	if !errors.Is(err, pg.ErrNoRows) { // Use errors.Is
		t.Errorf("expected pg.ErrNoRows, got %v", err)
	}
	if content != nil {
		t.Error("expected nil content, got content")
	}
}

func TestAdminCreateHomeContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		InsertFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*HomeContent); ok {
				// Simulate database setting an ID
				v.ID = "3"
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content := &HomeContent{Title: "New Home Content"} // ID will be set by the mock InsertFunc
	newContent, err := service.CreateHomeContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if newContent.ID != "3" {
		t.Errorf("expected ID 3, got %s", newContent.ID)
	}
}

func TestAdminUpdateHomeContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		UpdateFunc: func(modelDest any, dest ...interface{}) (pg.Result, error) {
			if v, ok := modelDest.(*HomeContent); ok {
				// Simulate database returning the updated object
				*v = HomeContent{ID: v.ID, Title: v.Title} // The passed 'content' already has the updated title
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content := &HomeContent{ID: "1", Title: "Updated Home Content"}
	updatedContent, err := service.UpdateHomeContent(content)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if updatedContent.Title != "Updated Home Content" {
		t.Errorf("expected title 'Updated Home Content', got %s", updatedContent.Title)
	}
}

func TestAdminDeleteHomeContent(t *testing.T) {
	mockDB := &testutils.MockPgDB{}
	service := NewService(mockDB)

	err := service.DeleteHomeContent("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAdminGetAllGrooveJrContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]GrooveJrContent); ok {
				*v = []GrooveJrContent{{ID: "1", Title: "Test GrooveJr Content"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]GrooveJrContent); ok {
					*v = []GrooveJrContent{{ID: "1", Title: "Test GrooveJr Content"}}
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

func TestAdminGetGrooveJrContentByID(t *testing.T) {
	// Test case: GrooveJr content found
	foundMockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*GrooveJrContent); ok {
				*v = GrooveJrContent{ID: "1", Title: "Test GrooveJr Content"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*GrooveJrContent); ok {
					*v = GrooveJrContent{ID: "1", Title: "Test GrooveJr Content"}
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
	if content.ID != "1" || content.Title != "Test GrooveJr Content" {
		t.Errorf("expected content with ID '1' and Title 'Test GrooveJr Content', got %v", content)
	}

	// Test case: GrooveJr content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

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
			if v, ok := modelDest.(*GrooveJrContent); ok {
				// Simulate database setting an ID
				v.ID = "3"
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)
	content := &GrooveJrContent{Title: "New GrooveJr Content"} // ID will be set by the mock InsertFunc
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
			if v, ok := modelDest.(*GrooveJrContent); ok {
				// Simulate database returning the updated object
				*v = GrooveJrContent{ID: v.ID, Title: v.Title} // The passed 'content' already has the updated title
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)
	content := &GrooveJrContent{ID: "1", Title: "Updated GrooveJr Content"}
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
	service := NewService(mockDB)

	err := service.DeleteGrooveJrContent("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAdminGetAllAboutContent(t *testing.T) {
	mockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*[]AboutContent); ok {
				*v = []AboutContent{{ID: "1", Title: "Test About Content"}}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*[]AboutContent); ok {
					*v = []AboutContent{{ID: "1", Title: "Test About Content"}}
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

func TestAdminGetAboutContentByID(t *testing.T) {
	// Test case: About content found
	foundMockQuery := &testutils.MockPgQuery{
		SelectFunc: func(modelDest any, dest ...interface{}) error {
			if v, ok := modelDest.(*AboutContent); ok {
				*v = AboutContent{ID: "1", Title: "Test About Content"}
			} else if len(dest) > 0 {
				if v, ok := dest[0].(*AboutContent); ok {
					*v = AboutContent{ID: "1", Title: "Test About Content"}
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
	if content.ID != "1" || content.Title != "Test About Content" {
		t.Errorf("expected content with ID '1' and Title 'Test About Content', got %v", content)
	}

	// Test case: About content not found
	notFoundMockQuery := &testutils.MockPgQuery{} // SelectFunc can be nil, default behavior handles ErrNoRows
	notFoundMockDB := &testutils.MockPgDB{MockQuery: notFoundMockQuery}
	notFoundService := NewService(notFoundMockDB)

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
			if v, ok := modelDest.(*AboutContent); ok {
				// Simulate database setting an ID
				v.ID = "3"
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content := &AboutContent{Title: "New About Content"} // ID will be set by the mock InsertFunc
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
			if v, ok := modelDest.(*AboutContent); ok {
				// Simulate database returning the updated object
				*v = AboutContent{ID: v.ID, Title: v.Title} // The passed 'content' already has the updated title
			}
			return &testutils.MockPgResult{NumRowsAffected: 1}, nil
		},
	}
	mockDB := &testutils.MockPgDB{MockQuery: mockQuery}
	service := NewService(mockDB)

	content := &AboutContent{ID: "1", Title: "Updated About Content"}
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
	service := NewService(mockDB)

	err := service.DeleteAboutContent("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}