package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auditor struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age,omitempty" json:"age,omitempty"`
}

func (a Auditor) GetHandshake() string {
	return "Se consultó información del usuario " + a.Name
}
