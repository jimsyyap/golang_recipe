package main

import (
	"fmt"
	"log"
	"math/rand"
	"syscall"
	"time"

	seccomp "github.com/seccomp/libseccomp-golang"
)

var (
	whitelist = []string{
		"getcwd", "exit_group", "rt_sigreturn", "mkdirat", "write",
	}
	filter      *seccomp.ScmpFilter
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	log.Printf("Starting app")

	if err := configureSeccomp(); err != nil {
		log.Printf(fmt.Sprintf("Failed to load seccomp filter: %v", err))
		return
	}

	dirPath := "/tmp/" + randomString(8)
	if err := syscall.Mkdir(dirPath, 0600); err != nil {
		log.Printf("Failed creating directory: %v", err)
		return
	}
	log.Printf("Directory %s created successfully", dirPath)

	// Trying to run non whitelisted syscall
	log.Println("Trying to get current working directory")
	wd, err := syscall.Getwd()
	if err != nil {
		log.Printf("Failed getting current working directory: %v", err)
		return
	}
	log.Printf("Current working directory is: %s", wd)
}

func configureSeccomp() error {
	var err error
	filter, err = seccomp.NewFilter(seccomp.ActErrno)
	if err != nil {
		return err
	}
	for _, name := range whitelist {
		syscallID, err := seccomp.GetSyscallFromName(name)
		if err != nil {
			return err
		}
		err = filter.AddRule(syscallID, seccomp.ActAllow)
		if err != nil {
			return err
		}
	}

	err = filter.Load()
	filter.Release()
	return err
}

/*
* key points
* the code uses a security feature called seccomp to restrict what the program can do
* the whitelist defines which specific actions the program is allowed to perform
* by limiting the program's actions, you increase its security and protect your system from potential harm
* this kind of security is important for protecting your computer from malicious programs. By limiting what a program can do, you reduce the risk of it doing something bad.
 */
