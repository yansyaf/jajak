package poll

import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

const (
	CollName = "polls"
)

type IPollRepository interface {
	GetPolls() ([]Poll, error)
	GetPollById(id string) (Poll, error)
}

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

func (r *PollRepository) GetPollById(id string) (Poll, error) {
	singlePoll := Poll{}
	err := r.db.C(CollName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&singlePoll)
	return singlePoll, err
}
