package websocket

import (
	"fmt"
	"log"
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
    Conn *net.Conn
    Pool *Pool
    Name ClientName
}

type Message struct {
    ClientId uint `json:"clientId"`
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
        message := Message{ClientId: c.ID, Body: string(msg)}
        c.Pool.Broadcast <- message
        fmt.Printf("Message Received: %+v\n", message)
    }
}


