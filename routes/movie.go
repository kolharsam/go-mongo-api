package routes

// all functions handlers for the route "/movie"

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"some-rest-api/funcerr"
	"some-rest-api/logger"
	"some-rest-api/mongo"
	"some-rest-api/response"
	"some-rest-api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add is for adding a particular movie to the db
func Add(w http.ResponseWriter, r *http.Request) {
	var movie mongo.MovieSchema
	json.NewDecoder(r.Body).Decode(&movie)
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	res, err := mongo.GetCollection().InsertOne(context, movie)
	movie.ID = res.InsertedID.(primitive.ObjectID)
	if !funcerr.IsErrored(err, "There was error in setting the data") {
		response.SetStatus(w, http.StatusOK)
		response.SendJSON(w, movie, "success", "none")
	} else {
		response.SetStatus(w, http.StatusBadRequest)
		response.SendJSON(w, movie, "error", "Check the object that you've been trying to add")
	}
}

// GetOne := for getting a particular movie from the database
func GetOne(w http.ResponseWriter, r *http.Request) {
	movieID, err := utils.GetID(r, "id")
	var movie mongo.MovieSchema
	if !funcerr.IsErrored(err, "Problem with ID: "+movieID.String()) {
		context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		mongo.GetCollection().FindOne(context, mongo.MovieSchema{ID: movieID}).Decode(&movie)

		if primitive.ObjectID.IsZero(movie.ID) {
			response.SendJSON(w, movie, "error", "Invalid / Incorrect ID passed!")
			logger.Message(("Invalid ID: " + utils.GetIDString(movieID)), "error")
		} else {
			response.SendJSON(w, movie, "success", "none")
			logger.Message(("Got the movie: " + utils.GetIDString(movie.ID)), "info")
		}

	}
}

// Update := to update a particular data on mongo
func Update(w http.ResponseWriter, r *http.Request) {
	var movie, newMovie mongo.MovieSchema

	json.NewDecoder(r.Body).Decode(&newMovie)

	movieID := newMovie.ID

	context, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	mongo.GetCollection().FindOneAndReplace(context, mongo.MovieSchema{ID: movieID}, newMovie).Decode(&movie)

	if primitive.ObjectID.IsZero(movie.ID) {
		response.SendJSON(w, movie, "error", "Illegal update operation")
		logger.Message(("Invalid ID: " + utils.GetIDString(movieID)), "error")
	} else {
		response.SendJSON(w, newMovie, "success", "")
		logger.Message(("Updated document: " + utils.GetIDString(movie.ID)), "info")
	}
}

// Delete := function to delete a document
func Delete(w http.ResponseWriter, r *http.Request) {
	movieID, err := utils.GetID(r, "id")

	if !funcerr.IsErrored(err, "Problem with ID: "+movieID.String()) {
		context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		deleted, errr := mongo.GetCollection().DeleteOne(context, mongo.MovieSchema{ID: movieID})
		log.Println(deleted, errr)

		if errr == nil {
			log.Println(deleted.DeletedCount)
			response.SendJSON(w, nil, "success", "Successfully deleted the movie"+movieID.String())
			logger.Message("Deleted the movie: "+movieID.String(), "info")
		} else {
			log.Println(errr)
			response.SendJSON(w, nil, "error", "There was an error deleting the document")
			logger.Message("Couldn't delete the movie: "+movieID.String(), "error")
		}
	}
}
