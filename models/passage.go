package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Passage struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title      string             `bson:"title,omitempty" json:"title,omitempty"`
	Category   string             `bson:"category,omitempty" json:"category,omitempty"`
	Difficulty int                `bson:"difficulty,omitempty" json:"difficulty,omitempty"`
	Sentences  []string           `bson:"sentences,omitempty" json:"sentences,omitempty"`
	Questions  []Question         `bson:"questions,omitempty" json:"questions,omitempty"`
}
