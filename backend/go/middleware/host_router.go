package middleware

import (
	"net/http"
	"os"
	"strings"
)

type HostRouter struct {
	AdminMux     *http.ServeMux
	OldSiteMux   *http.ServeMux
	MarketingMux *http.ServeMux
}

func (hr *HostRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := strings.ToLower(r.Host)
	domainBase := strings.ToLower(os.Getenv("DOMAIN_BASE"))
	if domainBase == "" {
		domainBase = "localhost:8080"
	}

	// Route based on subdomain
	if strings.HasPrefix(host, "admin."+domainBase) || host == "admin.localhost:8080" {
		hr.AdminMux.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(host, "old-site."+domainBase) || host == "old-site.localhost:8080" {
		hr.OldSiteMux.ServeHTTP(w, r)
		return
	}

	// Fallback to Marketing (root domain)
	hr.MarketingMux.ServeHTTP(w, r)
}
