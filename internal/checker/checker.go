package checker

import (
	"encoding/json"
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

func Check(baseURL string) (ok bool, status string) {
	httpClient := &http.Client{
		Timeout: time.Second * 1,
	}

	resp, err := httpClient.Get(baseURL + "/sessions")
	if err != nil {
		return false, err.Error()
	}
	var sessions Sessions
	json.NewDecoder(resp.Body).Decode(&sessions)

	if len(sessions.Session) == 0 {
		return false, "no sessions"
	}

	sessionURLURL := baseURL + "/session/" + sessions.Session[0].Id + "/url"
	resp, err = http.Get(sessionURLURL)
	if err != nil {
		return false, err.Error()
	}

	if resp.StatusCode != 200 {
		return false, "not 200"
	}

	return true, "ok"
}
