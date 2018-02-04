package interfaces

import "gopkg.in/mgo.v2"

type ICollection interface {
	Insert(...interface{}) error
	Update(selector interface{}, update interface{}) error
	Find(query interface{}) *mgo.Query
}
