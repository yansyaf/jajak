package handlers

import (
	"net/http"

	"github.com/toshim45/jajak/domains/poll"

	"gopkg.in/mgo.v2"
)

type PollHandler struct {
	s *poll.PollService
}

func NewPollHandler(database *mgo.Database) *PollHandler {
	service := poll.NewPollService(database)
	return &PollHandler{s: service}
}

func (h *PollHandler) GetPolls(w http.ResponseWriter, r *http.Request) {
	ReplyOk(w, h.s.GetPolls())
}
