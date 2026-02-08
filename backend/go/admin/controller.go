package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
)

type AdminController struct {
	Log     *zerolog.Logger
	Service Service
}

func NewAdminController(log *zerolog.Logger, service Service) *AdminController {
	return &AdminController{
		Log:     log,
		Service: service,
	}
}

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
func (ac *AdminController) GetAllBlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllBlogPostsHandler Hit")
	page, limit := getPaginationParams(r)
	posts, err := ac.Service.GetAllBlogPosts(page, limit)
	if err != nil {
		http.Error(w, "Error fetching blog posts", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (ac *AdminController) GetBlogPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostByIDHandler Hit")
	id := r.PathValue("id")
	post, err := ac.Service.GetBlogPostByID(id)
	if err != nil {
		http.Error(w, "Error fetching blog post", http.StatusInternalServerError)
		return
	}

	if post == nil {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (ac *AdminController) GetBlogPostsByTagHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostsByTagHandler Hit")
	tag := r.PathValue("tag")
	page, limit := getPaginationParams(r)

	posts, err := ac.Service.GetBlogPostsByTag(tag, page, limit)
	if err != nil {
		http.Error(w, "Error fetching blog posts by tag", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func (ac *AdminController) CreateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateBlogPostHandler Hit")
	var post BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newPost, err := ac.Service.CreateBlogPost(&post)
	if err != nil {
		http.Error(w, "Error creating blog post", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPost)
}

func (ac *AdminController) UpdateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateBlogPostHandler Hit")
	id := r.PathValue("id")
	var post BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	post.ID = id

	updatedPost, err := ac.Service.UpdateBlogPost(&post)
	if err != nil {
		http.Error(w, "Error updating blog post", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedPost)
}

func (ac *AdminController) DeleteBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteBlogPostHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteBlogPost(id); err != nil {
		http.Error(w, "Error deleting blog post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Home
func (ac *AdminController) GetAllHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllHomeContentHandler Hit")
	page, limit := getPaginationParams(r)
	content, err := ac.Service.GetAllHomeContent(page, limit)
	if err != nil {
		http.Error(w, "Error fetching home content", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(content)
}

func (ac *AdminController) GetHomeContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetHomeContentByIDHandler Hit")
	id := r.PathValue("id")
	content, err := ac.Service.GetHomeContentByID(id)
	if err != nil {
		http.Error(w, "Error fetching home content", http.StatusInternalServerError)
		return
	}

	if content == nil {
		http.Error(w, "Home content not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(content)
}

func (ac *AdminController) CreateHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateHomeContentHandler Hit")
	var content HomeContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newContent, err := ac.Service.CreateHomeContent(&content)
	if err != nil {
		http.Error(w, "Error creating home content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newContent)
}

func (ac *AdminController) UpdateHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateHomeContentHandler Hit")
	id := r.PathValue("id")
	var content HomeContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	content.ID = id

	updatedContent, err := ac.Service.UpdateHomeContent(&content)
	if err != nil {
		http.Error(w, "Error updating home content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedContent)
}

func (ac *AdminController) DeleteHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteHomeContentHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteHomeContent(id); err != nil {
		http.Error(w, "Error deleting home content", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GrooveJr
func (ac *AdminController) GetAllGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllGrooveJrContentHandler Hit")
	page, limit := getPaginationParams(r)
	content, err := ac.Service.GetAllGrooveJrContent(page, limit)
	if err != nil {
		http.Error(w, "Error fetching groove-jr content", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(content)
}

func (ac *AdminController) GetGrooveJrContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetGrooveJrContentByIDHandler Hit")
	id := r.PathValue("id")
	content, err := ac.Service.GetGrooveJrContentByID(id)
	if err != nil {
		http.Error(w, "Error fetching groove-jr content", http.StatusInternalServerError)
		return
	}

	if content == nil {
		http.Error(w, "GrooveJr content not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(content)
}

func (ac *AdminController) CreateGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateGrooveJrContentHandler Hit")
	var content GrooveJrContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newContent, err := ac.Service.CreateGrooveJrContent(&content)
	if err != nil {
		http.Error(w, "Error creating groove-jr content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newContent)
}

func (ac *AdminController) UpdateGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateGrooveJrContentHandler Hit")
	id := r.PathValue("id")
	var content GrooveJrContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	content.ID = id

	updatedContent, err := ac.Service.UpdateGrooveJrContent(&content)
	if err != nil {
		http.Error(w, "Error updating groove-jr content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedContent)
}

func (ac *AdminController) DeleteGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteGrooveJrContentHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteGrooveJrContent(id); err != nil {
		http.Error(w, "Error deleting groove-jr content", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// About
func (ac *AdminController) GetAllAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllAboutContentHandler Hit")
	page, limit := getPaginationParams(r)
	content, err := ac.Service.GetAllAboutContent(page, limit)
	if err != nil {
		http.Error(w, "Error fetching about content", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(content)
}

func (ac *AdminController) GetAboutContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAboutContentByIDHandler Hit")
	id := r.PathValue("id")
	content, err := ac.Service.GetAboutContentByID(id)
	if err != nil {
		http.Error(w, "Error fetching about content", http.StatusInternalServerError)
		return
	}

	if content == nil {
		http.Error(w, "About content not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(content)
}

func (ac *AdminController) CreateAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateAboutContentHandler Hit")
	var content AboutContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newContent, err := ac.Service.CreateAboutContent(&content)
	if err != nil {
		http.Error(w, "Error creating about content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newContent)
}

func (ac *AdminController) UpdateAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateAboutContentHandler Hit")
	id := r.PathValue("id")
	var content AboutContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	content.ID = id

	updatedContent, err := ac.Service.UpdateAboutContent(&content)
	if err != nil {
		http.Error(w, "Error updating about content", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedContent)
}

func (ac *AdminController) DeleteAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteAboutContentHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteAboutContent(id); err != nil {
		http.Error(w, "Error deleting about content", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AdminFileServer serves static files for the admin site.
func (ac *AdminController) AdminFileServer() http.Handler {
	fsAdmin := http.FileServer(http.Dir("build/admin/browser"))
	return http.StripPrefix("/admin/", auth.LogClientIp("/admin/", ac.Log, fsAdmin))
}
