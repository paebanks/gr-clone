package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Rating  int32              `json:"rating" bson:"rating"`
	Comment string             `json:"comment" bson:"comment"`
}
