//Package models provide models for solution
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Audit struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Auditor    Auditor            `bson:"auditor" json:"auditor"`
	Enterprise Enterprise         `bson:"enterprise" json:"enterprise"`
	Goals      []Goal             `bson:"goals" json:"goals"`
	Findings   []Finding          `bson:"findings" json:"findings"`
	StartDate  string             `bson:"startDate" json:"startDate"`
	EndDate    string             `bson:"endDate" json:"endDate"`
}
