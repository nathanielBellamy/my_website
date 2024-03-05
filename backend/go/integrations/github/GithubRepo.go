package github

import (
	"encoding/json"
	"fmt"
  "sort"
)

type GithubRepo struct {
  ColorData GithubColorData                   `json:"color_data"`
  CommitData GithubCommits                    `json:"commit_data"`
  CreatedAt string                            `json:"created_at"`
  Description string                          `json:"description"`
  HtmlUrl string                              `json:"html_url"`
  Language string                             `json:"language"`
  LanguageData GithubLanguageData             `json:"language_data"`
  Name string                                 `json:"name"`
  PushedAt string                             `json:"pushed_at"`
  UpdatedAt string                            `json:"updated_at"`
}

func (gr *GithubRepo) FetchLanguageBreakdown(client *GithubClient, colors *map[string]string) {
  url := fmt.Sprintf("%s/repos/%s/%s/languages", client.BaseUrl, client.Username, gr.Name)
  resp, err := client.HandleRequest(url)
  if err != nil {
    client.Log.Error().
               Err(err).
               Str("caller", "GithubRepo#FetchLanguageBreakdown").
               Msg("Github Api Error - /languages")
  }
  defer resp.Body.Close()

  var languageBreakdown GithubLanguageBreakdown
  json_err := json.NewDecoder(resp.Body).Decode(&languageBreakdown)
  if json_err != nil {
    client.Log.Error().
               Err(err).
               Str("caller", "GithubRepo#FetchLanguageBreakdown").
               Msg("Error Parsing Github JSON - /languages")
  }

  gr.ProcessLanguageBreakdown(client, languageBreakdown, colors)
}

func (gr *GithubRepo) ProcessLanguageBreakdown(
  client *GithubClient,
  languageBreakdown GithubLanguageBreakdown,
  colors *map[string]string,
) {
  // process language data to be ingested by Apache echarts
  var languageData GithubLanguageData
  for k, v := range languageBreakdown {
    languageData = append(languageData, GithubLanguage{Name: k, Value: v})
  }

  sort.Slice(languageData, func(i, j int) bool {
    return languageData[i].Value > languageData[j].Value
  })
  gr.LanguageData = languageData

  var colorData GithubColorData
  for _, githubLanguage := range languageData {
    color := (*colors)[githubLanguage.Name]
    colorData = append(colorData, color)
  }
  gr.ColorData = colorData
}

func (gr *GithubRepo) FetchCommits(client *GithubClient) {
  url := fmt.Sprintf("%s/repos/%s/%s/commits", client.BaseUrl, client.Username, gr.Name)
  resp, err := client.HandleRequest(url)
  if err != nil {
    client.Log.Error().
               Err(err).
               Str("caller", "GithubRepo#FetchCommits").
               Msg("Github Api Error - /commits")
  }
  defer resp.Body.Close()

  var commits GithubCommitsRaw
  json_err := json.NewDecoder(resp.Body).Decode(&commits)
  if json_err != nil {
    client.Log.Error().
               Err(err).
               Str("caller", "GithubRepo#FetchCommits").
               Msg("Error Parsing Github JSON - /commits")
  }
  gr.ProcessCommits(client, commits)
}

func (gr *GithubRepo) ProcessCommits(client *GithubClient, commitsRaw GithubCommitsRaw) {
  var commitData GithubCommits
  for _, commitRaw := range commitsRaw {
    commitData = append(
      commitData,
      GithubCommit{
        Sha: commitRaw.Sha,
        HtmlUrl: commitRaw.HtmlUrl,
        Date: commitRaw.Commit.Author.Date,
      },
    )
  }
  gr.CommitData = commitData
}
