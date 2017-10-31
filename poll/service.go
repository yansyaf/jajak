package poll

import (
	"github.com/toshim45/jajak/httputil"

	"github.com/satori/go.uuid"
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

func (s *Service) GetPollById(id uuid.UUID) Poll {
	model, err := s.r.GetPollById(id)
	if err == mgo.ErrNotFound {
		return Poll{}
	}
	httputil.ThrowPanic(err)
	return model
}

func (s *Service) StorePoll(in Poll) (out Poll, err error) {
	in.ID = uuid.NewV4()
	err = s.r.StorePoll(in)
	out = in
	return
}
