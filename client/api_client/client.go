package quikclient

import "net/http"

type QuikClient struct {
	HTTPClient http.Client
	API_URL    string
}
