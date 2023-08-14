package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/nathanielBellamy/my_website/backend/go/auth"
    "github.com/nathanielBellamy/my_website/backend/go/env"
    "github.com/nathanielBellamy/my_website/backend/go/websocket"
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

func setupDevAuth(cookieJar *auth.CookieJar) {
  fs_auth := http.FileServer(http.Dir("./../../auth/dev_auth/dist"))
  http.Handle("/", fs_auth)
  
  http.HandleFunc("/dev-auth", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("\n dev-auth %v \n", r.Method)
    auth.HandleDev(&w, r, cookieJar)
    http.Redirect(w, r, "/dev", 301)
  })

  http.HandleFunc("/dev", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("\n dev %v \n", r.Method)
    if (*cookieJar).ValidateSessionCookie(r) {
      fmt.Printf("\n valid foo cookie \n")
    }
  })
}

func setupRemotedevRoutes(cookieJar *auth.CookieJar) {
  setupDevAuth(cookieJar)

}

func setupLocalhostRoutes(cookieJar *auth.CookieJar) {
  setupDevAuth(cookieJar)
}

func setupProdRoutes() {
  // TODO: maybe set cookie when user goes through ep warning
}

func redirectToDevAuth(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w,r,"/dev-auth", 301)
}

func setupBaseRoutes(runtime_env env.Env, cookieJar *auth.CookieJar) {
  if runtime_env.IsProd() {
    fs := http.FileServer(http.Dir("./../../frontend/dist"))
    http.Handle("/", fs)
  }

  feedPool := websocket.NewPool()
  wasmPool := websocket.NewPool()
  go feedPool.StartFeed()
  go wasmPool.StartWasm()
  // TODO: verify client cookie before serving in remotedev MODE
  http.HandleFunc("/public-square-feed-ws", func(w http.ResponseWriter, r *http.Request) {
    if !runtime_env.IsProd() {
      // localhost and remote dev require basic login
      res := (*cookieJar).ValidateSessionCookie(r)
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
      res := (*cookieJar).ValidateSessionCookie(r)
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

func setupRoutes(runtime_env env.Env, cookieJar *auth.CookieJar) {
    switch runtime_env.Mode {
    case "localhost":
      setupLocalhostRoutes(cookieJar)
    case "prod":
      setupProdRoutes()
    case "remotedev":
      setupRemotedevRoutes(cookieJar)
    }

    setupBaseRoutes(runtime_env, cookieJar)

    // fs := http.FileServer(http.Dir("./../../frontend/dist"))
    // http.Handle("/", fs)

    // feedPool := websocket.NewPool()
    // wasmPool := websocket.NewPool()
    // go feedPool.StartFeed()
    // go wasmPool.StartWasm()
    // // TODO: verify client cookie before serving in remotedev MODE
    // http.HandleFunc("/public-square-feed-ws", func(w http.ResponseWriter, r *http.Request) {
    //   serveFeedWs(feedPool, w, r)
    // })
    // http.HandleFunc("/public-square-wasm-ws", func(w http.ResponseWriter, r *http.Request) {
    //   serveWasmWs(wasmPool, w, r)
    // })
}

func main() {
    fmt.Printf("Starting server on 8080 \n")
    
    mode := os.Getenv("MODE")
    runtime_env := env.Env {
      Mode: string(mode),
    }
    var cookieJar auth.CookieJar
    setupRoutes(runtime_env, &cookieJar)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("UnableToServe")
        log.Fatal(err)
    }
}
