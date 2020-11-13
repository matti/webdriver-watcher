package checker

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type Session struct {
	Id string `json:"id"`
}

type Sessions struct {
	Status  int       `json:"status"`
	Session []Session `json:"value"`
}

func Check(baseURL string) (ok bool, maybe bool, status string) {
	httpClient := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := httpClient.Get(baseURL + "/sessions")
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return false, false, err.Error()
	} else if err != nil {
		return true, false, err.Error()
	}

	var sessions Sessions
	json.NewDecoder(resp.Body).Decode(&sessions)

	if len(sessions.Session) == 0 {
		return true, false, "no sessions"
	}

	sessionURLURL := baseURL + "/session/" + sessions.Session[0].Id + "/url"
	resp, err = http.Get(sessionURLURL)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return false, false, err.Error()
	} else if err != nil {
		return true, false, err.Error()
	}

	if resp.StatusCode != 200 {
		return false, false, "not 200"
	}

	return true, true, "ok"
}
