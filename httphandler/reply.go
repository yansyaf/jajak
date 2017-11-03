package httphandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ReplyOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if data == nil {
		w.WriteHeader(201)
		return
	}
	js, err := json.Marshal(data)
	if err != nil {
		resp := map[string]interface{}{
			"error": fmt.Sprintf("%v", err),
		}
		js, _ = json.Marshal(resp)
	}
	w.WriteHeader(200)
	w.Write(js)
}

func ReplyFail(w http.ResponseWriter, status int, err error) {
	if status/1e2 == 4 {
		log.Printf("%v", err)
	} else {
		log.Panicf("%v", err)
	}

	resp := map[string]interface{}{
		"error": err.Error(),
	}
	js, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
