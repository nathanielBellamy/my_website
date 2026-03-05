package admin

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/nathanielBellamy/my_website/backend/go/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/models"
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

func getFilterOptions(r *http.Request) models.FilterOptions {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	status := r.URL.Query().Get("status")
	sortField := r.URL.Query().Get("sort")
	sortOrder := r.URL.Query().Get("order")
	tagsStr := r.URL.Query().Get("tags")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	if status != "current" && status != "inactive" && status != "past" && status != "future" {
		status = "current"
	}

	var tags []string
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}

	return models.FilterOptions{
		Page:      page,
		Limit:     limit,
		Status:    status,
		SortField: sortField,
		SortOrder: sortOrder,
		Tags:      tags,
	}
}

func (ac *AdminController) sendJSON(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		ac.Log.Error().Err(err).Msg("Error encoding response")
	}
}

// Blog
func (ac *AdminController) GetAllBlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllBlogPostsHandler Hit")
	opts := getFilterOptions(r)
	posts, total, err := ac.Service.GetAllBlogPosts(opts)
	if err != nil {
		ac.Log.Error().Err(err).Msg("Error fetching blog posts")
		utils.HandleDBError(w, err, "Error fetching blog posts")
		return
	}
	resp := models.PaginatedResponse[models.BlogPost]{
		Data:  posts,
		Total: total,
		Page:  opts.Page,
		Limit: opts.Limit,
	}
	ac.sendJSON(w, resp)
}

func (ac *AdminController) GetBlogPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostByIDHandler Hit")
	id := r.PathValue("id")
	ac.Log.Info().Str("id", id).Msg("Fetching blog post by ID")
	post, err := ac.Service.GetBlogPostByID(id)
	if err != nil {
		ac.Log.Error().Err(err).Str("id", id).Msg("Error fetching blog post")
		utils.HandleDBError(w, err, "Error fetching blog post")
		return
	}

	if post == nil {
		ac.Log.Warn().Str("id", id).Msg("Blog post not found")
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	ac.Log.Info().Str("id", id).Interface("post", post).Msg("Successfully fetched blog post")
	ac.sendJSON(w, post)
}

func (ac *AdminController) GetBlogPostsByTagHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetBlogPostsByTagHandler Hit")
	tag := r.PathValue("tag")
	opts := getFilterOptions(r)

	posts, err := ac.Service.GetBlogPostsByTag(tag, opts.Page, opts.Limit)
	if err != nil {
		utils.HandleDBError(w, err, "Error fetching blog posts by tag")
		return
	}

	ac.sendJSON(w, posts)
}

func (ac *AdminController) GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetTagsHandler Hit")
	search := r.URL.Query().Get("search")
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}

	tags, err := ac.Service.GetTags(search, limit)
	if err != nil {
		ac.Log.Error().Err(err).Msg("Error fetching tags")
		utils.HandleDBError(w, err, "Error fetching tags")
		return
	}

	ac.sendJSON(w, tags)
}

