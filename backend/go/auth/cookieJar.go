package auth

import (
	// "errors"
	// "strings"
  "fmt"
  "net/http"
	// "golang.org/x/crypto/bcrypt"

	"github.com/orcaman/concurrent-map/v2"
)


type CookieJar struct {
  cookies cmap.ConcurrentMap[string, bool]
}


func (cj *CookieJar) ValidateSessionCookie(r *http.Request) (bool) {
  incoming_cookie, err := r.Cookie("__Secure-nbs-dev")
  fmt.Printf(" \n \n incoming req: %v \n \n ", *r)
  fmt.Printf(" \n \n incoming cooking: %v \n \n ", incoming_cookie.Value)
  if err != nil {
    fmt.Printf(" \n \n incoming cooking: UH OH! \n \n ")
    // Handle error (e.g., no cookie found, or expired session)
    return false
  }


  fmt.Printf(" \n \n cookiejar: %v \n \n ", cj)
  active, present := cj.cookies.Get(incoming_cookie.Value)
  if !present {
    fmt.Printf(" here here 123 ")
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

