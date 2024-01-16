package web

import (
	"net/http"
)

// SetupGlobalMiddleware configures and returns the global middleware around the main handler.
func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return ApplyMiddleware(
		handler,
		LoggingMiddleware,    // Stub for where you would insert a logging middleware
		RecoveryMiddleware,   // Stub for a panic recovery middleware
	)
}

// ApplyMiddleware applies the given middleware and returns the final handler.
func ApplyMiddleware(handler http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		handler = m(handler)
	}
	return handler
}