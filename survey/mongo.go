package survey

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionName = "surveys"
)

type MongoRepository struct {
	SurveyRepository
	db *mgo.Database
}

func NewMongoRepository(database *mgo.Database) *MongoRepository {
	return &MongoRepository{db: database}
}

func (r *MongoRepository) GetSurveys() ([]Survey, error) {
	models := []Survey{}
	err := r.db.C(CollectionName).Find(nil).All(&models)
	return models, err
}

func (r *MongoRepository) GetSurveyById(id uuid.UUID) (Survey, error) {
	model := Survey{}
	//	err := r.db.C(CollectionName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&model)
	err := r.db.C(CollectionName).Find(bson.M{"_id": id}).One(&model)
	return model, err
}

func (r *MongoRepository) StoreSurvey(in Survey) (err error) {
	err = r.db.C(CollectionName).Insert(in)
	return
}

func (r *MongoRepository) StorePoll(id uuid.UUID, in Poll) (err error) {
	return
}
