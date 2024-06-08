package main

import (
    "fmt"
    "net"
)

func main() {
    routingTable := []struct {
        subnetmask net.IP
        network    net.IP
        nextHop    net.IP
    }{
        {net.IP{255, 255, 255, 240}, net.IP{192, 17, 7, 208}, net.IP{192, 12, 7, 15}},
        {net.IP{255, 255, 255, 240}, net.IP{192, 17, 7, 144}, net.IP{192, 12, 7, 67}},
        {net.IP{255, 255, 255, 0}, net.IP{192, 17, 7, 0}, net.IP{192, 12, 7, 251}},
        {net.IP{0, 0, 0, 0}, net.IP{0, 0, 0, 0}, net.IP{10, 10, 10, 10}},
    }
    incomingPacketsToRoute := []struct {
        sourceAddr net.IP
        destinationAddr net.IP
        data string
    }{

    }
