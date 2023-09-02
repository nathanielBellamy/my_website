package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
  "github.com/joho/godotenv"
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

    cookieJar := cmap.New[bool]()

    log.Info().
        Msg("Establishing Routes")

    SetupRoutes(&cookieJar, &log)

    if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal().
          Msg("UnableToServe")
    }

    log.Info().
        Msg("Now serving on 8080")
}

func SetupRoutes(cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
    mode := os.Getenv("MODE")
    if env.IsProd(mode) {
      SetupProdRoutes()
    } else if env.IsRemotedev(mode) {
      SetupRemotedevRoutes(cookieJar, log)
    } else {
      SetupLocalhostRoutes(cookieJar, log)
    }

    SetupBaseRoutes(cookieJar, log)
}

func SetupBaseRoutes(cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
  mode := os.Getenv("MODE")
  if env.IsProd(mode) {
    fs := http.FileServer(http.Dir("frontend"))
    http.Handle("/", auth.LogClientIp("/", log, fs) )
  }

  // setup recaptcha
  http.HandleFunc("/recaptcha", func (w http.ResponseWriter, r *http.Request) {
    // auth.LogClientIp("/recaptcha", log, http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {}))
    ip := auth.GetClientIpAddr(r)
    log.Info().
        Str("ip", ip).
        Msg("Recaptcha Endpoint Hit")
    res := auth.ValidateRecaptcha(r, log)
    log.Info().
        Str("ip", ip).
        Bool("res", res).
        Msg("ValidateRecaptcha res")
  })

  feedPool := websocket.NewPool(log)
  wasmPool := websocket.NewPool(log)
  go feedPool.StartFeed()
  go wasmPool.StartWasm()
  http.HandleFunc("/public-square-feed-ws", func(w http.ResponseWriter, r *http.Request) {
    if !env.IsProd(mode) {
      // localhost and remote dev require basic login
      res := auth.HasValidCookie(r, cookieJar, log)
      if res {
        websocket.ServeFeedWs(feedPool, w, r, log)
      } else {
        auth.RedirectToDevAuth(w, r, log)
      }
    } else {
      // prod is public
      websocket.ServeFeedWs(feedPool, w, r, log)
    }
  })
  http.HandleFunc("/public-square-wasm-ws", func(w http.ResponseWriter, r *http.Request) {
    mode := os.Getenv("MODE")
    if !env.IsProd(mode) {
      // localhost and remote dev require basic login
      res := auth.HasValidCookie(r, cookieJar, log)
      if res {
        websocket.ServeWasmWs(wasmPool, w, r, log)
      } else {
        auth.RedirectToDevAuth(w, r, log)
      }
    } else {
      // prod is public
      websocket.ServeWasmWs(wasmPool, w, r, log)
    }
  })
}

func SetupRemotedevRoutes(cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
  auth.SetupDevAuth(cookieJar, log)
}

func SetupLocalhostRoutes(cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
  auth.SetupDevAuth(cookieJar, log)
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
    handler.ServeHTTP(w,r)
  })
}
