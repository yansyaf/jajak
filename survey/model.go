package survey

import (
	"github.com/satori/go.uuid"
)

// Survey model
type Survey struct {
	ID      uuid.UUID         `bson:"_id" json:"id"`
	Title   string            `bson:"title" json:"title"`
	Creator string            `bson:"creator" json:"creator"`
	Options []string          `bson:"options" json:"options"`
	Polls   map[string]string `bson:"polls" json:"polls"`
}
