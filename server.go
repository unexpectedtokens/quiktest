//go:build server

package main

import "github.com/unexpectedtokens/api-tester/server"

func init() {
	Run = func() {
		server.RunServer()
	}
}
