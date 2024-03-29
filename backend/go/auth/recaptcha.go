package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/env"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

type RecaptchaData struct {
  ClientIP string
  ProjectID string
  Token string
  RecaptchaAction string
  RecaptchaSiteKey string
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
  ip := GetClientIpAddr(r)
  var jsonData FrontendJson
  
  err := json.NewDecoder(r.Body).Decode(&jsonData)
  if err != nil {
    log.Error().
        Err(err).
        Str("ip", ip).
        Msg("Error Decoding Recaptcha Payload")
    return false
  }

  log.Info().
      Str("Action", jsonData.Action).
      Str("Token", jsonData.Token).
      Str("ip", ip).
      Msg("Recaptcha Frontend Json")

  projectId := os.Getenv("RECAPTCHA_PROJECT_ID")
  siteKey := os.Getenv("RECAPTCHA_SITE_KEY")
  
  log.Info().
      Str("projectId", projectId).
      Str("siteKey", siteKey).
      Str("ip", ip).
      Msg("Recaptcha Env Vars")

  rData := RecaptchaData {
    ClientIP: ip,
    ProjectID: projectId,
    RecaptchaSiteKey: siteKey,
    Token: jsonData.Token,
    RecaptchaAction: jsonData.Action,
  }
  
  return CreateAssessment(rData, log)
}

func CreateAssessment(rData RecaptchaData, log *zerolog.Logger) bool {
  log.Info().
      Str("clientIP", rData.ClientIP).
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

  log.Info().
      Str("ip", rData.ClientIP).
      Str("url", url).
      Msg("Recaptcha URL")

  assessmentEvent := AssessmentEvent {
    Token: rData.Token,
    SiteKey: rData.RecaptchaSiteKey,
    ExpectedAction: rData.RecaptchaAction,
  }
  assessmentBody := AssessmentBody { Event: assessmentEvent }
  
  log.Info().
      Str("ip", rData.ClientIP).
      Any("assessmentBody", assessmentBody).
      Msg("CreateAssessment")

  jsonBody, err := json.Marshal(assessmentBody)

  if err != nil {
    log.Error().
        Err(err).
        Str("ip", rData.ClientIP).
        Msg("Recaptcha Assessment Json")
    return false
  }

  b := bytes.NewBuffer(jsonBody)

  client := http.Client{}
  response, err := client.Post(url, `json`, b)
  if err != nil {
    log.Error().
        Err(err).
        Str("ip", rData.ClientIP).
        Msg("Recaptcha CreateAssesment Client Resp")
    return false
  }
  defer response.Body.Close()

  var assessmentResp AssessmentResp
  assessmentErr := json.NewDecoder(response.Body).Decode(&assessmentResp)
  if assessmentErr != nil {
    log.Error().
        Str("ip", rData.ClientIP).
        Err(assessmentErr).
        Msg("Recaptcha Assement Response")
    return false
  }

  log.Info().
      Str("ip", rData.ClientIP).
      Any("response", assessmentResp).
      Msg("Recaptcha Assessment Resp")

  return assessmentResp.RiskAnalysis.Score > 0.5 // closer to 1.0 is safer/more likely legit
}

func SetRecaptchaCookieOnClient(
  w http.ResponseWriter, 
  cookieJar *cmap.ConcurrentMap[string, Cookie], 
  log *zerolog.Logger,
) {
  mode := os.Getenv("MODE")
  isLocalhost := env.IsLocalhost(mode)
  var name string 
  if isLocalhost {
    name = "nbs-psr"
  } else {
    name = "__Secure-nbs-psr"
  }
  // set cookie on client
  var h Hash
  sessionToken, err := h.Generate(time.Now().String())
  if err != nil {
    log.Error().
        Err(err).
        Msg("Generate PS Cookie Token")
    return 
  }

  c := http.Cookie {
    Name: name,
    Value: sessionToken,
    Path: "/",
    MaxAge: 60 * 60 * 2, // two hours or whenever the server restarts as cookieJar is in-memory
    Secure: !isLocalhost, // https only
    HttpOnly: true, // don't let JS touch it
    SameSite: http.SameSiteLaxMode,
  }

  cookieJar.SetIfAbsent(sessionToken, Cookie{Valid: true, Type: CTPSR})
  http.SetCookie(w, &c)
}
