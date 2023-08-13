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

  _, ok := cj.cookies.Get(incoming_cookie.Value)
  if !ok {
    valid_err := incoming_cookie.Valid()
    if valid_err != nil {
      cj.cookies.Set(incoming_cookie.Value, false)
    }
    return false
  }

  // valid_err := cookie.Valid()
  // if valid_err != nil {
  //   delete(cj.cookies, incoming_cookie.Value);
  //   return false
  // }

  return true
}