func (ac *AdminController) CreateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateBlogPostHandler Hit")

	var dto models.CreateBlogPostDTO
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		ac.Log.Error().Err(err).Msg("Error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(bytes.NewReader(bodyBytes))
	if err := decoder.Decode(&dto); err != nil {
		ac.Log.Error().Err(err).RawJSON("body", bodyBytes).Msg("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Map DTO to model
	post := &models.BlogPost{
		Title:         dto.Title,
		Ordering:      dto.Order,
		Content:       dto.Content,
		ActivatedAt:   dto.ActivatedAt,
		DeactivatedAt: dto.DeactivatedAt,
	}
	if dto.Author != nil {
		post.Author = &models.Author{Name: dto.Author.Name}
	}
	if len(dto.Tags) > 0 {
		post.Tags = make([]*models.Tag, len(dto.Tags))
		for i, tagDTO := range dto.Tags {
			post.Tags[i] = &models.Tag{Name: tagDTO.Name}
		}
	}

	newPost, err := ac.Service.CreateBlogPost(post)
	if err != nil {
		utils.HandleDBError(w, err, "Error creating blog post")
		return
	}

	ac.sendJSON(w, newPost)
}

func (ac *AdminController) UpdateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateBlogPostHandler Hit")
	id := r.PathValue("id")
	var post models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	post.ID = id

	updatedPost, err := ac.Service.UpdateBlogPost(&post)
	if err != nil {
		utils.HandleDBError(w, err, "Error updating blog post")
		return
	}

	ac.sendJSON(w, updatedPost)
}

func (ac *AdminController) DeleteBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteBlogPostHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteBlogPost(id); err != nil {
		utils.HandleDBError(w, err, "Error deleting blog post")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Home
func (ac *AdminController) GetAllHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllHomeContentHandler Hit")
	opts := getFilterOptions(r)
	content, total, err := ac.Service.GetAllHomeContent(opts)
	if err != nil {
		utils.HandleDBError(w, err, "Error fetching home content")
		return
	}
	resp := models.PaginatedResponse[models.HomeContent]{
		Data:  content,
		Total: total,
		Page:  opts.Page,
		Limit: opts.Limit,
	}
	ac.sendJSON(w, resp)
}

func (ac *AdminController) GetHomeContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetHomeContentByIDHandler Hit")
	id := r.PathValue("id")
	content, err := ac.Service.GetHomeContentByID(id)
	if err != nil {
		utils.HandleDBError(w, err, "Error fetching home content")
		return
	}

	if content == nil {
		http.Error(w, "Home content not found", http.StatusNotFound)
		return
	}

	ac.sendJSON(w, content)
}

func (ac *AdminController) CreateHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateHomeContentHandler Hit")
	var content models.HomeContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newContent, err := ac.Service.CreateHomeContent(&content)
	if err != nil {
		utils.HandleDBError(w, err, "Error creating home content")
		return
	}

	ac.sendJSON(w, newContent)
}

func (ac *AdminController) UpdateHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateHomeContentHandler Hit")
	id := r.PathValue("id")
	var content models.HomeContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	content.ID = id

	updatedContent, err := ac.Service.UpdateHomeContent(&content)
	if err != nil {
		utils.HandleDBError(w, err, "Error updating home content")
		return
	}

	ac.sendJSON(w, updatedContent)
}

func (ac *AdminController) DeleteHomeContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteHomeContentHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteHomeContent(id); err != nil {
		utils.HandleDBError(w, err, "Error deleting home content")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GrooveJr
func (ac *AdminController) GetAllGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllGrooveJrContentHandler Hit")
	opts := getFilterOptions(r)
	content, total, err := ac.Service.GetAllGrooveJrContent(opts)
	if err != nil {
		utils.HandleDBError(w, err, "Error fetching groove-jr content")
		return
	}
	resp := models.PaginatedResponse[models.GrooveJrContent]{
		Data:  content,
		Total: total,
		Page:  opts.Page,
		Limit: opts.Limit,
	}
	ac.sendJSON(w, resp)
}

func (ac *AdminController) GetGrooveJrContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetGrooveJrContentByIDHandler Hit")
	id := r.PathValue("id")
	content, err := ac.Service.GetGrooveJrContentByID(id)
	if err != nil {
		utils.HandleDBError(w, err, "Error fetching groove-jr content")
		return
	}

	if content == nil {
		http.Error(w, "GrooveJr content not found", http.StatusNotFound)
		return
	}

	ac.sendJSON(w, content)
}

func (ac *AdminController) CreateGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateGrooveJrContentHandler Hit")
	var content models.GrooveJrContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newContent, err := ac.Service.CreateGrooveJrContent(&content)
	if err != nil {
		utils.HandleDBError(w, err, "Error creating groove-jr content")
		return
	}

	ac.sendJSON(w, newContent)
}

func (ac *AdminController) UpdateGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateGrooveJrContentHandler Hit")
	id := r.PathValue("id")
	var content models.GrooveJrContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	content.ID = id

	updatedContent, err := ac.Service.UpdateGrooveJrContent(&content)
	if err != nil {
		utils.HandleDBError(w, err, "Error updating groove-jr content")
		return
	}

	ac.sendJSON(w, updatedContent)
}

