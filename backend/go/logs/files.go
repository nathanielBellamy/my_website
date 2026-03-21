package logs

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
)

type LogFileInfo struct {
	Path string `json:"path"`
	Date string `json:"date"`
	Size int64  `json:"size"`
}

type LogFilesResponse struct {
	Files []LogFileInfo `json:"files"`
}

// GetLogFilesHandler returns a list of available log files
func (lc *LogsController) GetLogFilesHandler(w http.ResponseWriter, r *http.Request) {
	lc.Log.Info().
		Str("ip", auth.GetClientIpAddr(r)).
		Msg("GetLogFilesHandler Hit")

	var files []LogFileInfo

	err := filepath.Walk(lc.LogDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			lc.Log.Warn().Err(err).Str("path", path).Msg("Error walking log directory")
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), "-log.txt") {
			return nil
		}

		relPath, err := filepath.Rel(lc.LogDir, path)
		if err != nil {
			relPath = path
		}

		// Extract date from filename: YYYY-MM-DDTHH-MM-SSZ-log.txt -> YYYY-MM-DD
		date := extractDateFromFilename(info.Name())

		files = append(files, LogFileInfo{
			Path: relPath,
			Date: date,
			Size: info.Size(),
		})

		return nil
	})

	if err != nil {
		lc.Log.Error().Err(err).Msg("Error walking log directory")
		http.Error(w, "Error reading log files", http.StatusInternalServerError)
		return
	}

	// Sort by date descending (newest first)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Date > files[j].Date
	})

	response := LogFilesResponse{
		Files: files,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		lc.Log.Error().Err(err).Msg("Error encoding log files response")
	}
}

func extractDateFromFilename(name string) string {
	// Filename format: YYYY-MM-DDTHH-MM-SSZ-log.txt
	if len(name) < 10 {
		return ""
	}
	return name[:10]
}
