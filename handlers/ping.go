package handlers

import (
	"net/http"

	"github.com/gorilla/schema"
	"gopkg.in/mgo.v2"
)

// swagger:parameters GetPing
type ping struct {
	// Required: false
	Message string `json:"message" schema:"message"`
}

type PingHandler struct {
	session *mgo.Session
}

func NewPingHandler(s *mgo.Session) *PingHandler {
	return &PingHandler{session: s}
}

func (h *PingHandler) GetPing(w http.ResponseWriter, r *http.Request) {
	err := h.session.Ping()
	if err != nil {
		ReplyFail(w, 500, err)
	}

	message := "hello-jajak"

	param := new(ping)
	err = schema.NewDecoder().Decode(param, r.URL.Query())
	if err != nil {
		ReplyFail(w, 400, err)
	}

	if param.Message != "" {
		message = param.Message
	}

	ReplyOk(w, ping{Message: message})
}
