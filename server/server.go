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
	checkErr(err)
	defer db.Close()

	err = db.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	// Test query
	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)

	fmt.Println(" id | username ")
	fmt.Println("----|---------")
	for rows.Next() {
		var id int
		var username string
		err = rows.Scan(&id, &username)
		checkErr(err)
		fmt.Printf("%3v |%8v \n", id, username)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
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
