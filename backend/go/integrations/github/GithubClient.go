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
    return nil, err
  }

  return resp, err
}

func (gc GithubClient) FetchRepos() (GithubRepos, error) {
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
    return nil, err
  }

  defer resp.Body.Close()

  var githubRepos GithubRepos
  json_err := json.NewDecoder(resp.Body).Decode(&githubRepos)
  if json_err != nil {
    gc.Log.Error().
           Err(json_err).
           Msg("Error Decoding Github Repos Payload")
    return nil, json_err
  }

  colors := LoadGithubColors()
  gc.FetchRefinedRepoData(&githubRepos, &colors)

  return githubRepos, nil
}

func (gc GithubClient) FetchRefinedRepoData(githubRepos *GithubRepos, colors *map[string]string) {
  var wg sync.WaitGroup

  for idx, _ := range *githubRepos {
    wg.Add(1)
    go func(idx int) {
      defer wg.Done()
      (*githubRepos)[idx].FetchLanguageBreakdown(&gc, colors)
      (*githubRepos)[idx].FetchCommits(&gc)
    }(idx)
  }

  wg.Wait()
}

