package db

import (
	"database/sql"
	"fmt"
	"log"
)

// User struct contains user data
type User struct {
	ID   int    `json:"userid"`
	Name string `json:"username"`
}

// Users contains multiple user data
var Users []User

// ConnectToDb opens a connection to a psql db
func ConnectToDb() *sql.DB {
	const (
		DB_USER     = "postgres"
		DB_PASSWORD = ""
		DB_NAME     = "test"
	)

	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	pg, err := sql.Open("postgres", dbConfig)
	checkErr(err)

	err = pg.Ping()
	checkErr(err)

	return pg
}

// GetUsers fetches users from db
func GetUsers(db *sql.DB) []User {
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

// CreateUser inserts new user into db
func CreateUser(db *sql.DB, name string) {
	fmt.Println("#createUser()")

	sqlStatement := `INSERT INTO test (name) VALUES ($1);`
	_, err := db.Exec(sqlStatement, name)
	checkErr(err)
	fmt.Printf("Added user %s\n", name)
	GetUsers(db)
}

// UpdateUser updates user in db
func UpdateUser(db *sql.DB, id int, name string) {
	fmt.Println("#updateUser()")
	sqlStatement := `
UPDATE test 
SET "name" = $1 
WHERE id = $2;`

	_, err := db.Exec(sqlStatement, name, id)
	checkErr(err)
	fmt.Printf("Updated user id %d's name to %s\n", id, name)
	GetUsers(db)
}

// DeleteUser deletes user from db
func DeleteUser(db *sql.DB, id int) {
	fmt.Println("#deleteUser()")
	sqlStatement := `
DELETE FROM test  
WHERE id = $1;`

	_, err := db.Exec(sqlStatement, id)
	checkErr(err)
	fmt.Printf("Deleted user id %d\n", id)
	GetUsers(db)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
