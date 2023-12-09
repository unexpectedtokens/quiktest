package quikclient

import (
	"fmt"
	"log"
	"net/http"
)

func (q QuikClient) Ping() error {
	url := fmt.Sprintf("%s/ping", q.API_URL)

	log.Printf("Pinging Quiktest api at %s", url)
	resp, err := q.HTTPClient.Get(url)

	if err != nil {
		return fmt.Errorf("error pinging %s: %s", url, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error pinging %s: expecting 200 but got %d", q.API_URL, resp.StatusCode)
	}

	return nil
}
