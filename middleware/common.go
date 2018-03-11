package middleware

import "net/http"

type (
	// Middleware type defines general type for middleware
	Middleware func(h http.Handler) http.Handler
)

// Chain allows to apply multiple middlewares at a time
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware :=range middlewares {
		h = middleware(h)
	}
	return h
}
