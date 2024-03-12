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
      log.Error().
          Err(repos_err).
          Str("caller", "GithubController#RegisterReposRoute").
          Msg("Error Connecting To Github")
      w.WriteHeader(http.StatusBadGateway)
      return
    }

    userLanguageSummary := github.GenerateUserLanguageSummary(repos)

    response := github.GithubReposResponse{Repos: repos, UserLanguageSummary: userLanguageSummary}
    responseJSON, json_err := json.Marshal(response)
    if json_err != nil {
      log.Error().
          Err(json_err).
          Str("caller", "GithubController#RegisterReposRoute").
          Msg("Error JSON-ifying GithubReposResponse")
      w.WriteHeader(http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJSON)
  })
}
