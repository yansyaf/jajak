package survey

//import "gopkg.in/mgo.v2"

//func New(database *mgo.Database) *Service {
//	repository := NewMongoRepository(database)
//	service := NewService(repository)
//	return service
//}

import "github.com/jmoiron/sqlx"

func New(database *sqlx.DB) *Service {
	repository := NewMySQLRepository(database)
	service := NewService(repository)
	return service
}
