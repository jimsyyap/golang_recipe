package main

import (
	"log"
	"time"
)

func main() {
	serviceBDone := make(chan bool, 1)
	alldone := make(chan bool, 1)

	go serviceB(serviceBDone)
	go serviceA(serviceBDone, alldone)

	<-alldone
}

func serviceB(serviceBDone chan bool) {
	log.Println("Service B started")
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
	}

	serviceBDone <- true
	log.Println("Service B done")
}

func serviceA(serviceBDone chan bool, finish chan bool) {
	<-serviceBDone
	log.Println("Service A started")
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
	}
	log.Println("Service A done")
	finish <- true
}
