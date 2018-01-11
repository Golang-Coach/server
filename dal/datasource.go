package dal

import (
	"github.com/Golang-Coach/server/models"
	"github.com/Golang-Coach/server/interfaces"
	"gopkg.in/mgo.v2/bson"
)

type DataStore struct {
	collection interfaces.ICollection
}

func NewDataStore(collection interfaces.ICollection) interfaces.IDataStore {
	return DataStore{
		collection: collection,
	}
}

func (store DataStore) FindPackage(query interface{}) (*models.Repository, error) {
	// find package with limit
	repositoryInfo := &models.Repository{}
	err := store.collection.Find(query).All(repositoryInfo)
	return repositoryInfo, err
}

func (store DataStore) FindPackageWithinLimit(query string, skip int, limit int) (*[]models.Repository, error) {
	// find package with limit
	repositories := &[]models.Repository{}
	result := store.collection.Find(bson.M{})
	if limit > 0 {
		result = result.Limit(limit)
	}

	if skip > 0 {
		result = result.Skip(skip)
	}
	err :=  result.All(repositories)
	return repositories, err
}


