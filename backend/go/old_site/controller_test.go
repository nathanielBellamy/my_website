package old_site

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/websocket"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/rs/zerolog"
)

// MockLogger is a mock implementation of zerolog.Logger for testing.
type MockLogger struct {
	Buf bytes.Buffer
}

func (m *MockLogger) Write(p []byte) (n int, err error) {
	return m.Buf.Write(p)
}

// MockPool is a mock implementation of websocket.Pool for testing.
type MockPool struct{}

func (m *MockPool) StartFeed() {}
func (m *MockPool) StartWasm() {}

func TestRecaptchaHandler(t *testing.T) {
	// Setup
	cookieJar := cmap.New[auth.Cookie]()
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	osc := NewOldSiteController(&cookieJar, &log, &websocket.Pool{}, &websocket.Pool{})

	// Save original functions and defer their restoration
	origValidateRecaptcha := auth.ValidateRecaptcha
	origSetRecaptchaCookieOnClient := auth.SetRecaptchaCookieOnClient
	t.Cleanup(func() {
		auth.ValidateRecaptcha = origValidateRecaptcha
		auth.SetRecaptchaCookieOnClient = origSetRecaptchaCookieOnClient
		t.Log(mockLogOutput.Buf.String()) // Log the buffer content
	})

	// Test Case 1: Successful Recaptcha validation
	req, err := http.NewRequest("POST", "/old-site/recaptcha", strings.NewReader("g-recaptcha-response=mock_success_token"))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	auth.ValidateRecaptcha = func(r *http.Request, l *zerolog.Logger) bool { return true }
	auth.SetRecaptchaCookieOnClient = func(w http.ResponseWriter, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) {
		http.SetCookie(w, &http.Cookie{Name: "recaptcha", Value: "valid"})
	}

	osc.RecaptchaHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if body := rr.Body.String(); body != "OK" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, "OK")
	}
	if cookie := rr.Header().Get("Set-Cookie"); !strings.Contains(cookie, "recaptcha=valid") {
		t.Errorf("handler did not set recaptcha cookie")
	}

	// Test Case 2: Failed Recaptcha validation
	req, err = http.NewRequest("POST", "/old-site/recaptcha", strings.NewReader("g-recaptcha-response=mock_fail_token"))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	auth.ValidateRecaptcha = func(r *http.Request, l *zerolog.Logger) bool { return false }
	auth.SetRecaptchaCookieOnClient = func(w http.ResponseWriter, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) {
		// Should not be called on failure
	}

	osc.RecaptchaHandler(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
	if body := rr.Body.String(); body != "NOT OK" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, "NOT OK")
	}
	if cookie := rr.Header().Get("Set-Cookie"); strings.Contains(cookie, "recaptcha=") {
		t.Errorf("handler set recaptcha cookie on failure")
	}
}

