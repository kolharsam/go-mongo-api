package response

import (
	"some-rest-api/mongo"
	"time"
)

// Response - if a query goes through, respond with the following schema
type Response struct {
	Success     bool              `json:"success"`
	Error       string            `json:"error"`
	RespondedAt time.Time         `json:"respondedAt"`
	Data        mongo.MovieSchema `json:"data"`
}
