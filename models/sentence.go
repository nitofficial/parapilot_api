package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sentence struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Text      string             `bson:"text,omitempty" json:"text,omitempty"`
	Order     int                `bson:"order,omitempty" json:"order,omitempty"`
	PassageID primitive.ObjectID `bson:"passageId,omitempty" json:"passageId,omitempty"`
}
