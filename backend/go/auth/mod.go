package auth

import (
	"fmt"
	// "encoding/json"
	"net/http"
	"strings"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
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

func HasValidCookie(w http.ResponseWriter, r *http.Request, cookieJar *cmap.ConcurrentMap[string, bool]) bool {
  fmt.Printf("\n wow zow \n")
  res := true  

  fmt.Printf("\n %v \n ", r)
  // Try to get the cookie
  cookie, err := r.Cookie("__Secure-nbs-dev")
  // Check if there was an error (e.g., cookie not found)
  if err != nil {
      // Handle the error
      fmt.Printf("\n Error getting the cookie:: ::%v", err)
      return false
  }


  fmt.Printf("\n wow zow 2222\n")

  fmt.Printf("\n coooookie::::: %v \n", cookie.Value)

  return res

  if cookieJar.Has(cookie.Value) {
    val, err := cookieJar.Get(cookie.Value)
    if err {
      res = false
    }
    res = val
  } else {
    cookieJar.SetIfAbsent(cookie.Value, false)
    res = false
  }

  return res
}

func ValidateDev (w http.ResponseWriter, r *http.Request, cookieJar *cmap.ConcurrentMap[string, bool]) bool {
  err := r.ParseForm()
  if err != nil {
    fmt.Printf("\n 1 Big Oh No This Time 1 \n")
    http.Error(w, err.Error(), http.StatusBadRequest)
    return false
  }
  
  clientSentPassword := r.Form.Get("pw")

  var h Hash
  var res bool
  res = h.Compare(clientSentPassword)

  if !res {
    fmt.Printf(" \n :: Incorrect Password :: \n")
    return false
  }

  sessionToken, err := h.Generate(time.Now().String())
  if err != nil {
    fmt.Printf("\n 2 Big Oh No This Time 2 \n")
    return false
  }

  fmt.Printf("session token: %v \n ", sessionToken)
  
  // set cookie on client
  c := http.Cookie {
    Name: "__Secure-nbs-dev",
    Value: sessionToken,
    MaxAge: 60 * 60 * 48, // two days
    Secure: true, // https only
    HttpOnly: true, // don't let JS touch it
  }

  fmt.Printf("cookie: %v \n \n \n", c)

  http.SetCookie(w, &c)
  fmt.Printf("writer: %v \n \n \n", w)

  cookieJar.SetIfAbsent(sessionToken, true)

  active, present := cookieJar.Get(sessionToken)
  if !present {
    fmt.Printf(" \n sessionToken not saved \n")
    return false
  } else {
    fmt.Printf(" \n token in cookieJar: %v \n \n ", active)
    return true
  }
}
