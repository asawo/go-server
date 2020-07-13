package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User struct contains user data
type User struct {
	ID   int    `json:"userid"`
	Name string `json:"username"`
}

// Users contains multiple user data
var Users []User

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request", r.Method)
	fmt.Println("writer", w)
	fmt.Fprintf(w, "Welcome to Home!")
	fmt.Println("Req made to endpoint: home")
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Users)
	fmt.Println("Req made to endpoint: users")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", getUsers)
}

func main() {
	handleRequests()

	Users = []User{
		User{ID: 1, Name: "Arthur"},
		User{ID: 2, Name: "Testmothy"},
	}

	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
