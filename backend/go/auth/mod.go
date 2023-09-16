package auth

import (
	"net/http"
  "os"
	"strings"

	"github.com/nathanielBellamy/my_website/backend/go/env"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

func GetClientIpAddr(r *http.Request) string {
  var res string

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

func HasValidCookie(
  r *http.Request, 
  cookieType CookieType, 
  cookieJar *cmap.ConcurrentMap[string, Cookie], 
  log *zerolog.Logger,
) bool {
  res := true
  ip := GetClientIpAddr(r)

  var cookieName string
  mode := os.Getenv("MODE")
  if env.IsLocalhost(mode) {
    switch cookieType {
    case CTDEV:
      cookieName = "nbs-dev"
    case CTPSR:
      cookieName = "nbs-psr"
    }
  } else {
    switch cookieType {
    case CTDEV:
      cookieName = "__Secure-nbs-dev"
    case CTPSR:
      cookieName = "__Secure-nbs-psr"
    }
  }
  // Try to get the cookie
  cookie, err := r.Cookie(cookieName)
  // Check if there was an error (e.g., cookie not found)
  if err != nil {
      // Handle the error
      log.Error().
          Err(err).
          Str("ip", ip).
          Int32("Type", int32(cookieType)).
          Str("cookieName", cookieName).
          Msg("Error Retrieving Cookie")
      return false
  }

  if cookieJar.Has(cookie.Value) {
    cookieFromJar, err := cookieJar.Get(cookie.Value)
    if err {
      log.Error().
          Str("ip", ip).
          Int32("Type", int32(cookieType)).
          Msg("Cookie Not Found in Jar")
      res = false
    }
    res = cookieFromJar.Valid && cookieFromJar.Type == cookieType
  } else {
    // poison invalid token
    cookieJar.SetIfAbsent(
      cookie.Value, 
      Cookie{ Valid: false, Type: cookieType },
    )
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

