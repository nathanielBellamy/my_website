package github

type GithubRepos = []GithubRepo

type GithubColorsData = []string

type GithubLanguageBreakdown = map[string]int

type GithubLanguage struct {
  // lower cased names to match frontend usage in Apache echarts
  Name  string `json:"name"`
  Value int    `json:"value"`
}

type GithubLanguageData = []GithubLanguage
