package routes

import (
	"net/http"
	"some-rest-api/response"
)

// Home := for route "/"
func Home(w http.ResponseWriter, r *http.Request) {
	response.SendText(w, "You've hit home!")
}
