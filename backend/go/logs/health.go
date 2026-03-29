package logs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/nathanielBellamy/my_website/backend/go/interfaces"
	"github.com/nathanielBellamy/my_website/backend/go/models"
	"github.com/rs/zerolog"
)

type HealthInfo struct {
	Uptime        string         `json:"uptime"`
	UptimeSeconds float64        `json:"uptimeSeconds"`
	GoRoutines    int            `json:"goRoutines"`
	MemAllocMB    float64        `json:"memAllocMb"`
	MemSysMB      float64        `json:"memSysMb"`
	NumGC         uint32         `json:"numGc"`
	DbConnected   bool           `json:"dbConnected"`
	GoVersion     string         `json:"goVersion"`
	NumCPU        int            `json:"numCpu"`
}

type HealthController struct {
	Log     *zerolog.Logger
	StartAt time.Time
	DB      interfaces.PgxDB
}

func NewHealthController(log *zerolog.Logger, startAt time.Time, db interfaces.PgxDB) *HealthController {
	return &HealthController{
		Log:     log,
		StartAt: startAt,
		DB:      db,
	}
}

// GetHealthHandler returns system health information
func (hc *HealthController) GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	hc.Log.Info().
		Str("ip", auth.GetClientIpAddr(r)).
		Msg("GetHealthHandler Hit")

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	uptime := time.Since(hc.StartAt)

	// Test database connectivity
	dbConnected := hc.checkDBConnection()

	health := HealthInfo{
		Uptime:        formatDuration(uptime),
		UptimeSeconds: uptime.Seconds(),
		GoRoutines:    runtime.NumGoroutine(),
		MemAllocMB:    float64(memStats.Alloc) / 1024 / 1024,
		MemSysMB:      float64(memStats.Sys) / 1024 / 1024,
		NumGC:         memStats.NumGC,
		DbConnected:   dbConnected,
		GoVersion:     runtime.Version(),
		NumCPU:        runtime.NumCPU(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(health); err != nil {
		hc.Log.Error().Err(err).Msg("Error encoding health response")
	}
}

func (hc *HealthController) checkDBConnection() bool {
	// Try a simple query to verify DB connectivity
	var result []models.WorkContent
	err := hc.DB.Model(&result).Limit(1).Select()
	if err != nil {
		hc.Log.Warn().Err(err).Msg("DB health check failed")
		return false
	}
	return true
}

func formatDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}
