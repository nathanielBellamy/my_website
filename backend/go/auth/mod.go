package auth

import (
	"net/http"
	"strings"

	"github.com/nathanielBellamy/my_website/backend/go/env"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

func GetClientIpAddr(r *http.Request) string {
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
	} else {
    res = r.RemoteAddr
  }
  return res
}

type IoPassword struct {
  Password string
}

func HasValidCookie(runtime_env env.Env, r *http.Request, cookieJar *cmap.ConcurrentMap[string, bool], log *zerolog.Logger) bool {
  res := true
  ip := GetClientIpAddr(r)

  var cookieName string
  if runtime_env.IsLocalhost() {
    cookieName = "nbs-dev"
  } else {
    cookieName = "__Secure-nbs-dev"
  }
  // Try to get the cookie
  cookie, err := r.Cookie(cookieName)
  // Check if there was an error (e.g., cookie not found)
  if err != nil {
      // Handle the error
      log.Error().
          Str("ip", ip).
          Err(err).
          Msg("Error Retrieving Cookie")
      return false
  }

  if cookieJar.Has(cookie.Value) {
    val, err := cookieJar.Get(cookie.Value)
    if err {
      log.Warn().
          Str("ip", ip).
          Msg("Cookie Not Valid")
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

func LogClientIp(url string, log *zerolog.Logger, handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      log.Info().
          Str("ip", GetClientIpAddr(r)).
          Str("url", url).
          Msg("URL HIT")
      
      handler.ServeHTTP(w, r)
      return  
  })
}