func TestPublicSquareFeedWsHandler(t *testing.T) {
	// Save original functions and defer their restoration
	origHasValidCookie := auth.HasValidCookie
	origRedirectToDevAuth := auth.RedirectToDevAuth
	origRedirectToHome := auth.RedirectToHome
	origServeFeedWs := websocket.ServeFeedWs
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).With().Logger().Level(zerolog.DebugLevel)
	t.Cleanup(func() {
		auth.HasValidCookie = origHasValidCookie
		auth.RedirectToDevAuth = origRedirectToDevAuth
		auth.RedirectToHome = origRedirectToHome
		websocket.ServeFeedWs = origServeFeedWs
		t.Log(mockLogOutput.Buf.String()) // Log the buffer content
	})

	cookieJar := cmap.New[auth.Cookie]()
	osc := NewOldSiteController(&cookieJar, &log, &websocket.Pool{}, &websocket.Pool{})

	// Test Case 1: Non-prod, valid dev and recaptcha
	t.Setenv("MODE", "localhost") // Changed from "dev" to "localhost"
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return true
	}
	redirectCalled := false
	auth.RedirectToDevAuth = func(w http.ResponseWriter, r *http.Request, l *zerolog.Logger) { redirectCalled = true }
	auth.RedirectToHome = func(w http.ResponseWriter, r *http.Request, l *zerolog.Logger) { redirectCalled = true }
	serveWsCalled := false
	websocket.ServeFeedWs = func(p *websocket.Pool, w http.ResponseWriter, r *http.Request, l *zerolog.Logger) {
		serveWsCalled = true
	}

	req, _ := http.NewRequest("GET", "/old-site/public-square-feed-ws", nil)
	rr := httptest.NewRecorder()
	osc.PublicSquareFeedWsHandler(rr, req)

	if redirectCalled {
		t.Errorf("Non-prod, valid dev/recaptcha: RedirectToDevAuth or RedirectToHome unexpectedly called")
	}
	if !serveWsCalled {
		t.Errorf("Non-prod, valid dev/recaptcha: ServeFeedWs not called")
	}

	// Test Case 2: Non-prod, invalid dev
	t.Setenv("MODE", "localhost") // Changed from "dev" to "localhost"
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return ct != auth.CTDEV
	}
	redirectCalled = false
	serveWsCalled = false

	req, _ = http.NewRequest("GET", "/old-site/public-square-feed-ws", nil)
	rr = httptest.NewRecorder()
	osc.PublicSquareFeedWsHandler(rr, req)

	if !redirectCalled {
		t.Errorf("Non-prod, invalid dev: RedirectToDevAuth not called")
	}
	if serveWsCalled {
		t.Errorf("Non-prod, invalid dev: ServeFeedWs unexpectedly called")
	}

	// Test Case 3: Prod, valid recaptcha
	t.Setenv("MODE", "prod")
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return true
	}
	redirectCalled = false
	serveWsCalled = false

	req, _ = http.NewRequest("GET", "/old-site/public-square-feed-ws", nil)
	rr = httptest.NewRecorder()
	osc.PublicSquareFeedWsHandler(rr, req)

	if redirectCalled {
		t.Errorf("Prod, valid recaptcha: RedirectToHome unexpectedly called")
	}
	if !serveWsCalled {
		t.Errorf("Prod, valid recaptcha: ServeFeedWs not called")
	}

	// Test Case 4: Prod, invalid recaptcha
	t.Setenv("MODE", "prod")
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return false
	}
	redirectCalled = false
	serveWsCalled = false

	req, _ = http.NewRequest("GET", "/old-site/public-square-feed-ws", nil)
	rr = httptest.NewRecorder()
	osc.PublicSquareFeedWsHandler(rr, req)

	if !redirectCalled {
		t.Errorf("Prod, invalid recaptcha: RedirectToHome not called")
	}
	if serveWsCalled {
		t.Errorf("Prod, invalid recaptcha: ServeFeedWs unexpectedly called")
	}
}

