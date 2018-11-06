package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {}

func main() {
	router := mux.NewRouter()

	fmt.Println("Starting the apllication...")
	router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
