package marketing

import "time"

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
