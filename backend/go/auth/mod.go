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

func HandleDev (w *http.ResponseWriter, r *http.Request, cookieJar *CookieJar) {
  err := r.ParseForm()
  if err != nil {
    http.Error(*w, err.Error(), http.StatusBadRequest)
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

    fmt.Printf("session token: %v \n ", sessionToken)

    // save cookie on server
    
    // set cookie on client
    c := http.Cookie {
      Name: "__Secure-nbs-dev",
      Value: sessionToken,
      MaxAge: 60 * 60 * 48, // two days
      Secure: true, // https only
      HttpOnly: true, // don't let JS touch it
    }

    fmt.Printf("cookie: %v \n \n \n", c)

    http.SetCookie(*w, &c)
    fmt.Printf("writer: %v \n \n \n", *w)
    // cookieJar.cookies.Set(sessionToken, true)

    fmt.Printf(" \n End Handle Dev \n \n ")
  }
}
