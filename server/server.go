package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// User struct contains user data
type User struct {
	ID   int    `json:"userid"`
	Name string `json:"username"`
}

// Users contains multiple user data
var Users []User

func home(w http.ResponseWriter, r *http.Request) {
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

func connectToDb() {
	const (
		DB_USER     = "postgres"
		DB_PASSWORD = ""
		DB_NAME     = "test"
	)

	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	var myUser User
	query := `SELECT * FROM test WHERE id = $1`
	err = db.QueryRow(query, 1).Scan(&myUser.ID, &myUser.Name)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Hi %s, welcome back!\n", myUser.Name)
}

func main() {
	fmt.Println("Server is running on localhost:8080")

	connectToDb()

	handleRequests()

	Users = []User{
		User{ID: 1, Name: "Arthur"},
		User{ID: 2, Name: "Testmothy"},
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
