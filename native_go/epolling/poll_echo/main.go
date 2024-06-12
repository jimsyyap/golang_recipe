package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
)

func echo(fd int) {
	defer syscall.Close(fd)
	var buf [1024]byte
	for {
		nbytes, e := syscall.Read(fd, buf[:])
		if nbytes > 0 {
			fmt.Printf("Receive : %s", buf)
			syscall.Write(fd, buf[:nbytes])
			fmt.Printf("Sending : %s", buf)
		}
		if e != nil {
			break
		}
	}
}

func main() {
	var event syscall.EpollEvent
	var events [4]syscall.EpollEvent

	log.Printf("\nApp PID : %d\n", os.Getpid())

	fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Socket err : ", err)
		os.Exit(1)
	}
	defer syscall.Close(fd)

	if err = syscall.SetNonblock(fd, true); err != nil {
		fmt.Println("SetNonblock err : ", err)
		os.Exit(1)
	}

	addr := syscall.SockaddrInet4{Port: 9999}
	copy(addr.Addr[:], net.ParseIP("127.0.0.1").To4())

	err = syscall.Bind(fd, &addr)
	if err != nil {
		fmt.Println("syscall.Bind err : ", err)
		os.Exit(1)
	}

	err = syscall.Listen(fd, 10)
	if err != nil {
		fmt.Println("syscall.Listen err : ", err)
		os.Exit(1)
	}

	epfd, e := syscall.EpollCreate1(0)
	if e != nil {
		log.Println("EpollCreate1 err : ", e)
		os.Exit(1)
	}
	defer syscall.Close(epfd)

	event.Events = syscall.EPOLLIN
	event.Fd = int32(fd)
	if e = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &event); e != nil {
		log.Println("EpollCtl EPOLL_CTL_ADD err : ", e)
		os.Exit(1)
	}

	for {
		n, err := syscall.EpollWait(epfd, events[:], -1)
		if err != nil {
			fmt.Println("EpollWait err : ", err)
			break
		}

		for ev := 0; ev < n; ev++ {
			if int(events[ev].Fd) == fd {
				connFd, _, err := syscall.Accept(fd)
				if err != nil {
					log.Println("Accept err : ", err)
					continue
				}

				syscall.SetNonblock(fd, true)
				event.Events = syscall.EPOLLIN
				event.Fd = int32(connFd)

				if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connFd, &event); err != nil {
					log.Print("EpollCtl err : ", connFd, err)
					os.Exit(1)
				}
			} else {
				go echo(int(events[ev].Fd))
			}
		}
	}
}

