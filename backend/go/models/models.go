package models

import (
	"time"
)

// Author represents an author of a blog post.
type Author struct {
	tableName     struct{}   `pg:"authors"`
	ID            string     `json:"id" pg:",pk,type:uuid"`
	Name          string     `json:"name"`
	ActivatedAt   *time.Time `json:"activatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
}

// Tag represents a tag for a blog post.
type Tag struct {
	tableName     struct{}   `pg:"tags"`
	ID            string     `json:"id" pg:",pk,type:uuid"`
	Name          string     `json:"name"`
	ActivatedAt   *time.Time `json:"activatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
}

type TagWithUsage struct {
	tableName struct{} `pg:"tags,alias:tag"`
	Tag
	UsageCount int `json:"usageCount"`
}

// BlogPost represents a blog post entry.
type BlogPost struct {
	tableName     struct{}   `pg:"blog_posts"`
	ID            string     `json:"id" pg:",pk,type:uuid"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	AuthorID      string     `pg:"author_id,type:uuid"`
	Author        *Author    `json:"author" pg:"rel:has-one"`
	Tags          []*Tag     `json:"tags" pg:"many2many:blog_post_tags"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	ActivatedAt   *time.Time `json:"activatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
	Ordering      int        `json:"order" pg:"ordering"`
}

type BlogPostTag struct {
	tableName  struct{} `pg:"blog_post_tags"`
	BlogPostID string   `json:"blogPostId" pg:",pk,type:uuid"`
	TagID      string   `json:"tagId" pg:",pk,type:uuid"`
}

// HomeContent represents content for the home page.
type HomeContent struct {
	tableName     struct{}   `pg:"home_contents"`
	ID            string     `json:"id" pg:",pk,type:uuid"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	ActivatedAt   *time.Time `json:"activatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
	Ordering      int        `json:"order" pg:"ordering"`
}

// GrooveJrContent represents content for the groove-jr page.
type GrooveJrContent struct {
	tableName     struct{}   `pg:"groove_jr_contents"`
	ID            string     `json:"id" pg:",pk,type:uuid"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	ActivatedAt   *time.Time `json:"activatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
	Ordering      int        `json:"order" pg:"ordering"`
}

// AboutContent represents content for the about page.
type AboutContent struct {
	tableName     struct{}   `pg:"about_contents"`
	ID            string     `json:"id" pg:",pk,type:uuid"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	ActivatedAt   *time.Time `json:"activatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
	Ordering      int        `json:"order" pg:"ordering"`
}

// Image represents an uploaded image metadata.
type Image struct {
	tableName    struct{}  `pg:"images"`
	ID           string    `json:"id" pg:",pk,type:uuid"`
	Filename     string    `json:"filename" pg:",unique"`
	OriginalName string    `json:"originalName"`
	AltText      string    `json:"altText"`
	CreatedAt    time.Time `json:"createdAt"`
}

// TrackerData represents tracking information.
type TrackerData struct {
	IP string `json:"ip"`
}
