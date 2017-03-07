package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Println("Message received: ", message)

	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	fmt.Fprint(conn, replyMessage)
}

// RunGuest takes an ip as argument
// and looks for connection
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port

	conn, dailErr := net.Dial("tcp", ipAndPort)
	if dailErr != nil {
		log.Fatal("Error: ", dailErr)
	}

	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Fprint(conn, message)
}