func TestPublicSquareWasmWsHandler(t *testing.T) {
	// Save original functions and defer their restoration
	origHasValidCookie := auth.HasValidCookie
	origRedirectToDevAuth := auth.RedirectToDevAuth
	origRedirectToHome := auth.RedirectToHome
	origServeWasmWs := websocket.ServeWasmWs
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).With().Logger().Level(zerolog.DebugLevel)
	t.Cleanup(func() {
		auth.HasValidCookie = origHasValidCookie
		auth.RedirectToDevAuth = origRedirectToDevAuth
		auth.RedirectToHome = origRedirectToHome
		websocket.ServeWasmWs = origServeWasmWs
		t.Log(mockLogOutput.Buf.String()) // Log the buffer content
	})

	cookieJar := cmap.New[auth.Cookie]()
	osc := NewOldSiteController(&cookieJar, &log, &websocket.Pool{}, &websocket.Pool{})

	// Test Case 1: Non-prod, valid dev and recaptcha
	t.Setenv("MODE", "localhost") // Changed from "dev" to "localhost"
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return true
	}
	redirectCalled := false
	auth.RedirectToDevAuth = func(w http.ResponseWriter, r *http.Request, l *zerolog.Logger) { redirectCalled = true }
	auth.RedirectToHome = func(w http.ResponseWriter, r *http.Request, l *zerolog.Logger) { redirectCalled = true }
	serveWsCalled := false
	websocket.ServeWasmWs = func(p *websocket.Pool, w http.ResponseWriter, r *http.Request, l *zerolog.Logger) {
		serveWsCalled = true
	}

	req, _ := http.NewRequest("GET", "/old-site/public-square-wasm-ws", nil)
	rr := httptest.NewRecorder()
	osc.PublicSquareWasmWsHandler(rr, req)

	if redirectCalled {
		t.Errorf("Non-prod, valid dev/recaptcha: RedirectToDevAuth or RedirectToHome unexpectedly called")
	}
	if !serveWsCalled {
		t.Errorf("Non-prod, valid dev/recaptcha: ServeWasmWs not called")
	}

	// Test Case 2: Non-prod, invalid dev
	t.Setenv("MODE", "localhost") // Changed from "dev" to "localhost"
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return ct != auth.CTDEV
	}
	redirectCalled = false
	serveWsCalled = false

	req, _ = http.NewRequest("GET", "/old-site/public-square-wasm-ws", nil)
	rr = httptest.NewRecorder()
	osc.PublicSquareWasmWsHandler(rr, req)

	if !redirectCalled {
		t.Errorf("Non-prod, invalid dev: RedirectToDevAuth not called")
	}
	if serveWsCalled {
		t.Errorf("Non-prod, invalid dev: ServeWasmWs unexpectedly called")
	}

	// Test Case 3: Prod, valid recaptcha
	t.Setenv("MODE", "prod")
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return true
	}
	redirectCalled = false
	serveWsCalled = false

	req, _ = http.NewRequest("GET", "/old-site/public-square-wasm-ws", nil)
	rr = httptest.NewRecorder()
	osc.PublicSquareWasmWsHandler(rr, req)

	if redirectCalled {
		t.Errorf("Prod, valid recaptcha: RedirectToHome unexpectedly called")
	}
	if !serveWsCalled {
		t.Errorf("Prod, valid recaptcha: ServeWasmWs not called")
	}

	// Test Case 4: Prod, invalid recaptcha
	t.Setenv("MODE", "prod")
	auth.HasValidCookie = func(r *http.Request, ct auth.CookieType, cj *cmap.ConcurrentMap[string, auth.Cookie], l *zerolog.Logger) bool {
		return false
	}
	redirectCalled = false
	serveWsCalled = false

	req, _ = http.NewRequest("GET", "/old-site/public-square-wasm-ws", nil)
	rr = httptest.NewRecorder()
	osc.PublicSquareWasmWsHandler(rr, req)

	if !redirectCalled {
		t.Errorf("Prod, invalid recaptcha: RedirectToHome not called")
	}
	if serveWsCalled {
		t.Errorf("Prod, invalid recaptcha: ServeWasmWs unexpectedly called")
	}
}

func TestOldSiteFileServer(t *testing.T) {
	// Setup
	cookieJar := cmap.New[auth.Cookie]()
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()
	osc := NewOldSiteController(&cookieJar, &log, &websocket.Pool{}, &websocket.Pool{})
	t.Cleanup(func() {
		t.Log(mockLogOutput.Buf.String()) // Log the buffer content
	})

	handler := osc.OldSiteFileServer()

	// Ensure the returned handler is not nil
	if handler == nil {
		t.Fatal("OldSiteFileServer returned a nil handler")
	}

	// You might want to do more extensive testing here,
	// but for basic unit testing, ensuring it returns a handler is a good start.
	// For actual file serving tests, integration tests would be more appropriate.
}
