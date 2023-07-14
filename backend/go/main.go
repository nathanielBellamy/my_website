package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gobwas/ws"
    "github.com/gobwas/ws/wsutil"
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, _, _, err := ws.UpgradeHTTP(r, w)
    if err != nil {
      fmt.Printf("UpgradeHTTP Error")
    }
    go func() {
      defer conn.Close()
      centralMem := []byte("CENT MEM")
      for {
        msg, op, err := wsutil.ReadClientData(conn)
        if err != nil {
          fmt.Printf("ReadClientData Error")
          conn.Close()
          return
        }
        centralMem = append(centralMem, []byte("+")...)
        msg = append(msg, centralMem...)
        err = wsutil.WriteServerMessage(conn, op, msg)
        if err != nil {
          fmt.Printf("WriteServerMessage Error")
          conn.Close()
          return
        }
      }
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
