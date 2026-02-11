package models

import (
	"time"
)

// Author represents an author of a blog post.
type Author struct {
	tableName struct{} `pg:"authors"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
}

// Tag represents a tag for a blog post.
type Tag struct {
	tableName struct{} `pg:"tags"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
}

// BlogPost represents a blog post entry.
type BlogPost struct {
	tableName struct{} `pg:"blog_posts"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `pg:"author_id"`
	Author    *Author   `json:"author" pg:"rel:has-one"`
	Tags      []*Tag    `json:"tags" pg:"many2many:blog_post_tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BlogPostTag struct {
	tableName  struct{} `pg:"blog_post_tags"`
	BlogPostID string   `json:"blogPostId"`
	TagID      string   `json:"tagId"`
}

// HomeContent represents content for the home page.
type HomeContent struct {
	tableName struct{} `pg:"home_contents"`
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
}

// GrooveJrContent represents content for the groove-jr page.
type GrooveJrContent struct {
	tableName struct{} `pg:"groove_jr_contents"`
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
}

// AboutContent represents content for the about page.
type AboutContent struct {
	tableName struct{} `pg:"about_contents"`
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
}

// TrackerData represents tracking information.
type TrackerData struct {
	IP string `json:"ip"`
}
