package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Author    string             `bson:"author,omitempty"`
	ISBN      string             `bson:"isbn,omitempty"`
	Publisher string             `bson:"publisher,omitempty"`
	Year      int                `bson:"year,omitempty"`
}
