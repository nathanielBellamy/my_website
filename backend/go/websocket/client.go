package websocket

import (
	"math/big"
	"net"

	"github.com/gobwas/ws/wsutil"
)

type ClientName struct {
	Alias  string
	First  string
	Second string
}

type Client struct {
	ID   big.Int
	IP   string
	Conn *net.Conn
	Pool *Pool
	Name ClientName
}

type Message struct {
	ClientId uint   `json:"clientId"`
	Body     string `json:"body"`
	System   bool   `json:"system"`
}

var ReadFeed = func(c *Client) {
	defer func() {
		c.Pool.Unregister <- c
		_ = (*c.Conn).Close()
	}()

	for {
		msg, _, err := wsutil.ReadClientData(*c.Conn)
		if err != nil {
			c.Pool.Log.Error().
				Err(err).
				Msg("ReadClientData FEED Error")
			return
		}
		msg_str := string(msg)
		c.Pool.Log.Info().
			Uint("client_id_feed", uint(c.ID.Uint64())).
			Str("ip", c.IP).
			Str("val", msg_str).
			Msg("USER MESSAGE")
		message := Message{ClientId: uint(c.ID.Uint64()), Body: msg_str}
		c.Pool.Broadcast <- message
	}
}

var ReadWasm = func(c *Client) {
	defer func() {
		c.Pool.Unregister <- c
		_ = (*c.Conn).Close()
	}()

	for {
		msg, _, err := wsutil.ReadClientData(*c.Conn)
		if err != nil {
			c.Pool.Log.Error().
				Err(err).
				Msg("ReadClientData WASM Error")
			return
		} else {
			c.Pool.BroadcastSettings <- msg
		}
	}
}
