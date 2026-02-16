package marketing

import (
	"encoding/json"
	"net/http"
	"strconv"

	"os"

	_ "github.com/lib/pq"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
)

// MarketingController holds dependencies for marketing-related handlers.

type MarketingController struct {
	Log     *zerolog.Logger
	Service Service
}

// NewMarketingController creates and returns a new MarketingController.

func NewMarketingController(log *zerolog.Logger, service Service) *MarketingController {

	return &MarketingController{
		Log:     log,
		Service: service,
	}

}

func (mc *MarketingController) sendJSON(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		mc.Log.Error().Err(err).Msg("Error encoding response")
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
	posts, err := mc.Service.GetAllBlogPosts(page, limit)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching blog posts")
		http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
		return
	}
	mc.sendJSON(w, posts)
}

// GetBlogPostByIDHandler handles fetching a single blog post by ID.
func (mc *MarketingController) GetBlogPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostByIDHandler Hit")
	id := r.PathValue("id") // Assuming mux or similar router that extracts path variables
	mc.Log.Debug().Str("idStr", id).Msg("GetBlogPostByIDHandler: PathValue 'id'")
	post, err := mc.Service.GetBlogPostByID(id)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching blog post")
		http.Error(w, "Error fetching blog post", http.StatusInternalServerError)
		return
	}

	if post == nil {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	mc.sendJSON(w, post)
}

// GetBlogPostsByTagHandler handles fetching blog posts by tag.

func (mc *MarketingController) GetBlogPostsByTagHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostsByTagHandler Hit")
	tag := r.PathValue("tag") // Assuming mux or similar router
	page, limit := getPaginationParams(r)

	posts, err := mc.Service.GetBlogPostsByTag(tag, page, limit)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching blog posts by tag")
		http.Error(w, "Error fetching blog posts by tag", http.StatusInternalServerError)
		return
	}

	mc.sendJSON(w, posts)
}

// Home

// GetAllHomeContentHandler handles fetching all home content.

func (mc *MarketingController) GetAllHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllHomeContentHandler Hit")
	page, limit := getPaginationParams(r)

	content, err := mc.Service.GetAllHomeContent(page, limit)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching home content")
		http.Error(w, "Error fetching home content", http.StatusInternalServerError)
		return
	}

	mc.sendJSON(w, content)
}

// GetHomeContentByIDHandler handles fetching home content by ID.

func (mc *MarketingController) GetHomeContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetHomeContentByIDHandler Hit")
	id := r.PathValue("id")

	content, err := mc.Service.GetHomeContentByID(id)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching home content")
		http.Error(w, "Error fetching home content", http.StatusInternalServerError)
		return
	}

	if content == nil {
		http.Error(w, "Home content not found", http.StatusNotFound)
		return
	}

	mc.sendJSON(w, content)
}

// GrooveJr

// GetAllGrooveJrContentHandler handles fetching all groove-jr content.

func (mc *MarketingController) GetAllGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllGrooveJrContentHandler Hit")
	page, limit := getPaginationParams(r)

	content, err := mc.Service.GetAllGrooveJrContent(page, limit)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching groove-jr content")
		http.Error(w, "Error fetching groove-jr content", http.StatusInternalServerError)
		return
	}

	mc.sendJSON(w, content)
}

// GetGrooveJrContentByIDHandler handles fetching groove-jr content by ID.

func (mc *MarketingController) GetGrooveJrContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetGrooveJrContentByIDHandler Hit")
	id := r.PathValue("id")

	content, err := mc.Service.GetGrooveJrContentByID(id)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching groove-jr content")
		http.Error(w, "Error fetching groove-jr content", http.StatusInternalServerError)
		return
	}

	if content == nil {
		http.Error(w, "GrooveJr content not found", http.StatusNotFound)
		return
	}

	mc.sendJSON(w, content)
}

// About

// GetAllAboutContentHandler handles fetching all about content.

func (mc *MarketingController) GetAllAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllAboutContentHandler Hit")
	page, limit := getPaginationParams(r)

	content, err := mc.Service.GetAllAboutContent(page, limit)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching about content")
		http.Error(w, "Error fetching about content", http.StatusInternalServerError)
		return
	}

	mc.sendJSON(w, content)
}

// GetAboutContentByIDHandler handles fetching about content by ID.

func (mc *MarketingController) GetAboutContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAboutContentByIDHandler Hit")
	id := r.PathValue("id")

	content, err := mc.Service.GetAboutContentByID(id)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching about content")
		http.Error(w, "Error fetching about content", http.StatusInternalServerError)
		return
	}

	if content == nil {
		http.Error(w, "About content not found", http.StatusNotFound)
		return
	}

	mc.sendJSON(w, content)
}

func GetMarketingFileServerNoAuth() http.Handler {
	root := http.Dir("build/marketing/browser")
	fs := http.FileServer(root)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the file exists
		if _, err := root.Open(r.URL.Path); os.IsNotExist(err) {
			// If not, serve index.html
			http.ServeFile(w, r, "build/marketing/browser/index.html")
			return
		}
		// Otherwise, serve the file
		fs.ServeHTTP(w, r)
	})
}
