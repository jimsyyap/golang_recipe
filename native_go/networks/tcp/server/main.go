package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	t := net.JoinHostPort("localhost", "3333")
	l, err := net.Listen("tcp", t)
	if err != nil {
		fmt.Println("error listening", err.Error())
	}
	defer l.Close()
	log.Println("listening on port 3333")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("error accepting ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error reading", err.Error())
	}
	log.Println("received", string(buf[:len]))
	conn.Write([]byte(fmt.Sprintf("msg recv: %s", buf[:len])))
	conn.Close()
}
