package monitoring

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rs/zerolog"
)

type GrafanaProxy struct {
	Log          *zerolog.Logger
	GrafanaURL   string
	ReverseProxy *httputil.ReverseProxy
}

func NewGrafanaProxy(log *zerolog.Logger, grafanaURL string) *GrafanaProxy {
	target, err := url.Parse(grafanaURL)
	if err != nil {
		log.Fatal().Err(err).Str("url", grafanaURL).Msg("Invalid Grafana URL")
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Set auth proxy header so Grafana trusts the request
		req.Header.Set("X-WEBAUTH-USER", "admin")
		// Keep /grafana/ prefix — Grafana uses serve_from_sub_path to handle it
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Error().Err(err).Str("path", r.URL.Path).Msg("Grafana proxy error")
		http.Error(w, "Grafana is not available", http.StatusBadGateway)
	}

	return &GrafanaProxy{
		Log:          log,
		GrafanaURL:   grafanaURL,
		ReverseProxy: proxy,
	}
}

func (gp *GrafanaProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gp.Log.Debug().
		Str("path", r.URL.Path).
		Msg("Proxying to Grafana")
	gp.ReverseProxy.ServeHTTP(w, r)
}
