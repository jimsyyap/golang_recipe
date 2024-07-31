package main

import (
	"enocding/json"
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

// keys are held in global variables
var (
	verifyKey, signKey []byte
)

// read the key files before starting http handlers
func init() {
	var err error
	signKey, err = os.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
		return
	}
}

// reads the login credentials, checks them and creates
