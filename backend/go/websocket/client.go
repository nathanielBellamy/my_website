package websocket

import (
	"fmt"
	"log"
	"net"
	// "sync"

	// "github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Client struct {
    ID   string
    Conn *net.Conn
    Pool *Pool
}

type Message struct {
    Type int    `json:"type"`
    Body string `json:"body"`
}

func (c *Client) Read() {
    defer func() {
        c.Pool.Unregister <- c
        (*c.Conn).Close()
    }()

    for {
        msg, _, err := wsutil.ReadClientData(*c.Conn)
        if err != nil {
            log.Println(err)
            return
        }
        message := Message{Type: 1, Body: string(msg)}
        c.Pool.Broadcast <- message
        fmt.Printf("Message Received: %+v\n", message)
    }
}
