package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	crand "crypto/rand"

	"github.com/nathanielBellamy/my_website/backend/go/env"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pquerna/otp/totp"
	"github.com/rs/zerolog"
	"github.com/wneessen/go-mail"
)

type OtpRequest struct {
	// Email is no longer required from the client as we default to ADMIN_EMAIL
}

type OtpVerify struct {
	Otp string `json:"otp"`
}

var pendingOtps = cmap.New[string]()
var pendingChallenges = cmap.New[string]()
var validPreAuthTokens = cmap.New[string]()

func sendEmail(to, subject, body string) error {
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	if host == "" || portStr == "" || user == "" || pass == "" {
		return fmt.Errorf("SMTP configuration missing")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}

	// go-mail handles header sanitization automatically
	msg := mail.NewMsg()
	if err := msg.From(user); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	if err := msg.To(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}
	msg.Subject(subject)
	msg.SetBodyString(mail.TypeTextPlain, body)

	client, err := mail.NewClient(host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(user),
		mail.WithPassword(pass),
	)
	if err != nil {
		return fmt.Errorf("failed to create mail client: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}

	return nil
}

func SetupAdminAuthV2(adminMux, oldSiteMux, marketingMux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, Cookie], log *zerolog.Logger, oldSiteFileServer http.Handler, adminFileServer http.Handler, marketingFileServer http.Handler) {
	// Debug: Log build directory structure
	log.Info().Msg("Checking build directory structure...")
	buildDirs := []string{
		"build/auth/admin",
		"build/auth/admin/browser",
		"build/admin/browser",
		"build/marketing/browser",
	}
	for _, dir := range buildDirs {
		absPath, _ := filepath.Abs(dir)
		if _, err := os.Stat(dir); err == nil {
			log.Info().Str("path", dir).Str("abs", absPath).Msg("Directory EXISTS")
			files, _ := os.ReadDir(dir)
			for _, f := range files {
				log.Debug().Str("parent", dir).Str("file", f.Name()).Msg("Found file")
			}
		} else {
			log.Warn().Str("path", dir).Str("abs", absPath).Msg("Directory NOT FOUND")
		}
	}

	// Serve the new Angular auth app using SpaHandler for client-side routing
	authRoot := "build/auth/admin/browser"
	fs_auth := SpaHandler(authRoot, "index.html")

	adminMux.Handle("/auth/", LogClientIp("/auth/", log, http.StripPrefix("/auth/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Str("path", r.URL.Path).Msg("Auth App Request")
		fs_auth.ServeHTTP(w, r)
	}))))

	// Admin app protected by new auth
	adminMux.Handle("/", RequireAdminAuthV2(cookieJar, log, WithSecurityHeaders(adminFileServer)))

	// Old site
	oldSiteMux.Handle("/", WithSecurityHeaders(oldSiteFileServer))

	// Marketing site
	marketingMux.Handle("/", WithSecurityHeaders(marketingFileServer))

	// Challenge Endpoint
	adminMux.HandleFunc("GET /v1/api/auth/admin/challenge", func(w http.ResponseWriter, r *http.Request) {
		b := make([]byte, 32)
		if _, err := crand.Read(b); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		challenge := hex.EncodeToString(b)

		idBytes := make([]byte, 16)
		if _, err := crand.Read(idBytes); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		challengeID := hex.EncodeToString(idBytes)

		pendingChallenges.Set(challengeID, challenge)

		http.SetCookie(w, &http.Cookie{ // #nosec G124 -- Secure, HttpOnly, SameSite are all set; Secure is conditionally false only for localhost development
			Name:     "nbs-auth-challenge",
			Value:    challengeID,
			Path:     "/v1/api/auth/admin",
			HttpOnly: true,
			Secure:   !env.IsLocalhost(os.Getenv("MODE")),
			SameSite: http.SameSiteStrictMode,
			MaxAge:   300, // 5 minutes
		})

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{"challenge": challenge}); err != nil {
			log.Error().Err(err).Msg("Error encoding challenge response")
		}
	})

	// Password Validation Endpoint
	adminMux.HandleFunc("POST /v1/api/auth/admin/password", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Hash string `json:"hash"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		cookie, err := r.Cookie("nbs-auth-challenge")
		if err != nil {
			http.Error(w, "Missing Challenge Cookie", http.StatusBadRequest)
			return
		}

		challenge, ok := pendingChallenges.Get(cookie.Value)
		if !ok {
			http.Error(w, "Invalid or Expired Challenge", http.StatusBadRequest)
			return
		}
		pendingChallenges.Remove(cookie.Value) // One-time use

		adminPw := os.Getenv("ADMIN_PW")
		if adminPw == "" {
			log.Error().Msg("ADMIN_PW not set")
			http.Error(w, "Configuration Error", http.StatusInternalServerError)
			return
		}

		// Expected: SHA256(ADMIN_PW + challenge)
		hasher := sha256.New()
		hasher.Write([]byte(adminPw + challenge))
		expected := hex.EncodeToString(hasher.Sum(nil))

		if req.Hash != expected {
			log.Warn().Str("ip", GetClientIpAddr(r)).Msg("Password validation failed")
			http.Error(w, "Invalid Password", http.StatusUnauthorized)
			return
		}

		// Success - Set Pre-Auth Cookie
		preAuthToken := make([]byte, 16)
		if _, err := crand.Read(preAuthToken); err != nil {
			log.Error().Err(err).Msg("Failed to generate pre-auth token")
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		tokenStr := hex.EncodeToString(preAuthToken)

		validPreAuthTokens.Set(tokenStr, "valid")

		http.SetCookie(w, &http.Cookie{ // #nosec G124 -- Secure, HttpOnly, SameSite are all set; Secure is conditionally false only for localhost development
			Name:     "nbs-pre-auth",
			Value:    tokenStr,
			Path:     "/v1/api/auth/admin",
			HttpOnly: true,
			Secure:   !env.IsLocalhost(os.Getenv("MODE")),
			SameSite: http.SameSiteStrictMode,
			MaxAge:   300, // 5 minutes
		})

		w.WriteHeader(http.StatusOK)
	})

	// OTP Routes
	adminMux.HandleFunc("POST /v1/api/auth/admin/otp/request", func(w http.ResponseWriter, r *http.Request) {
		// Verify Pre-Auth
		cookie, err := r.Cookie("nbs-pre-auth")
		if err != nil {
			log.Warn().Str("ip", GetClientIpAddr(r)).Msg("OTP Request Missing Pre-Auth")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if _, ok := validPreAuthTokens.Get(cookie.Value); !ok {
			log.Warn().Str("ip", GetClientIpAddr(r)).Msg("OTP Request Invalid Pre-Auth")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// We ignore the body for the request, as we only send to ADMIN_EMAIL

		adminEmail := os.Getenv("ADMIN_EMAIL")
		if adminEmail == "" {
			log.Error().Msg("ADMIN_EMAIL env var not set")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Generate 6-digit OTP
		n, err := crand.Int(crand.Reader, big.NewInt(1000000))
		if err != nil {
			log.Error().Err(err).Msg("Failed to generate OTP")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		otp := fmt.Sprintf("%06d", n.Int64())
		pendingOtps.Set(adminEmail, otp)

		// Email delivery
		log.Info().
			Str("email", adminEmail).
			Str("otp", otp).
			Msg("ADMIN OTP GENERATED")

		fmt.Printf("\n--- OTP FOR %s: %s ---\n\n", adminEmail, otp)

		// Send real email
		err = sendEmail(adminEmail, "Your Admin OTP", fmt.Sprintf("Your OTP is: %s", otp))
		if err != nil {
			log.Error().Err(err).Msg("Failed to send OTP email")
			// We might not want to fail the request if email fails, but for now let's just log it.
			// Ideally we return an error so the UI knows.
			// For localhost, if SMTP isn't set up, this will error.
		} else {
			log.Info().Msg("OTP Email sent successfully")
		}

		w.WriteHeader(http.StatusOK)
	})

	adminMux.HandleFunc("POST /v1/api/auth/admin/otp/verify", func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("OTP Verify Endpoint Hit")
		var req OtpVerify
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error().Err(err).Msg("Error decoding OTP verify request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		adminEmail := os.Getenv("ADMIN_EMAIL")
		log.Info().Str("email", adminEmail).Msg("Verifying OTP")

		mode := os.Getenv("MODE")
		isLocalhost := env.IsLocalhost(mode)

		// Check TOTP first if secret exists
		totpSecret := os.Getenv("TOTP_SECRET")
		if totpSecret != "" {
			if totp.Validate(req.Otp, totpSecret) {
				log.Info().Str("email", adminEmail).Msg("TOTP validation SUCCESS")
				issueSession(w, r, cookieJar, log, isLocalhost)
				return
			}
			log.Debug().Msg("TOTP validation failed or not used")
		}

		// Check Email OTP
		if storedOtp, ok := pendingOtps.Get(adminEmail); ok && storedOtp == req.Otp {
			log.Info().Str("email", adminEmail).Msg("Email OTP validation SUCCESS")
			pendingOtps.Remove(adminEmail)
			issueSession(w, r, cookieJar, log, isLocalhost)
			return
		}

		log.Warn().Str("email", adminEmail).Msg("Invalid OTP attempt")
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
	})
}

func issueSession(w http.ResponseWriter, r *http.Request, cookieJar *cmap.ConcurrentMap[string, Cookie], log *zerolog.Logger, isLocalhost bool) {
	var h Hash
	sessionToken, _ := h.Generate(time.Now().String())

	var name string
	if isLocalhost {
		name = "nbs-admin"
	} else {
		name = "__Secure-nbs-admin"
	}

	c := http.Cookie{ // #nosec G124 -- Secure, HttpOnly, SameSite are all set; Secure is conditionally false only for localhost development
		Name:     name,
		Value:    sessionToken,
		Path:     "/",
		MaxAge:   60 * 60 * 48,
		Secure:   !isLocalhost,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	cookieJar.Set(sessionToken, Cookie{Valid: true, Type: CTADMIN})
	http.SetCookie(w, &c)

	log.Info().Msg("Admin login success")
	w.WriteHeader(http.StatusOK)
}

func RequireAdminAuthV2(cookieJar *cmap.ConcurrentMap[string, Cookie], log *zerolog.Logger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mode := os.Getenv("MODE")
		enableAuthLocal := os.Getenv("ENABLE_AUTH_LOCAL") == "true"

		log.Debug().
			Str("mode", mode).
			Bool("enableAuthLocal", enableAuthLocal).
			Bool("isLocalhost", env.IsLocalhost(mode)).
			Msg("RequireAdminAuthV2 check")

		if env.IsLocalhost(mode) && !enableAuthLocal {
			log.Debug().Msg("Bypassing auth for localhost")
			handler.ServeHTTP(w, r)
			return
		}

		if HasValidCookieV2(r, CTADMIN, cookieJar, log) {
			handler.ServeHTTP(w, r)
			return
		}

		log.Warn().Str("ip", GetClientIpAddr(r)).Msg("Admin auth required")
		http.Redirect(w, r, "/auth/?return_to="+r.URL.Path, http.StatusSeeOther)
	})
}

func HasValidCookieV2(r *http.Request, cookieType CookieType, cookieJar *cmap.ConcurrentMap[string, Cookie], log *zerolog.Logger) bool {
	var cookieName string
	mode := os.Getenv("MODE")
	if env.IsLocalhost(mode) {
		cookieName = "nbs-admin"
	} else {
		cookieName = "__Secure-nbs-admin"
	}

	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return false
	}

	if cookieJar.Has(cookie.Value) {
		cookieFromJar, ok := cookieJar.Get(cookie.Value)
		return ok && cookieFromJar.Valid && cookieFromJar.Type == cookieType
	}

	return false
}

// RedirectToAdminAuthV2 redirects to the new admin auth page.
var RedirectToAdminAuthV2 = func(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
	log.Warn().
		Str("ip", GetClientIpAddr(r)).
		Msg("REDIRECT To Admin Auth")
	http.Redirect(w, r, "/auth/?return_to="+r.URL.Path, http.StatusSeeOther)
}

// RedirectToHome redirects to the home page.
var RedirectToHome = func(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
	log.Warn().
		Str("ip", GetClientIpAddr(r)).
		Msg("REDIRECT To Home")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func WithSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Security Headers
		// Note: strict-transport-security is handled by Nginx in prod, but good to have if we ever serve HTTPS directly
		// w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		// CSP: Allow Google Recaptcha, YouTube embeds, self, and inline styles/scripts (for Angular)
		// We allow ws: for localhost development
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval' https://www.google.com/recaptcha/ https://www.gstatic.com/recaptcha/; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self' data:; frame-src https://www.google.com/recaptcha/ https://www.youtube.com/; connect-src 'self' ws: wss: https://www.google.com/recaptcha/; frame-ancestors 'none'; base-uri 'self';")

		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")

		next.ServeHTTP(w, r)
	})
}
