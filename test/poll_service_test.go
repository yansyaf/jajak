package test

import (
	"testing"

	"github.com/toshim45/jajak/domains/poll"
)

type MockPollRepository struct {
	models []poll.Poll
}

func (r *MockPollRepository) GetPolls() ([]poll.Poll, error) {
	return r.models, nil
}
func (r *MockPollRepository) GetPollById(id string) (poll.Poll, error) {
	if len(r.models) == 0 {
		return poll.Poll{}, nil
	}
	return r.models[0], nil
}

type MockPollErrorRepository struct {
	models []poll.Poll
}

func TestGetPolls(t *testing.T) {
	polls := []poll.Poll{{Title: "test-1", Creator: "test-1@creator"}}

	s := poll.NewPollService(&MockPollRepository{polls})

	if s.GetPolls() == nil {
		t.Errorf("poll service should return mock array poll")
	}
}

func TestGetPollById(t *testing.T) {
	polls := []poll.Poll{{Title: "test-2", Creator: "test-2@creator"}}

	s := poll.NewPollService(&MockPollRepository{polls})

	if s.GetPollById("id-test-2").Title == "" {
		t.Errorf("poll service should return mock single poll")
	}
}
