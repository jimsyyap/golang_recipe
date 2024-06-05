package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	IP := "scanme.nmap.org"
	port := 80

	address := fmt.Sprintf("%s:%d", IP, port)
	fmt.Println("Connecting to " + address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Couldn't connect to %s: %v", address, err)
	}
	defer conn.Close()

	fmt.Println("Connected to " + address)
}
