package marketing

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
)

// MarketingController holds dependencies for marketing-related handlers.
type MarketingController struct {
	Log *zerolog.Logger
}

// NewMarketingController creates and returns a new MarketingController.
func NewMarketingController(log *zerolog.Logger) *MarketingController {
	return &MarketingController{
		Log: log,
	}
}

// getPaginationParams extracts page and limit from request queries.
func getPaginationParams(r *http.Request) (int, int) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	return page, limit
}

// Blog
// GetAllBlogPostsHandler handles fetching all blog posts.
func (mc *MarketingController) GetAllBlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllBlogPostsHandler Hit")
	page, limit := getPaginationParams(r)

	// Placeholder data
	posts := make([]BlogPost, 0)
	for i := 1; i <= 25; i++ {
		posts = append(posts, BlogPost{
			ID:        "blog-post-id-" + strconv.Itoa(i),
			Title:     "Blog Post " + strconv.Itoa(i),
			Content:   "This is the content of blog post " + strconv.Itoa(i) + ".",
			Author:    "Nate",
			Tags:      []string{"go", "backend", "testing"},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	start := (page - 1) * limit
	if start > len(posts) {
		start = len(posts)
	}
	end := start + limit
	if end > len(posts) {
		end = len(posts)
	}

	paginatedPosts := posts[start:end]

	json.NewEncoder(w).Encode(paginatedPosts)
}

// GetBlogPostByIDHandler handles fetching a single blog post by ID.
func (mc *MarketingController) GetBlogPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostByIDHandler Hit")
	id := r.PathValue("id") // Assuming mux or similar router that extracts path variables
	mc.Log.Debug().Str("idStr", id).Msg("GetBlogPostByIDHandler: PathValue 'id'")

	// Placeholder data
	if id == "blog-id-1" {
		json.NewEncoder(w).Encode(BlogPost{ID: "blog-id-1", Title: "My First Blog Post", Content: "This is the content of my first blog post.", Author: "Nate", Tags: []string{"go", "backend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	} else if id == "blog-id-2" {
		json.NewEncoder(w).Encode(BlogPost{ID: "blog-id-2", Title: "My Second Blog Post", Content: "This is the content of my second blog post.", Author: "Nate", Tags: []string{"angular", "frontend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	} else {
		http.Error(w, "Blog post not found", http.StatusNotFound)
	}
}

// GetBlogPostsByTagHandler handles fetching blog posts by tag.
func (mc *MarketingController) GetBlogPostsByTagHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostsByTagHandler Hit")
	tag := r.PathValue("tag") // Assuming mux or similar router
	page, limit := getPaginationParams(r)
	_ = page  // Dummy usage
	_ = limit // Dummy usage
	_ = tag   // Dummy usage

	// Placeholder data
	if tag == "go" {
		posts := []BlogPost{
			{ID: "blog-id-1", Title: "My First Blog Post", Content: "This is the content of my first blog post.", Author: "Nate", Tags: []string{"go", "backend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		json.NewEncoder(w).Encode(posts)
	} else if tag == "angular" {
		posts := []BlogPost{
			{ID: "blog-id-2", Title: "My Second Blog Post", Content: "This is the content of my second blog post.", Author: "Nate", Tags: []string{"angular", "frontend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		json.NewEncoder(w).Encode(posts)
	} else {
		json.NewEncoder(w).Encode([]BlogPost{}) // No posts for this tag
	}
}

// GetBlogPostsByDateHandler handles fetching blog posts by date.
func (mc *MarketingController) GetBlogPostsByDateHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostsByDateHandler Hit")
	dateStr := r.PathValue("date") // Assuming mux or similar router
	page, limit := getPaginationParams(r)
	_ = page    // Dummy usage
	_ = limit   // Dummy usage
	_ = dateStr // Dummy usage

	// Placeholder data
	// For simplicity, let's hardcode a date match
	if dateStr == "2026-01-23" {
		posts := []BlogPost{
			{ID: "blog-id-1", Title: "My First Blog Post", Content: "This is the content of my first blog post.", Author: "Nate", Tags: []string{"go", "backend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		json.NewEncoder(w).Encode(posts)
	} else {
		json.NewEncoder(w).Encode([]BlogPost{}) // No posts for this date
	}
}

// Home
// GetAllHomeContentHandler handles fetching all home content.
func (mc *MarketingController) GetAllHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllHomeContentHandler Hit")
	page, limit := getPaginationParams(r)

	// Placeholder data
	content := make([]HomeContent, 0)
	for i := 1; i <= 25; i++ {
		content = append(content, HomeContent{
			ID:      "home-id-" + strconv.Itoa(i),
			Title:   "Home Content " + strconv.Itoa(i),
			Content: "This is home content " + strconv.Itoa(i) + ".",
		})
	}

	start := (page - 1) * limit
	if start > len(content) {
		start = len(content)
	}
	end := start + limit
	if end > len(content) {
		end = len(content)
	}

	paginatedContent := content[start:end]

	json.NewEncoder(w).Encode(paginatedContent)
}

// GetHomeContentByIDHandler handles fetching home content by ID.
func (mc *MarketingController) GetHomeContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetHomeContentByIDHandler Hit")
	id := r.PathValue("id")

	// Placeholder data
	if id == "home-id-1" {
		json.NewEncoder(w).Encode(HomeContent{ID: "home-id-1", Title: "Welcome Home", Content: "This is the content for the home page."})
	} else {
		http.Error(w, "Home content not found", http.StatusNotFound)
	}
}

// GrooveJr
// GetAllGrooveJrContentHandler handles fetching all groove-jr content.
func (mc *MarketingController) GetAllGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllGrooveJrContentHandler Hit")
	page, limit := getPaginationParams(r)

	// Placeholder data
	content := make([]GrooveJrContent, 0)
	for i := 1; i <= 25; i++ {
		content = append(content, GrooveJrContent{
			ID:      "gj-content-" + strconv.Itoa(i),
			Title:   "GrooveJr Content " + strconv.Itoa(i),
			Content: "This is GrooveJr content " + strconv.Itoa(i) + ".",
		})
	}

	start := (page - 1) * limit
	if start > len(content) {
		start = len(content)
	}
	end := start + limit
	if end > len(content) {
		end = len(content)
	}

	paginatedContent := content[start:end]

	json.NewEncoder(w).Encode(paginatedContent)
}

// GetGrooveJrContentByIDHandler handles fetching groove-jr content by ID.
func (mc *MarketingController) GetGrooveJrContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetGrooveJrContentByIDHandler Hit")
	id := r.PathValue("id")

	// Placeholder data
	if id == "gj-id-1" {
		json.NewEncoder(w).Encode(GrooveJrContent{ID: "gj-id-1", Title: "GrooveJr Item 1", Content: "Content for GrooveJr item 1."})
	} else {
		http.Error(w, "GrooveJr content not found", http.StatusNotFound)
	}
}

// About
// GetAllAboutContentHandler handles fetching all about content.
func (mc *MarketingController) GetAllAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllAboutContentHandler Hit")
	page, limit := getPaginationParams(r)

	// Placeholder data
	content := make([]AboutContent, 0)
	for i := 1; i <= 25; i++ {
		content = append(content, AboutContent{
			ID:      "about-id-" + strconv.Itoa(i),
			Title:   "About Content " + strconv.Itoa(i),
			Content: "This is about content " + strconv.Itoa(i) + ".",
		})
	}

	start := (page - 1) * limit
	if start > len(content) {
		start = len(content)
	}
	end := start + limit
	if end > len(content) {
		end = len(content)
	}

	paginatedContent := content[start:end]

	json.NewEncoder(w).Encode(paginatedContent)
}

// GetAboutContentByIDHandler handles fetching about content by ID.
func (mc *MarketingController) GetAboutContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAboutContentByIDHandler Hit")
	id := r.PathValue("id")

	// Placeholder data
	if id == "about-id-1" {
		json.NewEncoder(w).Encode(AboutContent{ID: "about-id-1", Title: "About Me", Content: "This is the content for the about page."})
	} else {
		http.Error(w, "About content not found", http.StatusNotFound)
	}
}
