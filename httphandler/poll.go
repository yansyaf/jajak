package httphandler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toshim45/jajak/poll"
)

type Poll struct {
	s *poll.Service
}

func NewPoll(service *poll.Service) *Poll {
	return &Poll{s: service}
}

func (h *Poll) GetPolls(w http.ResponseWriter, r *http.Request) {
	ReplyOk(w, h.s.GetPolls())
}

func (h *Poll) GetPollById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	singlePoll := h.s.GetPollById(id)
	if singlePoll.ID == "" {
		ReplyFail(w, 404, fmt.Errorf("id %s not found", id))
		return
	}
	ReplyOk(w, singlePoll)
}
