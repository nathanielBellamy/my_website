package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/nathanielBellamy/my_website/backend/go/websocket"
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
      fmt.Printf("UpgradeHTTP Error")
      fmt.Println(err)
    }
    go func() {
      defer conn.Close()
      websocket.Reader(&conn)
    }()
}

func servWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("WebSocket Endpoint Hit")
      conn, err := websocket.Upgrade(w, r)
      if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
      }

      client := &websocket.Client{
        Conn: &conn,
        Pool: pool,
      }

      pool.Register <- client
      client.Read()
}

func setupRoutes() {
    fs := http.FileServer(http.Dir("./../../frontend/dist"))
    http.Handle("/", fs)

    src := http.HandlerFunc(handleWebSocket)
    http.Handle("/ws", src)
}

func main() {
    fmt.Printf("Starting server on 8080")
    
    setupRoutes()

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
