package response

import "time"

// Response - if a query goes through, respond with the following schema
type Response struct {
	Success      bool        `json:"success,omitempty" bson:"success,omitempty"`
	Message      string      `json:"message,omitempty" bson:"message,omitempty"`
	ResponseTime time.Time   `json:"response_time,omitempty" bson:"response_time,omitempty"`
	Data         interface{} `json:"data,omitempty" bson:"data,omitempty"`
}
