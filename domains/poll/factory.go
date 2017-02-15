package poll

import "gopkg.in/mgo.v2"

func NewService(database *mgo.Database) *PollService {
	repository := NewPollRepository(database)
	service := NewPollService(repository)
	return service
}
