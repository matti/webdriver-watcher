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

func Check(baseURL string) (ok bool, maybe bool, stage string, status string) {
	httpClient := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := httpClient.Get(baseURL + "/sessions")
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return false, false, "sessions", err.Error()
	} else if err != nil {
		return false, false, "sessions", err.Error()
	}

	var sessions Sessions
	json.NewDecoder(resp.Body).Decode(&sessions)

	if len(sessions.Session) == 0 {
		return false, false, "sessions", "no sessions"
	}

	sessionURLURL := baseURL + "/session/" + sessions.Session[0].Id + "/url"
	resp, err = http.Get(sessionURLURL)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return false, false, "session", err.Error()
	} else if err != nil {
		return false, false, "session", err.Error()
	}

	if resp.StatusCode != 200 {
		return false, false, "session", "not 200"
	}

	return true, true, "session", "ok"
}
