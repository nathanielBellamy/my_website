package auth

type CookieType int32

const (
	CTPSR   CookieType = 0 // Public Square Recaptcha
	CTDEV   CookieType = 1 // Dev Site Basic Login (LEGACY)
	CTADMIN CookieType = 2 // Admin Site OTP Login
)

type Cookie struct {
	Valid bool
	Type  CookieType
}
