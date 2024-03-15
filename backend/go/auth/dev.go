package auth

import (
	"net/http"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

func RequireDevAuth(cookieJar *cmap.ConcurrentMap[string, Cookie], log *zerolog.Logger, handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if HasValidCookie(r, CTDEV, cookieJar, log){
          handler.ServeHTTP(w, r)
          return
        } else {
          // http.Error(w, "Dev Not Validated", 503)
          log.Warn().
              Str("ip", GetClientIpAddr(r)).
              Msg("Valid Cookie NOT FOUND")
          RedirectToDevAuth(w, r, log)
          return
        }
    })
}

func ValidateDev (w http.ResponseWriter, r *http.Request, log *zerolog.Logger) (string, bool) {
  ip := GetClientIpAddr(r)
  err := r.ParseForm()
  if err != nil {
    log.Error().
        Str("ip", ip).
        Err(err).
        Msg("Error Parsing POST")
    http.Error(w, err.Error(), http.StatusBadRequest)
    return "", false
  }

  clientSentPassword := r.Form.Get("pw")

  var h Hash
  res := h.Compare(clientSentPassword)

  if !res {
    log.Warn().
        Str("ip", ip).
        Msg("Incorrect Password")
    return "", false
  }

  sessionToken, err := h.Generate(time.Now().String())
  if err != nil {
    log.Error().
        Str("ip", ip).
        Err(err).
        Msg("Error Generating Session Token")
    return "", false
  }

  return sessionToken, true
}

func RedirectToDevAuth(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
  log.Warn().
      Str("ip", GetClientIpAddr(r)).
      Msg("REDIRECT To Dev Auth")
  http.Redirect(w,r,"/auth/dev/", http.StatusSeeOther)
}

func RedirectToHome(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
  log.Warn().
      Str("ip", GetClientIpAddr(r)).
      Msg("REDIRECT To Home")
  http.Redirect(w,r,"/", http.StatusSeeOther)
}

