package survey

import (
	"fmt"

	"github.com/toshim45/jajak/httputil"

	"github.com/forestgiant/sliceutil"
	"github.com/satori/go.uuid"
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
	httputil.ThrowPanic(err)
	return model
}

func (s *Service) StoreSurvey(in Survey) (out Survey, err error) {
	in.ID = uuid.NewV4()
	err = s.r.StoreSurvey(in)
	out = in
	return
}

func (s *Service) StorePoll(id uuid.UUID, poll map[string]string) (err error) {
	var survey Survey
	if survey, err = s.r.GetSurveyById(id); err != nil {
		return
	}

	for _, value := range poll {
		if !sliceutil.Contains(survey.Options, value) {
			err = fmt.Errorf("option available: %v", survey.Options)
			return
		}
	}

	err = s.r.StorePoll(id, poll)
	return
}
