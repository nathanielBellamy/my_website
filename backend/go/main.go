package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-pg/pg/v10/orm"
	"github.com/nathanielBellamy/my_website/backend/go/admin"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/config"
	"github.com/nathanielBellamy/my_website/backend/go/db"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	appLogs "github.com/nathanielBellamy/my_website/backend/go/logs"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"
	appMetrics "github.com/nathanielBellamy/my_website/backend/go/metrics"
	"github.com/nathanielBellamy/my_website/backend/go/middleware"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/nathanielBellamy/my_website/backend/go/monitoring"
	"github.com/nathanielBellamy/my_website/backend/go/old_site"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

// MODE=<mode> ./main
func main() {
	startAt := time.Now()

	// Set up log file writer (tee to stdout + file for SSE streaming)
	logFile, err := openLogFile("log", startAt)
	if err != nil {
		// Fall back to stdout-only if log dir is not writable
		fmt.Fprintf(os.Stderr, "Warning: could not open log file: %v (falling back to stdout only)\n", err)
	}
	var logWriter io.Writer
	if logFile != nil {
		logWriter = io.MultiWriter(os.Stdout, logFile)
		defer logFile.Close()
	} else {
		logWriter = os.Stdout
	}

	log := zerolog.New(logWriter).With().Timestamp().Logger()

	// determine runtime env
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "localhost"
		if err := os.Setenv("MODE", mode); err != nil {
			log.Fatal().Err(err).Msg("Failed to set MODE env var")
		}
	}

	// read env file
	log.Info().
		Msg("Loading ENV")

	cfg, err := config.NewConfig(mode)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error loading config")
	}

	log.Info().
		Str("mode", mode).
		Msg("Runtime Env")

	dbClient, err := db.NewDBClient(cfg)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error creating DB client")
	}

	orm.RegisterTable((*models.BlogPostTag)(nil))

	cookieJar := cmap.New[auth.Cookie]()

	log.Info().
		Msg("Establishing Routes")

	feedPool := websocket.NewPool(&log)
	wasmPool := websocket.NewPool(&log)
	go feedPool.StartFeed()
	go wasmPool.StartWasm()

	pgAdapter := db.NewPgDBAdapter(dbClient)

	hostRouter := SetupRoutes(&cookieJar, &log, feedPool, wasmPool, pgAdapter, startAt)

	// Setup Global Rate Limiter: Allows 5 requests per second per IP with a burst of 10
	limiter := middleware.NewIPRateLimiter(rate.Limit(5), 10)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.RateLimitMiddleware(limiter, &log, hostRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal().
			Err(err).
			Msg("UnableToServe")
	}

	log.Info().
		Msg("Now serving on 8080")
}

func SetupRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, feedPool *websocket.Pool, wasmPool *websocket.Pool, db interfaces.PgxDB, startAt time.Time) http.Handler {
	mode := os.Getenv("MODE")
	oldSiteController := old_site.NewOldSiteController(cookieJar, log, feedPool, wasmPool)

	marketingService := marketing.NewService(db)
	marketingController := marketing.NewMarketingController(log, marketingService)

	adminService := admin.NewService(db, log)
	adminController := admin.NewAdminController(log, adminService)

	logsController, err := appLogs.NewLogsController(log, "log", startAt)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize logs controller")
	}
	healthController := appLogs.NewHealthController(log, startAt, db)

	grafanaProxy := monitoring.NewGrafanaProxy(log, "http://grafana:3000")

	adminMux := http.NewServeMux()
	oldSiteMux := http.NewServeMux()
	marketingMux := http.NewServeMux()

	SetupBaseRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController, marketingController, adminController, logsController, healthController, grafanaProxy)

	marketingFileServer := marketing.GetMarketingFileServerNoAuth(log)

	if env.IsProd(mode) {
		SetupProdRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, marketingController, adminController, oldSiteController, marketingFileServer)
	} else if env.IsRemotedev(mode) {
		SetupRemotedevRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController, adminController, marketingFileServer)
	} else {
		SetupLocalhostRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController, adminController, marketingFileServer)
	}

	return appMetrics.InstrumentHandler(log, &middleware.HostRouter{
		AdminMux:     adminMux,
		OldSiteMux:   oldSiteMux,
		MarketingMux: marketingMux,
	})
}

