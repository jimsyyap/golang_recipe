// This Go code is a small utility designed to check the status of AppArmor, a Linux kernel security module, by reading specific system files.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	appArmorEnabledPath = "/sys/module/apparmor/parameters/enabled"
	appArmorModePath    = "/sys/module/apparmor/parameters/mode"
)

func appArmorMode() (mode string) {
	content, err := os.ReadFile(appArmorModePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(content))
}

// appArmorMode reads the content of the file at appArmorModePath to the the current appArmor mode
// os.ReadFile reads the file content
// if an error occurs while reading the file, log.Fatal(err) logs the error and stops the program
// strings.TrimSpace(string(content)) trims any leading or trailing whitespace from the string

func appArmorEnabled() (support bool) {
	content, err := os.ReadFile(appArmorEnabledPath)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return
		}
		log.Fatal(err)
		return
	}
	return strings.TrimSpace(string(content)) == "y"
}

// appArmorEnabled reads the content of the file at appArmorEnabledPath to the the current appArmor mode
// os.ReadFile reads the file content
// if an error occurs while reading the file, log.Fatal(err) logs the error and stops the program
// strings.TrimSpace(string(content)) trims any leading or trailing whitespace from the string

func main() {
	fmt.Println("apparmor mode: ", appArmorMode())
	fmt.Println("apparmor is enabled: ", appArmorEnabled())
}

// the main function reads the content of the file at appArmorModePath to the the current appArmor mode
// the main function reads the content of the file at appArmorEnabledPath to the the current appArmor mode

//output
// [sudo] password for jim:
// apparmor mode:  enforce
// apparmor is enabled:  false
