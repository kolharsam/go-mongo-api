package routes

import (
	"net/http"
	"some-rest-api/response"
)

// Home := for route "/"
func Home(w http.ResponseWriter, r *http.Request) {
	response.Send(w, "Hello, from API!")
}
