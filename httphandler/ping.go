package httphandler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/toshim45/jajak/uptime"

	"github.com/gorilla/schema"
)

var IsShuttingDown bool
var ShuttingDownDelay time.Duration = 5 * time.Second

// swagger:parameters GetPing
type ping struct {
	// Required: false
	Message string `json:"message" schema:"message"`
}

type Ping struct {
	uptime      *uptime.Service
	persistence func() error
}

func NewPing(u *uptime.Service, p func() error) *Ping {
	return &Ping{uptime: u, persistence: p}
}

func (h *Ping) GetPing(w http.ResponseWriter, r *http.Request) {
	err := h.persistence()
	if err != nil {
		ReplyFail(w, 500, err)
	}

	message := fmt.Sprintf("[%s] server up since %s ago", time.Now(), h.uptime.GetDuration())

	if IsShuttingDown {
		ReplyFail(w, 500, fmt.Errorf("server is shutting down, delay %s s", ShuttingDownDelay.String()))
	}

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
