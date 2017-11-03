package survey

import (
	"github.com/satori/go.uuid"
)

type SurveyRepository interface {
	GetSurveys() ([]Survey, error)
	GetSurveyById(id uuid.UUID) (Survey, error)
	StoreSurvey(model Survey) error
	StorePoll(id uuid.UUID, poll map[string]string) error
}
