package main

import (
	"encoding/json"
	"fmt"
	"go-server/server/db"
	"log"
	"net/http"
	"strconv"
)

// Users contains multiple user data
var Users []db.User

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", users)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home!")
	fmt.Println("Req made to endpoint: home")
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pg := db.ConnectToDb()
	err := r.ParseForm()
	checkErr(err)

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		Users = pg.db.GetUsers()
	case "POST":
		w.WriteHeader(http.StatusCreated)
		name := r.FormValue("name")
		db.CreateUser(name)
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		s := r.FormValue("id")
		id, err := strconv.Atoi(s)
		checkErr(err)
		newName := r.FormValue("new name")
		db.UpdateUser(id, newName)
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		s := r.FormValue("id")
		id, err := strconv.Atoi(s)
		checkErr(err)
		db.DeleteUser(id)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

	json.NewEncoder(w).Encode(Users)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Server is running on localhost:8080")
	handleRequests()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
