package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sex int

const (
	Female Sex = iota
	Male
	Nonbinary
)

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Lastname    string             `json:"lastname" bson:"lastname"`
	Firstname   string             `json:"firstname" bson:"firstname"`
	Email       string             `json:"email" bson:"email"`
	Sex         Sex                `json:"sex" bson:"sex"`
	Dob         time.Time          `json:"dob" bson:"dob"`
	ReadingList []ReadingListItem  `json:"readingList" bson:"readingList"`
}
