package old_site

import (
	"net/http"
	"os"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/env"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

// OldSiteController holds dependencies for old-site handlers.
type OldSiteController struct {
	CookieJar *cmap.ConcurrentMap[string, auth.Cookie]
	Log       *zerolog.Logger
	FeedPool  *websocket.Pool
	WasmPool  *websocket.Pool
}

// NewOldSiteController creates and returns a new OldSiteController.
func NewOldSiteController(cookieJar *cmap.ConcurrentMap[string, auth.Cookie], log *zerolog.Logger, feedPool *websocket.Pool, wasmPool *websocket.Pool) *OldSiteController {
	return &OldSiteController{
		CookieJar: cookieJar,
		Log:       log,
		FeedPool:  feedPool,
		WasmPool:  wasmPool,
	}
}

// RecaptchaHandler handles Recaptcha validation for the old site.
func (osc *OldSiteController) RecaptchaHandler(w http.ResponseWriter, r *http.Request) {
	ip := auth.GetClientIpAddr(r)
	osc.Log.Info().
		Str("ip", ip).
		Msg("Recaptcha Endpoint Hit")

	res := auth.ValidateRecaptcha(r, osc.Log)
	osc.Log.Info().
		Str("ip", ip).
		Bool("res", res).
		Msg("ValidateRecaptcha")

	if res {
		auth.SetRecaptchaCookieOnClient(w, osc.CookieJar, osc.Log)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("NOT OK"))
	}
}

// PublicSquareFeedWsHandler handles the WebSocket connection for the public square feed.
func (osc *OldSiteController) PublicSquareFeedWsHandler(w http.ResponseWriter, r *http.Request) {
	ip := auth.GetClientIpAddr(r)
	mode := os.Getenv("MODE")
	if !env.IsProd(mode) {
		// localhost and remote dev require basic login
		validAdmin := auth.HasValidCookie(r, auth.CTADMIN, osc.CookieJar, osc.Log)
		validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, osc.CookieJar, osc.Log)
		osc.Log.Debug().
			Bool("debug_validAdmin", validAdmin).
			Bool("debug_validRecaptcha", validRecaptcha).
			Msg("DEBUG HasValidCookie values")
		osc.Log.Info().
			Str("ip", ip).
			Bool("validAdmin", validAdmin).
			Bool("validRecaptcha", validRecaptcha).
			Msg("PS FEED WS")

		if validAdmin && validRecaptcha {
			websocket.ServeFeedWs(osc.FeedPool, w, r, osc.Log)
		} else {
			auth.RedirectToAdminAuthV2(w, r, osc.Log)
		}
	} else {
		// prod is public
		// but protected by recaptcha
		validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, osc.CookieJar, osc.Log)
		osc.Log.Info().
			Str("ip", ip).
			Bool("validRecaptcha", validRecaptcha).
			Msg("PS FEED WS")

		if validRecaptcha {
			websocket.ServeFeedWs(osc.FeedPool, w, r, osc.Log)
		} else {
			auth.RedirectToHome(w, r, osc.Log)
		}
	}
}

// PublicSquareWasmWsHandler handles the WebSocket connection for the public square WASM.
func (osc *OldSiteController) PublicSquareWasmWsHandler(w http.ResponseWriter, r *http.Request) {
	ip := auth.GetClientIpAddr(r)
	mode := os.Getenv("MODE")
	if !env.IsProd(mode) {
		// localhost and remote dev require basic login
		validAdmin := auth.HasValidCookie(r, auth.CTADMIN, osc.CookieJar, osc.Log)
		validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, osc.CookieJar, osc.Log)
		osc.Log.Debug().
			Bool("debug_validAdmin", validAdmin).
			Bool("debug_validRecaptcha", validRecaptcha).
			Msg("DEBUG HasValidCookie values")
		osc.Log.Info().
			Str("ip", ip).
			Bool("validAdmin", validAdmin).
			Bool("validRecaptcha", validRecaptcha).
			Msg("PS WASM WS")

		if validAdmin && validRecaptcha {
			websocket.ServeWasmWs(osc.WasmPool, w, r, osc.Log)
		} else {
			auth.RedirectToAdminAuthV2(w, r, osc.Log)
		}
	} else {
		// prod is public
		// but protected by recaptcha
		validRecaptcha := auth.HasValidCookie(r, auth.CTPSR, osc.CookieJar, osc.Log)
		osc.Log.Info().
			Str("ip", ip).
			Bool("validRecaptcha", validRecaptcha).
			Msg("PS WASM WS")

		if validRecaptcha {
			websocket.ServeWasmWs(osc.WasmPool, w, r, osc.Log)
		} else {
			auth.RedirectToHome(w, r, osc.Log)
		}
	}
}

// OldSiteFileServer serves static files for the old site.
func (osc *OldSiteController) OldSiteFileServer() http.Handler {
	rootPath := "build/old-site"
	root := http.Dir(rootPath)
	fs := http.FileServer(root)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to open file
		f, err := root.Open(r.URL.Path)
		if os.IsNotExist(err) {
			// Not found, serve index.html for SPA routing
			http.ServeFile(w, r, rootPath+"/index.html")
			return
		}
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// If it exists, let FileServer handle it
		fs.ServeHTTP(w, r)
	})

	return auth.LogClientIp("/", osc.Log, handler)
}
