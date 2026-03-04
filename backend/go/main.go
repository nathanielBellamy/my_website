package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-pg/pg/v10/orm"
	"github.com/nathanielBellamy/my_website/backend/go/admin"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/config"
	"github.com/nathanielBellamy/my_website/backend/go/db"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"
	"github.com/nathanielBellamy/my_website/backend/go/middleware"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/nathanielBellamy/my_website/backend/go/old_site"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

// MODE=<mode> ./main
func main() {
	// init log
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

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

	hostRouter := SetupRoutes(&cookieJar, &log, feedPool, wasmPool, db.NewPgDBAdapter(dbClient))

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

func SetupRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, feedPool *websocket.Pool, wasmPool *websocket.Pool, db interfaces.PgxDB) http.Handler {
	mode := os.Getenv("MODE")
	oldSiteController := old_site.NewOldSiteController(cookieJar, log, feedPool, wasmPool)

	marketingService := marketing.NewService(db)
	marketingController := marketing.NewMarketingController(log, marketingService)

	adminService := admin.NewService(db, log)
	adminController := admin.NewAdminController(log, adminService)

	adminMux := http.NewServeMux()
	oldSiteMux := http.NewServeMux()
	marketingMux := http.NewServeMux()

	SetupBaseRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController, marketingController, adminController)

	marketingFileServer := marketing.GetMarketingFileServerNoAuth(log)

	if env.IsProd(mode) {
		SetupProdRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, marketingController, adminController, oldSiteController, marketingFileServer)
	} else if env.IsRemotedev(mode) {
		SetupRemotedevRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController, adminController, marketingFileServer)
	} else {
		SetupLocalhostRoutes(adminMux, oldSiteMux, marketingMux, cookieJar, log, oldSiteController, adminController, marketingFileServer)
	}

	return &middleware.HostRouter{
		AdminMux:     adminMux,
		OldSiteMux:   oldSiteMux,
		MarketingMux: marketingMux,
	}
}

func SetupBaseRoutes(adminMux, oldSiteMux, marketingMux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, marketingController *marketing.MarketingController, adminController *admin.AdminController) {
	log.Info().
		Msg("Setting up BaseRoutes")

	// old-site routes
	oldSiteMux.HandleFunc("POST /recaptcha", oldSiteController.RecaptchaHandler)
	oldSiteMux.HandleFunc("GET /public-square-feed-ws", oldSiteController.PublicSquareFeedWsHandler)
	oldSiteMux.HandleFunc("GET /public-square-wasm-ws", oldSiteController.PublicSquareWasmWsHandler)

	// marketing routes
	// Blog
	marketingMux.HandleFunc("GET /api/marketing/blog", marketingController.GetAllBlogPostsHandler)
	marketingMux.HandleFunc("GET /api/marketing/blog/{id}", marketingController.GetBlogPostByIDHandler)
	marketingMux.HandleFunc("GET /api/marketing/blog/tag/{tag}", marketingController.GetBlogPostsByTagHandler)
	marketingMux.HandleFunc("GET /api/marketing/tags", marketingController.GetTagsHandler)

	// Home
	marketingMux.HandleFunc("GET /api/marketing/home", marketingController.GetAllHomeContentHandler)
	marketingMux.HandleFunc("GET /api/marketing/home/{id}", marketingController.GetHomeContentByIDHandler)

	// GrooveJr
	marketingMux.HandleFunc("GET /api/marketing/groovejr", marketingController.GetAllGrooveJrContentHandler)
	marketingMux.HandleFunc("GET /api/marketing/groovejr/{id}", marketingController.GetGrooveJrContentByIDHandler)

	// About
	marketingMux.HandleFunc("GET /api/marketing/about", marketingController.GetAllAboutContentHandler)
	marketingMux.HandleFunc("GET /api/marketing/about/{id}", marketingController.GetAboutContentByIDHandler)

	// Sitemap
	marketingMux.HandleFunc("GET /sitemap.xml", marketingController.SitemapHandler)
	// Robots.txt
	marketingMux.HandleFunc("GET /robots.txt", marketingController.RobotsTxtHandler)

	// Images
	marketingMux.HandleFunc("GET /api/images/{filename}", marketingController.ImageServingHandler)

	// admin routes
	// Blog
	adminMux.Handle("GET /api/admin/blog", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllBlogPostsHandler)))
	adminMux.Handle("GET /api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostByIDHandler)))
	adminMux.Handle("GET /api/admin/blog/tag/{tag}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostsByTagHandler)))
	adminMux.Handle("GET /api/admin/tags", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetTagsHandler)))
	adminMux.Handle("POST /api/admin/blog", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateBlogPostHandler)))
	adminMux.Handle("PUT /api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateBlogPostHandler)))
	adminMux.Handle("DELETE /api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteBlogPostHandler)))

	// Home
	adminMux.Handle("GET /api/admin/home", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllHomeContentHandler)))
	adminMux.Handle("GET /api/admin/home/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetHomeContentByIDHandler)))
	adminMux.Handle("POST /api/admin/home", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateHomeContentHandler)))
	adminMux.Handle("PUT /api/admin/home/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateHomeContentHandler)))
	adminMux.Handle("DELETE /api/admin/home/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteHomeContentHandler)))

	// GrooveJr
	adminMux.Handle("GET /api/admin/groovejr", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllGrooveJrContentHandler)))
	adminMux.Handle("GET /api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetGrooveJrContentByIDHandler)))
	adminMux.Handle("POST /api/admin/groovejr", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateGrooveJrContentHandler)))
	adminMux.Handle("PUT /api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateGrooveJrContentHandler)))
	adminMux.Handle("DELETE /api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteGrooveJrContentHandler)))

	// About
	adminMux.Handle("GET /api/admin/about", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllAboutContentHandler)))
	adminMux.Handle("GET /api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAboutContentByIDHandler)))
	adminMux.Handle("POST /api/admin/about", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateAboutContentHandler)))
	adminMux.Handle("PUT /api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateAboutContentHandler)))
	adminMux.Handle("DELETE /api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteAboutContentHandler)))

	// CSV
	adminMux.Handle("GET /api/admin/csv/{entity}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.ExportCSVHandler)))
	adminMux.Handle("POST /api/admin/csv/{entity}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.ImportCSVHandler)))

	// Images
	adminMux.Handle("POST /api/admin/upload", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UploadImageHandler)))
	adminMux.Handle("GET /api/admin/images", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.ListImagesHandler)))
	adminMux.Handle("DELETE /api/admin/images/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteImageHandler)))
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

