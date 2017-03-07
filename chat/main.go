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
	ipAndPort := ip + ":" + port

	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}
	fmt.Println("Listen to: ", ipAndPort)

	conn, acceptError := listener.Accept()
	if acceptError != nil {
		log.Fatal("Error: ", acceptError)
	}

	fmt.Println("New Connection Accepted")
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
