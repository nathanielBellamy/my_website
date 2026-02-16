package main

import (
	"net/http"
	"os"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/admin"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/config"
	"github.com/nathanielBellamy/my_website/backend/go/db"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/nathanielBellamy/my_website/backend/go/old_site"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
	"github.com/go-pg/pg/v10/orm"
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
			log.Warn().Err(err).Msg("Failed to set MODE environment variable")
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

	SetupRoutes(http.DefaultServeMux, &cookieJar, &log, feedPool, wasmPool, db.NewPgDBAdapter(dbClient))

	// Create HTTP server with proper timeouts to prevent resource exhaustion
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // Use DefaultServeMux
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Info().
		Msg("Starting server on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal().
			Err(err).
			Msg("UnableToServe")
	}
}

func SetupRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, feedPool *websocket.Pool, wasmPool *websocket.Pool, db interfaces.PgxDB) {
	mode := os.Getenv("MODE")
	oldSiteController := old_site.NewOldSiteController(cookieJar, log, feedPool, wasmPool)

	marketingService := marketing.NewService(db)
	marketingController := marketing.NewMarketingController(log, marketingService)

	adminService := admin.NewService(db, log)
	adminController := admin.NewAdminController(log, adminService)

	SetupBaseRoutes(mux, cookieJar, log, oldSiteController, marketingController, adminController)

	if env.IsProd(mode) {
		SetupProdRoutes(mux, cookieJar, log, marketingController, adminController, oldSiteController, marketing.GetMarketingFileServerNoAuth())
	} else if env.IsRemotedev(mode) {
		SetupRemotedevRoutes(mux, cookieJar, log, oldSiteController, adminController, marketing.GetMarketingFileServerNoAuth())
	} else {
		SetupLocalhostRoutes(mux, cookieJar, log, oldSiteController, adminController, marketing.GetMarketingFileServerNoAuth())
	}
}

func SetupBaseRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, marketingController *marketing.MarketingController, adminController *admin.AdminController) {
	log.Info().
		Msg("Setting up BaseRoutes")

	// old-site routes
	mux.HandleFunc("/old-site/recaptcha", oldSiteController.RecaptchaHandler)
	mux.HandleFunc("/old-site/public-square-feed-ws", oldSiteController.PublicSquareFeedWsHandler)
	mux.HandleFunc("/old-site/public-square-wasm-ws", oldSiteController.PublicSquareWasmWsHandler)

	// marketing routes
	// Blog
	mux.HandleFunc("GET /api/marketing/blog", marketingController.GetAllBlogPostsHandler)
	mux.HandleFunc("GET /api/marketing/blog/{id}", marketingController.GetBlogPostByIDHandler)
	mux.HandleFunc("GET /api/marketing/blog/tag/{tag}", marketingController.GetBlogPostsByTagHandler)

	// Home
	mux.HandleFunc("GET /api/marketing/home", marketingController.GetAllHomeContentHandler)
	mux.HandleFunc("GET /api/marketing/home/{id}", marketingController.GetHomeContentByIDHandler)

	// GrooveJr
	mux.HandleFunc("GET /api/marketing/groovejr", marketingController.GetAllGrooveJrContentHandler)
	mux.HandleFunc("GET /api/marketing/groovejr/{id}", marketingController.GetGrooveJrContentByIDHandler)

	// About
	mux.HandleFunc("GET /api/marketing/about", marketingController.GetAllAboutContentHandler)
	mux.HandleFunc("GET /api/marketing/about/{id}", marketingController.GetAboutContentByIDHandler)

	// admin routes
	// Blog
	mux.Handle("GET /api/admin/blog", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllBlogPostsHandler)))
	mux.Handle("GET /api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostByIDHandler)))
	mux.Handle("GET /api/admin/blog/tag/{tag}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostsByTagHandler)))
	mux.Handle("POST /api/admin/blog", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateBlogPostHandler)))
	mux.Handle("PUT /api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateBlogPostHandler)))
	mux.Handle("DELETE /api/admin/blog/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteBlogPostHandler)))

	// Home
	mux.Handle("GET /api/admin/home", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllHomeContentHandler)))
	mux.Handle("GET /api/admin/home/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetHomeContentByIDHandler)))
	mux.Handle("POST /api/admin/home", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateHomeContentHandler)))
	mux.Handle("PUT /api/admin/home/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateHomeContentHandler)))
	mux.Handle("DELETE /api/admin/home/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteHomeContentHandler)))

	// GrooveJr
	mux.Handle("GET /api/admin/groovejr", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllGrooveJrContentHandler)))
	mux.Handle("GET /api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetGrooveJrContentByIDHandler)))
	mux.Handle("POST /api/admin/groovejr", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateGrooveJrContentHandler)))
	mux.Handle("PUT /api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateGrooveJrContentHandler)))
	mux.Handle("DELETE /api/admin/groovejr/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteGrooveJrContentHandler)))

	// About
	mux.Handle("GET /api/admin/about", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAllAboutContentHandler)))
	mux.Handle("GET /api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.GetAboutContentByIDHandler)))
	mux.Handle("POST /api/admin/about", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.CreateAboutContentHandler)))
	mux.Handle("PUT /api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.UpdateAboutContentHandler)))
	mux.Handle("DELETE /api/admin/about/{id}", auth.RequireAdminAuthV2(cookieJar, log, http.HandlerFunc(adminController.DeleteAboutContentHandler)))
}

func SetupRemotedevRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, adminController *admin.AdminController, marketingFileServer http.Handler) {
	auth.SetupAdminAuthV2(mux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer(), marketingFileServer)
}

func SetupLocalhostRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, adminController *admin.AdminController, marketingFileServer http.Handler) {
	auth.SetupAdminAuthV2(mux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer(), marketingFileServer)
}

func SetupProdRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, marketingController *marketing.MarketingController, adminController *admin.AdminController, oldSiteController *old_site.OldSiteController, marketingFileServer http.Handler) {
	log.Info().
		Msg("Setting up ProdRoutes")

	// Admin App (auth-protected, /admin/ prefix), Old Site (/old-site/), and Marketing App (/)
	// are all handled by SetupAdminAuth for consistency across environments.
	auth.SetupAdminAuthV2(mux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer(), auth.LogClientIp("/", log, marketingFileServer))
}

func _SetHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// this method wound up being superfluous for what we needed at the time of writing
		// but it's nice to have the infrastructure established

		// w.Header().Set("Content-Type", "text/javascript")
		// w.Header().Set("Content-Type", "text/html, text/css")
		handler.ServeHTTP(w, r)
	})
}
