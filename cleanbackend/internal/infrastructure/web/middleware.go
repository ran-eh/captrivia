package web

import (
	"net/http"
)

// LoggingMiddleware is a placeholder for request logging middleware.
// Add logger parameter and actions to log requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logging logic goes here.

		next.ServeHTTP(w, r)
	})
}

// RecoveryMiddleware is a placeholder for recovery middleware to recover from panics.
// Add actions required to recover from panics, like logging the error, etc.
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Recovery logic goes here.

		next.ServeHTTP(w, r)
	})
}