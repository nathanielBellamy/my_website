package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func getClientIpAddr(r *http.Request) (string, error) {
  var res string
  // TODO: 
  //  this will have to be conditional on runtime_env
  //  are we behind Nginx or not?

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

func HasValidCookie(r *http.Request, cookieJar *cmap.ConcurrentMap[string, bool]) bool {
  res := true  

  // Try to get the cookie
  cookie, err := r.Cookie("nbs-dev")
  // Check if there was an error (e.g., cookie not found)
  if err != nil {
      // Handle the error
      fmt.Printf("\n :: Error getting the cookie:: ::%v", err)
      return false
  }

  if cookieJar.Has(cookie.Value) {
    val, err := cookieJar.Get(cookie.Value)
    if err {
      res = false
    }
    res = val
  } else {
    // poison invalid token
    cookieJar.SetIfAbsent(cookie.Value, false)
    res = false
  }

  return res
}

func ValidateDev (w http.ResponseWriter, r *http.Request) (string, bool) {
  err := r.ParseForm()
  if err != nil {
    fmt.Printf(" \n :: Error Parsing POST :: \n")
    http.Error(w, err.Error(), http.StatusBadRequest)
    return "", false
  }
  
  clientSentPassword := r.Form.Get("pw")

  var h Hash
  res := h.Compare(clientSentPassword)

  if !res {
    fmt.Printf(" \n :: Incorrect Password :: \n")
    return "", false
  }

  sessionToken, err := h.Generate(time.Now().String())
  if err != nil {
    fmt.Printf(" \n :: Error Generating Session Token :: \n")
    return "", false
  }
  
  return sessionToken, true
}
