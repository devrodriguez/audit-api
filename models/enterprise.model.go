package models

type Enterprise struct {
	Name     string `bson:"name"`
	Address  string `bson:"address"`
	Nit      string `bson:"nit"`
	Business string `bson:"business"`
	Status   bool   `bson:"status"`
}
