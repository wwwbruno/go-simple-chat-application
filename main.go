package main

import (
	"bufio"
	"flag"
	"fmt"
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
	listener, listenErr := net.Listen("tcp", ip+":"+port)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}

	conn, acceptError := listener.Accept()
	if acceptError != nil {
		log.Fatal("Error: ", acceptError)
	}

	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message received: ", message)
}

func runGuest(ip string) {
}
