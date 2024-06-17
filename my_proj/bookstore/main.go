package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	ServerAddress string
	DbUser        string
	DbPassword    string
	DbHost        string
	DbName        string
}

func loadConfiguration() Configuration {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return Configuration{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DbUser:        os.Getenv("DB_USER"),
		DbPassword:    os.Getenv("DB_PASSWORD"),
		DbHost:        os.Getenv("DB_HOST"),
		DbName:        os.Getenv("DB_NAME"),
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	config := loadConfiguration()
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}
