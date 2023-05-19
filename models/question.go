package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Text          string             `bson:"text,omitempty" json:"text,omitempty"`
	Difficulty    int                `bson:"difficulty,omitempty" json:"difficulty,omitempty"`
	Category      string             `bson:"category,omitempty" json:"category,omitempty"`
	PassageID     primitive.ObjectID `bson:"passageId,omitempty" json:"passageId,omitempty"`
	Options       []Option           `bson:"options,omitempty" json:"options,omitempty"`
	CorrectAnswer int                `bson:"correctAnswer,omitempty" json:"correctAnswer,omitempty"`
}
