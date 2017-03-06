package main

import (
	"flag"
	"os"

	"github.com/wwwbruno/chat-application/chat"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listen on the specified ip address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		chat.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		chat.RunGuest(connIP)
	}
}
