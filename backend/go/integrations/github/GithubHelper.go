package github

import (
  "encoding/json"
  "fmt"
  "os"
  // "slices"
  "sort"
)

func GenerateUserLanguageSummary(githubRepos GithubRepos) UserLanguageSummary {
  var languageData LanguageData

  for _, repo := range githubRepos {
    for _, githubLanguage := range repo.LanguageData {
      // TODO:
      // - troubleshoot importing slices
      // langIdx := slices.IndexFunc(languageData, func(gl GithubLanguage) bool {
      //   gl.Name == githubLanguage.Name
      // })
      langIdx := -1
      for idx, lang := range languageData {
        if lang.Name == githubLanguage.Name {
          langIdx = idx
        }
      }
      if langIdx > -1 {
        languageData[langIdx].Value += githubLanguage.Value
      } else {
        languageData = append(languageData, githubLanguage)
      }
    }
  }

  sort.Slice(languageData, func(i, j int) bool {
    return languageData[i].Value > languageData[j].Value
  })

  colors := LoadGithubColors()
  colorData := ColorDataFromLanguageData(languageData, &colors)

  return UserLanguageSummary{ColorData: colorData, LanguageData: languageData}
}

func ColorDataFromLanguageData(languageData LanguageData, colors *map[string]string) ColorData {
  var colorData GithubColorData
  for _, githubLanguage := range languageData {
    color := (*colors)[githubLanguage.Name]
    colorData = append(colorData, color)
  }
  return colorData
}

func LoadGithubColors() map[string]string {
  // load language color data
  colorsRaw, loadErr := os.ReadFile("fixtures/GithubColors.json")
  if loadErr != nil {
    fmt.Printf("OS LOAD ERROR: %+v\n", loadErr)
  }
  var colors map[string]string
  colors_err := json.Unmarshal(colorsRaw, &colors)

  if colors_err != nil {
    fmt.Printf("could not load colors")
  }
  return colors
}
