package main

import (
  "fmt"
  "net/http"
  "os"

  "github.com/joho/godotenv"
  "github.com/nathanielBellamy/my_website/backend/go/auth"
  "github.com/nathanielBellamy/my_website/backend/go/controllers"
  "github.com/nathanielBellamy/my_website/backend/go/env"
  cmap "github.com/orcaman/concurrent-map/v2"
  "github.com/rs/zerolog"
)

// MODE=<mode> ./main
func main() {
    // init log
    file, err := os.Create("log.txt")
    if err != nil {
      fmt.Printf("Failed creating log file: %s", err)
    }
    log := zerolog.New(file).With().Timestamp().Logger()

    // determine runtime env
    mode := os.Getenv("MODE")
    if mode == "" {
      mode = "localhost"
    }

    // read env file
    log.Info().
        Msg("Loading ENV")

    envErr := godotenv.Load(".env." + mode)
    if envErr != nil {
      log.Fatal().
          Err(envErr).
          Msg("Error loading .env file")
    }

    log.Info().
        Str("mode", mode).
        Msg("Runtime Env")

    cookieJar := cmap.New[auth.Cookie]()

    log.Info().
        Msg("Establishing Routes")

    SetupRoutes(&cookieJar, &log)

    if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal().
          Msg("UnableToServe")
    }

    log.Info().
        Msg("Now serving on 8080")
}

func SetupRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger) {
    mode := os.Getenv("MODE")
    if env.IsProd(mode) {
      SetupProdRoutes()
    } else if env.IsRemotedev(mode) {
      SetupRemotedevRoutes(cookieJar, log)
    } else {
      SetupLocalhostRoutes(cookieJar, log)
    }

    SetupBaseRoutes(cookieJar, log)
}

func SetupBaseRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger) {
  mode := os.Getenv("MODE")
  if env.IsProd(mode) {
    fs := http.FileServer(http.Dir("frontend"))
    http.Handle("/", auth.LogClientIp("/", log, fs) )
  }

  RegisterControllers(cookieJar, log)
}

func RegisterControllers(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger) {
  var cs [3]controllers.Controller
  cs[0] = controllers.RecaptchaController{Route: "recaptcha"}
  cs[1] = controllers.PublicSquareController{
    FeedWebsocketRoute: "public-square-feed-ws",
    WasmWebsocketRoute: "public-square-wasm-ws",
  }
  cs[2] = controllers.GithubController {
    ReposRoute: "api/github/repos",
  }

  for _, c := range cs {
    c.RegisterController(cookieJar, log)
  }
}

func SetupRemotedevRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger) {
  auth.SetupDevAuth(cookieJar, log)
}

func SetupLocalhostRoutes(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger) {
  auth.SetupDevAuth(cookieJar, log)
}

func SetupProdRoutes() {
  // TODO: maybe Set cookie when user goes through ep warning
}

func _SetHeaders(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // this method wound up being superfluous for what we needed at the time of writing
    // but it's nice to have the infrastructure established

    // w.Header().Set("Content-Type", "text/javascript")
    // w.Header().Set("Content-Type", "text/html, text/css")
    handler.ServeHTTP(w,r)
  })
}
