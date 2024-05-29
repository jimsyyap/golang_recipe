package main

import (
	"log"
	"math/rand"
	"os"
	"syscall"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	log.Printf("starting app")

	/*
		dirPath := "/tmp" + randomString(8)
		if err := syscall.Mkdir(dirPath, 0600); err != nil {
			log.Printf("failed to create dir: %v", err)
			return
		}
	*/

	homeDir, _ := os.UserHomeDir()
	dirPath := homeDir + "/my_temp_dir" + randomString(8)
	if err := syscall.Mkdir(dirPath, 0755); err != nil {
		log.Printf("failed to create dir: %v", err)
		return
	}
	log.Printf("directory %s created", dirPath)

	log.Printf("trying to get current working dir")
	wd, err := syscall.Getwd()
	if err != nil {
		log.Printf("failed to get current working dir: %v", err)
		return
	}
	log.Printf("current working dir: %s", wd)
}

/*
* key points
* this code creates a random string of letters
* it uses this random string to create a unique folder name in /tmp
* it tries to find the current working dir
* this code doesnt do anything dangerous, but it demonstrates how to:
* generate random strings
* create folders with unique names
* use system calls to get information about the computer's env
 */
