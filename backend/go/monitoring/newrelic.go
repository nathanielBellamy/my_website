package monitoring

import (
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
)

// NewRelicApp initializes and returns a New Relic application instance.
// Returns nil if initialization fails (allowing the server to run without New Relic).
func NewRelicApp(log *zerolog.Logger, appName string, licenseKey string) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(true),
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialize New Relic agent")
		return nil
	}
	log.Info().Str("appName", appName).Msg("New Relic agent initialized")
	return app
}

// NewRelicMiddleware wraps an http.Handler with New Relic transaction tracking.
func NewRelicMiddleware(app *newrelic.Application, next http.Handler) http.Handler {
	if app == nil {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txn := app.StartTransaction(r.Method + " " + r.URL.Path)
		defer txn.End()
		txn.SetWebRequestHTTP(r)
		w = txn.SetWebResponse(w)
		r = newrelic.RequestWithTransactionContext(r, txn)
		next.ServeHTTP(w, r)
	})
}
