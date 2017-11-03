package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/toshim45/jajak/survey"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type Survey struct {
	s *survey.Service
}

func NewSurvey(service *survey.Service) *Survey {
	return &Survey{s: service}
}

func (h *Survey) GetSurveys(w http.ResponseWriter, r *http.Request) {
	ReplyOk(w, h.s.GetSurveys())
}

func (h *Survey) GetSurveyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	model := h.s.GetSurveyById(uuid.FromStringOrNil(id))
	if model.ID == uuid.Nil {
		ReplyFail(w, 404, fmt.Errorf("id %s not found", id))
		return
	}
	ReplyOk(w, model)
}

func (h *Survey) StoreSurvey(w http.ResponseWriter, r *http.Request) {
	var model survey.Survey
	var err error
	if err = json.NewDecoder(r.Body).Decode(&model); err != nil {
		ReplyFail(w, 400, err)
		return
	}

	if model, err = h.s.StoreSurvey(model); err != nil {
		ReplyFail(w, 500, err)
		return
	}

	ReplyOk(w, model)
}

func (h *Survey) StorePoll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramId := vars["id"]
	var model map[string]string
	var err error
	var id uuid.UUID

	if id, err = uuid.FromString(paramId); err != nil {
		ReplyFail(w, 400, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&model); err != nil {
		ReplyFail(w, 400, err)
		return
	}

	if err = h.s.StorePoll(id, model); err != nil {
		ReplyFail(w, 500, err)
		return
	}

	ReplyOk(w, nil)
}
