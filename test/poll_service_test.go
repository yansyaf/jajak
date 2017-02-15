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

type MockPollErrorRepository struct {
	models []poll.Poll
}

func TestGetPolls(t *testing.T) {
	polls := []poll.Poll{{Title: "test-1", Creator: "test-1@creator"}}

	s := poll.NewPollService(&MockPollRepository{polls})

	if s.GetPolls() == nil {
		t.Errorf("poll service should return mock poll")
	}
}
