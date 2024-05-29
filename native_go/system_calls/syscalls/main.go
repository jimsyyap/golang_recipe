package main

import (
	"log"
	s "syscall"
)

func main() {
	c := make([]byte, 512)

	log.Println("Getpid: ", s.Getpid())
	log.Println("Getpgrp: ", s.Getpgrp())
	log.Println("Getppid: ", s.Getppid())
	log.Println("Gettid: ", s.Gettid())

	_, err := s.Getcwd(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Getcwd: ", string(c))
}

// Output:
// $ src/syscalls  » go run ./main.go
// 2024/05/23 15:26:51 Getpid:  59528
// 2024/05/23 15:26:51 Getpgrp:  59471
// 2024/05/23 15:26:51 Getppid:  59471
// 2024/05/23 15:26:51 Gettid:  59528
// 2024/05/23 15:26:51 Getcwd:  /home/jim/go/src/syscalls
