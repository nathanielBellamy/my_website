package auth

import (
    "fmt"
    // "encoding/json"
    "net/http"
    "strings"
    "time"
)

func getClientIpAddr(r *http.Request) (string, error) {
  var res string
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		// The header can contain multiple IPs, comma-separated.
		// The client's IP is typically the first one.
		ips := strings.Split(forwardedFor, ",")
		if len(ips) > 0 {
			res = strings.TrimSpace(ips[0])
		}
	}
  return res, nil
}

type IoPassword struct {
  Password string
}

func HandleDev (w http.ResponseWriter, r *http.Request, cookieJar *CookieJar) {
  fmt.Printf("Wow Zow \n")
  
  err := r.ParseForm()
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  
  clientSentPassword := r.Form.Get("pw")

  var h Hash
  var res bool
  res = h.Compare(clientSentPassword)

  if res {
    sessionToken, err := h.Generate(time.Now().String())
    if err != nil {
      return
    }

    // save cookie on server
    
    // set cookie on client
    c := http.Cookie {
      Name: "devdev-nbs-dev",
      Value: sessionToken,
      Domain: "devdev-nbs-dev.dev",
      MaxAge: 60 * 60 * 48, // two days
      Secure: true, // https only
      HttpOnly: true, // don't let JS touch it
    }

    (*cookieJar).cookies.Set(sessionToken, true)

    http.SetCookie(w, &c)
    http.ServeFile(w, r, "./../../frontend/dist")
  }
}
