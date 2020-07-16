package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// User struct contains user data
type User struct {
	ID   int    `json:"userid"`
	Name string `json:"username"`
}

// Users contains multiple user data
var Users []User

// PostgresDb is a sql.DB struct
type PostgresDb struct {
	db *sql.DB
}

var postgres PostgresDb

// ConnectToDb opens a connection to a psql db
func ConnectToDb() PostgresDb {
	const (
		DB_USER     = "postgres"
		DB_PASSWORD = ""
		DB_NAME     = "test"
	)

	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	var err error
	postgres.db, err = sql.Open("postgres", dbConfig)
	checkErr(err)

	err = postgres.db.Ping()
	checkErr(err)

	return postgres
}

// GetUsers fetches users from db
func (postgres *PostgresDb) GetUsers() []User {
	fmt.Println("#getUsers()")
	Users = []User{}

	rows, err := postgres.db.Query("SELECT * FROM test;")
	checkErr(err)
	defer rows.Close()
	defer postgres.db.Close()

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
func (postgres *PostgresDb) CreateUser(name string) {
	fmt.Println("#createUser()")

	sqlStatement := `INSERT INTO test (name) VALUES ($1);`
	_, err := postgres.db.Exec(sqlStatement, name)
	checkErr(err)
	fmt.Printf("Added user %s\n", name)
	postgres.GetUsers()
}

// UpdateUser updates user in db
func (postgres *PostgresDb) UpdateUser(id int, name string) {
	fmt.Println("#updateUser()")
	sqlStatement := `
UPDATE test 
SET "name" = $1 
WHERE id = $2;`

	_, err := postgres.db.Exec(sqlStatement, name, id)
	checkErr(err)
	fmt.Printf("Updated user id %d's name to %s\n", id, name)
	postgres.GetUsers()
}

// DeleteUser deletes user from db
func (postgres *PostgresDb) DeleteUser(id int) {
	fmt.Println("#deleteUser()")
	sqlStatement := `
DELETE FROM test  
WHERE id = $1;`

	_, err := postgres.db.Exec(sqlStatement, id)
	checkErr(err)
	fmt.Printf("Deleted user id %d\n", id)
	postgres.GetUsers()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
