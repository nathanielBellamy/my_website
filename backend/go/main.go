package main

import (
  "fmt"
  "net/http"
  "os"

  "github.com/joho/godotenv"
  "github.com/nathanielBellamy/my_website/backend/go/auth"
  "github.com/nathanielBellamy/my_website/backend/go/controllers"
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

    RegisterControllers(&cookieJar, &log)

    if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal().
          Msg("UnableToServe")
    }

    log.Info().
        Msg("Now serving on 8080")
}

func RegisterControllers(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger) {
  var cs []controllers.Controller
  cs = append(cs, controllers.AppController{
    HomeRoute: "",
  })
  cs = append(cs, controllers.AuthController{
    AuthRoute: "auth/dev",
    AuthLoginRoute: "auth/dev/dev-auth",
  })
  cs = append(cs, controllers.RecaptchaController{
    RecaptchaRoute: "recaptcha",
  })
  cs = append(cs, controllers.PublicSquareController{
    FeedWebsocketRoute: "public-square-feed-ws",
    WasmWebsocketRoute: "public-square-wasm-ws",
  })
  cs = append(cs, controllers.GithubController {
    ReposRoute: "api/github/repos",
  })

  for _, c := range cs {
    c.RegisterController(cookieJar, log)
  }
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
