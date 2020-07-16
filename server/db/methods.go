package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgresDb is a sql.DB struct
type PostgresDb struct {
	pg *sql.DB
}

// User struct contains user data
type User struct {
	ID   int    `json:"userid"`
	Name string `json:"username"`
}

// Users contains multiple user data
var Users []User

// ConnectToDb opens a connection to a psql db
func (db *PostgresDb) ConnectToDb() *PostgresDb {
	const (
		DB_USER     = "postgres"
		DB_PASSWORD = ""
		DB_NAME     = "test"
	)

	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	pg, err := db.pg.Open("postgres", dbConfig)
	checkErr(err)

	err = pg.Ping()
	checkErr(err)

	return pg
}

// GetUsers fetches users from db
func (db *PostgresDb) GetUsers() []User {
	fmt.Println("#getUsers()")
	Users = []User{}

	rows, err := db.pg.Query("SELECT * FROM test;")
	checkErr(err)
	defer rows.Close()
	defer db.pg.Close()

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
func (db *PostgresDb) CreateUser(name string) {
	fmt.Println("#createUser()")

	sqlStatement := `INSERT INTO test (name) VALUES ($1);`
	_, err := db.pg.Exec(sqlStatement, name)
	checkErr(err)
	fmt.Printf("Added user %s\n", name)
	db.GetUsers()
}

// UpdateUser updates user in db
func (db *PostgresDb) UpdateUser(id int, name string) {
	fmt.Println("#updateUser()")
	sqlStatement := `
UPDATE test 
SET "name" = $1 
WHERE id = $2;`

	_, err := db.pg.Exec(sqlStatement, name, id)
	checkErr(err)
	fmt.Printf("Updated user id %d's name to %s\n", id, name)
	db.GetUsers()
}

// DeleteUser deletes user from db
func (db *PostgresDb) DeleteUser(id int) {
	fmt.Println("#deleteUser()")
	sqlStatement := `
DELETE FROM test  
WHERE id = $1;`

	_, err := db.pg.Exec(sqlStatement, id)
	checkErr(err)
	fmt.Printf("Deleted user id %d\n", id)
	db.GetUsers()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
