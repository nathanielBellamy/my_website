package controllers

import (
	"fmt"
	"net/http"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

type RecaptchaController struct {
  Route string
}

func (rc RecaptchaController) RegisterController(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  rc.EstablishRecaptcha(cookieJar, log)
}

func (rc RecaptchaController) EstablishRecaptcha(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  route := fmt.Sprintf("/%s", rc.Route)
  http.HandleFunc(route, func (w http.ResponseWriter, r *http.Request) {
    ip := auth.GetClientIpAddr(r)
    log.Info().
        Str("ip", ip).
        Msg("Recaptcha Endpoint Hit")

    res := auth.ValidateRecaptcha(r, log)
    log.Info().
        Str("ip", ip).
        Bool("res", res).
        Msg("ValidateRecaptcha")

    if res {
      auth.SetRecaptchaCookieOnClient(w, cookieJar, log)

      w.WriteHeader(http.StatusOK)
      w.Write([]byte("OK"))
    } else {
      w.WriteHeader(http.StatusForbidden)
      w.Write([]byte("NOT OK"))
    }
  })
}

