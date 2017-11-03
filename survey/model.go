package survey

import (
	"github.com/satori/go.uuid"
)

// Survey model
type Survey struct {
	ID      uuid.UUID `bson:"_id" json:"id"`
	Title   string    `bson:"title" json:"title"`
	Creator string    `bson:"creator" json:"creator"`
	Options []string  `bson:"options" json:"options"`
	Polls   []Poll    `bson:"polls" json:"polls"`
}

// Poll model tight to Survey
type Poll struct {
	Key    string `bson:"key" json:"key"`
	Poller string `bson:"poller" json:"poller"`
}
