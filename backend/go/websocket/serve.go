package websocket

import (
	"net/http"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
)

func ServeFeedWs(pool *Pool, w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
      conn, err := Upgrade(w, r, log)
      if err != nil {
        log.Error().
            Err(err).
            Msg("FEED Upgrade ERROR")
      }

      client := &Client{
        ID: pool.NewClientId(),
        IP: auth.GetClientIpAddr(r),
        Conn: &conn,
        Pool: pool,
      }
      
      log.Info().
          Str("ip", client.IP).
          Uint("client_id_feed", client.ID).
          Msg("FEED Endpoint Hit")

      WriteMessage(client.Conn, Message{ClientId: client.ID, Body: "connected"}, log)

      pool.Register <- client
      client.ReadFeed()
}

func ServeWasmWs(pool *Pool, w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
      conn, err := Upgrade(w, r, log)
      if err != nil {
        log.Error().
            Err(err).
            Msg("WASM Upgrade ERROR")
      }

      client := &Client{
        ID: pool.NewClientId(),
        IP: auth.GetClientIpAddr(r),
        Conn: &conn,
        Pool: pool,
      }

      log.Info().
          Str("ip", client.IP).
          Uint("client_id_wasm", client.ID).
          Msg("WASM Endpoint Hit")

      pool.Register <- client
      client.ReadWasm()
}
