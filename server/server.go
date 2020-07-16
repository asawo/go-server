package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	db := connectToDb()
	err := r.ParseForm()
	checkErr(err)

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		getUsers(db)
	case "POST":
		w.WriteHeader(http.StatusCreated)
		name := r.FormValue("name")
		createUser(db, name)
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		s := r.FormValue("id")
		id, err := strconv.Atoi(s)
		checkErr(err)
		newName := r.FormValue("new name")
		updateUser(db, id, newName)
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		s := r.FormValue("id")
		id, err := strconv.Atoi(s)
		checkErr(err)
		deleteUser(db, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

	json.NewEncoder(w).Encode(Users)
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

	return db
}

func getUsers(db *sql.DB) []User {
	fmt.Println("#getUsers()")
	Users = []User{}

	rows, err := db.Query("SELECT * FROM test;")
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
	return Users
}

func createUser(db *sql.DB, name string) {
	fmt.Println("#createUser()")

	sqlStatement := `INSERT INTO test (name) VALUES ($1);`
	_, err := db.Exec(sqlStatement, name)
	checkErr(err)
	fmt.Printf("Added user %s\n", name)

	getUsers(db)
}

func updateUser(db *sql.DB, id int, name string) {
	fmt.Println("#updateUser()")
	sqlStatement := `
UPDATE test 
SET "name" = $1 
WHERE id = $2;`

	_, err := db.Exec(sqlStatement, name, id)
	checkErr(err)
	fmt.Printf("Updated user id %d's name to %s\n", id, name)

	getUsers(db)
}

func deleteUser(db *sql.DB, id int) {
	fmt.Println("#deleteUser()")
	sqlStatement := `
DELETE FROM test  
WHERE id = $1;`

	_, err := db.Exec(sqlStatement, id)
	checkErr(err)
	fmt.Printf("Deleted user id %d\n", id)

	getUsers(db)
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
