package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Goal struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Description      string             `bson:"description" json:"description"`
	ShortDescription string             `bson:"shortDescription" json:"shortDescription"`
}
