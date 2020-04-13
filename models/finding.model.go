package models

type Finding struct {
	Description string   `bson:"description" json:"description"`
	Files       []string `bson:"files" json:"files"`
}
