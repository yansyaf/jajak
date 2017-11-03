package test

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	url := "http://localhost:8071/ping"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Can't Ping %s: %v", url, err)
	}
	if res == nil {
		t.Error("Expect body not nil")
	}
}
