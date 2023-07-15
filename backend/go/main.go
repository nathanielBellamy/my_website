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

func main() {
    fmt.Printf("Starting server on 8080")

    fs := http.FileServer(http.Dir("./../../frontend/dist"))
    http.Handle("/", fs)

    src := http.HandlerFunc(handleWebSocket)
    http.Handle("/ws", src)
    // http.Handle("/public-square", handleWebSocket)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
