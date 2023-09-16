package auth

type CookieType int32

const (
    CTPSR CookieType = 0 // Public Square Recaptcha
    CTDEV CookieType = 1 // Dev Site Basic Login
)

type Cookie struct {
  Valid bool
  Type CookieType
}

