package checker

import (
	"net"
	"net/http"
	"time"
)

func Cdp(baseURL string) (ok bool, status string) {
	httpClient := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := httpClient.Get(baseURL + "/json")
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return false, err.Error()
	} else if err != nil {
		return false, err.Error()
	}

	if resp.StatusCode != 200 {
		return false, "not 200"
	}

	return true, "ok"
}
