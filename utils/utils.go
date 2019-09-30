package utils

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetID := get the ObjectID from the query params
func GetID(r *http.Request, param string) (primitive.ObjectID, error) {
	id := mux.Vars(r)[param]
	return primitive.ObjectIDFromHex(id)
}
