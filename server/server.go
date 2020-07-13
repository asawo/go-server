package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "test"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
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

	connectToDb()

	handleRequests()

	Users = []User{
		User{ID: 1, Name: "Arthur"},
		User{ID: 2, Name: "Testmothy"},
	}

	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
