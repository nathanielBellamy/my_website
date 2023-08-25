package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
  "github.com/rs/zerolog"
  // "github.com/rs/zerolog/log"
)

func main() {
    // setup logging
  	file, err := os.Create("log.txt")
    if err != nil {
      fmt.Printf("Failed creating file: %s", err)
    }
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log := zerolog.New(file)
    log.Info().Msg("Starting server on 8080 \n")
    
    mode := os.Getenv("MODE")
    runtime_env := env.Env {
      Mode: string(mode),
    }
    cookieJar := cmap.New[bool]()
    SetupRoutes(runtime_env, &cookieJar)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("UnableToServe")
        // log.Fatal(err)
    }
}

func SetupRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool]) {
    switch runtime_env.Mode {
    case "localhost":
      SetupLocalhostRoutes(runtime_env, cookieJar)
    case "prod":
      SetupProdRoutes()
    case "remotedev":
      SetupRemotedevRoutes(runtime_env, cookieJar)
    }

    SetupBaseRoutes(runtime_env, cookieJar)
}

func SetupBaseRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool]) {
  if runtime_env.IsProd() {
    fs := http.FileServer(http.Dir("frontend"))
    http.Handle("/", fs)
  }

  feedPool := websocket.NewPool()
  wasmPool := websocket.NewPool()
  go feedPool.StartFeed()
  go wasmPool.StartWasm()
  http.HandleFunc("/public-square-feed-ws", func(w http.ResponseWriter, r *http.Request) {
    if !runtime_env.IsProd() {
      // localhost and remote dev require basic login
      res := auth.HasValidCookie(runtime_env, r, cookieJar)
      if res {
        websocket.ServeFeedWs(feedPool, w, r)
      } else {
        auth.RedirectToDevAuth(w,r)
      }
    } else {
      // prod is public
      websocket.ServeFeedWs(feedPool, w, r)
    }
  })
  http.HandleFunc("/public-square-wasm-ws", func(w http.ResponseWriter, r *http.Request) {
    if !runtime_env.IsProd() {
      // localhost and remote dev require basic login
      res := auth.HasValidCookie(runtime_env, r, cookieJar)
      if res {
        websocket.ServeWasmWs(wasmPool, w, r)
      } else {
        auth.RedirectToDevAuth(w,r)
      }
    } else {
      // prod is public
      websocket.ServeWasmWs(wasmPool, w, r)
    }
  })
}

func SetupRemotedevRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool]) {
  auth.SetupDevAuth(runtime_env, cookieJar)

}

func SetupLocalhostRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool]) {
  auth.SetupDevAuth(runtime_env, cookieJar)
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


