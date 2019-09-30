package middleware

import "net/http"

// AddHeader := adds a header to the response
func AddHeader(response http.ResponseWriter, header string, value string) {
	response.Header().Add(header, value)
}
