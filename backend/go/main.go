package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/nathanielBellamy/my_website/backend/go/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("WebSocket Endpoint Hit")
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


    pool := websocket.NewPool()
    go pool.Start()
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
      serveWs(pool, w, r)
    })
}

func main() {
    fmt.Printf("Starting server on 8080")
    
    setupRoutes()

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
