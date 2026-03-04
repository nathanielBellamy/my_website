package middleware

import (
	"net/http"
	"sync"

	"github.com/nathanielBellamy/my_website/backend/go/auth"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

// IPRateLimiter stores rate limiters for each IP.
type IPRateLimiter struct {
	ips   map[string]*rate.Limiter
	mu    *sync.RWMutex
	r     rate.Limit
	b     int
}

// NewIPRateLimiter creates a new rate limiter that allows 'r' events per second with a burst of 'b'.
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips:   make(map[string]*rate.Limiter),
		mu:    &sync.RWMutex{},
		r:     r,
		b:     b,
	}
}

// AddIP creates a new rate limiter for an IP if it doesn't exist.
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

// GetLimiter returns the rate limiter for the provided IP.
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.RUnlock()
		return i.AddIP(ip)
	}

	i.mu.RUnlock()
	return limiter
}

// RateLimitMiddleware is an HTTP middleware that applies rate limiting based on the client's IP address.
func RateLimitMiddleware(limiter *IPRateLimiter, log *zerolog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := auth.GetClientIpAddr(r)
		limiterForIP := limiter.GetLimiter(ip)

		if !limiterForIP.Allow() {
			log.Warn().Str("ip", ip).Msg("Rate limit exceeded")
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
