package github

type GithubRepos = []GithubRepo

type GithubColorData = []string

type GithubLanguageBreakdown = map[string]int

type Author struct {
  Date string `json:"date"`
}
type Commit struct {
  Author    Author  `json:"author"`
}
type GithubCommitRaw struct {
  Sha       string   `json:"sha"`
  HtmlUrl   string   `json:"html_url"`
  Commit    Commit   `json:"commit"`
}

type GithubCommitsRaw = []GithubCommitRaw

type GithubCommit struct {
  Sha       string  `json:"sha"`
  HtmlUrl   string  `json:"html_url"`
  Date      string  `json:"date"`
}
type GithubCommits = []GithubCommit


type GithubLanguage struct {
  // lower cased names to match frontend usage in Apache echarts
  Name  string `json:"name"`
  Value int    `json:"value"`
}

type GithubLanguageData = []GithubLanguage