func (ac *AdminController) DeleteGrooveJrContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteGrooveJrContentHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteGrooveJrContent(id); err != nil {
		utils.HandleDBError(w, err, "Error deleting groove-jr content")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// About
func (ac *AdminController) GetAllAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAllAboutContentHandler Hit")
	opts := getFilterOptions(r)
	content, total, err := ac.Service.GetAllAboutContent(opts)
	if err != nil {
		utils.HandleDBError(w, err, "Error fetching about content")
		return
	}
	resp := models.PaginatedResponse[models.AboutContent]{
		Data:  content,
		Total: total,
		Page:  opts.Page,
		Limit: opts.Limit,
	}
	ac.sendJSON(w, resp)
}

func (ac *AdminController) GetAboutContentByIDHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("GetAboutContentByIDHandler Hit")
	id := r.PathValue("id")
	content, err := ac.Service.GetAboutContentByID(id)
	if err != nil {
		utils.HandleDBError(w, err, "Error fetching about content")
		return
	}

	if content == nil {
		http.Error(w, "About content not found", http.StatusNotFound)
		return
	}

	ac.sendJSON(w, content)
}

func (ac *AdminController) CreateAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("CreateAboutContentHandler Hit")
	var content models.AboutContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newContent, err := ac.Service.CreateAboutContent(&content)
	if err != nil {
		utils.HandleDBError(w, err, "Error creating about content")
		return
	}

	ac.sendJSON(w, newContent)
}

func (ac *AdminController) UpdateAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UpdateAboutContentHandler Hit")
	id := r.PathValue("id")
	var content models.AboutContent
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	content.ID = id

	updatedContent, err := ac.Service.UpdateAboutContent(&content)
	if err != nil {
		utils.HandleDBError(w, err, "Error updating about content")
		return
	}

	ac.sendJSON(w, updatedContent)
}

