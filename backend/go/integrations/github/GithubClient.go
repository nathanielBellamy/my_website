package github

import (
	"encoding/json"
	"fmt"
	"net/http"
  "os"
  "sync"

	"github.com/rs/zerolog"
)

type GithubClient struct {
  Username string
  BaseUrl string
  Log *zerolog.Logger
}

func (gc GithubClient) HandleRequest(url string) (*http.Response, error) {
  http_client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  token := os.Getenv("GITHUB_KEY")
  req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

  resp, err := http_client.Do(req)
  if (err != nil) {
    gc.Log.Error().
           Str("url", url).
           Err(err).
           Msg("GithubClient API Error When HandleRequest")
  }

  return resp, err
}

func (gc GithubClient) FetchRepos() (GithubRepos) {
  url := fmt.Sprintf("%s/%s/%s/%s",
    gc.BaseUrl,
    "users",
    gc.Username,
    "repos",
  )
  resp, err := gc.HandleRequest(url)
  if (err != nil) {
    gc.Log.Error().
           Err(err).
           Msg("Error Contacting Github When FetchRepos")
  }
  defer resp.Body.Close()

  var githubRepos GithubRepos

  json_err := json.NewDecoder(resp.Body).Decode(&githubRepos)
  if json_err != nil {
    gc.Log.Error().
           Err(json_err).
           Msg("Error Decoding Github Repos Payload")
  }

  gc.FetchRefinedRepoData(&githubRepos)
  return githubRepos
}

func (gc GithubClient) FetchRefinedRepoData(githubRepos *GithubRepos) {
  var wg sync.WaitGroup

  // load language color data
  colorsRaw, loadErr := os.ReadFile("fixtures/GithubColors.json")
  if loadErr != nil {
    fmt.Printf("OS LOAD ERROR: %+v\n", loadErr)
  }
  var colors map[string]string
  colors_err := json.Unmarshal(colorsRaw, &colors)

  if colors_err != nil {
    fmt.Printf("could not load colors")
    gc.Log.Error().
           Err(colors_err).
           Str("caller", "GithubClient#FetchRefinedRepoData").
           Msg("An Error Occurred Whie Parsing Colors Fixture")
  }

  for idx, _ := range *githubRepos {
    wg.Add(1)
    go func(idx int) {
      defer wg.Done()
      (*githubRepos)[idx].FetchLanguageBreakdown(&gc, &colors)
      (*githubRepos)[idx].FetchCommits(&gc)
    }(idx)
  }

  wg.Wait()
}

