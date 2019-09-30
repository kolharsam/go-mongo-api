package routes

// all functions handlers for the route "/movie"

import (
	"context"
	"encoding/json"
	"net/http"
	"some-rest-api/funcerr"
	"some-rest-api/logger"
	"some-rest-api/middleware"
	"some-rest-api/mongo"
	"some-rest-api/response"
	"some-rest-api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddMovie is for adding a particular movie to the db
func AddMovie(w http.ResponseWriter, r *http.Request) {
	middleware.AddHeader(w, "content-type", "application/JSON")
	var movie mongo.MovieSchema
	json.NewDecoder(r.Body).Decode(&movie)
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	res, err := mongo.GetCollection().InsertOne(context, movie)
	movie.ID = res.InsertedID.(primitive.ObjectID)
	if !funcerr.IsErrored(err, "There was error in setting the data") {
		response.SetStatus(w, http.StatusOK)
		response.SendJSON(w, movie, "success")
	} else {
		response.SetStatus(w, http.StatusBadRequest)
		response.SendJSON(w, movie, "error")
	}
}

// GetMovie := for getting a particular movie from the database
func GetMovie(w http.ResponseWriter, r *http.Request) {
	movieID, err := utils.GetID(r, "id")
	var movie mongo.MovieSchema
	if !funcerr.IsErrored(err, "Invalid ID: "+movieID.String()) {
		context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		mongo.GetCollection().FindOne(context, mongo.MovieSchema{ID: movieID}).Decode(&movie)
		response.SendJSON(w, movie, "success")
		logger.Message(("Got the movie: " + movie.ID.String()), "info")
	} else {
		response.SendJSON(w, movie, "error")
		logger.Message(("Invalid ID: " + movieID.String()), "info")
	}
}
