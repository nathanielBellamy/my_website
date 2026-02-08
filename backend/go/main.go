package main

import (
	"net/http"
	"os"

	// "github.com/go-pg/pg/v10" // Removed unused import
	"github.com/nathanielBellamy/my_website/backend/go/admin"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/config"
	"github.com/nathanielBellamy/my_website/backend/go/db"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"
	"github.com/nathanielBellamy/my_website/backend/go/old_site"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

// MODE=<mode> ./main
func main() {
	// init log
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// determine runtime env
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "localhost"
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

	cookieJar := cmap.New[auth.Cookie]()

	log.Info().
		Msg("Establishing Routes")

	feedPool := websocket.NewPool(&log)
	wasmPool := websocket.NewPool(&log)
	go feedPool.StartFeed()
	go wasmPool.StartWasm()

	SetupRoutes(http.DefaultServeMux, &cookieJar, &log, feedPool, wasmPool, db.NewPgDBAdapter(dbClient))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().
			Msg("UnableToServe")
	}

	log.Info().
		Msg("Now serving on 8080")
}

func SetupRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, feedPool *websocket.Pool, wasmPool *websocket.Pool, db interfaces.PgxDB) {
	mode := os.Getenv("MODE")
	oldSiteController := old_site.NewOldSiteController(cookieJar, log, feedPool, wasmPool)

	marketingService := marketing.NewService(db)
	marketingController := marketing.NewMarketingController(log, marketingService)

	adminService := admin.NewService(db)
	adminController := admin.NewAdminController(log, adminService)

	SetupBaseRoutes(mux, cookieJar, log, oldSiteController, marketingController, adminController)

	if env.IsProd(mode) {
		SetupProdRoutes()
	} else if env.IsRemotedev(mode) {
		SetupRemotedevRoutes(mux, cookieJar, log, oldSiteController, adminController)
	} else {
		SetupLocalhostRoutes(mux, cookieJar, log, oldSiteController, adminController)
	}
}

func SetupBaseRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, marketingController *marketing.MarketingController, adminController *admin.AdminController) {
	log.Info().
		Msg("Setting up BaseRoutes")
	mode := os.Getenv("MODE")
	if env.IsProd(mode) {
		fs := http.FileServer(http.Dir("old-site"))
		mux.Handle("/", auth.LogClientIp("/", log, fs))
	}

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
	mux.Handle("GET /api/admin/blog", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetAllBlogPostsHandler)))
	mux.Handle("GET /api/admin/blog/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostByIDHandler)))
	mux.Handle("GET /api/admin/blog/tag/{tag}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetBlogPostsByTagHandler)))
	mux.Handle("POST /api/admin/blog", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.CreateBlogPostHandler)))
	mux.Handle("PUT /api/admin/blog/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.UpdateBlogPostHandler)))
	mux.Handle("DELETE /api/admin/blog/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.DeleteBlogPostHandler)))

	// Home
	mux.Handle("GET /api/admin/home", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetAllHomeContentHandler)))
	mux.Handle("GET /api/admin/home/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetHomeContentByIDHandler)))
	mux.Handle("POST /api/admin/home", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.CreateHomeContentHandler)))
	mux.Handle("PUT /api/admin/home/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.UpdateHomeContentHandler)))
	mux.Handle("DELETE /api/admin/home/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.DeleteHomeContentHandler)))

	// GrooveJr
	mux.Handle("GET /api/admin/groovejr", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetAllGrooveJrContentHandler)))
	mux.Handle("GET /api/admin/groovejr/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetGrooveJrContentByIDHandler)))
	mux.Handle("POST /api/admin/groovejr", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.CreateGrooveJrContentHandler)))
	mux.Handle("PUT /api/admin/groovejr/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.UpdateGrooveJrContentHandler)))
	mux.Handle("DELETE /api/admin/groovejr/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.DeleteGrooveJrContentHandler)))

	// About
	mux.Handle("GET /api/admin/about", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetAllAboutContentHandler)))
	mux.Handle("GET /api/admin/about/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.GetAboutContentByIDHandler)))
	mux.Handle("POST /api/admin/about", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.CreateAboutContentHandler)))
	mux.Handle("PUT /api/admin/about/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.UpdateAboutContentHandler)))
	mux.Handle("DELETE /api/admin/about/{id}", auth.RequireDevAuth(cookieJar, log, http.HandlerFunc(adminController.DeleteAboutContentHandler)))
}

func SetupRemotedevRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, adminController *admin.AdminController) {
	auth.SetupDevAuth(mux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer())
}

func SetupLocalhostRoutes(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, adminController *admin.AdminController) {
	auth.SetupDevAuth(mux, cookieJar, log, oldSiteController.OldSiteFileServer(), adminController.AdminFileServer())
}

func SetupProdRoutes() {
	// TODO: maybe Set cookie when user goes through ep warning
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
