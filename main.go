package main

import (
	"flag"
	"log"
	"net"
	"os"
)

const port = "8080"

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listen on the specified ip address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		runGuest(connIP)
	}
}

func runHost(ip string) {
	_, listenErr := net.Listen("tcp", ip+":"+port)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}
}

func runGuest(ip string) {
}
