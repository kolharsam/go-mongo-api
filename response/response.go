package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"some-rest-api/logger"
	"some-rest-api/mongo"
	"time"
)

// SetStatus is to be used if you want to set the status code
func SetStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

// Send := is a function for sending a "text" response
func Send(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, str)
}

// SendJSON := to send responses
func SendJSON(w http.ResponseWriter, data mongo.MovieSchema, msgType string) {
	var resp Response
	switch msgType {
	case "success":
		resp = Response{
			Success:     true,
			Error:       "none",
			RespondedAt: time.Now(),
			Data:        data,
		}
		logger.Message(("Sucessfully added movie: " + data.ID.String()), "info")
	case "error":
		resp = Response{
			Success:     false,
			Error:       "There was an error in completing the request",
			RespondedAt: time.Now(),
			Data:        data,
		}
		logger.Message(("Couldn't add movie: " + data.ID.String()), "error")
	}
	json.NewEncoder(w).Encode(resp)
}
