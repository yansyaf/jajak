package poll

import (
	"github.com/toshim45/jajak/httputil"
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
	httputil.ThrowPanic(err)
	return polls
}

func (s *PollService) GetPollById(id string) Poll {
	singlePoll, err := s.r.GetPollById(id)
	if err == mgo.ErrNotFound {
		return Poll{}
	}
	httputil.ThrowPanic(err)
	return singlePoll
}
