package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website.80")
	if err != nil {
		log.Fatalln("no connection:", err)
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln("io.Copy:", err)
		}
	}()
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln("io.Copy:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("listen:", err)
	}
	for {
		src, err := listener.Accept()
		if err != nil {
			log.Fatalln("accept:", err)
		}
		go handle(src)
	}
}

// output: 2024/05/28 16:19:10 listen: listen tcp :80: bind: permission denied

/*
* this program acts as a simple tcp proxy server
* it listen for incoming connections on port 80
* when a client connects, it forwards the data to joescatcam.website on port 80 and sends the response back to the clientit uses goroutiens to handle multiple connections simultaneously, making it effecient and capable of serving multiple clients at once.
* in essence, this proxy server facilitates comm between clients and the specified remote server, acting as an intemediary
*/
exit status 1
