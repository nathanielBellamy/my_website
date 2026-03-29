package logs

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func setupTestLogDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	// Create a log directory structure: log/2026/03/
	logDir := filepath.Join(dir, "2026", "03")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		t.Fatalf("Failed to create log dir: %v", err)
	}

	// Write a sample log file
	logFile := filepath.Join(logDir, "2026-03-21T10-00-00Z-log.txt")
	lines := []string{
		`{"level":"info","time":"2026-03-21T10:00:01Z","message":"Loading ENV"}`,
		`{"level":"info","mode":"localhost","time":"2026-03-21T10:00:02Z","message":"Runtime Env"}`,
		`{"level":"error","ip":"127.0.0.1","time":"2026-03-21T10:00:03Z","message":"Connection refused"}`,
		`{"level":"warn","ip":"192.168.1.1","time":"2026-03-21T10:00:04Z","message":"Rate limit exceeded"}`,
		`{"level":"debug","time":"2026-03-21T10:00:05Z","message":"Debug message"}`,
	}

	content := ""
	for _, line := range lines {
		content += line + "\n"
	}

	if err := os.WriteFile(logFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write log file: %v", err)
	}

	return dir
}

func newTestLogsController(t *testing.T, logDir string) *LogsController {
	t.Helper()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	lc, err := NewLogsController(&log, logDir, time.Now())
	if err != nil {
		t.Fatalf("Failed to create logs controller: %v", err)
	}
	return lc
}

func TestParseLogLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected LogEntry
	}{
		{
			name: "valid info log",
			line: `{"level":"info","time":"2026-03-21T10:00:01Z","message":"Loading ENV"}`,
			expected: LogEntry{
				Level:   "info",
				Time:    "2026-03-21T10:00:01Z",
				Message: "Loading ENV",
			},
		},
		{
			name: "log with extra fields",
			line: `{"level":"error","ip":"127.0.0.1","time":"2026-03-21T10:00:03Z","message":"Connection refused"}`,
			expected: LogEntry{
				Level:   "error",
				Time:    "2026-03-21T10:00:03Z",
				Message: "Connection refused",
				Fields:  map[string]interface{}{"ip": "127.0.0.1"},
			},
		},
		{
			name: "invalid json",
			line: "this is not json",
			expected: LogEntry{
				Level:   "unknown",
				Message: "this is not json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := parseLogLine(tt.line)
			if entry.Level != tt.expected.Level {
				t.Errorf("Level: got %q, want %q", entry.Level, tt.expected.Level)
			}
			if entry.Time != tt.expected.Time {
				t.Errorf("Time: got %q, want %q", entry.Time, tt.expected.Time)
			}
			if entry.Message != tt.expected.Message {
				t.Errorf("Message: got %q, want %q", entry.Message, tt.expected.Message)
			}
		})
	}
}

func TestMatchesFilters(t *testing.T) {
	line := `{"level":"error","time":"2026-03-21T10:00:03Z","message":"Connection refused"}`

	tests := []struct {
		name     string
		level    string
		search   string
		expected bool
	}{
		{"no filter", "", "", true},
		{"matching level", "error", "", true},
		{"non-matching level", "info", "", false},
		{"matching search", "", "connection", true},
		{"non-matching search", "", "foobar", false},
		{"matching both", "error", "refused", true},
		{"matching level non-matching search", "error", "foobar", false},
		{"case insensitive level", "ERROR", "", true},
		{"case insensitive search", "", "CONNECTION", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchesFilters(line, tt.level, tt.search)
			if result != tt.expected {
				t.Errorf("matchesFilters(%q, %q): got %v, want %v", tt.level, tt.search, result, tt.expected)
			}
		})
	}
}

func TestExtractDateFromFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected string
	}{
		{"standard log file", "2026-03-21T10-00-00Z-log.txt", "2026-03-21"},
		{"short name", "abc", ""},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractDateFromFilename(tt.filename)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGetLogFilesHandler(t *testing.T) {
	logDir := setupTestLogDir(t)
	lc := newTestLogsController(t, logDir)

	req, err := http.NewRequest("GET", "/v1/api/admin/logs/files", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	lc.GetLogFilesHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Status: got %d, want %d", rr.Code, http.StatusOK)
	}

	var resp LogFilesResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(resp.Files) != 1 {
		t.Errorf("File count: got %d, want 1", len(resp.Files))
	}

	if resp.Files[0].Date != "2026-03-21" {
		t.Errorf("File date: got %q, want %q", resp.Files[0].Date, "2026-03-21")
	}
}

func TestGetLogHistoryHandler(t *testing.T) {
	logDir := setupTestLogDir(t)
	lc := newTestLogsController(t, logDir)

	t.Run("all entries", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/api/admin/logs/history?page=1&limit=50", nil)
		rr := httptest.NewRecorder()
		lc.GetLogHistoryHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Status: got %d, want %d", rr.Code, http.StatusOK)
		}

		var resp PaginatedLogResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if resp.Total != 5 {
			t.Errorf("Total: got %d, want 5", resp.Total)
		}
	})

	t.Run("filter by level", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/api/admin/logs/history?level=error", nil)
		rr := httptest.NewRecorder()
		lc.GetLogHistoryHandler(rr, req)

		var resp PaginatedLogResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if resp.Total != 1 {
			t.Errorf("Total errors: got %d, want 1", resp.Total)
		}
	})

	t.Run("search filter", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/api/admin/logs/history?search=rate+limit", nil)
		rr := httptest.NewRecorder()
		lc.GetLogHistoryHandler(rr, req)

		var resp PaginatedLogResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if resp.Total != 1 {
			t.Errorf("Total: got %d, want 1", resp.Total)
		}
	})

	t.Run("date filter", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/api/admin/logs/history?date=2026-03-21", nil)
		rr := httptest.NewRecorder()
		lc.GetLogHistoryHandler(rr, req)

		var resp PaginatedLogResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if resp.Total != 5 {
			t.Errorf("Total: got %d, want 5", resp.Total)
		}
	})

	t.Run("non-matching date filter", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/api/admin/logs/history?date=2099-01-01", nil)
		rr := httptest.NewRecorder()
		lc.GetLogHistoryHandler(rr, req)

		var resp PaginatedLogResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if resp.Total != 0 {
			t.Errorf("Total: got %d, want 0", resp.Total)
		}
	})

	t.Run("pagination", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/api/admin/logs/history?page=1&limit=2", nil)
		rr := httptest.NewRecorder()
		lc.GetLogHistoryHandler(rr, req)

		var resp PaginatedLogResponse
		if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if len(resp.Data) != 2 {
			t.Errorf("Page size: got %d, want 2", len(resp.Data))
		}
		if resp.Total != 5 {
			t.Errorf("Total: got %d, want 5", resp.Total)
		}
		if resp.Page != 1 {
			t.Errorf("Page: got %d, want 1", resp.Page)
		}
	})
}

func TestFindLatestLogFile(t *testing.T) {
	logDir := setupTestLogDir(t)
	lc := newTestLogsController(t, logDir)

	// Add a second log file with a later timestamp
	logDir2 := filepath.Join(logDir, "2026", "03")
	laterFile := filepath.Join(logDir2, "2026-03-21T12-00-00Z-log.txt")
	if err := os.WriteFile(laterFile, []byte(`{"level":"info","time":"2026-03-21T12:00:00Z","message":"Later log"}`+"\n"), 0644); err != nil {
		t.Fatalf("Failed to write later log file: %v", err)
	}

	latest, err := lc.findLatestLogFile()
	if err != nil {
		t.Fatalf("findLatestLogFile failed: %v", err)
	}

	if filepath.Base(latest) != "2026-03-21T12-00-00Z-log.txt" {
		t.Errorf("Latest file: got %q, want 2026-03-21T12-00-00Z-log.txt", filepath.Base(latest))
	}
}

func TestTailLogFile(t *testing.T) {
	logDir := setupTestLogDir(t)
	lc := newTestLogsController(t, logDir)

	logFile := filepath.Join(logDir, "2026", "03", "2026-03-21T10-00-00Z-log.txt")

	t.Run("tail 2 lines", func(t *testing.T) {
		lines, err := lc.tailLogFile(logFile, 2, "", "")
		if err != nil {
			t.Fatalf("tailLogFile failed: %v", err)
		}
		if len(lines) != 2 {
			t.Errorf("Lines: got %d, want 2", len(lines))
		}
	})

	t.Run("tail more than available", func(t *testing.T) {
		lines, err := lc.tailLogFile(logFile, 100, "", "")
		if err != nil {
			t.Fatalf("tailLogFile failed: %v", err)
		}
		if len(lines) != 5 {
			t.Errorf("Lines: got %d, want 5", len(lines))
		}
	})

	t.Run("tail with level filter", func(t *testing.T) {
		lines, err := lc.tailLogFile(logFile, 100, "error", "")
		if err != nil {
			t.Fatalf("tailLogFile failed: %v", err)
		}
		if len(lines) != 1 {
			t.Errorf("Lines: got %d, want 1", len(lines))
		}
	})
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{"seconds only", 45 * time.Second, "45s"},
		{"minutes and seconds", 5*time.Minute + 30*time.Second, "5m 30s"},
		{"hours minutes seconds", 2*time.Hour + 15*time.Minute + 10*time.Second, "2h 15m 10s"},
		{"days", 26*time.Hour + 30*time.Minute, "1d 2h 30m 0s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDuration(tt.duration)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