func SetupBaseRoutes(adminMux, oldSiteMux, marketingMux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, marketingController *marketing.MarketingController, adminController *admin.AdminController, logsController *appLogs.LogsController, healthController *appLogs.HealthController, grafanaProxy *monitoring.GrafanaProxy) {
	log.Info().
		Msg("Setting up BaseRoutes")

	// old-site routes
	oldSiteMux.HandleFunc("POST /v1/recaptcha", oldSiteController.RecaptchaHandler)
	oldSiteMux.HandleFunc("GET /v1/public-square-feed-ws", oldSiteController.PublicSquareFeedWsHandler)
	oldSiteMux.HandleFunc("GET /v1/public-square-wasm-ws", oldSiteController.PublicSquareWasmWsHandler)

	// marketing routes
	// Blog
	marketingMux.HandleFunc("GET /v1/api/marketing/blog", marketingController.GetAllBlogPostsHandler)
	marketingMux.HandleFunc("GET /v1/api/marketing/blog/{id}", marketingController.GetBlogPostByIDHandler)
	marketingMux.HandleFunc("GET /v1/api/marketing/blog/tag/{tag}", marketingController.GetBlogPostsByTagHandler)
	marketingMux.HandleFunc("GET /v1/api/marketing/tags", marketingController.GetTagsHandler)

	// Work
	marketingMux.HandleFunc("GET /v1/api/marketing/work", marketingController.GetAllWorkContentHandler)
	marketingMux.HandleFunc("GET /v1/api/marketing/work/{id}", marketingController.GetWorkContentByIDHandler)

	// GrooveJr
	marketingMux.HandleFunc("GET /v1/api/marketing/groovejr", marketingController.GetAllGrooveJrContentHandler)
	marketingMux.HandleFunc("GET /v1/api/marketing/groovejr/{id}", marketingController.GetGrooveJrContentByIDHandler)

	// About
	marketingMux.HandleFunc("GET /v1/api/marketing/about", marketingController.GetAllAboutContentHandler)
	marketingMux.HandleFunc("GET /v1/api/marketing/about/{id}", marketingController.GetAboutContentByIDHandler)

	// Sitemap
	marketingMux.HandleFunc("GET /sitemap.xml", marketingController.SitemapHandler)
	// Robots.txt
	marketingMux.HandleFunc("GET /robots.txt", marketingController.RobotsTxtHandler)

	// Images
	marketingMux.HandleFunc("GET /v1/api/images/{filename}", marketingController.ImageServingHandler)

	// admin routes
	// Blog
	adminMux.Handle("GET /v1/api/admin/blog", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllBlogPostsHandler)))
	adminMux.Handle("GET /v1/api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostByIDHandler)))
	adminMux.Handle("GET /v1/api/admin/blog/tag/{tag}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostsByTagHandler)))
	adminMux.Handle("GET /v1/api/admin/tags", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetTagsHandler)))
	adminMux.Handle("POST /v1/api/admin/blog", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateBlogPostHandler)))
	adminMux.Handle("PUT /v1/api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateBlogPostHandler)))
	adminMux.Handle("DELETE /v1/api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteBlogPostHandler)))

	// Work
	adminMux.Handle("GET /v1/api/admin/work", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllWorkContentHandler)))
	adminMux.Handle("GET /v1/api/admin/work/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetWorkContentByIDHandler)))
	adminMux.Handle("POST /v1/api/admin/work", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateWorkContentHandler)))
	adminMux.Handle("PUT /v1/api/admin/work/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateWorkContentHandler)))
	adminMux.Handle("DELETE /v1/api/admin/work/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteWorkContentHandler)))

	// GrooveJr
	adminMux.Handle("GET /v1/api/admin/groovejr", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllGrooveJrContentHandler)))
	adminMux.Handle("GET /v1/api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetGrooveJrContentByIDHandler)))
	adminMux.Handle("POST /v1/api/admin/groovejr", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateGrooveJrContentHandler)))
	adminMux.Handle("PUT /v1/api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateGrooveJrContentHandler)))
	adminMux.Handle("DELETE /v1/api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteGrooveJrContentHandler)))

	// About
	adminMux.Handle("GET /v1/api/admin/about", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllAboutContentHandler)))
	adminMux.Handle("GET /v1/api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAboutContentByIDHandler)))
	adminMux.Handle("POST /v1/api/admin/about", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateAboutContentHandler)))
	adminMux.Handle("PUT /v1/api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateAboutContentHandler)))
	adminMux.Handle("DELETE /v1/api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteAboutContentHandler)))

	// CSV
	adminMux.Handle("GET /v1/api/admin/csv/{entity}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.ExportCSVHandler)))
	adminMux.Handle("POST /v1/api/admin/csv/{entity}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.ImportCSVHandler)))

	// Images
	adminMux.Handle("POST /v1/api/admin/upload", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UploadImageHandler)))
	adminMux.Handle("GET /v1/api/admin/images", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.ListImagesHandler)))
	adminMux.Handle("DELETE /v1/api/admin/images/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteImageHandler)))

	// Logs & Health
	adminMux.Handle("GET /v1/api/admin/logs/stream", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(logsController.StreamLogsHandler)))
	adminMux.Handle("GET /v1/api/admin/logs/history", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(logsController.GetLogHistoryHandler)))
	adminMux.Handle("GET /v1/api/admin/logs/files", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(logsController.GetLogFilesHandler)))
	adminMux.Handle("GET /v1/api/admin/health", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(healthController.GetHealthHandler)))

	// Prometheus Metrics (auth-protected for external access)
	adminMux.Handle("GET /v1/api/admin/metrics", auth.RequireAdminAuthV2(cookieJar, log, promhttp.Handler()))
	// Internal metrics endpoint for Prometheus scraper (no auth — only accessible within Docker network)
	// Registered on both adminMux and marketingMux so Prometheus can reach it via Host: backend:8080
	adminMux.Handle("GET /internal/metrics", promhttp.Handler())
	marketingMux.Handle("GET /internal/metrics", promhttp.Handler())

	// Grafana Proxy
	adminMux.Handle("/grafana/", auth.RequireAdminAuthV2(cookieJar, log, grafanaProxy))
}

