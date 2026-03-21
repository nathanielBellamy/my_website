package logs

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
)

type LogEntry struct {
	Level   string                 `json:"level"`
	Time    string                 `json:"time"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

type PaginatedLogResponse struct {
	Data  []LogEntry `json:"data"`
	Total int        `json:"total"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
}

// GetLogHistoryHandler returns paginated historical log entries
func (lc *LogsController) GetLogHistoryHandler(w http.ResponseWriter, r *http.Request) {
	lc.Log.Info().
		Str("ip", auth.GetClientIpAddr(r)).
		Msg("GetLogHistoryHandler Hit")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	levelFilter := r.URL.Query().Get("level")
	searchFilter := r.URL.Query().Get("search")
	dateFilter := r.URL.Query().Get("date") // YYYY-MM-DD

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50
	}
	if limit > 500 {
		limit = 500
	}

	// Collect matching log files
	logFiles, err := lc.findLogFilesByDate(dateFilter)
	if err != nil {
		lc.Log.Error().Err(err).Msg("Error finding log files")
		http.Error(w, "Error reading logs", http.StatusInternalServerError)
		return
	}

	// Read all matching entries from the files (newest first)
	var allEntries []LogEntry
	for i := len(logFiles) - 1; i >= 0; i-- {
		entries, err := lc.readLogFileEntries(logFiles[i], levelFilter, searchFilter)
		if err != nil {
			lc.Log.Warn().Err(err).Str("file", logFiles[i]).Msg("Error reading log file")
			continue
		}
		// Prepend so newest entries are first
		allEntries = append(entries, allEntries...)
	}

	// Reverse so newest is first
	for i, j := 0, len(allEntries)-1; i < j; i, j = i+1, j-1 {
		allEntries[i], allEntries[j] = allEntries[j], allEntries[i]
	}

	total := len(allEntries)

	// Paginate
	start := (page - 1) * limit
	if start >= total {
		start = total
	}
	end := start + limit
	if end > total {
		end = total
	}

	response := PaginatedLogResponse{
		Data:  allEntries[start:end],
		Total: total,
		Page:  page,
		Limit: limit,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		lc.Log.Error().Err(err).Msg("Error encoding log history response")
	}
}

func (lc *LogsController) findLogFilesByDate(dateFilter string) ([]string, error) {
	type logFile struct {
		path    string
		modTime time.Time
	}
	var files []logFile

	err := filepath.Walk(lc.LogDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), "-log.txt") {
			return nil
		}

		if dateFilter != "" {
			// dateFilter is YYYY-MM-DD, log filenames are YYYY-MM-DDTHH-MM-SSZ-log.txt
			if !strings.HasPrefix(info.Name(), dateFilter) {
				return nil
			}
		}

		files = append(files, logFile{path: path, modTime: info.ModTime()})
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort by modification time (oldest first)
	sort.Slice(files, func(i, j int) bool {
		return files[i].modTime.Before(files[j].modTime)
	})

	result := make([]string, len(files))
	for i, f := range files {
		result[i] = f.path
	}
	return result, nil
}

func (lc *LogsController) readLogFileEntries(filePath, levelFilter, searchFilter string) ([]LogEntry, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if !matchesFilters(line, levelFilter, searchFilter) {
			continue
		}

		entry := parseLogLine(line)
		entries = append(entries, entry)
	}

	return entries, scanner.Err()
}

func parseLogLine(line string) LogEntry {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(line), &raw); err != nil {
		return LogEntry{
			Level:   "unknown",
			Message: line,
		}
	}

	entry := LogEntry{}
	if lvl, ok := raw["level"].(string); ok {
		entry.Level = lvl
		delete(raw, "level")
	}
	if t, ok := raw["time"].(string); ok {
		entry.Time = t
		delete(raw, "time")
	}
	if msg, ok := raw["message"].(string); ok {
		entry.Message = msg
		delete(raw, "message")
	}

	if len(raw) > 0 {
		entry.Fields = raw
	}

	return entry
}
