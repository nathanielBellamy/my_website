package controllers

import (
  "fmt"
  "net/http"
  "os"

  "github.com/nathanielBellamy/my_website/backend/go/auth"
  "github.com/nathanielBellamy/my_website/backend/go/env"
  cmap "github.com/orcaman/concurrent-map/v2"
  "github.com/rs/zerolog"
)

type AppController struct {
  HomeRoute string      // ""
}

func (ac AppController) RegisterController(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  mode := os.Getenv("MODE")
  if env.IsProd(mode) {
    ac.RegisterPublicProdHome(log)
  } else {
    ac.RegisterPrivateDevHome(cookieJar, log)
  }
}

func (ac AppController) RegisterPublicProdHome(log *zerolog.Logger) {
  route := fmt.Sprintf("/%s", ac.HomeRoute)
  fs := http.FileServer(http.Dir("frontend"))
  http.Handle(route, auth.LogClientIp(route, log, fs))
}

func (ac AppController) RegisterPrivateDevHome(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  route := fmt.Sprintf("/%s", ac.HomeRoute)
  fs_frontend := http.FileServer(http.Dir("frontend"))
  http.Handle(
    route,
    http.StripPrefix("/",
      auth.LogClientIp("/", log,
        auth.RequireDevAuth(cookieJar, log, fs_frontend),
      ),
    ),
  )
}
