package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// acceptedHeader
var acceptedHeader = []string{
	"Accept",
	"Accept-Encoding",
	"Access-Control-Allow-Credentials",
	"Access-Control-Allow-Headers",
	"Access-Control-Allow-Methods",
	"Access-Control-Allow-Origin",
	"Access-Control-Expose-Headers",
	"Access-Control-Max-Age",
	"Access-Control-Request-Headers",
	"Access-Control-Request-Method",
	"Allow",
	"Authorization",
	"Content-Disposition",
	"Content-Encoding",
	"Content-Length",
	"Content-Security-Policy",
	"Content-Security-Policy-Report-Only",
	"Content-Type",
	"Cookie",
	"If-Modified-Since",
	"Last-Modified",
	"Location",
	"Origin",
	"Referrer-Policy",
	"Server",
	"Set-Cookie",
	"Strict-Transport-Security",
	"Upgrade",
	"Vary",
	"WWW-Authenticate",
	"X-CSRF-Token",
	"X-Content-Type-Options",
	"X-Forwarded-For",
	"X-Forwarded-Proto",
	"X-Forwarded-Protocol",
	"X-Forwarded-Ssl",
	"X-Frame-Options",
	"X-HTTP-Method-Override",
	"X-Real-IP",
	"X-Request-ID",
	"X-Requested-With",
	"X-Url-Scheme",
	"X-XSS-Protection",
}

// acceptedMethods
var acceptedMethods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

// Cors
func CORS(origins []string) func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: origins,
		AllowedMethods: acceptedMethods,
		AllowedHeaders: acceptedHeader,
	})
}

// SetHeader
func SetHeader(key, value string) func(next http.Handler) http.Handler {
	return middleware.SetHeader(key, value)
}

// Gzip
func Gzip(level int) func(next http.Handler) http.Handler {
	return middleware.Compress(level, "gzip")
}

// Timeout
func Timeout(dur time.Duration) func(next http.Handler) http.Handler {
	return middleware.Timeout(dur)
}

// RealIP
func RealIP(h http.Handler) http.Handler {
	return middleware.RealIP(h)
}

// StripSlashes
func StripSlashes(h http.Handler) http.Handler {
	return middleware.StripSlashes(h)
}
