package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	ID 				primitive.ObjectID	`bson:"_id,omitempty"`
	OriginURL		string				`bson:"originurl,omitempty"`
	ShortenedURL	string				`bson:"shortenedurl,omitempty"`
}
