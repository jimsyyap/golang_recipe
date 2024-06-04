package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 3333,
		IP:   net.ParseIP("0.0.0.0"),
	})

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("listening on %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 512)
		_, remoteAddr, err := conn.ReadFromUDP(message[:])
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Millisecond)
		conn.WriteToUDP([]byte("\n"), remoteAddr)
	}
}
