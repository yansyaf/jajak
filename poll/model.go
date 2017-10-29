package poll

import "gopkg.in/mgo.v2/bson"

type Poll struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Title   string        `bson:"title" json:"title"`
	Creator string        `bson:"creator" json:"creator"`
	Items   []string      `bson:"items" json:"items"`
}
