package httphandler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toshim45/jajak/poll"
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

func (h *PollHandler) GetPollById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	singlePoll := h.s.GetPollById(id)
	if singlePoll.ID == "" {
		ReplyFail(w, 404, fmt.Errorf("id %s not found", id))
		return
	}
	ReplyOk(w, singlePoll)
}
