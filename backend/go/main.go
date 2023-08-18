package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
)

func serveFeedWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("WebSocket FEED Endpoint Hit")
      conn, err := websocket.Upgrade(w, r)
      if err != nil {
        fmt.Printf("FEED Upgrade ERROR \n")
        fmt.Fprintf(w, "%+v\n", err)
      }

      client := &websocket.Client{
        ID: pool.NewClientId(),
        Conn: &conn,
        Pool: pool,
      }

      websocket.WriteMessage(client.Conn, websocket.Message{ClientId: client.ID, Body: "connected"})

      pool.Register <- client
      client.ReadFeed()
}

func serveWasmWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("WebSocket WASM Endpoint Hit")
      conn, err := websocket.Upgrade(w, r)
      if err != nil {
        fmt.Printf("WASM Upgrade ERROR \n")
        fmt.Fprintf(w, "%+v\n", err)
      }

      client := &websocket.Client{
        ID: pool.NewClientId(),
        Conn: &conn,
        Pool: pool,
      }

      pool.Register <- client
      client.ReadWasm()
}

func setupDevAuth(cookieJar *cmap.ConcurrentMap[string, bool]) {
  fs_auth := http.FileServer(http.Dir("auth/dev"))
  http.Handle("/auth/dev",  http.StripPrefix("/auth/dev", fs_auth))
  
  fs_frontend := http.FileServer(http.Dir("frontend"))
  http.Handle("/", requireDevAuth(cookieJar, fs_frontend))
  

  http.HandleFunc("/auth/dev/dev-auth", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("\n dev-auth %v \n", r.Method)
    if auth.ValidateDev(w, r, cookieJar) {
      fmt.Printf("yoooo")
      http.Redirect(w, r, "./../../", 301)
    } else {
      http.Error(w, "Invalid Password", 503)
    }
  })
}

func setHeaders(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // this method wound up being superfluous for what we needed at the time of writing
    // but it's nice to have the infrastructure established
    
    // w.Header().Set("Content-Type", "text/javascript")
    // w.Header().Set("Content-Type", "text/html, text/css")
    handler.ServeHTTP(w,r)
  })
}

func requireDevAuth(cookieJar *cmap.ConcurrentMap[string, bool], handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if true {// auth.ValidateDev(w, r, cookieJar){
          handler.ServeHTTP(w, r)
        } else {
          // http.Error(w, "Dev Not Validated", 503)
          // fmt.Printf("fooooo")
          redirectToDevAuth(w, r)
          return
        }
    })
}

func setupRemotedevRoutes(cookieJar *cmap.ConcurrentMap[string, bool]) {
  setupDevAuth(cookieJar)

}

func setupLocalhostRoutes(cookieJar *cmap.ConcurrentMap[string, bool]) {
  setupDevAuth(cookieJar)
}

func setupProdRoutes() {
  // TODO: maybe set cookie when user goes through ep warning
}

func redirectToDevAuth(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("\n redirectToDevAuth \n")
  http.Redirect(w,r,"/dev/auth", 301)
}

func setupBaseRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool]) {
  if runtime_env.IsProd() {
    fs := http.FileServer(http.Dir("./../../frontend/dist"))
    http.Handle("/", fs)
  }

  feedPool := websocket.NewPool()
  wasmPool := websocket.NewPool()
  go feedPool.StartFeed()
  go wasmPool.StartWasm()
  http.HandleFunc("/public-square-feed-ws", func(w http.ResponseWriter, r *http.Request) {
    if !runtime_env.IsProd() {
      // localhost and remote dev require basic login
      res := auth.ValidateSessionCookie(r, cookieJar)
      if res {
        serveFeedWs(feedPool, w, r)
      } else {
        redirectToDevAuth(w,r)
      }
    } else {
      // prod is public
      serveFeedWs(feedPool, w, r)
    }
  })
  http.HandleFunc("/public-square-wasm-ws", func(w http.ResponseWriter, r *http.Request) {
    if !runtime_env.IsProd() {
      // localhost and remote dev require basic login
      res := auth.ValidateSessionCookie(r, cookieJar)
      if res {
        serveWasmWs(wasmPool, w, r)
      } else {
        redirectToDevAuth(w,r)
      }
    } else {
      // prod is public
      serveWasmWs(wasmPool, w, r)
    }
  })
}

func setupRoutes(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool]) {
    switch runtime_env.Mode {
    case "localhost":
      setupLocalhostRoutes(cookieJar)
    case "prod":
      setupProdRoutes()
    case "remotedev":
      setupRemotedevRoutes(cookieJar)
    }

    setupBaseRoutes(runtime_env, cookieJar)
}

func main() {
    fmt.Printf("Starting server on 8080 \n")
    
    mode := os.Getenv("MODE")
    runtime_env := env.Env {
      Mode: string(mode),
    }
    cookieJar := cmap.New[bool]()
    setupRoutes(runtime_env, &cookieJar)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("UnableToServe")
        log.Fatal(err)
    }
}
