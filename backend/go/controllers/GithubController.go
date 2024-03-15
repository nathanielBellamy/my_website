package controllers

import (
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/nathanielBellamy/my_website/backend/go/auth"
  "github.com/nathanielBellamy/my_website/backend/go/integrations/github"
  cmap "github.com/orcaman/concurrent-map/v2"
  "github.com/rs/zerolog"
)

type GithubController struct {
  ReposRoute string
}

type GithubControllerError struct {
  Msg string
}

func (gce GithubControllerError) Error() (string) {
  return "GithubControllerError: " + gce.Msg
}

func (gc GithubController) RegisterController(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  gc.RegisterReposRoute(cookieJar, log)
}

func (gc GithubController) RegisterReposRoute(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  route := fmt.Sprintf("/%s", gc.ReposRoute)
  http.HandleFunc(route, func (w http.ResponseWriter, r *http.Request) {
    ip := auth.GetClientIpAddr(r)
    log.Info().
        Str("ip", ip).
        Str("route", gc.ReposRoute).
        Msg("Github Controller Hit")

    client := github.GithubClient{Username: "nathanielBellamy", BaseUrl: "https://api.github.com", Log: log}

    repos, repos_err := client.FetchRepos()
    if repos_err != nil {
      gcErr := GithubControllerError{Msg: repos_err.Error()}
      log.Error().
          Err(gcErr).
          Str("caller", "GithubController#RegisterReposRoute").
          Msg("Error Connecting To Github")
      w.WriteHeader(http.StatusBadGateway)
      return
    }

    userLanguageSummary, langSumErr := github.GenerateUserLanguageSummary(repos)
    if langSumErr != nil {
      gcErr := GithubControllerError{Msg: langSumErr.Error()}
      log.Error().
          Err(gcErr).
          Str("caller", "GithubController#RegisterReposRoute").
          Msg("Error Generating Language Summary")
      w.WriteHeader(http.StatusInternalServerError)
      return
    }

    response := github.GithubReposResponse{Repos: repos, UserLanguageSummary: userLanguageSummary}
    responseJSON, json_err := json.Marshal(response)
    if json_err != nil {
      gcErr := GithubControllerError{Msg: json_err.Error()}
      log.Error().
          Err(gcErr).
          Str("caller", "GithubController#RegisterReposRoute").
          Msg("Error JSON-ifying GithubReposResponse")
      w.WriteHeader(http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJSON)
  })
}
