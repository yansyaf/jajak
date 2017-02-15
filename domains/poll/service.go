package poll

import "github.com/toshim45/jajak/utils"

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
