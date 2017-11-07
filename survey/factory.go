package survey

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/mgo.v2"
)

func NewMongoService(database *mgo.Database) *Service {
	repository := NewMongoRepository(database)
	service := NewService(repository)
	return service
}

func NewMySQLService(database *sqlx.DB) *Service {
	repository := NewMySQLRepository(database)
	service := NewService(repository)
	return service
}
