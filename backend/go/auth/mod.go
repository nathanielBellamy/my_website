package auth

import (
    // "fmt"
    "encoding/json"
    // "log"
    "net/http"
    "os"
)

func HandleDev (w http.ResponseWriter, r *http.Request) {
  var clientSentPassword string
  err := json.NewDecoder(r.Body).Decode(&clientSentPassword)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  correctPassword, err := os.ReadFile("/dev_pw")
  if err != nil {
    return
  }

  var h Hash

  var res bool
  res = h.Compare(correctPassword, clientSentPassword)

  if res {
    // TODO: set token in client
    http.ServeFile(w, r, "./../../frontend/dist")
  }
}
