package main

import (
	"github.com/blackhat-go/ch-6/smb/smb"
	"log"
)

func main() {
	host := "172.16.248.192"
	options := smb.Options{
		Host:        host,
		Port:        445,
		User:        "alice",
		Domain:      "corp",
		Workstation: "",
		Password:    "Password123!",
	}
	debug := false
	session, err := smb.NewSession(options, debug)
	if err != nil {
		log.Fatalln("[!] Error:", err)
	}
	defer session.Close()

	if session.IsSigningRequired {
		log.Println("[*] Signing is required")
	} else {
		log.Println("[*] Signing is not required")
	}

	if session.IsAuthenticated {
		log.Println("[*] Authentication successful")
	} else {
		log.Println("[!] Authentication failed")
	}

	if err != nil {
		log.Fatalln("[!] Error:", err)
	}
}
