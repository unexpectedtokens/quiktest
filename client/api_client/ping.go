package quikclient

import (
	"fmt"
	"net/http"
)

func (q QuikClient) Ping() error {
	resp, err := q.HTTPClient.Get(q.API_URL)

	if err != nil {
		return fmt.Errorf("error pinging %s: %s", q.API_URL, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error pinging %s: expecting 200 but got %d", q.API_URL, resp.StatusCode)
	}

	return nil
}
