package poll

import (
	"github.com/satori/go.uuid"
)

type IPollRepository interface {
	GetPolls() ([]Poll, error)
	GetPollById(id uuid.UUID) (Poll, error)
	StorePoll(model Poll) error
}
