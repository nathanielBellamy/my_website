package websocket

import (
	"net/http"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
)

var ServeFeedWs = func(pool *Pool, w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
	conn, err := Upgrade(w, r, log)
	if err != nil {
		log.Error().
			Err(err).
			Msg("FEED Upgrade ERROR")
		return
	}

	id, idErr := pool.NewClientId()
	if idErr != nil {
		log.Error().Err(idErr).Msg("Websocket Pool, NewClientId ERROR")
	}

	client := &Client{
		ID:   id,
		IP:   auth.GetClientIpAddr(r),
		Conn: &conn,
		Pool: pool,
	}

	log.Info().
		Str("ip", client.IP).
		Uint("client_id_feed", uint(client.ID.Uint64())).
		Msg("FEED Endpoint Hit")

	WriteMessage(client.Conn, Message{ClientId: uint(client.ID.Uint64()), Body: "connected"}, log)

	pool.Register <- client
	ReadFeed(client)
}

var ServeWasmWs = func(pool *Pool, w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
	conn, err := Upgrade(w, r, log)
	if err != nil {
		log.Error().
			Err(err).
			Msg("WASM Upgrade ERROR")
	}

	id, idErr := pool.NewClientId()
	if idErr != nil {
		log.Error().Err(idErr).Msg("Websocket Pool, NewClientId ERROR")
	}

	client := &Client{
		ID:   id,
		IP:   auth.GetClientIpAddr(r),
		Conn: &conn,
		Pool: pool,
	}

	log.Info().
		Str("ip", client.IP).
		Uint("client_id_wasm", uint(client.ID.Uint64())).
		Msg("WASM Endpoint Hit")

	pool.Register <- client
	ReadWasm(client)
}
