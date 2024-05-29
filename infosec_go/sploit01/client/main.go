package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jimsyyap/sploit01/rpc"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	if host == "" || pass == "" {
		log.Fatalln("missing required env MSFHOST or MSFPASS")
	}

	msf, err := rpc.New(host, user, pass)
	if err != nil {
		log.Fatalln(err)
	}
	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("sessions")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
