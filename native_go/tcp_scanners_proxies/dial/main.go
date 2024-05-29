package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("connection successful")
	}
}

// output is "connection successful"

/*
* this program attempts to establish a TCP connection to 'localhost' on port 8080. If it fails to make the connection, it prints the error message to the console. If it succeeds, prints output
*
 */
