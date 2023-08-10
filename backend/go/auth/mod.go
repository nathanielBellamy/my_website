package dev_auth

import (
    // "fmt"
    "encoding/json"
    // "log"
    "net/http"
    "os"
)

type devAuth struct {}

func (c *devAuth) handleDevAuth (w http.ResponseWriter, r *http.Request) {

  // TODO: read headers
  // if request is not from dev-site return

  var passwordCompareResult bool

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

  passwordsMatch := h.Compare(correctPassword, clientSentPassword)

  if res == nil {
    // serve the SPA
  }
}
