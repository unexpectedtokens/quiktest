package main

import (
	"os"

	"github.com/unexpectedtokens/api-tester/client"
	"github.com/unexpectedtokens/api-tester/server"
)

func main() {
	if len(os.Args) == 2 {
		server.RunServer()
	} else {
		client.RunClient()
	}
}
