package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Enterprise struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Address  string             `bson:"address" json:"address,omitempty"`
	Nit      string             `bson:"nit" json:"nit,omitempty"`
	Business string             `bson:"business" json:"business,omitempty"`
	Status   bool               `bson:"status" json:"status,omitempty"`
}
