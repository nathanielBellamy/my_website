package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/nathanielBellamy/my_website/backend/go/websocket"
    "github.com/nathanielBellamy/my_website/backend/go/auth"
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

func setupRoutes() {
    if os.Getenv("MODE") == "remotedev" {
      http.HandleFunc("/dev-auth", auth.HandleDev)
    }

    fs := http.FileServer(http.Dir("./../../frontend/dist"))
    http.Handle("/", fs)

    feedPool := websocket.NewPool()
    wasmPool := websocket.NewPool()
    go feedPool.StartFeed()
    go wasmPool.StartWasm()
    // TODO: verify client token before serving in remotedev MODE
    http.HandleFunc("/public-square-feed-ws", func(w http.ResponseWriter, r *http.Request) {
      serveFeedWs(feedPool, w, r)
    })
    http.HandleFunc("/public-square-wasm-ws", func(w http.ResponseWriter, r *http.Request) {
      serveWasmWs(wasmPool, w, r)
    })
}

func main() {
    fmt.Printf("Starting server on 8080 \n")
    
    setupRoutes()

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("UnableToServe")
        log.Fatal(err)
    }
}
