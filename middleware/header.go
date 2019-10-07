package middleware

import "net/http"

// AddHeaders := adds a header to the response
func AddHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		w.Header().Add("Access-Control-Max-Age", "3600")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With, remember-me")
		next.ServeHTTP(w, r)
	})
}

// AddContentHeader := to attach the content-type header of JSON to all responses
func AddContentHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
