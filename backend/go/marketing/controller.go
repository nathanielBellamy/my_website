package marketing

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v10" // Needed for PgxDB interface's Model method return type
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
	_ "github.com/lib/pq"
)

// MarketingController holds dependencies for marketing-related handlers.
type MarketingController struct {
	Log *zerolog.Logger
	DB  PgxDB // Use the interface here
}

// NewMarketingController creates and returns a new MarketingController.
func NewMarketingController(log *zerolog.Logger, db PgxDB) *MarketingController { // Use the interface here
	return &MarketingController{
		Log: log,
		DB:  db,
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

	var posts []BlogPost
	err := mc.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Limit(limit).
		Offset((page - 1) * limit).
		Select()

	if err != nil {
		http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

// GetBlogPostByIDHandler handles fetching a single blog post by ID.
func (mc *MarketingController) GetBlogPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostByIDHandler Hit")
	id := r.PathValue("id") // Assuming mux or similar router that extracts path variables
	mc.Log.Debug().Str("idStr", id).Msg("GetBlogPostByIDHandler: PathValue 'id'")

	var post BlogPost
	err := mc.DB.Model(&post).
		Where("id = ?", id).
		Relation("Author").
		Relation("Tags").
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			http.Error(w, "Blog post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching blog post", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(post)
}

// GetBlogPostsByTagHandler handles fetching blog posts by tag.
func (mc *MarketingController) GetBlogPostsByTagHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostsByTagHandler Hit")
	tag := r.PathValue("tag") // Assuming mux or similar router
	page, limit := getPaginationParams(r)

	var posts []BlogPost
	err := mc.DB.Model(&posts).
		Relation("Author").
		Relation("Tags").
		Join("JOIN blog_post_tags AS bpt ON bpt.blog_post_id = blog_post.id").
		Join("JOIN tags AS t ON t.id = bpt.tag_id").
		Where("t.name = ?", tag).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()

	if err != nil {
		http.Error(w, "Error fetching blog posts by tag", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

// Home
// GetAllHomeContentHandler handles fetching all home content.
func (mc *MarketingController) GetAllHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllHomeContentHandler Hit")
	page, limit := getPaginationParams(r)

	var content []HomeContent
	err := mc.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()

	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching home content")
		http.Error(w, "Error fetching home content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(content)
}

// GetHomeContentByIDHandler handles fetching home content by ID.
func (mc *MarketingController) GetHomeContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetHomeContentByIDHandler Hit")
	id := r.PathValue("id")

	var content HomeContent
	err := mc.DB.Model(&content).
		Where("id = ?", id).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			http.Error(w, "Home content not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching home content", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(content)
}

// GrooveJr
// GetAllGrooveJrContentHandler handles fetching all groove-jr content.
func (mc *MarketingController) GetAllGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllGrooveJrContentHandler Hit")
	page, limit := getPaginationParams(r)

	var content []GrooveJrContent
	err := mc.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()

	if err != nil {
		http.Error(w, "Error fetching groove-jr content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(content)
}

// GetGrooveJrContentByIDHandler handles fetching groove-jr content by ID.
func (mc *MarketingController) GetGrooveJrContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetGrooveJrContentByIDHandler Hit")
	id := r.PathValue("id")

	var content GrooveJrContent
	err := mc.DB.Model(&content).
		Where("id = ?", id).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			http.Error(w, "GrooveJr content not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching groove-jr content", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(content)
}

// About
// GetAllAboutContentHandler handles fetching all about content.
func (mc *MarketingController) GetAllAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllAboutContentHandler Hit")
	page, limit := getPaginationParams(r)

	var content []AboutContent
	err := mc.DB.Model(&content).
		Limit(limit).
		Offset((page - 1) * limit).
		Select()

	if err != nil {
		http.Error(w, "Error fetching about content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(content)
}

// GetAboutContentByIDHandler handles fetching about content by ID.
func (mc *MarketingController) GetAboutContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAboutContentByIDHandler Hit")
	id := r.PathValue("id")

	var content AboutContent
	err := mc.DB.Model(&content).
		Where("id = ?", id).
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			http.Error(w, "About content not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching about content", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(content)
}