package poll

import (
	"github.com/toshim45/jajak/utils"
	"gopkg.in/mgo.v2"
)

type PollService struct {
	r IPollRepository
}

func NewPollService(repository IPollRepository) *PollService {
	return &PollService{r: repository}
}

func (s *PollService) GetPolls() []Poll {
	polls, err := s.r.GetPolls()
	utils.ThrowPanic(err)
	return polls
}

func (s *PollService) GetPollById(id string) Poll {
	singlePoll, err := s.r.GetPollById(id)
	if err == mgo.ErrNotFound {
		return Poll{}
	}
	utils.ThrowPanic(err)
	return singlePoll
}
