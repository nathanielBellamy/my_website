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
	_ = page // Dummy usage to avoid "declared and not used" error
	_ = limit // Dummy usage to avoid "declared and not used" error

	// Placeholder data
	posts := []BlogPost{
		{ID: 1, Title: "My First Blog Post", Content: "This is the content of my first blog post.", Author: "Nate", Tags: []string{"go", "backend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Title: "My Second Blog Post", Content: "This is the content of my second blog post.", Author: "Nate", Tags: []string{"angular", "frontend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	json.NewEncoder(w).Encode(posts)
}

// GetBlogPostByIDHandler handles fetching a single blog post by ID.
func (mc *MarketingController) GetBlogPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostByIDHandler Hit")
	idStr := r.PathValue("id") // Assuming mux or similar router that extracts path variables
	mc.Log.Debug().Str("idStr", idStr).Msg("GetBlogPostByIDHandler: PathValue 'id'")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		mc.Log.Error().Err(err).Str("idStr", idStr).Msg("GetBlogPostByIDHandler: Error converting idStr to int")
		http.Error(w, "Invalid blog post ID", http.StatusBadRequest)
		return
	}

	// Placeholder data
	if id == 1 {
		json.NewEncoder(w).Encode(BlogPost{ID: 1, Title: "My First Blog Post", Content: "This is the content of my first blog post.", Author: "Nate", Tags: []string{"go", "backend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	} else if id == 2 {
		json.NewEncoder(w).Encode(BlogPost{ID: 2, Title: "My Second Blog Post", Content: "This is the content of my second blog post.", Author: "Nate", Tags: []string{"angular", "frontend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	} else {
		http.Error(w, "Blog post not found", http.StatusNotFound)
	}
}

// GetBlogPostsByTagHandler handles fetching blog posts by tag.
func (mc *MarketingController) GetBlogPostsByTagHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostsByTagHandler Hit")
	tag := r.PathValue("tag") // Assuming mux or similar router
	page, limit := getPaginationParams(r)
	_ = page // Dummy usage
	_ = limit // Dummy usage
	_ = tag // Dummy usage

	// Placeholder data
	if tag == "go" {
		posts := []BlogPost{
			{ID: 1, Title: "My First Blog Post", Content: "This is the content of my first blog post.", Author: "Nate", Tags: []string{"go", "backend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		json.NewEncoder(w).Encode(posts)
	} else if tag == "angular" {
		posts := []BlogPost{
			{ID: 2, Title: "My Second Blog Post", Content: "This is the content of my second blog post.", Author: "Nate", Tags: []string{"angular", "frontend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
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
	_ = page // Dummy usage
	_ = limit // Dummy usage
	_ = dateStr // Dummy usage

	// Placeholder data
	// For simplicity, let's hardcode a date match
	if dateStr == "2026-01-23" {
		posts := []BlogPost{
			{ID: 1, Title: "My First Blog Post", Content: "This is the content of my first blog post.", Author: "Nate", Tags: []string{"go", "backend"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
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
	_ = page // Dummy usage
	_ = limit // Dummy usage

	// Placeholder data
	content := []HomeContent{
		{ID: 1, Title: "Welcome Home", Content: "This is the content for the home page."},
	}
	json.NewEncoder(w).Encode(content)
}

// GetHomeContentByIDHandler handles fetching home content by ID.
func (mc *MarketingController) GetHomeContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetHomeContentByIDHandler Hit")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid home content ID", http.StatusBadRequest)
		return
	}

	// Placeholder data
	if id == 1 {
		json.NewEncoder(w).Encode(HomeContent{ID: 1, Title: "Welcome Home", Content: "This is the content for the home page."})
	} else {
		http.Error(w, "Home content not found", http.StatusNotFound)
	}
}

// GrooveJr
// GetAllGrooveJrContentHandler handles fetching all groove-jr content.
func (mc *MarketingController) GetAllGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllGrooveJrContentHandler Hit")
	page, limit := getPaginationParams(r)
	_ = page // Dummy usage
	_ = limit // Dummy usage

	// Placeholder data
	content := []GrooveJrContent{
		{ID: 1, Title: "GrooveJr Item 1", Content: "Content for GrooveJr item 1."},
	}
	json.NewEncoder(w).Encode(content)
}

// GetGrooveJrContentByIDHandler handles fetching groove-jr content by ID.
func (mc *MarketingController) GetGrooveJrContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetGrooveJrContentByIDHandler Hit")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid GrooveJr content ID", http.StatusBadRequest)
		return
	}

	// Placeholder data
	if id == 1 {
		json.NewEncoder(w).Encode(GrooveJrContent{ID: 1, Title: "GrooveJr Item 1", Content: "Content for GrooveJr item 1."})
	} else {
		http.Error(w, "GrooveJr content not found", http.StatusNotFound)
	}
}

// About
// GetAllAboutContentHandler handles fetching all about content.
func (mc *MarketingController) GetAllAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllAboutContentHandler Hit")
	page, limit := getPaginationParams(r)
	_ = page // Dummy usage
	_ = limit // Dummy usage

	// Placeholder data
	content := []AboutContent{
		{ID: 1, Title: "About Me", Content: "This is the content for the about page."},
	}
	json.NewEncoder(w).Encode(content)
}

// GetAboutContentByIDHandler handles fetching about content by ID.
func (mc *MarketingController) GetAboutContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAboutContentByIDHandler Hit")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid about content ID", http.StatusBadRequest)
		return
	}

	// Placeholder data
	if id == 1 {
		json.NewEncoder(w).Encode(AboutContent{ID: 1, Title: "About Me", Content: "This is the content for the about page."})
	} else {
		http.Error(w, "About content not found", http.StatusNotFound)
	}
}

// Tracker
// PostTrackerDataHandler handles receiving tracker data.
func (mc *MarketingController) PostTrackerDataHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("PostTrackerDataHandler Hit")
	// For now, just acknowledge the request.
	// In a real scenario, you'd parse the request body and save tracking info.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tracker data received"))
}