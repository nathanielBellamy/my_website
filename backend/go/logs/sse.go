package logs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
)

type LogsController struct {
	Log     *zerolog.Logger
	LogDir  string
	LogRoot *os.Root
	StartAt time.Time
}

func NewLogsController(log *zerolog.Logger, logDir string, startAt time.Time) (*LogsController, error) {
	root, err := os.OpenRoot(logDir)
	if err != nil {
		return nil, fmt.Errorf("failed to open log root directory %s: %w", logDir, err)
	}
	return &LogsController{
		Log:     log,
		LogDir:  logDir,
		LogRoot: root,
		StartAt: startAt,
	}, nil
}

// StreamLogsHandler streams log entries via SSE
func (lc *LogsController) StreamLogsHandler(w http.ResponseWriter, r *http.Request) {
	lc.Log.Info().
		Str("ip", auth.GetClientIpAddr(r)).
		Msg("StreamLogsHandler Hit")

	flusher, ok := w.(http.Flusher)
	if !ok {
		lc.Log.Error().Msg("Streaming not supported by ResponseWriter")
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Disable the server's write timeout for this long-lived SSE connection
	rc := http.NewResponseController(w)
	if err := rc.SetWriteDeadline(time.Time{}); err != nil {
		lc.Log.Warn().Err(err).Msg("Could not disable write deadline for SSE")
	}

	levelFilter := r.URL.Query().Get("level")
	searchFilter := r.URL.Query().Get("search")
	linesStr := r.URL.Query().Get("lines")
	backfillLines := 100
	if linesStr != "" {
		if n, err := strconv.Atoi(linesStr); err == nil && n >= 0 {
			backfillLines = n
		}
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")

	// Find the latest log file for backfill
	latestFile, err := lc.findLatestLogFile()
	if err != nil {
		lc.Log.Error().Err(err).Msg("Error finding latest log file")
		sendSSEEvent(w, "error", `{"error":"no log files found"}`)
		flusher.Flush()
		return
	}

	// Backfill: send last N lines from the latest log file
	backfillEntries, err := lc.tailLogFile(latestFile, backfillLines, levelFilter, searchFilter)
	if err != nil {
		lc.Log.Warn().Err(err).Msg("Error reading backfill lines")
	}

	for _, entry := range backfillEntries {
		sendSSEEvent(w, "log", entry)
	}
	sendSSEEvent(w, "backfill-complete", `{"count":`+strconv.Itoa(len(backfillEntries))+`}`)
	flusher.Flush()

	// Poll for new log lines
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	currentFile := latestFile
	var lastOffset int64

	// Start reading from end of file
	if relCurrent, relErr := filepath.Rel(lc.LogDir, currentFile); relErr == nil {
		if info, statErr := lc.LogRoot.Stat(relCurrent); statErr == nil {
			lastOffset = info.Size()
		}
	}

	ctx := r.Context()
	for {
		select {
		case <-ctx.Done():
			lc.Log.Info().
				Str("ip", auth.GetClientIpAddr(r)).
				Msg("SSE client disconnected")
			return
		case <-ticker.C:
			// Check if a newer log file has appeared
			newestFile, err := lc.findLatestLogFile()
			if err != nil {
				continue
			}

			if newestFile != currentFile {
				// Read remaining lines from old file
				lc.streamNewLines(w, flusher, currentFile, &lastOffset, levelFilter, searchFilter)
				currentFile = newestFile
				lastOffset = 0
			}

			lc.streamNewLines(w, flusher, currentFile, &lastOffset, levelFilter, searchFilter)
		}
	}
}

func (lc *LogsController) streamNewLines(w http.ResponseWriter, flusher http.Flusher, filePath string, lastOffset *int64, levelFilter, searchFilter string) {
	relPath, err := filepath.Rel(lc.LogDir, filePath)
	if err != nil {
		lc.Log.Error().Err(err).Str("file", filePath).Msg("Error computing relative log path")
		return
	}

	info, err := lc.LogRoot.Stat(relPath)
	if err != nil {
		return
	}

	if info.Size() <= *lastOffset {
		return
	}

	f, err := lc.LogRoot.Open(relPath)
	if err != nil {
		lc.Log.Error().Err(err).Str("file", filePath).Msg("Error opening log file for streaming")
		return
	}
	defer f.Close()

	if _, err := f.Seek(*lastOffset, 0); err != nil {
		lc.Log.Error().Err(err).Msg("Error seeking in log file")
		return
	}

	scanner := bufio.NewScanner(f)
	hadLines := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if matchesFilters(line, levelFilter, searchFilter) {
			sendSSEEvent(w, "log", line)
			hadLines = true
		}
	}

	*lastOffset = info.Size()
	if hadLines {
		flusher.Flush()
	}
}

func (lc *LogsController) findLatestLogFile() (string, error) {
	type logFile struct {
		path    string
		modTime time.Time
	}
	var files []logFile
	err := filepath.Walk(lc.LogDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // skip errors
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "-log.txt") {
			files = append(files, logFile{path: path, modTime: info.ModTime()})
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error walking log directory: %w", err)
	}

	if len(files) == 0 {
		return "", fmt.Errorf("no log files found in %s", lc.LogDir)
	}

	// Sort by modification time (most recent last)
	sort.Slice(files, func(i, j int) bool {
		return files[i].modTime.Before(files[j].modTime)
	})
	return files[len(files)-1].path, nil
}

func (lc *LogsController) tailLogFile(filePath string, n int, levelFilter, searchFilter string) ([]string, error) {
	relPath, err := filepath.Rel(lc.LogDir, filePath)
	if err != nil {
		return nil, fmt.Errorf("error computing relative log path: %w", err)
	}
	f, err := lc.LogRoot.Open(relPath)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %w", err)
	}
	defer f.Close()

	var allLines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if matchesFilters(line, levelFilter, searchFilter) {
			allLines = append(allLines, line)
		}
	}

	if len(allLines) <= n {
		return allLines, nil
	}
	return allLines[len(allLines)-n:], nil
}

func matchesFilters(line, levelFilter, searchFilter string) bool {
	if levelFilter != "" {
		var entry map[string]interface{}
		if err := json.Unmarshal([]byte(line), &entry); err == nil {
			if lvl, ok := entry["level"].(string); ok {
				if !strings.EqualFold(lvl, levelFilter) {
					return false
				}
			}
		}
	}

	if searchFilter != "" {
		if !strings.Contains(strings.ToLower(line), strings.ToLower(searchFilter)) {
			return false
		}
	}

	return true
}

func sendSSEEvent(w http.ResponseWriter, event, data string) {
	fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, data)
}
