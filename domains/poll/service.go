package poll

import (
	"github.com/toshim45/jajak/utils"
	"gopkg.in/mgo.v2"
)

type PollService struct {
	r *PollRepository
}

func NewPollService(database *mgo.Database) *PollService {
	repository := NewPollRepository(database)
	return &PollService{r: repository}
}

func (s *PollService) GetPolls() []Poll {
	polls, err := s.r.GetPolls()
	utils.ThrowPanic(err)
	return polls
}
