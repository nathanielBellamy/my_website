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

type AuthController struct {
  AuthRoute string      // "auth/dev"
  AuthLoginRoute string // "auth/dev/dev-auth"
}

func (ac AuthController) RegisterController(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
){
  mode := os.Getenv("MODE")
  if !env.IsProd(mode) {
    ac.RegisterAuthHome(log)
    ac.RegisterAuthLoginRoute(cookieJar, log)
  }
}

func (ac AuthController) RegisterAuthHome(
  log *zerolog.Logger,
) {
  fs_auth := http.FileServer(http.Dir(ac.AuthRoute))
  route := fmt.Sprintf("/%s/", ac.AuthRoute)
  http.Handle(
    route,
    auth.LogClientIp("/auth/dev", log,
      http.StripPrefix("/auth/dev/", fs_auth),
    ),
  )
}

func (ac AuthController) RegisterAuthLoginRoute(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
){
  // TODO:
  // - set up salt route
  // - client requests salt
  // - server sends
  // - once client receives they can attempt login
  // - password is hashed on both client and server
  // - sent for comparison
  // - validate or don't

  route := fmt.Sprintf("/%s", ac.AuthLoginRoute)
  http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
    mode := os.Getenv("MODE")
    ip := auth.GetClientIpAddr(r)
    log.Info().
        Str("ip", ip).
        Msg("Dev Auth LOGIN ATTEMPT")
    sessionToken, success := auth.ValidateDev(w, r, log)
    if success {
      isLocalhost := env.IsLocalhost(mode)
      var name string
      if isLocalhost {
        name = "nbs-dev"
      } else {
        name = "__Secure-nbs-dev"
      }
      // set cookie on client
      c := http.Cookie {
        Name: name,
        Value: sessionToken,
        Path: "/",
        MaxAge: 60 * 60 * 48, // two days or whenever the server restarts as cookieJar is in-memory
        Secure: !isLocalhost, // https only
        HttpOnly: true, // don't let JS touch it
        SameSite: http.SameSiteLaxMode,
      }

      cookieJar.SetIfAbsent(sessionToken, auth.Cookie{Valid: true, Type: auth.CTDEV})
      http.SetCookie(w, &c)

      log.Info().
          Str("ip", ip).
          Msg("Dev Auth LOGIN SUCCESS")

      // TODO: wait for the browser to set the cookie
      //  - avoid redirect loop
      //  - Firefox takes at least 3 seconds
      //  - hopefully deters bots a bit
      //  - naive time.Sleep didn't solve it 100% so it's not worth it
      //  - for now, we'll eat some polluted logs due to redirect loop

      http.Redirect(w, r, "/", http.StatusFound)
      return
    } else {
      log.Warn().
          Str("ip", ip).
          Msg("Dev Auth LOGIN FAILURE")
      http.Error(w, "Invalid Password", http.StatusServiceUnavailable)
      return
    }
  })
}
