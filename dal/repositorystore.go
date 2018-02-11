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

func (store RepositoryStore) FindId(id string) (*models.RepositoryInfo, error) {
	// find package with limit
	repositoryInfo := &models.RepositoryInfo{}
	err := store.GetCollection().FindId(bson.ObjectIdHex(id)).One(&repositoryInfo)
	return repositoryInfo, err
}

func (store RepositoryStore) FindPackageWithinLimit(query string, skip int, limit int) (*[]models.RepositoryInfo, error) {
	// find package with limit
	repositories := &[]models.RepositoryInfo{}
	var pipeline []bson.M
	if query != "" {
		pipeline = []bson.M{
			{"$match": bson.M{"$text": bson.M{"$search": query}, "processed": true}},
			{"$sort": bson.M{"score": bson.M{"$meta": "textScore"}}},
			{"$skip": skip},
			{"$limit": limit},
			{"$project": bson.M{"readme": 0}},
		}
	} else {
		pipeline = []bson.M{
			{"$match": bson.M{"processed": true}},
			{"$skip": skip},
			{"$limit": limit},
			{"$project": bson.M{"readme": 0}},
		}
	}

	result := store.GetCollection().
		Pipe(pipeline)

	err := result.All(repositories)
	return repositories, err
}
