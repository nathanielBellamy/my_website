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
  projectID string
  recaptchaSiteKey string
  token string
  recaptchaAction string
}

type AssessmentBody struct {
  event AssessmentEvent
}

type AssessmentEvent struct {
  token string
  siteKey string
  expectedAction string
}

func ValidateRecaptcha(token string, recaptchaAction string, log *zerolog.Logger) {
  
  rData := RecaptchaData {
    projectID: os.Getenv("RECAPTCHA_PROJECT_ID"),
    recaptchaSiteKey: os.Getenv("RECAPTCHA_SITE_KEY"),
    token: token,
    recaptchaAction: recaptchaAction,
  }
  
  CreateAssessment(rData, log)
}

func CreateAssessment(rData RecaptchaData, log *zerolog.Logger) {
  url := fmt.Sprintf(
    "https://recaptchaenterprise.googleapis.com/v1/projects/%s/assessments?key=%s",
    rData.projectID,
    rData.recaptchaSiteKey,
  )
  assessmentEvent := AssessmentEvent {
    token: rData.token,
    siteKey: rData.recaptchaSiteKey,
    expectedAction: rData.recaptchaAction,
  }
  assessmentBody := AssessmentBody { event: assessmentEvent }
  jsonBody, err := json.Marshal(assessmentBody)
  if err != nil {
    log.Error().
        Err(err).
        Msg("Recaptcha Assessment Json")
  }

  b := bytes.NewBuffer(jsonBody)

  req, err := http.NewRequest(
    http.MethodPost, 
    url, 
    b,
  )

  res, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Error().
        Err(err).
        Msg("Recaptcha Response")
  }
  
  fmt.Printf("Recaptcha Res: %v", res)
}
