package poll

import (
	"github.com/satori/go.uuid"
)

type Poll struct {
	ID      uuid.UUID `bson:"id" json:"id"`
	Title   string    `bson:"title" json:"title"`
	Creator string    `bson:"creator" json:"creator"`
	Items   []string  `bson:"items" json:"items"`
}
