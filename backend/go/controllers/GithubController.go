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
  RepoLanguagesRoute string
  RepoCommitsRoute string
}

func (gc GithubController) RegisterController(
  cookieJar *cmap.ConcurrentMap[string, auth.Cookie],
  log *zerolog.Logger,
) {
  gc.RegisterReposRoute(cookieJar, log)
  // gc.RepoLanguagesRoute(cookieJar, log)
  // gc.RepoCommitsRoute(cookieJar, log)
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

    repos := client.FetchRepos()
    userLanguageSummary := github.GenerateUserLanguageSummary(repos)

    response := github.GithubReposResponse{Repos: repos, UserLanguageSummary: userLanguageSummary}
    responseJSON, err := json.Marshal(response)
    if err != nil {
      log.Error().
          Err(err).
          Str("caller", "GithubController#RegisterReposRoute").
          Msg("Error JSON-ifying GithubReposResponse")
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(responseJSON)
  })
}

// func (gc GithubController) RepoLanguagesRoute(cookieJar, log) {

// }

// func (gc GithubController) RepoCommitsRoute(cookieJar, log) {

// }
