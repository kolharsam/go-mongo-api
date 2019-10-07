package routes

import (
	"context"
	"log"
	"net/http"
	"some-rest-api/logger"
	"some-rest-api/mongo"
	"some-rest-api/response"
	"time"
)

// GetAll := get all documents in the collection
func GetAll(w http.ResponseWriter, r *http.Request) {
	var allMovies []mongo.MovieSchema

	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	result, err := mongo.GetCollection().Find(context, mongo.MovieSchema{})

	if err == nil {
		result.All(context, &allMovies)
		response.SendJSON(w, allMovies, "success", "Fetched all movies!")
		logger.Message("Fetched all movies", "info")
	} else {
		log.Println(err)
		logger.Message("There was an error", "error")
		response.SendJSON(w, nil, "error", "There was an error completing the request")
	}
}
