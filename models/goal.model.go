package models

type Goal struct {
	Description      string `bson:"description" json:"description"`
	ShortDescription string `bson:"shortDescription" json:"shortDescription"`
}
