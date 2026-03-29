package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockNetError is a mock implementation of net.Error for testing timeouts.
type mockNetError struct {
	timeout   bool
	temporary bool
	msg       string
}

func (m *mockNetError) Error() string   { return m.msg }
func (m *mockNetError) Timeout() bool   { return m.timeout }
func (m *mockNetError) Temporary() bool { return m.temporary }

func TestHandleDBError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		defaultMsg     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Connection Refused Error",
			err:            errors.New("dial tcp 127.0.0.1:5432: connect: connection refused"),
			defaultMsg:     "Default DB Error",
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   "Service Unavailable\n",
		},
		{
			name:           "I/O Timeout Error",
			err:            errors.New("read tcp 127.0.0.1:5432->127.0.0.1:5434: i/o timeout"),
			defaultMsg:     "Default DB Error",
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   "Service Unavailable\n",
		},
		{
			name:           "EOF Error",
			err:            errors.New("unexpected EOF"),
			defaultMsg:     "Default DB Error",
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   "Service Unavailable\n",
		},
		{
			name:           "Network Timeout Interface Error",
			err:            &mockNetError{timeout: true, msg: "custom net timeout"},
			defaultMsg:     "Default DB Error",
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   "Service Unavailable\n",
		},
		{
			name:           "Standard DB Syntax Error",
			err:            errors.New("syntax error at or near \"SELECT\""),
			defaultMsg:     "Query Failed",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Query Failed\n",
		},
		{
			name:           "Nil Error",
			err:            nil,
			defaultMsg:     "Something went wrong",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Something went wrong\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			
			HandleDBError(recorder, tt.err, tt.defaultMsg)
			
			if recorder.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, recorder.Code)
			}
			
			if recorder.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, recorder.Body.String())
			}
		})
	}
}
