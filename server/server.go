package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		db := connectToDb()
		getUsers(db)
		json.NewEncoder(w).Encode(Users)
	case "POST":
		w.WriteHeader(http.StatusCreated)
		err := r.ParseForm()
		checkErr(err)

		reqBody, err := ioutil.ReadAll(r.Body)
		checkErr(err)
		fmt.Printf("reqBody: %s /n", reqBody)

		db := connectToDb()
		// postUser(db, reqBody.name)
		getUsers(db)
		json.NewEncoder(w).Encode(Users)
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

func getUsers(db *sql.DB) {
	fmt.Println("# GET")

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
		Users = append(Users, dbUser)
		fmt.Printf("%3v |%8v \n", dbUser.ID, dbUser.Name)
	}
}

func postUser(db *sql.DB, name string) {
	fmt.Println("# POST")

	sqlStatement := `INSERT INTO test (name) VALUES ($1)`
	_, err := db.Exec(sqlStatement, name)
	checkErr(err)
	fmt.Printf("Added user %s", name)
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
