package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/scrypt"
)

const (
	passwordFile      = "passwords.json"
	keyDerivationSalt = "some-random-salt"
	passwordLength    = 16
	numDigits         = 4
	numSymbols        = 4
	noUpper           = false
	allowRepeat       = false
)

type PasswordEntry struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordManager struct {
	Entries []PasswordEntry `json:"entries"`
}

func deriveKey(password string) ([]byte, error) {
	return scrypt.Key([]byte(password), []byte(keyDerivationSalt), 16384, 8, 1, 32)
}

func encrypt(data, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(encodedCipherText string, key []byte) ([]byte, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(encodedCipherText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func (pm *PasswordManager) load(password string) error {
	key, err := deriveKey(password)
	if err != nil {
		return err
	}

	encryptedData, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		if os.IsNotExist(err) {
			pm.Entries = []PasswordEntry{}
			return nil
		}
		return err
	}

	data, err := decrypt(string(encryptedData), key)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &pm)
}

func (pm *PasswordManager) save(password string) error {
	key, err := deriveKey(password)
	if err != nil {
		return err
	}

	data, err := json.Marshal(pm)
	if err != nil {
		return err
	}

	encryptedData, err := encrypt(data, key)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(passwordFile, []byte(encryptedData), 0600)
}

func (pm *PasswordManager) addEntry(service, username, password string) {
	pm.Entries = append(pm.Entries, PasswordEntry{Service: service, Username: username, Password: password})
}

func (pm *PasswordManager) listEntries() {
	for _, entry := range pm.Entries {
		fmt.Printf("Service: %s, Username: %s, Password: %s\n", entry.Service, entry.Username, entry.Password)
	}
}

func (pm *PasswordManager) getPassword(service string) {
	for _, entry := range pm.Entries {
		if entry.Service == service {
			fmt.Printf("Service: %s, Username: %s, Password: %s\n", entry.Service, entry.Username, entry.Password)
			return
		}
	}
	fmt.Printf("No entry found for service: %s\n", service)
}

func generateStrongPassword() (string, error) {
	return password.Generate(passwordLength, numDigits, numSymbols, noUpper, allowRepeat)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter master password: ")
	masterPassword, _ := reader.ReadString('\n')
	masterPassword = strings.TrimSpace(masterPassword)

	var pm PasswordManager
	if err := pm.load(masterPassword); err != nil {
		fmt.Println("Error loading passwords:", err)
		return
	}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a password")
		fmt.Println("2. List all passwords")
		fmt.Println("3. Get a password")
		fmt.Println("4. Generate a strong password")
		fmt.Println("5. Exit")

		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter service: ")
			service, _ := reader.ReadString('\n')
			service = strings.TrimSpace(service)

			fmt.Print("Enter username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Print("Enter password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			pm.addEntry(service, username, password)
			if err := pm.save(masterPassword); err != nil {
				fmt.Println("Error saving password:", err)
			} else {
				fmt.Println("Password saved successfully.")
			}

		case "2":
			pm.listEntries()

		case "3":
			fmt.Print("Enter service: ")
			service, _ := reader.ReadString('\n')
			service = strings.TrimSpace(service)
			pm.getPassword(service)

		case "4":
			password, err := generateStrongPassword()
			if err != nil {
				fmt.Println("Error generating strong password:", err)
			} else {
				fmt.Printf("Generated strong password: %s\n", password)
			}

		case "5":
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
