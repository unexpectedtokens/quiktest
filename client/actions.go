package client

type Location string

const (
	HEADER Location = "header"
)

type Action struct {
	Location Location
}
