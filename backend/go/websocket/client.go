package websocket

import (
	"net"

	"github.com/gobwas/ws/wsutil"
)

type ClientName struct {
    Alias string
    First string
    Second string
}

type Client struct {
    // ID = 0 indicates system
    ID   uint
    IP   string
    Conn *net.Conn
    Pool *Pool
    Name ClientName
}

type Message struct {
    ClientId uint `json:"clientId"`
    Body string `json:"body"`
    System bool `json:"system"`
}

func (c *Client) ReadFeed() {
    defer func() {
        c.Pool.Unregister <- c
        (*c.Conn).Close()
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
                   Uint("client_id_feed", c.ID).
                   Str("ip", c.IP).
                   Str("val", msg_str).
                   Msg("USER MESSAGE")
        message := Message{ClientId: c.ID, Body: msg_str}
        c.Pool.Broadcast <- message
    }
}

func (c *Client) ReadWasm() {
    defer func() {
        c.Pool.Unregister <- c
        (*c.Conn).Close()
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



