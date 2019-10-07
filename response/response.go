package response

import (
	"encoding/json"
	"fmt"
	"net/http"
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
func SendJSON(w http.ResponseWriter, data interface{}, msgType string, msg string) {
	var resp Response
	switch msgType {
	case "success":
		resp = Response{
			Success:     true,
			Message:     msg,
			RespondedAt: time.Now(),
			Data:        data,
		}
	case "error":
		resp = Response{
			Success:     false,
			Message:     msg,
			RespondedAt: time.Now(),
		}
	}
	json.NewEncoder(w).Encode(resp)
}
