package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type RepositoryInfo struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name          string        `json:"name"`
	Owner         string        `json:"owner"`
	FullName      string        `json:"fullName"`
	Description   string        `json:"description"`
	Stars         int           `json:"stars"`
	Forks         int           `json:"forks"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	LastUpdatedBy string        `json:"lastUpdatedBy"`
	ReadMe        string        `json:"readme"`
	Tags          []string      `json:"tags"`
	Categories    []string      `json:"categories"`
	User          User          `json:"user"`
	Processed     bool          `json:"processed"`
	ProcessedAt   time.Time     `json:"processedAt"`
	License       string        `json:"license"`
}
