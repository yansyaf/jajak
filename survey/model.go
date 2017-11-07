package survey

import (
	"github.com/satori/go.uuid"
)

// Survey model
type Survey struct {
	ID      uuid.UUID         `bson:"_id" json:"id" db:"id"`
	Title   string            `bson:"title" json:"title" db:"title"`
	Creator string            `bson:"creator" json:"creator" db:"creator"`
	Options []string          `bson:"options" json:"options" db:"_"`
	Polls   map[string]string `bson:"polls" json:"polls" db:"_"`
}

// Poll model
type Poll struct {
	Key   string `json:"key" db:"key"`
	Value string `json:"value" db:"value"`
}
