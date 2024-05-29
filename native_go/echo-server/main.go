package main

import (
	"bufio"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data.
/*
func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data.
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err != nil && err != io.EOF {
			log.Println("Unexpected error")
			break
		}

		if err == io.EOF && size == 0 {
			log.Println("Client disconnected")
			break
		}

		log.Printf("Received %d bytes: %s", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}
*/

// using bufio package that wraps reader and writer to create a buffered IO mechanism
func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("Unable to read data")
	}

	log.Printf("Received %d bytes: %s", len(s), s)
	log.Println("writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	if err := writer.Flush(); err != nil {
		log.Fatalln("Unable to flush data")
	}
}

func main() {
	// Listen on TCP port 20080
	listener, err := net.Listen("tcp", ":20080")

	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening on 0.0.0.0:20080")

	for {
		// Accept connections
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle connections in a new goroutine.
		go echo(conn)
	}
}

/*
* this program is a simple TCP echo server
* it listens for incoming connections on port 20080
* and echos received data back to the client
* it can handle multiple connections at the same time using goroutines
* this type of server is often used for testing network connections and debugging purposes
 */
