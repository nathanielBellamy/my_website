package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gobwas/ws"
	  "github.com/gobwas/ws/wsutil"
)

func main() {
    fmt.Printf("Starting server on 8080")

    fs := http.FileServer(http.Dir("./../../frontend/dist/"))
    http.Handle("/", fs)

    ws := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      conn, _, _, err := ws.UpgradeHTTP(r, w)
      if err != nil {
        fmt.Printf("UpgradeHTTP Error")
      }
      go func() {
        defer conn.Close()

        for {
          msg, op, err := wsutil.ReadClientData(conn)
          if err != nil {
            fmt.Printf("ReadClientData Error")
          }
          err = wsutil.WriteServerMessage(conn, op, msg)
          if err != nil {
            fmt.Printf("WriteServerMessage Error")
          }
        }
      }()
    })
    http.Handle("/public-square", ws)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
