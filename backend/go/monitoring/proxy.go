package monitoring

import (
	"bytes"
	"io"
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
		// Set Host to the target so Grafana sees its own hostname
		req.Host = target.Host
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

	proxy.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode >= 400 {
			body, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				log.Error().
					Int("status", resp.StatusCode).
					Err(readErr).
					Str("path", resp.Request.URL.Path).
					Msgf("Grafana error %d on %s (could not read body)", resp.StatusCode, resp.Request.URL.Path)
			} else {
				log.Error().
					Int("status", resp.StatusCode).
					Str("body", truncate(string(body), 500)).
					Str("path", resp.Request.URL.Path).
					Msgf("Grafana error %d on %s: %s", resp.StatusCode, resp.Request.URL.Path, truncate(string(body), 200))
				resp.Body = io.NopCloser(bytes.NewReader(body))
			}
		}
		return nil
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Error().Err(err).Str("path", r.URL.Path).Msg("Grafana proxy connection error")
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
		Str("originalPath", r.URL.Path).
		Msg("Proxying to Grafana")
	gp.ReverseProxy.ServeHTTP(w, r)
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
