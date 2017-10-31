package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/toshim45/jajak/poll"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
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
	singlePoll := h.s.GetPollById(uuid.FromStringOrNil(id))
	if singlePoll.ID == uuid.Nil {
		ReplyFail(w, 404, fmt.Errorf("id %s not found", id))
		return
	}
	ReplyOk(w, singlePoll)
}

func (h *Poll) StorePoll(w http.ResponseWriter, r *http.Request) {
	var model poll.Poll
	var err error
	if err = json.NewDecoder(r.Body).Decode(&model); err != nil {
		ReplyFail(w, 400, err)
	}

	if model, err = h.s.StorePoll(model); err != nil {
		ReplyFail(w, 500, err)
	}

	ReplyOk(w, model)
}
