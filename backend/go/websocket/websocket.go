package websocket

// https://tutorialedge.net/projects/chat-system-in-go-and-react/part-4-handling-multiple-clients/

import (
    "encoding/json"
    "fmt"
    // "io"
    // "log"
    "net"
    "net/http"

    "github.com/gobwas/ws"
    "github.com/gobwas/ws/wsutil"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
    conn, _, _, err := ws.UpgradeHTTP(r, w)
    if err != nil {
      fmt.Printf("UpgradeHTTP Error \n")
      return conn, err
    }
    return conn, nil
}

func Reader(conn *net.Conn) ([]byte, error){
    for {
        msg, _, err := wsutil.ReadClientData(*conn)
        if err != nil {
          fmt.Printf("ReadClientData Error \n")
          (*conn).Close()
          return []byte(""), err
        }
        
        fmt.Printf(string(msg))
        msg = append(msg, []byte(" !GO! ")...)
        Writer(conn, msg)
    }
}

func WriteMessage(conn *net.Conn, msg Message) {
    j, err := json.Marshal(msg)
    if err != nil {
      fmt.Printf("JSON MESSAGE ERROR \n")
      return
    }

    Writer(conn, []byte(j))
}

func WriteSlice(conn *net.Conn, slice []uint8) {
  err := wsutil.WriteServerBinary(*conn, []byte(slice))
  if err != nil {
    fmt.Printf("WriteServerSlice Error \n")
    (*conn).Close()
    return
  }
}

func Writer(conn *net.Conn, msg []byte) {
    // fmt.Printf("Writer with msg: %v", msg)
    err := wsutil.WriteServerMessage(*conn, ws.OpText, []byte(msg))
    if err != nil {
      fmt.Printf("WriteServerMessage Error \n")
      (*conn).Close()
      return
    }
}
