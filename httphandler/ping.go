package httphandler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/toshim45/jajak/uptime"

	"github.com/gorilla/schema"
	"gopkg.in/mgo.v2"
)

// swagger:parameters GetPing
type ping struct {
	// Required: false
	Message string `json:"message" schema:"message"`
}

type Ping struct {
	session *mgo.Session
	uptime  *uptime.Service
}

func NewPing(s *mgo.Session, u *uptime.Service) *Ping {
	return &Ping{session: s, uptime: u}
}

func (h *Ping) GetPing(w http.ResponseWriter, r *http.Request) {
	err := h.session.Ping()
	if err != nil {
		ReplyFail(w, 500, err)
	}

	message := fmt.Sprintf("[%s] server up since %s ago", time.Now(), h.uptime.GetDuration())

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
