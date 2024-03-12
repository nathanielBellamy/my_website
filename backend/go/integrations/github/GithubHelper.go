package github

import (
  "encoding/json"
  "os"
  // "slices"
  "sort"
)

type GithubHelperError struct {
  Msg string
}

func (ghe GithubHelperError) Error() (string) {
  return "GithubHelperError: " + ghe.Msg
}

func GenerateUserLanguageSummary(githubRepos GithubRepos) (UserLanguageSummary, error) {
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
          break
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

  colors, err := LoadGithubColors()
  if err != nil {
    return UserLanguageSummary{}, GithubHelperError{Msg: err.Error()}
  }
  colorData := ColorDataFromLanguageData(languageData, &colors)

  return UserLanguageSummary{ColorData: colorData, LanguageData: languageData}, nil
}

func ColorDataFromLanguageData(languageData LanguageData, colors *map[string]string) ColorData {
  var colorData GithubColorData
  for _, githubLanguage := range languageData {
    color := (*colors)[githubLanguage.Name]
    colorData = append(colorData, color)
  }
  return colorData
}

func LoadGithubColors() (map[string]string, error) {
  // load language color data
  colorsRaw, loadErr := os.ReadFile("fixtures/GithubColors.json")
  if loadErr != nil {
    return nil, GithubHelperError{Msg: " LoadGithubColors: " + loadErr.Error()}
  }
  var colors map[string]string
  jsonErr := json.Unmarshal(colorsRaw, &colors)

  if jsonErr != nil {
    return nil, GithubHelperError{Msg: " LoadGithubColors: " + jsonErr.Error()}
  }
  return colors, nil
}
