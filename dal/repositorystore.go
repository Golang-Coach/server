package dal

import (
	"github.com/Golang-Coach/server/db"
	"github.com/Golang-Coach/server/interfaces"
	"github.com/Golang-Coach/server/models"
	"gopkg.in/mgo.v2/bson"
)

type RepositoryStore struct {
	db.DataStore
}

func NewRepositoryStore(dataStore *db.DataStore) interfaces.IRepositoryStore {
	return RepositoryStore{
		*dataStore,
	}
}

func (store RepositoryStore) FindPackage(query interface{}) (*models.Repository, error) {
	// find package with limit
	repositoryInfo := &models.Repository{}
	err := store.GetCollection().Find(query).All(repositoryInfo)
	return repositoryInfo, err
}

func (store RepositoryStore) FindPackageWithinLimit(query string, skip int, limit int) (*[]models.Repository, error) {
	// find package with limit
	repositories := &[]models.Repository{}

	result := store.GetCollection().
		Find(bson.M{}).
		Select(bson.M{"readme": 0})

	if limit > 0 {
		result = result.Limit(limit)
	}

	if skip > 0 {
		result = result.Skip(skip)
	}
	err := result.All(repositories)
	return repositories, err
}
