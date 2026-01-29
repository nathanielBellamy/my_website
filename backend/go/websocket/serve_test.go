package websocket

import (
	"bytes"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
)

// MockLogger is a mock implementation of zerolog.Logger for testing.
type MockLogger struct {
	Buf bytes.Buffer
}

func (m *MockLogger) Write(p []byte) (n int, err error) {
	return m.Buf.Write(p)
}

// MockNetConn is a mock implementation of net.Conn for testing.
type MockNetConn struct {
	ReadBuf  bytes.Buffer
	WriteBuf bytes.Buffer
	CloseErr error
	Closed   bool
	once     sync.Once
}

