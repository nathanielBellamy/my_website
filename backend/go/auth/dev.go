package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/env"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

func SetupDevAuth(cookieJar *cmap.ConcurrentMap[string, Cookie], log *zerolog.Logger) {
  fs_frontend := http.FileServer(http.Dir("frontend"))
  http.Handle("/", http.StripPrefix("/", LogClientIp("/", log, RequireDevAuth(cookieJar, log, fs_frontend))))
  fs_auth := http.FileServer(http.Dir("auth/dev"))
  http.Handle("/auth/dev/",  LogClientIp("/auth/dev", log, http.StripPrefix("/auth/dev/", fs_auth)))

  // TODO:
  // - set up salt route
  // - client requests salt
  // - server sends
  // - once client receives they can attempt login
  // - password is hashed on both client and server
  // - sent for comparison
  // - validate or don't

  http.HandleFunc("/auth/dev/dev-auth", func(w http.ResponseWriter, r *http.Request) {
    mode := os.Getenv("MODE")
    ip := GetClientIpAddr(r)
    log.Info().
        Str("ip", ip).
        Msg("Dev Auth LOGIN ATTEMPT")
    sessionToken, success := ValidateDev(w, r, log)
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

      cookieJar.SetIfAbsent(sessionToken, Cookie{Valid: true, Type: CTDEV})
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

