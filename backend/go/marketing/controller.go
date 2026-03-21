package marketing

import (
	"encoding/json"
	"encoding/xml"
	"github.com/nathanielBellamy/my_website/backend/go/utils"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"os"
	"path/filepath"

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

	tagsStr := r.URL.Query().Get("tags")
	var tags []string
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}

	posts, err := mc.Service.GetAllBlogPosts(page, limit, tags)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching blog posts")
		utils.HandleDBError(w, err, "Error fetching blog posts")
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
		utils.HandleDBError(w, err, "Error fetching blog post")
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
		utils.HandleDBError(w, err, "Error fetching blog posts by tag")
		return
	}

	mc.sendJSON(w, posts)
}

// GetTagsHandler handles fetching tags.
func (mc *MarketingController) GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetTagsHandler Hit")
	search := r.URL.Query().Get("search")
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	tags, err := mc.Service.GetTags(search, limit)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching tags")
		utils.HandleDBError(w, err, "Error fetching tags")
		return
	}

	mc.sendJSON(w, tags)
}

// Work

// GetAllWorkContentHandler handles fetching all work content.

func (mc *MarketingController) GetAllWorkContentHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllWorkContentHandler Hit")
	page, limit := getPaginationParams(r)

	content, err := mc.Service.GetAllWorkContent(page, limit)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching work content")
		utils.HandleDBError(w, err, "Error fetching work content")
		return
	}

	mc.sendJSON(w, content)
}

// GetWorkContentByIDHandler handles fetching work content by ID.

func (mc *MarketingController) GetWorkContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetWorkContentByIDHandler Hit")
	id := r.PathValue("id")

	content, err := mc.Service.GetWorkContentByID(id)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching work content")
		utils.HandleDBError(w, err, "Error fetching work content")
		return
	}

	if content == nil {
		http.Error(w, "Work content not found", http.StatusNotFound)
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
		utils.HandleDBError(w, err, "Error fetching groove-jr content")
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
		utils.HandleDBError(w, err, "Error fetching groove-jr content")
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
		utils.HandleDBError(w, err, "Error fetching about content")
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
		utils.HandleDBError(w, err, "Error fetching about content")
		return
	}

	if content == nil {
		http.Error(w, "About content not found", http.StatusNotFound)
		return
	}

	mc.sendJSON(w, content)
}

func (mc *MarketingController) determineBaseURL(r *http.Request) string {
	// Default to production
	baseUrl := "https://nateschieber.dev"

	if host := r.Header.Get("Host"); host != "" {
		// Allow localhost for development
		if strings.HasPrefix(host, "localhost") || strings.HasPrefix(host, "127.0.0.1") {
			return "http://" + host
		}
		// Allow www subdomain
		if host == "www.nateschieber.dev" {
			return "https://" + host
		}
		// Allow root domain
		if host == "nateschieber.dev" {
			return "https://" + host
		}
	}
	return baseUrl
}

// SitemapHandler generates the XML sitemap.
func (mc *MarketingController) SitemapHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("SitemapHandler Hit")

	posts, err := mc.Service.GetSitemapData()
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error fetching sitemap data")
		utils.HandleDBError(w, err, "Error fetching sitemap data")
		return
	}

	baseUrl := mc.determineBaseURL(r)

	type URL struct {
		Loc        string `xml:"loc"`
		LastMod    string `xml:"lastmod,omitempty"`
		ChangeFreq string `xml:"changefreq,omitempty"`
		Priority   string `xml:"priority,omitempty"`
	}

	type URLSet struct {
		XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
		URLs    []URL    `xml:"url"`
	}

	var urls []URL

	// Static Pages
	pages := []string{"", "focus", "latest-posts", "about", "groovejr", "blog", "privacy-policy"}
	for _, page := range pages {
		urlStr := baseUrl
		if page != "" {
			urlStr += "/" + page
		}
		urls = append(urls, URL{
			Loc:        urlStr,
			ChangeFreq: "weekly",
			Priority:   "0.8",
		})
	}

	// Dynamic Blog Posts
	for _, post := range posts {
		urlStr := baseUrl + "/blog/" + strings.ReplaceAll(post.ID, "-", "")
		lastMod := post.UpdatedAt.Format("2006-01-02")
		urls = append(urls, URL{
			Loc:        urlStr,
			LastMod:    lastMod,
			ChangeFreq: "monthly",
			Priority:   "0.6",
		})
	}

	w.Header().Set("Content-Type", "application/xml")
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		mc.Log.Error().Err(err).Msg("Error writing XML header")
		return
	}
	if err := xml.NewEncoder(w).Encode(URLSet{URLs: urls}); err != nil {
		mc.Log.Error().Err(err).Msg("Error encoding sitemap")
	}
}

// RobotsTxtHandler serves the robots.txt file.
func (mc *MarketingController) RobotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("RobotsTxtHandler Hit")

	baseUrl := mc.determineBaseURL(r)

	w.Header().Set("Content-Type", "text/plain")

	const robotsTemplate = `User-agent: *
Allow: /

Sitemap: {{.BaseUrl}}/sitemap.xml
`
	tmpl, err := template.New("robots").Parse(robotsTemplate)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error parsing robots.txt template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		BaseUrl string
	}{
		BaseUrl: baseUrl,
	}

	if err := tmpl.Execute(w, data); err != nil {
		mc.Log.Error().Err(err).Msg("Error executing robots.txt template")
	}
}

// ImageServingHandler serves images from the uploads directory.
func (mc *MarketingController) ImageServingHandler(w http.ResponseWriter, r *http.Request) {
	mc.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("ImageServingHandler Hit")
	filename := r.PathValue("filename")

	// Sanitize filename to prevent directory traversal
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}

	uploadDir := "uploads/images"
	filePath := filepath.Join(uploadDir, filepath.Clean(filename))

	// Use os.Root to scope file access under uploadDir
	root, err := os.OpenRoot(uploadDir)
	if err != nil {
		mc.Log.Error().Err(err).Msg("Error opening upload root directory")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	// Check if file exists
	if _, err := root.Stat(filename); err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	// Set caching headers (1 year)
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

	http.ServeFile(w, r, filePath)
}

func GetMarketingFileServerNoAuth(log *zerolog.Logger) http.Handler {
	return auth.SpaHandler("build/marketing/browser", "index.html")
}
