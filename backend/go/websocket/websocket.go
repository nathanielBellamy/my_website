package websocket

// https://tutorialedge.net/projects/chat-system-in-go-and-react/part-4-handling-multiple-clients/
import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog"
)

func Upgrade(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) (net.Conn, error) {
    conn, _, _, err := ws.UpgradeHTTP(r, w)
    if err != nil {
      log.Error().
          Err(err).
          Msg("UpgradeHTTP Error")
      return conn, err
    }
    return conn, nil
}

func WriteMessage(conn *net.Conn, msg Message, log *zerolog.Logger) {
    j, err := json.Marshal(msg)
    if err != nil {
      log.Error().
          Err(err).
          Msg("JSON MESSAGE ERROR")
      return
    }

    Writer(conn, []byte(j), log)
}

func WriteSlice(conn *net.Conn, slice []uint8, log *zerolog.Logger) {
  err := wsutil.WriteServerBinary(*conn, []byte(slice))
  if err != nil {
    log.Error().
        Err(err).
        Msg("WriteServer SLICE Error")
    (*conn).Close()
    return
  }
}

func Writer(conn *net.Conn, msg []byte, log *zerolog.Logger) {
    err := wsutil.WriteServerMessage(*conn, ws.OpText, []byte(msg))
    if err != nil {
      log.Error().
          Err(err).
          Msg("WriteServer MESSAGE Error")
      (*conn).Close()
      return
    }
}
