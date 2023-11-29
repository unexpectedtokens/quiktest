//go:build cli

package main

import (
	"github.com/unexpectedtokens/api-tester/client"
)

func init() {
	Run = func() {
		client.RunClient()
	}
}
