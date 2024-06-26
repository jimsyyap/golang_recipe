package main

import (
	"log"
	"net"
	"runtime"
	"time"
)

func listen(i int, connection *net.UDPConn, quit chan struct{}) {
	buffer := make([]byte, 1024)
	for {
		_, remote, err := connection.ReadFromUDP(buffer)
		if err != nil {
			break
		}

		time.Sleep(10 * time.Millisecond)
		connection.WriteToUDP([]byte("\n"), remote)
	}
	quit <- struct{}{}
}

func main() {
	addr := net.UDPAddr{
		Port: 3333,
	}
	connection, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	quit := make(chan struct{})
	for i := 0; i < runtime.NumCPU(); i++ {
		log.Println("Starting worker", i)
		id := i
		go listen(id, connection, quit)
	}
	<-quit
}
