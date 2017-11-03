package test

import (
	"testing"

	"github.com/toshim45/jajak/survey"

	"github.com/satori/go.uuid"
)

type MockSurveyRepository struct {
	models []survey.Survey
}

type MockSsurveyErrorRepository struct {
	models []survey.Survey
}

func (r *MockSurveyRepository) GetSurveys() ([]survey.Survey, error) {
	return r.models, nil
}
func (r *MockSurveyRepository) GetSurveyById(id uuid.UUID) (survey.Survey, error) {
	if len(r.models) == 0 {
		return survey.Survey{}, nil
	}
	return r.models[0], nil
}

func (r *MockSurveyRepository) StoreSurvey(model survey.Survey) error {
	return nil
}

func (r *MockSurveyRepository) StorePoll(id uuid.UUID, poll map[string]string) error {
	return nil
}

func TestGetSurveys(t *testing.T) {
	surveys := []survey.Survey{{Title: "test-1", Creator: "test-1@creator"}}

	s := survey.NewService(&MockSurveyRepository{surveys})

	if s.GetSurveys() == nil {
		t.Errorf("survey service should return mock array survey")
	}
}

func TestGetSurveyById(t *testing.T) {
	surveys := []survey.Survey{{Title: "test-2", Creator: "test-2@creator"}}

	s := survey.NewService(&MockSurveyRepository{surveys})

	if s.GetSurveyById(uuid.NewV4()).Title == "" {
		t.Errorf("survey service should return mock single survey")
	}
}
