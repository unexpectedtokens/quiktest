package test_request

import (
	"fmt"
	"net/http"
)

func Ping(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error pinging %s: %s", url, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error pinging %s: expecting 200 but got %d", url, resp.StatusCode)
	}

	return nil
}
