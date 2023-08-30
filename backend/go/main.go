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
)

func main() {
    // setup logging
  	file, err := os.Create("log.txt")
    if err != nil {
      fmt.Printf("Failed creating log file: %s", err)
    }
    log := zerolog.New(file).With().Timestamp().Logger()
    log.Info().
        Msg("Starting server on 8080")
    
    mode := os.Getenv("MODE")
    runtime_env := env.Env {
      Mode: string(mode),
    }
    cookieJar := cmap.New[bool]()
    SetupRoutes(runtime_env, &cookieJar, &log)

    if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal().
          Msg("UnableToServe")
    }
}

func SetupRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
    switch runtime_env.Mode {
    case "localhost":
      SetupLocalhostRoutes(runtime_env, cookieJar, log)
    case "prod":
      SetupProdRoutes()
    case "remotedev":
      SetupRemotedevRoutes(runtime_env, cookieJar, log)
    }

    SetupBaseRoutes(runtime_env, cookieJar, log)
}

func SetupBaseRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
  if runtime_env.IsProd() {
    fs := http.FileServer(http.Dir("frontend"))
    http.Handle("/", auth.LogClientIp("/", log, fs) )
  }

  if !runtime_env.IsLocalhost() {
    // setup recaptcha
    http.HandleFunc("/ps-recaptcha", func(w http.ResponseWriter, r *http.Request) {
      fmt.Printf("recaptcha endpoint hit")
      // TODO
      // Read token from body
      // Verify token
    })
  }

  feedPool := websocket.NewPool(log)
  wasmPool := websocket.NewPool(log)
  go feedPool.StartFeed()
  go wasmPool.StartWasm()
  http.HandleFunc("/public-square-feed-ws", func(w http.ResponseWriter, r *http.Request) {
    if !runtime_env.IsProd() {
      // localhost and remote dev require basic login
      res := auth.HasValidCookie(runtime_env, r, cookieJar, log)
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
    if !runtime_env.IsProd() {
      // localhost and remote dev require basic login
      res := auth.HasValidCookie(runtime_env, r, cookieJar, log)
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

func SetupRemotedevRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
  auth.SetupDevAuth(runtime_env, cookieJar, log)

}

func SetupLocalhostRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) {
  auth.SetupDevAuth(runtime_env, cookieJar, log)
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
