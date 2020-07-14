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

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		// w.Write([]byte(`{"message": "get called"}`,))
		json.NewEncoder(w).Encode(Users)
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", users)
}

func connectToDb() *sql.DB {
	const (
		DB_USER     = "postgres"
		DB_PASSWORD = ""
		DB_NAME     = "test"
	)

	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbConfig)
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	fmt.Println("Successfully connected!")

	return db
}

func getQuery(db *sql.DB) {
	fmt.Println("# Querying")

	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)
	defer rows.Close()
	defer db.Close()

	fmt.Println(" id | username ")
	fmt.Println("----|---------")
	for rows.Next() {
		var dbUser User
		err = rows.Scan(&dbUser.ID, &dbUser.Name)
		checkErr(err)
		fmt.Printf("%3v |%8v \n", dbUser.ID, dbUser.Name)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Server is running on localhost:8080")

	db := connectToDb()
	getQuery(db)

	handleRequests()

	Users = []User{
		User{ID: 1, Name: "Arthur"},
		User{ID: 2, Name: "Testmothy"},
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
