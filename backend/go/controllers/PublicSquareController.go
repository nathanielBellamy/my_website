package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

type PublicSquareController struct {
  FeedWebsocketRoute string
  WasmWebsocketRoute string
}

func (psc PublicSquareController) RegisterController(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  psc.EstablishFeedWebsocketPool(cookieJar, log)
  psc.EstablishWasmWebsocketPool(cookieJar, log)
}

func (psc PublicSquareController) EstablishFeedWebsocketPool(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  feedPool := websocket.NewPool(log)
  go feedPool.StartFeed()

  route := fmt.Sprintf("/%s", psc.FeedWebsocketRoute)
  http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
    ip := auth.GetClientIpAddr(r)
    mode := os.Getenv("MODE")
    if !env.IsProd(mode) {
      // localhost and remote dev require basic login
      validDev := auth.HasValidCookie(r, auth.CTPSR, cookieJar, log)
      validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, cookieJar, log)
      log.Info().
          Str("ip", ip).
          Bool("validDev", validDev).
          Bool("validRecaptcha", validRecaptcha).
          Msg("PS FEED WS")

      if validDev && validRecaptcha {
        websocket.ServeFeedWs(feedPool, w, r, log)
      } else {
        auth.RedirectToDevAuth(w, r, log)
      }
    } else {
      // prod is public
      // but protected by recaptcha
      validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, cookieJar, log)
      log.Info().
          Str("ip", ip).
          Bool("validRecaptcha", validRecaptcha).
          Msg("PS FEED WS")

      if validRecaptcha {
        websocket.ServeFeedWs(feedPool, w, r, log)
      } else {
        auth.RedirectToHome(w, r, log)
      }
    }
  })
}

func (psc PublicSquareController) EstablishWasmWebsocketPool(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  wasmPool := websocket.NewPool(log)
  go wasmPool.StartWasm()
  route := fmt.Sprintf("/%s", psc.WasmWebsocketRoute)
  http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
    ip := auth.GetClientIpAddr(r)
    mode := os.Getenv("MODE")
    if !env.IsProd(mode) {
      // localhost and remote dev require basic login
      validDev := auth.HasValidCookie(r, auth.CTPSR, cookieJar, log)
      validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, cookieJar, log)
      log.Info().
          Str("ip", ip).
          Bool("validDev", validDev).
          Bool("validRecaptcha", validRecaptcha).
          Msg("PS WASM WS")

      if validDev && validRecaptcha {
        websocket.ServeWasmWs(wasmPool, w, r, log)
      } else {
        auth.RedirectToDevAuth(w, r, log)
      }
    } else {
      // prod is public
      // but protected by recaptcha
      validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, cookieJar, log)
      log.Info().
          Str("ip", ip).
          Bool("validRecaptcha", validRecaptcha).
          Msg("PS WASM WS")

      if validRecaptcha {
        websocket.ServeWasmWs(wasmPool, w, r, log)
      } else {
        auth.RedirectToHome(w, r, log)
      }
    }
  })
}
