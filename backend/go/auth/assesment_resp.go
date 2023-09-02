package auth

import (
  // "encoding/json"
)

// Example response from Google's docs
// {
//  "event":{
//     "expectedAction":"EXPECTED_ACTION",
//     "hashedAccountId":"ACCOUNT_ID",
//     "siteKey":"SITE_KEY",
//     "token":"TOKEN",
//     "userAgent":"(USER-PROVIDED STRING)",
//     "userIpAddress":"USER_PROVIDED_IP_ADDRESS"
//  },
//  "name":"ASSESSMENT_ID",
//  "riskAnalysis":{
//    "reasons":[],
//    "score":"SCORE"
//  },
//  "tokenProperties":{
//    "action":"USER_INTERACTION",
//    "createTime":"TIMESTAMP",
//    "hostname":"HOSTNAME",
//    "invalidReason":"(ENUM)",
//    "valid":(BOOLEAN)
//  }
// }

type AssessmentResp struct {
  Event AssessmentRespEvent `json:"event"`
  Name string `json:"name"`
  RiskAnalysis AssessmentRiskAnalysis `json:"riskAnalysis"`
  TokenProperties AssessmentTokenProperties `json:"tokenProperties"`
}

type AssessmentRespEvent struct {
  ExpectedAction string `json:"expectedAction"`
  HashedAccountId string `json:"hashedAccountId"`
  SiteKey string `json:"siteKey"`
  Token string `json:"token"`
  UserAgent string `json:"userAgent"`
  UserIpAddress string `json:"userIpAddress`
}

type AssessmentRiskAnalysis struct {
   Reasons []string `json:"reasons"`
   Score float32 `json:"score"`
}

type AssessmentTokenProperties struct {
  Action string `json:"action"`
  CreateTime string `json:"createTime"`
  Hostname string `json:"hostname"`
  InvalidREason string `json:"invalidReason"`
  valid bool `json:"valid"`
}


