package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name string
	UserName string
	ProfileUrl string
}

type Repository struct{
	ID bson.ObjectId  `bson:"_id,omitempty"`
	Name string
	Description string
	LatestRelease string
	PublishedAt time.Time
	UpdatedAt time.Time
	Owner string
	StarsCount int
	ForksCount int
	LastUpdatedBy User
	ReadMe string
	Tags []string
	Categories []string
}

