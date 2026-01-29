package marketing

import (
	"time"
	// "github.com/go-pg/pg/v10" // Removed unused import
)

// PgxQuerySeter defines the query operations required for building database queries.
type PgxQuerySeter interface {
	Relation(name string) PgxQuerySeter
	Limit(count int) PgxQuerySeter
	Offset(offset int) PgxQuerySeter
	Where(query string, params ...interface{}) PgxQuerySeter
	Join(join string, params ...interface{}) PgxQuerySeter
	Select(dest ...interface{}) error
}

// PgxDB defines the database operations required by the MarketingController.
type PgxDB interface {
	Model(model ...interface{}) PgxQuerySeter // Changed return type back to PgxQuerySeter
}

// Author represents an author of a blog post.
type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Tag represents a tag for a blog post.
type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BlogPost represents a blog post entry.
type BlogPost struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    *Author   `json:"author"`
	Tags      []*Tag    `json:"tags" pg:"many2many:blog_post_tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// HomeContent represents content for the home page.
type HomeContent struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GrooveJrContent represents content for the groove-jr page.
type GrooveJrContent struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// AboutContent represents content for the about page.
type AboutContent struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// TrackerData represents tracking information.
type TrackerData struct {
	IP string `json:"ip"`
}
