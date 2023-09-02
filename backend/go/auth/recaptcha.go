package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

type RecaptchaData struct {
  ProjectID string
  RecaptchaSiteKey string
  Token string
  RecaptchaAction string
}

type AssessmentBody struct {
  Event AssessmentEvent `json:"event"`
}

type AssessmentEvent struct {
  Token string `json:"token"`
  SiteKey string `json:"siteKey"`
  ExpectedAction string `json:"expectedAction"`
}

type FrontendJson struct {
  Action string `json:"action"`
  Token  string `json:"token"`
}

func ValidateRecaptcha(r *http.Request, log *zerolog.Logger) bool {
  var jsonData FrontendJson
  
  err := json.NewDecoder(r.Body).Decode(&jsonData)
  if err != nil {
    log.Error().
        Err(err).
        Msg("Error Decoding Recaptcha Payload")
    return false
  }

  log.Info().
      Str("Action", jsonData.Action).
      Str("Token", jsonData.Token).
      Msg("Recaptcha Frontend Json")

  projectId := os.Getenv("RECAPTCHA_PROJECT_ID")
  siteKey := os.Getenv("RECAPTCHA_SITE_KEY")
  
  log.Info().
      Str("projectId", projectId).
      Str("siteKey", siteKey).
      Msg("Recaptcha Env Vars")

  rData := RecaptchaData {
    ProjectID: projectId,
    RecaptchaSiteKey: siteKey,
    Token: jsonData.Token,
    RecaptchaAction: jsonData.Action,
  }
  
  return CreateAssessment(rData, log)
}

func CreateAssessment(rData RecaptchaData, log *zerolog.Logger) bool {
  log.Info().
      Str("projectId", rData.ProjectID).
      Str("siteKey", rData.RecaptchaSiteKey).
      Str("recaptchaAction", rData.RecaptchaAction).
      Str("token", rData.Token).
      Msg("CreateAssessment rData")

  apiKey := os.Getenv("GOOGLE_API_KEY")
  url := fmt.Sprintf(
    "https://recaptchaenterprise.googleapis.com/v1/projects/%s/assessments?key=%s",
    rData.ProjectID,
    apiKey,
  )
  assessmentEvent := AssessmentEvent {
    Token: rData.Token,
    SiteKey: rData.RecaptchaSiteKey,
    ExpectedAction: rData.RecaptchaAction,
  }
  assessmentBody := AssessmentBody { Event: assessmentEvent }
  
  log.Info().
      Any("assessmentBody", assessmentBody).
      Msg("CreateAssessment")

  jsonBody, err := json.Marshal(assessmentBody)

  if err != nil {
    log.Error().
        Err(err).
        Msg("Recaptcha Assessment Json")
    return false
  }

  b := bytes.NewBuffer(jsonBody)

  client := http.Client{}
  response, err := client.Post(url, `json`, b)
  if err != nil {
    log.Error().
        Err(err).
        Msg("Recaptcha CreateAssesment Client Resp")
    return false
  }
  defer response.Body.Close()

  var assessmentResp AssessmentResp
  assessmentErr := json.NewDecoder(response.Body).Decode(&assessmentResp)
  if assessmentErr != nil {
    log.Error().
        Err(assessmentErr).
        Msg("Recaptcha Assement Response")
    return false
  }

  log.Info().
      Any("response", assessmentResp).
      Msg("Recaptcha Assessment Resp")
  return true
}
