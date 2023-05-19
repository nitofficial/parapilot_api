package models

type Option struct {
	ID   int    `bson:"id,omitempty" json:"id,omitempty"`
	Text string `bson:"text,omitempty" json:"text,omitempty"`
}
