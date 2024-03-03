package controllers

import (
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
)

type Controller interface {
  RegisterController(
    cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
    log *zerolog.Logger,
  )
}
