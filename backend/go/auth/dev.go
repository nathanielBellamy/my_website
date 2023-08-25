package auth

import (
  "fmt"
	"net/http"
	"time"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/nathanielBellamy/my_website/backend/go/env"
)

func SetupDevAuth(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool]) {
  fs_frontend := http.FileServer(http.Dir("frontend"))
  http.Handle("/", http.StripPrefix("/", RequireDevAuth(runtime_env, cookieJar, fs_frontend)))
  
  fs_auth := http.FileServer(http.Dir("auth/dev"))
  http.Handle("/auth/dev/",  http.StripPrefix("/auth/dev/", fs_auth))

  http.HandleFunc("/auth/dev/dev-auth", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("\n :: dev-auth :: \n")
    sessionToken, success := ValidateDev(w, r)
    if success {
      isLocalhost := runtime_env.IsLocalhost()
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

      cookieJar.SetIfAbsent(sessionToken, true)
      http.SetCookie(w, &c)

      http.Redirect(w, r, "/", http.StatusFound)
      return
    } else {
      http.Error(w, "Invalid Password", http.StatusServiceUnavailable)
      return
    }
  })
}

func RequireDevAuth(runtime_env env.Env, cookieJar *cmap.ConcurrentMap[string, bool], handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if HasValidCookie(runtime_env, r, cookieJar){
          handler.ServeHTTP(w, r)
          return
        } else {
          // http.Error(w, "Dev Not Validated", 503)
          fmt.Printf("\n :: REDIRECT To Dev Auth :: \n")
          RedirectToDevAuth(w, r)
          return
        }
    })
}

func ValidateDev (w http.ResponseWriter, r *http.Request) (string, bool) {
  err := r.ParseForm()
  if err != nil {
    fmt.Printf(" \n :: Error Parsing POST :: \n")
    http.Error(w, err.Error(), http.StatusBadRequest)
    return "", false
  }
  
  clientSentPassword := r.Form.Get("pw")

  var h Hash
  res := h.Compare(clientSentPassword)

  if !res {
    fmt.Printf(" \n :: Incorrect Password :: \n")
    return "", false
  }

  sessionToken, err := h.Generate(time.Now().String())
  if err != nil {
    fmt.Printf(" \n :: Error Generating Session Token :: \n")
    return "", false
  }
  
  return sessionToken, true
}

func RedirectToDevAuth(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("\n redirectToDevAuth \n")
  http.Redirect(w,r,"/auth/dev/", http.StatusSeeOther)
}

