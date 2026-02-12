package auth

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/pquerna/otp/totp"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

type OtpRequest struct {
	Email string `json:"email"`
}

type OtpVerify struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

var pendingOtps = cmap.New[string]()

func SetupAdminAuthV2(mux *http.ServeMux, cookieJar *cmap.ConcurrentMap[string, Cookie], log *zerolog.Logger, oldSiteFileServer http.Handler, adminFileServer http.Handler, marketingFileServer http.Handler) {
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
	
	mux.Handle("/auth/admin/", LogClientIp("/auth/admin/", log, http.StripPrefix("/auth/admin/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Str("path", r.URL.Path).Msg("Auth App Request")
		fs_auth.ServeHTTP(w, r)
	}))))

	// Admin app protected by new auth
	mux.Handle("/admin/", RequireAdminAuthV2(cookieJar, log, adminFileServer))

	// Old site
	mux.Handle("/old-site/", oldSiteFileServer)

	// Marketing site
	mux.Handle("/", marketingFileServer)

	// OTP Routes
	mux.HandleFunc("POST /api/auth/admin/otp/request", func(w http.ResponseWriter, r *http.Request) {
		var req OtpRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		adminEmail := os.Getenv("ADMIN_EMAIL")
		if req.Email != adminEmail {
			log.Warn().Str("email", req.Email).Msg("Unauthorized OTP request attempt")
			// Return success anyway to avoid email enumeration
			w.WriteHeader(http.StatusOK)
			return
		}

		// Generate 6-digit OTP
		otp := fmt.Sprintf("%06d", rand.Intn(1000000))
		pendingOtps.Set(req.Email, otp)

		// Mock email delivery
		log.Info().
			Str("email", req.Email).
			Str("otp", otp).
			Msg("ADMIN OTP GENERATED (Check logs for delivery)")
		
		fmt.Printf("\n--- OTP FOR %s: %s ---\n\n", req.Email, otp)

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("POST /api/auth/admin/otp/verify", func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("OTP Verify Endpoint Hit")
		var req OtpVerify
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error().Err(err).Msg("Error decoding OTP verify request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		log.Info().Str("email", req.Email).Msg("Verifying OTP")

		mode := os.Getenv("MODE")
		isLocalhost := env.IsLocalhost(mode)

		// Check TOTP first if secret exists
		totpSecret := os.Getenv("TOTP_SECRET")
		if totpSecret != "" {
			if totp.Validate(req.Otp, totpSecret) {
				log.Info().Str("email", req.Email).Msg("TOTP validation SUCCESS")
				issueSession(w, r, cookieJar, log, isLocalhost)
				return
			}
			log.Debug().Msg("TOTP validation failed or not used")
		}

		// Check Email OTP
		if storedOtp, ok := pendingOtps.Get(req.Email); ok && storedOtp == req.Otp {
			log.Info().Str("email", req.Email).Msg("Email OTP validation SUCCESS")
			pendingOtps.Remove(req.Email)
			issueSession(w, r, cookieJar, log, isLocalhost)
			return
		}

		log.Warn().Str("email", req.Email).Msg("Invalid OTP attempt")
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

	c := http.Cookie{
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
		http.Redirect(w, r, "/auth/admin/?return_to="+r.URL.Path, http.StatusSeeOther)
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
	http.Redirect(w, r, "/auth/admin/?return_to="+r.URL.Path, http.StatusSeeOther)
}

// RedirectToHome redirects to the home page.
var RedirectToHome = func(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
	log.Warn().
		Str("ip", GetClientIpAddr(r)).
		Msg("REDIRECT To Home")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
