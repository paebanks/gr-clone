package models

import "time"

type BookType int

const (
	Hardcover BookType = iota
	Paperback
	Ebook
	Audiobook
)

type ReadingListItem struct {
	Book     Book      `json:"book" bson:"book"`
	Type     BookType  `json:"type" bson:"type"`
	Started  time.Time `json:"started" bson:"started"`
	Finished time.Time `json:"finished" bson:"finished"`
}
