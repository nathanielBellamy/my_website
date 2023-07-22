package websocket

import (
	"fmt"
	// "log"
	"net"

	// "github.com/gobwas/ws"
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

func (c *Client) ReadFeed() {
    defer func() {
        c.Pool.Unregister <- c
        (*c.Conn).Close()
    }()

    for {
        msg, _, err := wsutil.ReadClientData(*c.Conn)
        if err != nil {
            fmt.Printf(" ReadClientData FEED Error \n")
            fmt.Printf(" err %v", err)
            return
        }
        message := Message{ClientId: c.ID, Body: string(msg)}
        c.Pool.Broadcast <- message
        fmt.Printf("Message Received: %+v\n", message)
    }
}

func (c *Client) ReadWasm() {
    defer func() {
        c.Pool.Unregister <- c
        (*c.Conn).Close()
    }()

    for {
        msg, opCode, err := wsutil.ReadClientData(*c.Conn)

        fmt.Printf(" ReadClientData WASM \n")
        fmt.Printf(" -----> %v \n", msg)
        fmt.Printf(" -----> %v \n", opCode)
        if err != nil {
            fmt.Printf(" ReadClientData WASM Error \n")
            fmt.Printf("err: %v", err)
            return
        } else {
          message := Message{ClientId: c.ID, Body: string(msg)}
          c.Pool.Broadcast <- message
          fmt.Printf("Message Received: %+v\n", message)
        }
    }
}