func (ac *AdminController) DeleteAboutContentHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteAboutContentHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteAboutContent(id); err != nil {
		utils.HandleDBError(w, err, "Error deleting about content")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Images

func (ac *AdminController) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("UploadImageHandler Hit")

	// Limit upload size to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	altText := r.FormValue("altText")

	// Create directory if it doesn't exist
	uploadDir := "uploads/images"
	if err := os.MkdirAll(uploadDir, 0750); err != nil {
		ac.Log.Error().Err(err).Msg("Error creating upload directory")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Generate unique filename
	ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Use os.Root to scope file access under uploadDir and prevent directory traversal
	root, err := os.OpenRoot(uploadDir)
	if err != nil {
		ac.Log.Error().Err(err).Msg("Error opening upload root directory")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	dst, err := root.Create(filename)
	if err != nil {
		ac.Log.Error().Err(err).Msg("Error saving file to disk")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		ac.Log.Error().Err(err).Msg("Error copying file content")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Save metadata to DB
	filePath := filepath.Join(uploadDir, filename)
	image, err := ac.Service.UploadImage(filename, header.Filename, altText)
	if err != nil {
		ac.Log.Error().Err(err).Msg("Error saving image metadata")
		// Clean up file if DB fails
		if removeErr := root.Remove(filename); removeErr != nil {
			ac.Log.Error().Err(removeErr).Str("filePath", filePath).Msg("Error cleaning up file after DB failure")
		}
		utils.HandleDBError(w, err, "Error saving image metadata")
		return
	}

	ac.sendJSON(w, image)
}

func (ac *AdminController) ListImagesHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("ListImagesHandler Hit")
	images, err := ac.Service.ListImages()
	if err != nil {
		utils.HandleDBError(w, err, "Error listing images")
		return
	}
	ac.sendJSON(w, images)
}

func (ac *AdminController) DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	ac.Log.Info().Str("ip", auth.GetClientIpAddr(r)).Msg("DeleteImageHandler Hit")
	id := r.PathValue("id")

	if err := ac.Service.DeleteImage(id); err != nil {
		utils.HandleDBError(w, err, "Error deleting image")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AdminFileServer serves static files for the admin site, using the SpaHandler to handle client-side routing.
func (ac *AdminController) AdminFileServer() http.Handler {
	rootPath := "build/admin/browser"
	root := http.Dir(rootPath)
	fs := http.FileServer(root)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to open file
		f, err := root.Open(r.URL.Path)
		if os.IsNotExist(err) {
			// Not found, serve index.html
			http.ServeFile(w, r, rootPath+"/index.html")
			return
		}
		defer f.Close()

		// If it exists, let FileServer handle it
		fs.ServeHTTP(w, r)
	})

	return auth.LogClientIp("/", ac.Log, handler)
}

// CSV Handlers

func (ac *AdminController) ExportCSVHandler(w http.ResponseWriter, r *http.Request) {
	entity := r.PathValue("entity")
	ac.Log.Info().Str("entity", entity).Msg("ExportCSVHandler Hit")

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", entity))

	writer := csv.NewWriter(w)
	defer writer.Flush()

	switch entity {
	case "blog":
		posts, err := ac.Service.ExportBlogPosts()
		if err != nil {
			utils.HandleDBError(w, err, "Error exporting blog posts")
			return
		}
		// Headers
		if err := writer.Write([]string{"id", "title", "content", "author_name", "author_id", "tags", "ordering", "created_at", "updated_at", "activated_at", "deactivated_at"}); err != nil {
			ac.Log.Error().Err(err).Msg("Error writing CSV header")
			return
		}
		// Data
		for _, p := range posts {
			tags := []string{}
			for _, t := range p.Tags {
				tags = append(tags, t.Name)
			}
			authorName := ""
			authorID := ""
			if p.Author != nil {
				authorName = p.Author.Name
				authorID = p.Author.ID
			}
			if err := writer.Write([]string{
				p.ID,
				p.Title,
				p.Content,
				authorName,
				authorID,
				strings.Join(tags, "|"),
				strconv.Itoa(p.Ordering),
				p.CreatedAt.Format(time.RFC3339),
				p.UpdatedAt.Format(time.RFC3339),
				formatTimePtr(p.ActivatedAt),
				formatTimePtr(p.DeactivatedAt),
			}); err != nil {
				ac.Log.Error().Err(err).Msg("Error writing CSV record")
				return
			}
		}
	case "home":
		content, err := ac.Service.ExportHomeContent()
		if err != nil {
			utils.HandleDBError(w, err, "Error exporting home content")
			return
		}
		if err := writer.Write([]string{"id", "title", "content", "ordering", "activated_at", "deactivated_at"}); err != nil {
			ac.Log.Error().Err(err).Msg("Error writing CSV header")
			return
		}
		for _, c := range content {
			if err := writer.Write([]string{
				c.ID,
				c.Title,
				c.Content,
				strconv.Itoa(c.Ordering),
				formatTimePtr(c.ActivatedAt),
				formatTimePtr(c.DeactivatedAt),
			}); err != nil {
				ac.Log.Error().Err(err).Msg("Error writing CSV record")
				return
			}
		}
	case "groove-jr":
		content, err := ac.Service.ExportGrooveJrContent()
		if err != nil {
			utils.HandleDBError(w, err, "Error exporting groove-jr content")
			return
		}
		if err := writer.Write([]string{"id", "title", "content", "ordering", "activated_at", "deactivated_at"}); err != nil {
			ac.Log.Error().Err(err).Msg("Error writing CSV header")
			return
		}
		for _, c := range content {
			if err := writer.Write([]string{
				c.ID,
				c.Title,
				c.Content,
				strconv.Itoa(c.Ordering),
				formatTimePtr(c.ActivatedAt),
				formatTimePtr(c.DeactivatedAt),
			}); err != nil {
				ac.Log.Error().Err(err).Msg("Error writing CSV record")
				return
			}
		}
	case "about":
		content, err := ac.Service.ExportAboutContent()
		if err != nil {
			utils.HandleDBError(w, err, "Error exporting about content")
			return
		}
		if err := writer.Write([]string{"id", "title", "content", "ordering", "activated_at", "deactivated_at"}); err != nil {
			ac.Log.Error().Err(err).Msg("Error writing CSV header")
			return
		}
		for _, c := range content {
			if err := writer.Write([]string{
				c.ID,
				c.Title,
				c.Content,
				strconv.Itoa(c.Ordering),
				formatTimePtr(c.ActivatedAt),
				formatTimePtr(c.DeactivatedAt),
			}); err != nil {
				ac.Log.Error().Err(err).Msg("Error writing CSV record")
				return
			}
		}
	case "tags":
		tags, err := ac.Service.ExportTags()
		if err != nil {
			utils.HandleDBError(w, err, "Error exporting tags")
			return
		}
		if err := writer.Write([]string{"id", "name", "activated_at", "deactivated_at"}); err != nil {
			ac.Log.Error().Err(err).Msg("Error writing CSV header")
			return
		}
		for _, t := range tags {
			if err := writer.Write([]string{
				t.ID,
				t.Name,
				formatTimePtr(t.ActivatedAt),
				formatTimePtr(t.DeactivatedAt),
			}); err != nil {
				ac.Log.Error().Err(err).Msg("Error writing CSV record")
				return
			}
		}
	case "authors":
		authors, err := ac.Service.ExportAuthors()
		if err != nil {
			utils.HandleDBError(w, err, "Error exporting authors")
			return
		}
		if err := writer.Write([]string{"id", "name", "activated_at", "deactivated_at"}); err != nil {
			ac.Log.Error().Err(err).Msg("Error writing CSV header")
			return
		}
		for _, a := range authors {
			if err := writer.Write([]string{
				a.ID,
				a.Name,
				formatTimePtr(a.ActivatedAt),
				formatTimePtr(a.DeactivatedAt),
			}); err != nil {
				ac.Log.Error().Err(err).Msg("Error writing CSV record")
				return
			}
		}
	default:
		http.Error(w, "Unknown entity", http.StatusBadRequest)
	}
}

func (ac *AdminController) ImportCSVHandler(w http.ResponseWriter, r *http.Request) {
	entity := r.PathValue("entity")
	ac.Log.Info().Str("entity", entity).Msg("ImportCSVHandler Hit")

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Error reading CSV", http.StatusBadRequest)
		return
	}

	if len(records) < 2 {
		http.Error(w, "Empty CSV or missing header", http.StatusBadRequest)
		return
	}

	header := records[0]
	// Basic map to index
	colIdx := make(map[string]int)
	for i, h := range header {
		colIdx[h] = i
	}

	switch entity {
	case "blog":
		var posts []models.BlogPost
		for _, row := range records[1:] {
			post := models.BlogPost{}
			if idx, ok := colIdx["id"]; ok {
				post.ID = row[idx]
			}
			if idx, ok := colIdx["title"]; ok {
				post.Title = row[idx]
			}
			if idx, ok := colIdx["content"]; ok {
				post.Content = row[idx]
			}
			if idx, ok := colIdx["ordering"]; ok {
				post.Ordering, _ = strconv.Atoi(row[idx])
			}
			if idx, ok := colIdx["created_at"]; ok {
				post.CreatedAt, _ = time.Parse(time.RFC3339, row[idx])
			}
			if idx, ok := colIdx["updated_at"]; ok {
				post.UpdatedAt, _ = time.Parse(time.RFC3339, row[idx])
			}
			if idx, ok := colIdx["activated_at"]; ok {
				post.ActivatedAt = parseTimePtr(row[idx])
			}
			if idx, ok := colIdx["deactivated_at"]; ok {
				post.DeactivatedAt = parseTimePtr(row[idx])
			}

			// Relations
			authorName := ""
			authorID := ""
			if idx, ok := colIdx["author_name"]; ok {
				authorName = row[idx]
			}
			if idx, ok := colIdx["author_id"]; ok {
				authorID = row[idx]
			}
			if authorName != "" || authorID != "" {
				post.Author = &models.Author{Name: authorName, ID: authorID}
			}

			if idx, ok := colIdx["tags"]; ok && row[idx] != "" {
				tagNames := strings.Split(row[idx], "|")
				for _, name := range tagNames {
					post.Tags = append(post.Tags, &models.Tag{Name: name})
				}
			}
			posts = append(posts, post)
		}
		if err := ac.Service.ImportBlogPosts(posts); err != nil {
			ac.Log.Error().Err(err).Msg("Error importing blog posts")
			utils.HandleDBError(w, err, "Error importing blog posts")
			return
		}

	case "home":
		var content []models.HomeContent
		for _, row := range records[1:] {
			c := models.HomeContent{}
			if idx, ok := colIdx["id"]; ok {
				c.ID = row[idx]
			}
			if idx, ok := colIdx["title"]; ok {
				c.Title = row[idx]
			}
			if idx, ok := colIdx["content"]; ok {
				c.Content = row[idx]
			}
			if idx, ok := colIdx["ordering"]; ok {
				c.Ordering, _ = strconv.Atoi(row[idx])
			}
			if idx, ok := colIdx["activated_at"]; ok {
				c.ActivatedAt = parseTimePtr(row[idx])
			}
			if idx, ok := colIdx["deactivated_at"]; ok {
				c.DeactivatedAt = parseTimePtr(row[idx])
			}
			content = append(content, c)
		}
		if err := ac.Service.ImportHomeContent(content); err != nil {
			utils.HandleDBError(w, err, "Error importing home content")
			return
		}

	case "groove-jr":
		var content []models.GrooveJrContent
		for _, row := range records[1:] {
			c := models.GrooveJrContent{}
			if idx, ok := colIdx["id"]; ok {
				c.ID = row[idx]
			}
			if idx, ok := colIdx["title"]; ok {
				c.Title = row[idx]
			}
			if idx, ok := colIdx["content"]; ok {
				c.Content = row[idx]
			}
			if idx, ok := colIdx["ordering"]; ok {
				c.Ordering, _ = strconv.Atoi(row[idx])
			}
			if idx, ok := colIdx["activated_at"]; ok {
				c.ActivatedAt = parseTimePtr(row[idx])
			}
			if idx, ok := colIdx["deactivated_at"]; ok {
				c.DeactivatedAt = parseTimePtr(row[idx])
			}
			content = append(content, c)
		}
		if err := ac.Service.ImportGrooveJrContent(content); err != nil {
			utils.HandleDBError(w, err, "Error importing groove-jr content")
			return
		}

	case "about":
		var content []models.AboutContent
		for _, row := range records[1:] {
			c := models.AboutContent{}
			if idx, ok := colIdx["id"]; ok {
				c.ID = row[idx]
			}
			if idx, ok := colIdx["title"]; ok {
				c.Title = row[idx]
			}
			if idx, ok := colIdx["content"]; ok {
				c.Content = row[idx]
			}
			if idx, ok := colIdx["ordering"]; ok {
				c.Ordering, _ = strconv.Atoi(row[idx])
			}
			if idx, ok := colIdx["activated_at"]; ok {
				c.ActivatedAt = parseTimePtr(row[idx])
			}
			if idx, ok := colIdx["deactivated_at"]; ok {
				c.DeactivatedAt = parseTimePtr(row[idx])
			}
			content = append(content, c)
		}
		if err := ac.Service.ImportAboutContent(content); err != nil {
			utils.HandleDBError(w, err, "Error importing about content")
			return
		}

	case "tags":
		var tags []models.Tag
		for _, row := range records[1:] {
			t := models.Tag{}
			if idx, ok := colIdx["id"]; ok {
				t.ID = row[idx]
			}
			if idx, ok := colIdx["name"]; ok {
				t.Name = row[idx]
			}
			if idx, ok := colIdx["activated_at"]; ok {
				t.ActivatedAt = parseTimePtr(row[idx])
			}
			if idx, ok := colIdx["deactivated_at"]; ok {
				t.DeactivatedAt = parseTimePtr(row[idx])
			}
			tags = append(tags, t)
		}
		if err := ac.Service.ImportTags(tags); err != nil {
			utils.HandleDBError(w, err, "Error importing tags")
			return
		}

	case "authors":
		var authors []models.Author
		for _, row := range records[1:] {
			a := models.Author{}
			if idx, ok := colIdx["id"]; ok {
				a.ID = row[idx]
			}
			if idx, ok := colIdx["name"]; ok {
				a.Name = row[idx]
			}
			if idx, ok := colIdx["activated_at"]; ok {
				a.ActivatedAt = parseTimePtr(row[idx])
			}
			if idx, ok := colIdx["deactivated_at"]; ok {
				a.DeactivatedAt = parseTimePtr(row[idx])
			}
			authors = append(authors, a)
		}
		if err := ac.Service.ImportAuthors(authors); err != nil {
			utils.HandleDBError(w, err, "Error importing authors")
			return
		}

	default:
		http.Error(w, "Unknown entity", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func formatTimePtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}

func parseTimePtr(s string) *time.Time {
	if s == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil
	}
	return &t
}
