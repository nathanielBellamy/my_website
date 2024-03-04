package github

import (
	"encoding/json"
	"fmt"
	"net/http"
  "os"

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

func (gc GithubClient) FetchRepos() {
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

  var jsonData interface{}

  json_err := json.NewDecoder(resp.Body).Decode(&jsonData)
  if json_err != nil {
    gc.Log.Error().
           Err(json_err).
           Msg("Error Decoding Github Repos Payload")
  }

  fmt.Printf("%+v\n", jsonData)
}
