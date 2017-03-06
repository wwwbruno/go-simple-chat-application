package main

import (
	"flag"
	"fmt"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listen on the specified ip address")
	flag.Parse()

	if isHost {
		fmt.Println("is host")
	} else {
		fmt.Println("is guest")
	}
}
