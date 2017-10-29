package poll

import (
	"github.com/toshim45/jajak/httputil"
	"gopkg.in/mgo.v2"
)

type Service struct {
	r IPollRepository
}

func NewService(repository IPollRepository) *Service {
	return &Service{r: repository}
}

func (s *Service) GetPolls() []Poll {
	polls, err := s.r.GetPolls()
	httputil.ThrowPanic(err)
	return polls
}

func (s *Service) GetPollById(id string) Poll {
	singlePoll, err := s.r.GetPollById(id)
	if err == mgo.ErrNotFound {
		return Poll{}
	}
	httputil.ThrowPanic(err)
	return singlePoll
}
