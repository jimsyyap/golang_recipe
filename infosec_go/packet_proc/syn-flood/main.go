package main

import (
	"fmt"
	"log"
	// "net"
	// "os"
	// "strings"
	// "time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snaplen  = int32(320)
	promisc  = true
	timeout  = pcap.BlockForever
	filter   = "tcp[13] == 0x11 or tcp[13] == 0x10 or tcp[13] == 0x18"
	devFound = false
	results  = make(map[string]int)
)

func capture(iface, target string) {
	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Capturing packets")
	for packet := range source.Packets() {
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			continue
		}
		transportLayer := packet.TransportLayer()
		if transportLayer == nil {
			continue
		}

		srcHost := networkLayer.NetworkFlow().Src().String()
		srcPort := transportLayer.TransportFlow().Src().String()

		if srcHost != target {
			continue
		}
		results[srcPort] += 1
	}
}
