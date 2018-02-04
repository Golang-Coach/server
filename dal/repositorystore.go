package dal

import (
	"github.com/Golang-Coach/server/db"
	"github.com/Golang-Coach/server/interfaces"
	"github.com/Golang-Coach/server/models"
	"github.com/globalsign/mgo/bson"
)

type RepositoryStore struct {
	db.DataStore
}

func NewRepositoryStore(dataStore *db.DataStore) interfaces.IRepositoryStore {
	return RepositoryStore{
		*dataStore,
	}
}

func (store RepositoryStore) FindPackage(query interface{}) (*models.RepositoryInfo, error) {
	// find package with limit
	repositoryInfo := &models.RepositoryInfo{}
	err := store.GetCollection().Find(query).All(repositoryInfo)
	return repositoryInfo, err
}

func (store RepositoryStore) FindPackageWithinLimit(query string, skip int, limit int) (*[]models.RepositoryInfo, error) {
	// find package with limit
	repositories := &[]models.RepositoryInfo{}

	result := store.GetCollection().
		Find(bson.M{"processed": true}).
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
