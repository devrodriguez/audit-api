package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Finding struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Description string             `bson:"description" json:"description"`
	Files       []string           `bson:"files" json:"files"`
}
