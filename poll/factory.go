package poll

import "gopkg.in/mgo.v2"

func New(database *mgo.Database) *Service {
	repository := NewMongoRepository(database)
	service := NewService(repository)
	return service
}
