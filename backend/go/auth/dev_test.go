package auth

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"



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

func TestMain(m *testing.M) {
	// Set MODE to localhost for testing purposes
	os.Setenv("MODE", "localhost")
	// Set a dummy password for testing hash comparison
	os.Setenv("PW", "testpassword")
	// Run tests
	code := m.Run()
	// Clean up
	os.Unsetenv("MODE")
	os.Unsetenv("PW")
	os.Exit(code)
}

func TestSetupDevAuth_DevAuthHandler(t *testing.T) {
	// Mock logger
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()

	// Mock dependencies
	cookieJar := cmap.New[Cookie]()
	mockOldSiteFileServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "old site content")
	})

	// Temporarily override ValidateDev and RedirectToDevAuth for isolation
	originalValidateDev := ValidateDev
	originalRedirectToDevAuth := RedirectToDevAuth
	defer func() {
		ValidateDev = originalValidateDev
		RedirectToDevAuth = originalRedirectToDevAuth
	}()

	// Mock ValidateDev for successful login
	ValidateDev = func(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) (string, bool) {
		return "test_session_token", true
	}
	// Mock RedirectToDevAuth to just write a status for testing purposes
	RedirectToDevAuth = func(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Redirected to Dev Auth")
	}

	// Create a new ServeMux for testing
	testMux := http.NewServeMux()

	// Call SetupDevAuth to register handlers
	SetupDevAuth(testMux, &cookieJar, &log, mockOldSiteFileServer)

	t.Run("successful login", func(t *testing.T) {
		form := url.Values{}
		form.Add("pw", "testpassword")
		req := httptest.NewRequest("POST", "/auth/dev/dev-auth", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		testMux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusFound { // Should redirect on success
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
		}
		if location := rr.Header().Get("Location"); location != "/" {
			t.Errorf("handler returned wrong redirect location: got %q want %q", location, "/")
		}
		// Check if cookie is set
		cookies := rr.Result().Cookies()
		if len(cookies) == 0 {
			t.Fatal("Expected cookie to be set, but got none")
		}
		if cookies[0].Name != "nbs-dev" {
			t.Errorf("Expected cookie name %q, got %q", "nbs-dev", cookies[0].Name)
		}
		if cookies[0].Value != "test_session_token" {
			t.Errorf("Expected cookie value %q, got %q", "test_session_token", cookies[0].Value)
		}
		if !cookieJar.Has("test_session_token") {
			t.Errorf("Expected session token to be in cookieJar, but it's not")
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset() // Clear buffer for next test
	})

	t.Run("failed login", func(t *testing.T) {
		ValidateDev = func(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) (string, bool) {
			return "", false // Simulate validation failure
		}

		form := url.Values{}
		form.Add("pw", "wrongpassword")
		req := httptest.NewRequest("POST", "/auth/dev/dev-auth", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		testMux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusServiceUnavailable { // Based on current implementation
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusServiceUnavailable)
		}
		body, _ := io.ReadAll(rr.Body)
		if string(body) != "Invalid Password\n" {
			t.Errorf("handler returned unexpected body: got %q want %q", string(body), "Invalid Password\n")
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset() // Clear buffer for next test
	})
}

func TestRequireDevAuth(t *testing.T) {
	// Mock logger
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()

	// Mock dependencies
	cookieJar := cmap.New[Cookie]()
	
	// Create a dummy handler that RequireDevAuth will protect
	protectedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Protected content")
	})

	// Temporarily override RedirectToDevAuth
	originalRedirectToDevAuth := RedirectToDevAuth
	defer func() {
		RedirectToDevAuth = originalRedirectToDevAuth
	}()
	RedirectToDevAuth = func(w http.ResponseWriter, r *http.Request, log *zerolog.Logger) {
		w.WriteHeader(http.StatusFound) // StatusFound (302) is what RedirectToDevAuth uses
		w.Header().Set("Location", "/auth/dev/")
		fmt.Fprint(w, "Redirected to Dev Auth")
	}

	// Create the middleware
	middleware := RequireDevAuth(&cookieJar, &log, protectedHandler)

	// Create a test server to handle requests through the middleware
	testMux := http.NewServeMux()
	testMux.Handle("/", middleware) // Register the middleware to handle all paths

	t.Run("with valid cookie", func(t *testing.T) {
		// Add a valid session token to the cookieJar
		sessionToken := "valid_session_token"
		cookieJar.Set(sessionToken, Cookie{Valid: true, Type: CTDEV})

		req := httptest.NewRequest("GET", "/protected", nil)
		req.AddCookie(&http.Cookie{Name: "nbs-dev", Value: sessionToken}) // Use "nbs-dev" as per SetupDevAuth
		rr := httptest.NewRecorder()

		testMux.ServeHTTP(rr, req) // Use the testMux

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		if body := rr.Body.String(); body != "Protected content" {
			t.Errorf("handler returned unexpected body: got %q want %q", body, "Protected content")
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset()
	})

	t.Run("with invalid/missing cookie", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/protected", nil)
		// No cookie added or an invalid one

		rr := httptest.NewRecorder()

		testMux.ServeHTTP(rr, req) // Use the testMux

		if status := rr.Code; status != http.StatusFound { // Expect redirect
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
		}
		if location := rr.Header().Get("Location"); location != "/auth/dev/" {
			t.Errorf("handler returned wrong redirect location: got %q want %q", location, "/auth/dev/")
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset()
	})
}

func TestValidateDev(t *testing.T) {
	// Mock logger
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()

	t.Run("correct password", func(t *testing.T) {
		form := url.Values{}
		form.Add("pw", "testpassword") // This matches os.Getenv("PW") set in TestMain
		req := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		sessionToken, success := ValidateDev(rr, req, &log)

		if !success {
			t.Error("Expected ValidateDev to succeed, but it failed")
		}
		if sessionToken == "" {
			t.Error("Expected session token to be generated, but it was empty")
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset()
	})

	t.Run("incorrect password", func(t *testing.T) {
		form := url.Values{}
		form.Add("pw", "wrongpassword")
		req := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		sessionToken, success := ValidateDev(rr, req, &log)

		if success {
			t.Error("Expected ValidateDev to fail, but it succeeded")
		}
		if sessionToken != "" {
			t.Errorf("Expected session token to be empty, got %q", sessionToken)
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset()
	})

	t.Run("error parsing form", func(t *testing.T) {
		// Create a request with a malformed form data to cause parse form error
		req := httptest.NewRequest("POST", "/", strings.NewReader("malformed%key=value"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // This will cause r.ParseForm() to fail
		rr := httptest.NewRecorder()

		sessionToken, success := ValidateDev(rr, req, &log)

		if success {
			t.Error("Expected ValidateDev to fail, but it succeeded")
		}
		if sessionToken != "" {
			t.Errorf("Expected session token to be empty, got %q", sessionToken)
		}
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset()
	})
}

func TestRedirectFunctions(t *testing.T) {
	// Mock logger
	mockLogOutput := &MockLogger{}
	log := zerolog.New(mockLogOutput).Level(zerolog.DebugLevel).With().Logger()

	t.Run("RedirectToDevAuth", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/somepath", nil)
		rr := httptest.NewRecorder()

		RedirectToDevAuth(rr, req, &log)

		if status := rr.Code; status != http.StatusSeeOther {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
		}
		if location := rr.Header().Get("Location"); location != "/auth/dev/" {
			t.Errorf("handler returned wrong redirect location: got %q want %q", location, "/auth/dev/")
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset()
	})

	t.Run("RedirectToHome", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/somepath", nil)
		rr := httptest.NewRecorder()

		RedirectToHome(rr, req, &log)

		if status := rr.Code; status != http.StatusSeeOther {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
		}
		if location := rr.Header().Get("Location"); location != "/" {
			t.Errorf("handler returned wrong redirect location: got %q want %q", location, "/")
		}
		t.Logf("Log output:\n%s", mockLogOutput.Buf.String())
		mockLogOutput.Buf.Reset()
	})
}
