package quikclient

import (
	"fmt"
	"net/http"
)

type QuikClient struct {
	HTTPClient http.Client
	API_URL    string
}

func (q QuikClient) formatUrl(segment string) string {
	return fmt.Sprintf("%s/api%s", q.API_URL, segment)
}
