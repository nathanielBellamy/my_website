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
    BroadcastSettings chan []uint8
}

func NewPool() *Pool {
    return &Pool{
        NextClientId: uint(randInRange(48593, 915740)),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan Message),
        BroadcastSettings:  make(chan []uint8),
    }
}

func (pool *Pool) NewClientId() (uint) {
    newId := pool.NextClientId
    pool.NextClientId += 123
  	return newId
}

func (pool *Pool) StartFeed() {
    for {
      select {
        case client := <-pool.Register:
            pool.Clients[client] = true
            id := client.ID
            fmt.Printf("Size of Connection Pool: %v \n", len(pool.Clients))
            for client, _ := range pool.Clients {
                fmt.Printf("ClientId %v Connected \n", id)
                messageBody := fmt.Sprintf("u-%v 👋", id)
                message := Message{ClientId: 0, Body: messageBody}
                WriteMessage(client.Conn, message)
            }
            break
        case client := <-pool.Unregister:
            id := client.ID
            messageBody := fmt.Sprintf("u-%v 🫡", id)
            delete(pool.Clients, client)
            fmt.Printf("Size of Connection Pool: %v \n", len(pool.Clients))
            for client, _ := range pool.Clients {
              WriteMessage(client.Conn, Message{ClientId: 0, Body: messageBody})
            }
            break
        case message := <-pool.Broadcast:
            for client, _ := range pool.Clients {
                WriteMessage(client.Conn, message)            
            }
      }
    }
}

func (pool *Pool) StartWasm() {
    settingsBlob := []uint8{31, 133, 43, 63, 82, 184, 30, 63, 205, 204, 76, 61, 0, 0, 128, 63, 215, 163, 48, 63, 31, 133, 43, 63, 10, 215, 35, 61, 0, 0, 128, 63, 143, 194, 53, 63, 143, 194, 53, 63, 10, 215, 163, 60, 0, 0, 128, 63, 72, 225, 58, 63, 92, 143, 66, 63, 10, 215, 35, 60, 0, 0, 128, 63, 0, 0, 64, 63, 41, 92, 79, 63, 0, 0, 0, 0, 0, 0, 128, 63, 10, 215, 35, 63, 154, 153, 89, 63, 143, 194, 117, 62, 0, 0, 128, 63, 20, 174, 7, 63, 102, 102, 102, 63, 72, 225, 250, 62, 0, 0, 128, 63, 61, 10, 215, 62, 51, 51, 115, 63, 164, 112, 61, 63, 0, 0, 128, 63, 0, 0, 0, 0, 20, 174, 71, 63, 0, 0, 128, 63, 0, 0, 128, 63, 205, 204, 204, 61, 225, 122, 20, 63, 195, 245, 104, 63, 0, 0, 128, 63, 205, 204, 76, 62, 20, 174, 199, 62, 225, 122, 84, 63, 0, 0, 128, 63, 154, 153, 153, 62, 92, 143, 66, 62, 0, 0, 64, 63, 0, 0, 128, 63, 205, 204, 204, 62, 0, 0, 0, 0, 31, 133, 43, 63, 0, 0, 128, 63, 31, 133, 235, 62, 0, 0, 0, 0, 31, 133, 235, 62, 0, 0, 128, 63, 184, 30, 5, 63, 0, 0, 0, 0, 0, 0, 128, 62, 0, 0, 128, 63, 225, 122, 20, 63, 0, 0, 0, 0, 10, 215, 35, 61, 0, 0, 128, 63, 2, 8, 2, 0, 14, 0, 0, 0, 14, 0, 0, 0, 8, 0, 0, 0, 154, 153, 89, 63, 10, 215, 35, 60, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 0, 112, 65, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 128, 62, 1, 0, 0, 0, 0, 0, 12, 66, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 215, 99, 63, 1, 0, 0, 0, 0, 163, 208, 63, 1, 0, 0, 0, 0, 174, 71, 63, 1, 0, 0, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 205, 204, 76, 61, 0, 0, 0, 0, 0, 0, 12, 66, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 63, 4, 0, 0, 0, 0, 0, 112, 65, 0, 0, 0, 0, 1, 1, 0, 0, 164, 112, 189, 62, 12, 0, 0, 63, 0, 0, 112, 65, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 128, 62, 9, 0, 112, 65, 0, 0, 12, 66, 0, 0, 0, 0, 1, 0, 128, 62, 18, 0, 0, 0, 10, 215, 99, 63, 0, 0, 0, 0, 215, 163, 208, 63, 236, 81, 184, 61, 0, 0, 0, 0, 10, 215, 35, 189, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
    for {
      select {
        case client := <-pool.Register:
            pool.Clients[client] = true
            fmt.Printf("Size of Connection Pool: %v \n", len(pool.Clients))
            // send the current settingsBlob to the new client
            WriteSlice (client.Conn, settingsBlob)
            break
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            break
        case arr := <-pool.BroadcastSettings:
            // fmt.Printf("Sending message to all clients in Pool \n")
            // fmt.Printf("%v", arr)
            // update public settingsBlob
            settingsBlob = arr
            // fmt.Printf("new_blob: %v \n", settingsBlob)
            // send updated blob to everyone
            for client, _ := range pool.Clients {
                WriteSlice(client.Conn, settingsBlob)
            }
      }
    }
}

