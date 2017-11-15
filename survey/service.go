package survey

import (
	"log"

	"github.com/forestgiant/sliceutil"
	"github.com/satori/go.uuid"
)

type Service struct {
	r SurveyRepository
}

func NewService(repository SurveyRepository) *Service {
	return &Service{r: repository}
}

func (s *Service) GetSurveys() (models []Survey, err error) {
	models, err = s.r.GetSurveys()
	return
}

func (s *Service) GetSurveyById(id uuid.UUID) (model Survey, err error) {
	model, err = s.r.GetSurveyById(id)
	return
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
			log.Printf("option available: %v, req poll: %s", survey.Options, value)
			err = ERROR_NOT_FOUND
			return
		}
	}

	err = s.r.StorePoll(id, poll)
	return
}
