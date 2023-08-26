package websocket

import (
	"net"

	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog"
)

type ClientName struct {
    Alias string
    First string
    Second string
}

type Client struct {
    // ID = 0 indicates system
    ID   uint
    Conn *net.Conn
    Pool *Pool
    Name ClientName
    Log *zerolog.Logger
}

type Message struct {
    ClientId uint `json:"clientId"`
    Body string `json:"body"`
}

func (c *Client) ReadFeed() {
    defer func() {
        c.Pool.Unregister <- c
        (*c.Conn).Close()
    }()

    for {
        msg, _, err := wsutil.ReadClientData(*c.Conn)
        if err != nil {
            c.Log.Error().
                  Err(err).
                  Msg("ReadClientData FEED Error")
            return
        }
        message := Message{ClientId: c.ID, Body: string(msg)}
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
            c.Log.Error().
                  Err(err).
                  Msg("ReadClientData WASM Error")
            return
        } else {
          c.Pool.BroadcastSettings <- msg
        }
    }
}



