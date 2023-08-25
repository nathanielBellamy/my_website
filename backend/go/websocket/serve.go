package websocket

import (
  "fmt"
	"net/http"
)

func ServeFeedWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("FEED Endpoint Hit")
      conn, err := Upgrade(w, r)
      if err != nil {
        fmt.Printf("\n :: FEED Upgrade ERROR :: \n")
        fmt.Fprintf(w, "%+v\n", err)
      }

      client := &Client{
        ID: pool.NewClientId(),
        Conn: &conn,
        Pool: pool,
      }

      WriteMessage(client.Conn, Message{ClientId: client.ID, Body: "connected"})

      pool.Register <- client
      client.ReadFeed()
}

func ServeWasmWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
      fmt.Println("WASM Endpoint Hit")
      conn, err := Upgrade(w, r)
      if err != nil {
        fmt.Printf("\n :: WASM Upgrade ERROR :: \n")
        fmt.Fprintf(w, "%+v\n", err)
      }

      client := &Client{
        ID: pool.NewClientId(),
        Conn: &conn,
        Pool: pool,
      }

      pool.Register <- client
      client.ReadWasm()
}
