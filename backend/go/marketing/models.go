package marketing

import "time"

// BlogPost represents a blog post entry.
type BlogPost struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
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
