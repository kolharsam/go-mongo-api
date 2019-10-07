package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MovieSchema := The Schema for the Movie API
type MovieSchema struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Rating float32            `json:"rating,omitempty" bson:"rating,omitempty"`
	Review string             `json:"review,omitempty" bson:"review,omitempty"`
}
