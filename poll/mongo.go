package poll

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
	models := []Poll{}
	err := r.db.C(CollName).Find(nil).All(&models)
	return models, err
}

func (r *MongoRepository) GetPollById(id uuid.UUID) (Poll, error) {
	model := Poll{}
	//	err := r.db.C(CollName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&model)
	err := r.db.C(CollName).Find(bson.M{"id": id}).One(&model)
	return model, err
}

func (r *MongoRepository) StorePoll(in Poll) (err error) {
	err = r.db.C(CollName).Insert(in)
	return
}
