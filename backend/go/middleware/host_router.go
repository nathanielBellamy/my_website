package middleware

import (
	"net/http"
	"strings"
)

type HostRouter struct {
	AdminMux     *http.ServeMux
	OldSiteMux   *http.ServeMux
	MarketingMux *http.ServeMux
}

func (hr *HostRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := strings.ToLower(r.Host)

	// Route based on subdomain prefix
	// This supports admin.localhost, admin.mydomain.dev, and admin.127.0.0.1.nip.io (for E2E)
	if strings.HasPrefix(host, "admin.") {
		hr.AdminMux.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(host, "old-site.") {
		hr.OldSiteMux.ServeHTTP(w, r)
		return
	}

	// Fallback to Marketing (root domain)
	hr.MarketingMux.ServeHTTP(w, r)
}
