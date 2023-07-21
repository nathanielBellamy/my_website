package websocket

import "fmt"
import "math/rand"
import "time"

func randInRange(min int, max int) (int){
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min + 1) + min
}

type Pool struct {
    NextClientId uint
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    Broadcast  chan Message
}

func NewPool() *Pool {
    return &Pool{
        NextClientId: uint(randInRange(48593, 915740)),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan Message),
    }
}

func (pool *Pool) NewClientId() (uint) {
    newId := pool.NextClientId
    pool.NextClientId += 123
  	return newId
}

func (pool *Pool) Start() {
    for {
      select {
        case client := <-pool.Register:
            pool.Clients[client] = true
            id := client.ID
            fmt.Printf("Size of Connection Pool: %v \n", len(pool.Clients))
            for client, _ := range pool.Clients {
                fmt.Printf("ClientId %v Connected \n", id)
                messageBody := fmt.Sprintf("👋 User %v Joined 👋", id)
                message := Message{ClientId: 0, Body: messageBody}
                WriteMessage(client.Conn, message)
            }
            break
        case client := <-pool.Unregister:
            id := client.ID
            messageBody := fmt.Sprintf("🫡 User %v Disconnected 🫡", id)
            delete(pool.Clients, client)
            fmt.Printf("Size of Connection Pool: %v \n", len(pool.Clients))
            for client, _ := range pool.Clients {
              WriteMessage(client.Conn, Message{ClientId: 0, Body: messageBody})
            }
            break
        case message := <-pool.Broadcast:
            fmt.Printf("Sending message to all clients in Pool \n")
            fmt.Printf("%v", message)
            for client, _ := range pool.Clients {
                WriteMessage(client.Conn, message)            
            }
      }
    }
}

