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

type GithubClientError struct {
  Msg string
}
func (gce GithubClientError) Error() (string) {
  return "GithubClientError: " + gce.Msg
}

func (gc GithubClient) HandleRequest(url string) (*http.Response, error) {
  http_client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    gcErr := GithubClientError{Msg: err.Error()}
    gc.Log.Error().
           Str("url", url).
           Err(gcErr).
           Str("caller", "GithubClient#HandleRequest").
           Msg("GithubClient http.NewRequest Error")

    return nil, gcErr
  }
  token := os.Getenv("GITHUB_KEY")
  req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

  resp, err := http_client.Do(req)
  if err != nil {
    gcErr := GithubClientError{Msg: err.Error()}
    gc.Log.Error().
           Str("url", url).
           Err(gcErr).
           Str("caller", "GithubClient#HandleRequest").
           Msg("GithubClient API Error When HandleRequest")
    return nil, gcErr
  }

  return resp, nil
}

func (gc GithubClient) FetchRepos() (GithubRepos, error) {
  url := fmt.Sprintf("%s/%s/%s/%s",
    gc.BaseUrl,
    "users",
    gc.Username,
    "repos",
  )
  resp, err := gc.HandleRequest(url)
  if err != nil {
    gcErr := GithubClientError{Msg: err.Error()}
    gc.Log.Error().
           Err(gcErr).
           Msg("Error Contacting Github When FetchRepos")
    return nil, gcErr
  }

  defer resp.Body.Close()

  var githubRepos GithubRepos
  json_err := json.NewDecoder(resp.Body).Decode(&githubRepos)
  if json_err != nil {
    gcErr := GithubClientError{Msg: json_err.Error()}
    gc.Log.Error().
           Err(gcErr).
           Msg("Error Decoding Github Repos Payload")
    return nil, gcErr
  }

  colors, load_err := LoadGithubColors()
  if load_err != nil {
    gcErr := GithubClientError{Msg: load_err.Error()}
    gc.Log.Error().
           Err(gcErr).
           Msg("Error Loading Github Colors")
    return nil, gcErr
  }

  refined_repo_data_err := gc.FetchRefinedRepoData(&githubRepos, &colors)
  if refined_repo_data_err != nil {
    gcErr := GithubClientError{Msg: refined_repo_data_err.Error()}
    gc.Log.Error().
           Err(gcErr).
           Msg("Error Fetching Refined Repo Data")
    return nil, gcErr
  }

  return githubRepos, nil
}

func (gc GithubClient) FetchRefinedRepoData(
  githubRepos *GithubRepos,
  colors *map[string]string,
) (error) {
  var wg sync.WaitGroup

  var errors []error
  for idx, _ := range *githubRepos {
    wg.Add(1)
    go func(idx int) {
      defer wg.Done()
      lang_err := (*githubRepos)[idx].FetchLanguageBreakdown(&gc, colors)
      if lang_err != nil {
        gc.Log.Error().
               Err(lang_err).
               Str("caller", "GithubClient#FetchRefinedRepoData").
               Msg("Error Fetching Repo Language Breakdown")
        errors = append(errors, lang_err)
      }

      commit_err := (*githubRepos)[idx].FetchCommits(&gc)
      if commit_err != nil {
        gc.Log.Error().
               Err(commit_err).
               Str("caller", "GithubClient#FetchRefinedRepoData").
               Msg("Error Fetching Repo Commits")
        errors = append(errors, commit_err)
      }
    }(idx)
  }

  wg.Wait()

  if len(errors) > 0 {
    gce := GithubClientError{Msg: ""}
    for _, err := range errors {
      gce.Msg += "\n " + err.Error()
    }

    return gce
  } else {
    return nil
  }
}

