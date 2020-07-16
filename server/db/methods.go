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

}

func deleteUser(db *sql.DB, id int) {
	fmt.Println("#deleteUser()")
	sqlStatement := `
DELETE FROM test  
WHERE id = $1;`

	_, err := db.Exec(sqlStatement, id)
	checkErr(err)
	fmt.Printf("Deleted user id %d\n", id)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
