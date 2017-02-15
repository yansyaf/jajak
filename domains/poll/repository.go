package poll

import "gopkg.in/mgo.v2"

const (
	CollName = "polls"
)

type PollRepository struct {
	db *mgo.Database
}

func NewPollRepository(database *mgo.Database) *PollRepository {
	return &PollRepository{db: database}
}

func (r *PollRepository) GetPolls() ([]Poll, error) {
	polls := []Poll{}
	err := r.db.C(CollName).Find(nil).All(&polls)
	return polls, err
}
