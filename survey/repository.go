package survey

import (
	"github.com/satori/go.uuid"
)

type SurveyRepository interface {
	GetSurveys() (models []Survey, err error)
	GetSurveyById(id uuid.UUID) (model Survey, err error)
	StoreSurvey(model Survey) (err error)
	StorePoll(id uuid.UUID, poll map[string]string) (err error)
}
