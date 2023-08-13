package auth

import (
	// "errors"
	// "strings"
  "net/http"
	// "golang.org/x/crypto/bcrypt"

	"github.com/orcaman/concurrent-map/v2"
)


type CookieJar struct {
  cookies cmap.ConcurrentMap[string, bool]
}


func (cj *CookieJar) ValidateSessionCookie(r *http.Request) (bool) {
  incoming_cookie, err := r.Cookie("session_token")
  if err != nil {
    // Handle error (e.g., no cookie found, or expired session)
    return false
  }

  active, present := cj.cookies.Get(incoming_cookie.Value)
  if !present {
    return false 
  } else {
    valid_err := incoming_cookie.Valid()
    if valid_err != nil {
      // deactivate
      cj.cookies.Set(incoming_cookie.Value, false)
      return false
    }
    return active 
  }
}

