package auth

import (
	// "errors"
	// "strings"
  "fmt"
  "net/http"
	// "golang.org/x/crypto/bcrypt"

	"github.com/orcaman/concurrent-map/v2"
)


// type CookieJar struct {
//   cookies cmap.ConcurrentMap[string, bool]
// }


func ValidateSessionCookie(r *http.Request, cookieJar *cmap.ConcurrentMap[string, bool]) (bool) {
  incoming_cookie, err := r.Cookie("__Secure-nbs-dev")
  fmt.Printf(" \n \n incoming req: %v \n \n ", *r)
  fmt.Printf(" \n \n incoming cooking: %v \n \n ", incoming_cookie.Value)
  if err != nil {
    fmt.Printf(" \n \n incoming cooking: UH OH! \n \n ")
    // Handle error (e.g., no cookie found, or expired session)
    return false
  }


  fmt.Printf(" \n \n cookiejar: %v \n \n ", cookieJar)
  active, present := cookieJar.Get(incoming_cookie.Value)
  if !present {
    fmt.Printf(" cookie not found ")
    return false 
  } else {
    valid_err := incoming_cookie.Valid()
    if valid_err != nil {
      // deactivate
      cookieJar.Set(incoming_cookie.Value, false)
      return false
    }
    return active 
  }
}

