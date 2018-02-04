package db

import (
	"github.com/globalsign/mgo"
)

type DataStore struct {
	Session *mgo.Session
}

func (store DataStore) GetCollection() *mgo.Collection {
	// get collection
	collection := store.GetDatabase().C("repositories")
	return collection
}

func (store DataStore) GetDatabase() *mgo.Database {
	session := store.Session.Clone()
	database := &mgo.Database{session, "golang-coach"} /// TODO get this from environment variable
	return database
}

func (store *DataStore) EnsureConnected() {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("Ping session")
			//store.session.Ping()
			//Your reconnect logic here.
		}
	}()

	//Ping panics if session is closed. (see mgo.Session.Panic())
	store.Session.Ping()
}
