package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SetStatus is to be used if you want to set the status code using the http package
func SetStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

// SendText := is a function for sending a "text" response
func SendText(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, str)
}

// Send := to send responses
func (resp *Response) Send(data interface{}) *Response {
	resp.Data = data
	return resp
}

// RSuccess := to send a success response
func (resp *Response) RSuccess(w http.ResponseWriter, msg string) {
	resp.Message = msg
	resp.Success = true
	resp.ResponseTime = time.Now()
	json.NewEncoder(w).Encode(resp)
}

// RError := to send a error response
func (resp *Response) RError(w http.ResponseWriter, msg string) {
	resp.Message = msg
	resp.Success = false
	resp.ResponseTime = time.Now()
	json.NewEncoder(w).Encode(resp)
}
