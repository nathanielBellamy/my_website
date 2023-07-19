package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/nathanielBellamy/my_website/backend/go/websocket"
)

func serveFeedWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("WebSocket FEED Endpoint Hit")
      conn, err := websocket.Upgrade(w, r)
      if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
      }

      client := &websocket.Client{
        ID: pool.NewClientId(),
        Conn: &conn,
        Pool: pool,
      }

      websocket.WriteMessage(client.Conn, websocket.Message{ClientId: client.ID, Body: "connected"})

      pool.Register <- client
      client.Read()
}

func serveWasmWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("WebSocket WASM Endpoint Hit")
      conn, err := websocket.Upgrade(w, r)
      if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
      }

      client := &websocket.Client{
        ID: pool.NewClientId(),
        Conn: &conn,
        Pool: pool,
      }

      pool.Register <- client
      client.Read()
}

func setupRoutes() {
    fs := http.FileServer(http.Dir("./../../frontend/dist"))
    http.Handle("/", fs)

    feedPool := websocket.NewPool()
    wasmPool := websocket.NewPool()
    go feedPool.Start()
    go wasmPool.Start()
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
        log.Fatal(err)
    }
}
