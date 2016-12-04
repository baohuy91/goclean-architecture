package controller

import "net/http"

// These are middleware to support http request
type Mdw interface {
	// Handle CORS request, especially for REST api
	CORS(handler http.Handler) http.Handler
	// Add proper header for json
	Header(handler http.Handler) http.Handler
	// Log to log file
	Log(handler http.Handler) http.Handler
}

// Factory function for new middleware
func NewMdw() Mdw {
	return &mdwImpl{}
}


type mdwImpl struct {
	// TODO: add necessary property here,
	// e.g. auth repository for validate token
}

// Handle CORS request fow API for client
func (m *mdwImpl) CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}
		// Lets Gorilla work
		handler.ServeHTTP(w, r)
	})
}

func (m *mdwImpl) Log(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request access to server here
		handler.ServeHTTP(w, r)
	})
}

// Add Header to all response
func (m *mdwImpl) Header(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		handler.ServeHTTP(w, r)
	})
}