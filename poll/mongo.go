package poll

import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

const (
	CollName = "polls"
)

type MongoRepository struct {
	IPollRepository
	db *mgo.Database
}

func NewMongoRepository(database *mgo.Database) *MongoRepository {
	return &MongoRepository{db: database}
}

func (r *MongoRepository) GetPolls() ([]Poll, error) {
	polls := []Poll{}
	err := r.db.C(CollName).Find(nil).All(&polls)
	return polls, err
}

func (r *MongoRepository) GetPollById(id string) (Poll, error) {
	singlePoll := Poll{}
	err := r.db.C(CollName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&singlePoll)
	return singlePoll, err
}
