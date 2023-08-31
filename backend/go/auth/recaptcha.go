package auth

import (
	"context"
	"fmt"
	"os"

	recaptcha "cloud.google.com/go/recaptchaenterprise/apiv1"
	recaptchapb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"

	"github.com/rs/zerolog"
)

type RecaptchaData struct {
  projectID string
  recaptchaSiteKey string
  token string
  recaptchaAction string
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

/**
* Create an assessment to analyze the risk of an UI action.
*
* @param projectID: GCloud Project ID
* @param recaptchaSiteKey: reCAPTCHA key obtained by registering a domain/app to use recaptcha services.
* @param token: The token obtained from the client on passing the recaptchaSiteKey.
* @param recaptchaAction: Action name corresponding to the token.
*/
func CreateAssessment(rData RecaptchaData, log *zerolog.Logger) {
 // Create the recaptcha client.
 // TODO: To avoid memory issues, move this client generation outside
 // of this example, and cache it (recommended) or call client.close()
 // before exiting this method.
  ctx := context.Background()
  client, err := recaptcha.NewClient(ctx)
  if err != nil {
    fmt.Printf("Error creating reCAPTCHA client\n")
  }
  defer client.Close()

  // Set the properties of the event to be tracked.
  event := &recaptchapb.Event{
    Token: rData.token,
    SiteKey: rData.recaptchaSiteKey,
  }

  assessment := &recaptchapb.Assessment{
    Event: event,
  }

  // Build the assessment request.
  request := &recaptchapb.CreateAssessmentRequest{
    Assessment: assessment,
    Parent: fmt.Sprintf("projects/%s", rData.projectID),
  }

  response, err := client.CreateAssessment(
    ctx,
    request,
  )

  if err != nil {
    fmt.Printf("%v", err.Error())
  }

  // Check if the token is valid.
  if response.TokenProperties.Valid == false {
    fmt.Printf("The CreateAssessment() call failed because the token"+
    " was invalid for the following reasons: %v",
    response.TokenProperties.InvalidReason)
    return
  }

  // Check if the expected action was executed.
  if response.TokenProperties.Action == rData.recaptchaAction {
    log.Info().
        Uint("score", response.RiskAnalysis.Score).
        Msg("reCAPTCHA Response Risk Analysis")
        
    // for _,reason := range response.RiskAnalysis.Reasons {
    //   fmt.Printf(reason.String()+"\n")
    // }

    return response.RiskAnalysis.Score > 50
  }

  log.Error().
      Msg("reCAPTCHA tag action attribute does not match the expected action to score")
}
