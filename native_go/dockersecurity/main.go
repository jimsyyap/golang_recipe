package main

import (
	"log"
	"runtime"
)

func main() {
	log.Println("hola from inside docker image")
	log.Println("build using go version ", runtime.Version())
}
