package handlers

import (
	"net/http"

	"github.com/toshim45/jajak/domains/poll"
)

type PollHandler struct {
	s *poll.PollService
}

func NewPollHandler(service *poll.PollService) *PollHandler {
	return &PollHandler{s: service}
}

func (h *PollHandler) GetPolls(w http.ResponseWriter, r *http.Request) {
	ReplyOk(w, h.s.GetPolls())
}
