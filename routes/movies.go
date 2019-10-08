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
	var resp response.Response
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	result, err := mongo.GetCollection().Find(context, mongo.MovieSchema{})

	if err == nil {
		result.All(context, &allMovies)
		resp.Send(allMovies).RSuccess(w, "")
		logger.Message("Fetched all movies", "info")
	} else {
		log.Println(err)
		logger.Message("There was an error"+err.Error(), "error")
		resp.Send(allMovies).RError(w, "There was an error"+err.Error())
	}
}
