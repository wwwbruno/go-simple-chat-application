package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const port = "8080"

// RunHost takes an ip as argument
// and listen for connection
func RunHost(ip string) {
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

// RunGuest takes an ip as argument
// and looks for connection
func RunGuest(ip string) {
}
