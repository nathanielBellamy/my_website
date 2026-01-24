package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/marketing"
	"github.com/nathanielBellamy/my_website/backend/go/old_site"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

// MODE=<mode> ./main
func main() {
	// init log
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Printf("Failed creating log file: %s", err)
	}
	log := zerolog.New(file).With().Timestamp().Logger()

	// determine runtime env
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "localhost"
	}

	// read env file
	log.Info().
		Msg("Loading ENV")

	envErr := godotenv.Load(".env." + mode)
	if envErr != nil {
		log.Fatal().
			Err(envErr).
			Msg("Error loading .env file")
	}

	log.Info().
		Str("mode", mode).
		Msg("Runtime Env")

	cookieJar := cmap.New[auth.Cookie]()

	log.Info().
		Msg("Establishing Routes")

	feedPool := websocket.NewPool(&log)
	wasmPool := websocket.NewPool(&log)
	go feedPool.StartFeed()
	go wasmPool.StartWasm()

	SetupRoutes(&cookieJar, &log, feedPool, wasmPool)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().
			Msg("UnableToServe")
	}

	log.Info().
		Msg("Now serving on 8080")
}

func SetupRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, feedPool *websocket.Pool, wasmPool *websocket.Pool) {
	mode := os.Getenv("MODE")
	oldSiteController := old_site.NewOldSiteController(cookieJar, log, feedPool, wasmPool)
	marketingController := marketing.NewMarketingController(log)

	if env.IsProd(mode) {
		SetupProdRoutes()
	} else if env.IsRemotedev(mode) {
		SetupRemotedevRoutes(cookieJar, log, oldSiteController)
	} else {
		SetupLocalhostRoutes(cookieJar, log, oldSiteController)
	}

	SetupBaseRoutes(cookieJar, log, oldSiteController, marketingController)
}

func SetupBaseRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController, marketingController *marketing.MarketingController) {
	log.Info().
		Msg("Setting up BaseRoutes")
	mode := os.Getenv("MODE")
	if env.IsProd(mode) {
		fs := http.FileServer(http.Dir("old-site"))
		http.Handle("/", auth.LogClientIp("/", log, fs))
	}

	// old-site routes
	http.HandleFunc("/old-site/recaptcha", oldSiteController.RecaptchaHandler)
	http.HandleFunc("/old-site/public-square-feed-ws", oldSiteController.PublicSquareFeedWsHandler)
	http.HandleFunc("/old-site/public-square-wasm-ws", oldSiteController.PublicSquareWasmWsHandler)

	// marketing routes
	// Blog
	http.HandleFunc("/api/marketing/blog", marketingController.GetAllBlogPostsHandler)
	http.HandleFunc("/api/marketing/blog/{id}", marketingController.GetBlogPostByIDHandler)
	http.HandleFunc("/api/marketing/blog/tag/{tag}", marketingController.GetBlogPostsByTagHandler)
	http.HandleFunc("/api/marketing/blog/date/{date}", marketingController.GetBlogPostsByDateHandler)

	// Home
	http.HandleFunc("/api/marketing/home", marketingController.GetAllHomeContentHandler)
	http.HandleFunc("/api/marketing/home/{id}", marketingController.GetHomeContentByIDHandler)

	// GrooveJr
	http.HandleFunc("/api/marketing/groovejr", marketingController.GetAllGrooveJrContentHandler)
	http.HandleFunc("/api/marketing/groovejr/{id}", marketingController.GetGrooveJrContentByIDHandler)

	// About
	http.HandleFunc("/api/marketing/about", marketingController.GetAllAboutContentHandler)
	http.HandleFunc("/api/marketing/about/{id}", marketingController.GetAboutContentByIDHandler)
}

func SetupRemotedevRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController) {
	auth.SetupDevAuth(cookieJar, log, oldSiteController.OldSiteFileServer())
}

func SetupLocalhostRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, oldSiteController *old_site.OldSiteController) {
	auth.SetupDevAuth(cookieJar, log, oldSiteController.OldSiteFileServer())
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
