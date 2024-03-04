package github

import (
	"encoding/json"
	"fmt"
)

type GithubRepo struct {
  // ColorData GithubColorsData
  // CommitData GithubCommitsData
  CreatedAt string                            `json:"created_at"`
  Description string                          `json:"description"`
  HtmlUrl string                              `json:"html_url"`
  Language string                             `json:"language"`
  LanguageData GithubLanguageData             `json:"langauge_data"`
  Name string                                 `json:"name"`
  PushedAt string                             `json:"pushed_at"`
  UpdatedAt string                            `json:"updated_at"`
}

func (gr *GithubRepo) FetchLanguageBreakdown(client GithubClient) {
  fmt.Printf("REPO NAME: %s", gr.Name)
  url := fmt.Sprintf("%s/repos/%s/%s/languages", client.BaseUrl, client.Username, gr.Name)
  resp, err := client.HandleRequest(url)
  if (err != nil) {
    client.Log.Error().
               Err(err).
               Msg("GithubRepo Api Error When FetchLanguageBreakdown")
  }
  defer resp.Body.Close()

  var languageBreakdown GithubLanguageBreakdown
  json_err := json.NewDecoder(resp.Body).Decode(&languageBreakdown)
  if (json_err != nil) {
    client.Log.Error().
               Err(err).
               Msg("Error Parsing JSON When FetchLanguageBreakdown")
  }

  var languageData GithubLanguageData
  for k, v := range languageBreakdown {
    languageData = append(languageData, GithubLanguage{Name: k, Value: v})
  }
  gr.LanguageData = languageData
  // fmt.Printf("REPO %+v\n", gr)
}
