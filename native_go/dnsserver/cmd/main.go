package main

import (
	"log"
	"net"
	"network.dns/server/pkg/dns"
)

type DNSConfig struct {
	dnsForwarder string
	port         int
}

func main() {
	dnsconfig := DNSConfig{
		dnsForwarder: "8.8.8.8:53",
		port:         53,
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: dnsconfig.port})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dnsFwdConn, err := net.Dial("udp", dnsconfig.dnsForwarder)
	if err != nil {
		panic(err)
	}
	defer dnsFwdConn.Close()

	if err != nil {
		panic(err)
	}

	dnsServer := dns.NewServer(conn, dns.NewUPResolver(dnsFwdConn))
	log.Printf("Starting DNS server on port %d\n", dnsconfig.port)
	dnsServer.Listen()
}
