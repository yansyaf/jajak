package survey

import (
	"github.com/toshim45/jajak/httputil"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

type Service struct {
	r SurveyRepository
}

func NewService(repository SurveyRepository) *Service {
	return &Service{r: repository}
}

func (s *Service) GetSurveys() []Survey {
	surveys, err := s.r.GetSurveys()
	httputil.ThrowPanic(err)
	return surveys
}

func (s *Service) GetSurveyById(id uuid.UUID) Survey {
	model, err := s.r.GetSurveyById(id)
	if err == mgo.ErrNotFound {
		return Survey{}
	}
	httputil.ThrowPanic(err)
	return model
}

func (s *Service) StoreSurvey(in Survey) (out Survey, err error) {
	in.ID = uuid.NewV4()
	err = s.r.StoreSurvey(in)
	out = in
	return
}

func (s *Service) StorePoll(id uuid.UUID, in Poll) (err error) {
	return s.r.StorePoll(id, in)
}
