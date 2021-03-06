package survey

import (
	"strings"

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

func (r *MongoRepository) GetSurveys() (models []Survey, err error) {
	err = r.db.C(CollectionName).Find(nil).All(&models)
	return
}

func (r *MongoRepository) GetSurveyById(id uuid.UUID) (model Survey, err error) {
	//	err = r.db.C(CollectionName).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&model)
	err = r.db.C(CollectionName).Find(bson.M{"_id": id}).One(&model)
	if err == mgo.ErrNotFound {
		err = ERROR_NOT_FOUND
	}
	return
}

func (r *MongoRepository) StoreSurvey(in Survey) (err error) {
	err = r.db.C(CollectionName).Insert(in)
	return
}

func (r *MongoRepository) StorePoll(id uuid.UUID, in map[string]string) (err error) {
	for k, v := range in {
		err = r.db.C(CollectionName).UpdateId(id, bson.M{"$set": bson.M{"polls." + strings.Split(k, ".")[0]: v}})
	}

	return
}
