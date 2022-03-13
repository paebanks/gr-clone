package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Title         string             `json:"title" bson:"title"`
	Author        string             `json:"author" bson:"author"`
	Publisher     string             `json:"published" bson:"publisher"`
	DatePublished time.Time          `json:"datePublished" bson:"datePublished"`
	Volume        int32              `json:"volume,omitempty" bson:"volume,omitempty"`
	Edition       int32              `json:"edition,omitempty" bson:"edition,omitempty"`
	Reviews       []Review           `json:"reviews" bson:"reviews"`
}
