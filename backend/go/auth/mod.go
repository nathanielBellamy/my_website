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
    // TODO: set cookie in client
    // type Cookie struct {
    //   Name  string
    //   Value string

    //   Path       string    // optional
    //   Domain     string    // optional
    //   Expires    time.Time // optional
    //   RawExpires string    // for reading cookies only

    //   // MaxAge=0 means no 'Max-Age' attribute specified.
    //   // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    //   // MaxAge>0 means Max-Age attribute present and given in seconds
    //   MaxAge   int
    //   Secure   bool
    //   HttpOnly bool
    //   SameSite SameSite
    //   Raw      string
    //   Unparsed []string // Raw text of unparsed attribute-value pairs
    // }
    // http.SetCookie(w, cookie)
    http.ServeFile(w, r, "./../../frontend/dist")
  }
}