func SetupRemotedevRoutes(adminMux, oldSiteMux, marketingMux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, adminController *admin.AdminController, marketingFileServer http.Handler) {
	auth.SetupAdminAuthV2(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer(), marketingFileServer)
}

func SetupLocalhostRoutes(adminMux, oldSiteMux, marketingMux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, adminController *admin.AdminController, marketingFileServer http.Handler) {
	auth.SetupAdminAuthV2(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer(), marketingFileServer)
}

func SetupProdRoutes(adminMux, oldSiteMux, marketingMux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, marketingController *marketing.MarketingController, adminController *admin.AdminController, oldSiteController *old_site.OldSiteController, marketingFileServer http.Handler) {
	log.Info().
		Msg("Setting up ProdRoutes")

	auth.SetupAdminAuthV2(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer(), auth.LogClientIp("/", log, marketingFileServer))
}

// openLogFile creates a timestamped log file in logDir/YYYY/MM/
func openLogFile(logDir string, startAt time.Time) (*os.File, error) {
	year := startAt.UTC().Format("2006")
	month := startAt.UTC().Format("01")
	dir := filepath.Join(logDir, year, month)

	if err := os.MkdirAll(dir, 0750); err != nil {
		return nil, fmt.Errorf("failed to create log directory %s: %w", dir, err)
	}

	filename := startAt.UTC().Format("2006-01-02T15-04-05Z") + "-log.txt"

	root, err := os.OpenRoot(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to open log root directory %s: %w", dir, err)
	}
	defer root.Close()

	f, err := root.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file %s: %w", filename, err)
	}
	return f, nil
}

