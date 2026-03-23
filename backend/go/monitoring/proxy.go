package monitoring

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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
		req.Header.Set("X-WEBAUTH-USER", "admin")
		// Strip /grafana prefix — Grafana serves at root, root_url's appSubUrl
		// tells the frontend to use /grafana/ for browser-side routing
		req.URL.Path = strings.TrimPrefix(req.URL.Path, "/grafana")
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}
		if req.URL.RawPath != "" {
			req.URL.RawPath = strings.TrimPrefix(req.URL.RawPath, "/grafana")
			if req.URL.RawPath == "" {
				req.URL.RawPath = "/"
			}
		}
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
