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
	var resp response.Response
	json.NewDecoder(r.Body).Decode(&movie)
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	res, err := mongo.GetCollection().InsertOne(context, movie)
	movie.ID = res.InsertedID.(primitive.ObjectID)
	if !funcerr.IsErrored(err, "There was error in setting the data") {
		response.SetStatus(w, http.StatusOK)
		resp.Send(movie).RSuccess(w, "Added ID: "+movie.ID.String())
	} else {
		response.SetStatus(w, http.StatusBadRequest)
		resp.Send(movie).RError(w, "Check the object that you've been trying to add")
	}
}

// GetOne := for getting a particular movie from the database
func GetOne(w http.ResponseWriter, r *http.Request) {
	movieID, err := utils.GetID(r, "id")
	var movie mongo.MovieSchema
	var resp response.Response
	if !funcerr.IsErrored(err, "Problem with ID: "+movieID.String()) {
		context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		mongo.GetCollection().FindOne(context, mongo.MovieSchema{ID: movieID}).Decode(&movie)

		if primitive.ObjectID.IsZero(movie.ID) {
			resp.Send(movie).RError(w, "Invalid / Incorrect ID passed!")
			logger.Message(("Invalid ID: " + utils.GetIDString(movieID)), "error")
		} else {
			resp.Send(movie).RSuccess(w, "none")
			logger.Message(("Got the movie: " + utils.GetIDString(movie.ID)), "info")
		}

	}
}

// Update := to update a particular data on mongo
func Update(w http.ResponseWriter, r *http.Request) {
	var movie, newMovie mongo.MovieSchema
	var resp response.Response

	json.NewDecoder(r.Body).Decode(&newMovie)

	movieID := newMovie.ID

	context, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	mongo.GetCollection().FindOneAndReplace(context, mongo.MovieSchema{ID: movieID}, newMovie).Decode(&movie)

	if primitive.ObjectID.IsZero(movie.ID) {
		resp.Send(movie).RError(w, "Illegal update operation")
		logger.Message(("Invalid ID: " + utils.GetIDString(movieID)), "error")
	} else {
		resp.Send(newMovie).RSuccess(w, "")
		logger.Message(("Updated document: " + utils.GetIDString(movie.ID)), "info")
	}
}

// Delete := function to delete a document
func Delete(w http.ResponseWriter, r *http.Request) {
	movieID, err := utils.GetID(r, "id")
	var resp response.Response
	var movie mongo.MovieSchema

	if !funcerr.IsErrored(err, "Problem with ID: "+movieID.String()) {
		context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		deleted, errr := mongo.GetCollection().DeleteOne(context, mongo.MovieSchema{ID: movieID})
		log.Println(deleted, errr)

		if errr == nil {
			movie.ID = movieID
			log.Println(deleted.DeletedCount)
			resp.Send(movie).RSuccess(w, "Successfully deleted the movie"+movieID.String())
			logger.Message("Deleted the movie: "+movieID.String(), "info")
		} else {
			log.Println(errr)
			resp.Send(movie).RError(w, "There was an error deleting the document"+movieID.String())
			logger.Message("Couldn't delete the movie: "+movieID.String(), "error")
		}
	}
}
