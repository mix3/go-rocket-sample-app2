package main

import (
	"net"

	"github.com/lestrrat/go-server-starter-listener"
	"github.com/mix3/go-rocket-sample-app2/webapp"
)

func main() {
	// pwd, _ := os.Getwd()

	listener, _ := ss.NewListener()
	if listener == nil {
		// Fallback if not running under Server::Starter
		var err error
		listener, err = net.Listen("tcp", ":8080")
		if err != nil {
			panic("Failed to listen to port 8080")
		}
	}

	webapp.Start(listener)
}
