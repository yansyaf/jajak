package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2"
)

type ping struct {
	Message string `json:"message"`
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
	if r.URL.Query().Get("message") != "" {
		message = r.URL.Query().Get("message")
	}
	ReplyOk(w, ping{Message: message})
}
